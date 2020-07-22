package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

// func makeResponJson(value []byte, w http.ResponseWriter) {
// 	res := string(value)
// 	var responRequired map[string]interface{}
// 	json.Unmarshal([]byte(res), &responRequired)
// 	w.Write(value)
// }

func responseJsonError(w http.ResponseWriter, message string) {
	// if err != nil {
	// 	JSON(w, statusCode, struct {
	// 		Error string `json:"error"`
	// 	}{
	// 		Error: err.Error(),
	// 	})
	// 	return
	// }
	// JSON(w, http.StatusBadRequest, nil)
	res := map[string]interface{}{
		"success": false,
		"error": map[string]interface{}{
			"message": message,
		},
	}
	r, _ := json.Marshal(res)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(r)
	return
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	var key = os.Getenv("KEY")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			responseJsonError(w, "Unauthorized")
			// ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

		if err != nil {
			responseJsonError(w, "token tidak di kenal")
			// ERROR(w, http.StatusBadRequest, errors.New("token tidak di kenal"))
			return
		}

		if token.Valid {
			next(w, r)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				responseJsonError(w, "token tidak di kenal 1")
				// ERROR(w, http.StatusUnauthorized, errors.New("token tidak di kenal 1"))
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				responseJsonError(w, "token expired")
				// ERROR(w, http.StatusUnauthorized, errors.New("token expired"))
				return
			} else {
				responseJsonError(w, "token tidak di kenal 2")
				// ERROR(w, http.StatusUnauthorized, errors.New("token tidak di kenal 2"))
				return
			}
		} else {
			responseJsonError(w, "token tidak di kenal 3")
			// ERROR(w, http.StatusUnauthorized, errors.New("token tidak di kenal 3"))
			return
		}
	}
}
