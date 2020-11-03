package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danangkonang/ceodeaja-go/helper"
)

func RefresToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorizationHeader := r.Header.Get("Authorization")
	respon := helper.DataToken(authorizationHeader)
	userId := respon["Id"]
	userEmail := respon["Email"]
	id := fmt.Sprintf("%v", userId)
	email := fmt.Sprintf("%v", userEmail)
	newToken := helper.MakeToken(id, email)
	res := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"token": newToken,
		},
	}
	o, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
