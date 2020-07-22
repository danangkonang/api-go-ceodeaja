package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danangkonang/rest-api/helper"
)

type ResetPassword struct {
	Email          string
	NewPassword    string
	RepeatPassword string
	Created_at     time.Time
	Updated_at     time.Time
}

func RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decode := json.NewDecoder(r.Body)
	var req ResetPassword
	err := decode.Decode(&req)
	if err != nil {
		panic(err)
	}

	msgRequired, isRequired := helper.Required(req.Email, "email")
	if !isRequired {
		helper.MakeResponJson(msgRequired, w)
		return
	}

	msgIsEmail, isEmail := helper.IsEmail(req.Email)
	if !isEmail {
		helper.MakeResponJson(msgIsEmail, w)
		return
	}

	msgEmailRegistered, isRegistered := helper.EmailRegistered(req.Email)
	if !isRegistered {
		helper.MakeResponJson(msgEmailRegistered, w)
		return
	}

	// sendLinkResetPassword()
}

func NewPassword(w http.ResponseWriter, r *http.Request) {

}

// func makeResponJson(value []byte, w http.ResponseWriter) {
// 	res := string(value)
// 	var responRequired map[string]interface{}
// 	json.Unmarshal([]byte(res), &responRequired)
// 	w.Write(value)
// }
