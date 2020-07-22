package helper

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/danangkonang/rest-api/config"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Id         string
	Username   string
	Email      string
	Password   string
	Active     bool
	Created_at time.Time
	Updated_at time.Time
}

func GetUserByEmail(email string) ([]byte, bool) {
	db := config.Connect()
	sqlStatement := `SELECT id,email,password FROM users WHERE email=$1;`
	var user LoginUser
	row := db.QueryRow(sqlStatement, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password)
	switch err {
	case sql.ErrNoRows:
		res := map[string]interface{}{
			"error": map[string]interface{}{
				"message": "email belum terdaftar",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	case nil:
		res := map[string]interface{}{
			"id":       user.Id,
			"email":    user.Email,
			"password": user.Password,
		}
		r, _ := json.Marshal(res)
		return r, true
	default:
		panic(err)
	}
}

func VeryfiPasswordDb(password, passwordHash string) ([]byte, bool) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	// return nil, err == nil
	if err != nil {
		res := map[string]interface{}{
			"error": map[string]interface{}{
				"message": "wrong password",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	} else {
		res := map[string]interface{}{
			"error": map[string]interface{}{
				"message": "true password",
			},
		}
		r, _ := json.Marshal(res)
		return r, true
	}
}
