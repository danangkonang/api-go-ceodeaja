package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danangkonang/rest-api/helper"
)

type UserLogin struct {
	Id         int
	Username   string
	Email      string
	Password   string
	Active     bool
	Created_at time.Time
	Updated_at time.Time
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decode := json.NewDecoder(r.Body)
	var t UserLogin
	err := decode.Decode(&t)
	if err != nil {
		res := map[string]interface{}{
			"success": false,
			"error": map[string]interface{}{
				"message": "form can't null",
			},
		}
		r, _ := json.Marshal(res)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(r)
		return
	}

	msgRequiredEmail, requiredEmail := helper.Required(t.Email, "email")
	if !requiredEmail {
		helper.MakeResponJson(msgRequiredEmail, w)
		return
	}
	msgRequiredPassword, requiredPassword := helper.Required(t.Password, "password")
	if !requiredPassword {
		helper.MakeResponJson(msgRequiredPassword, w)
		return
	}

	msgGetUserByEmail, isGetUserByEmail := helper.GetUserByEmail(t.Email)
	if !isGetUserByEmail {
		helper.MakeResponJson(msgGetUserByEmail, w)
		return
	}
	var result map[string]interface{}
	json.Unmarshal(msgGetUserByEmail, &result)
	id := fmt.Sprintf("%v", result["id"])
	email := fmt.Sprintf("%v", result["email"])
	passUser := t.Password
	passdDb := fmt.Sprintf("%v", result["password"])

	msgTruePassword, isTruePassword := helper.VeryfiPasswordDb(passUser, passdDb)
	if !isTruePassword {
		helper.MakeResponJson(msgTruePassword, w)
		return
	}

	token := helper.MakeToken(id, email)
	res := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"token": token,
		},
	}
	o, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
