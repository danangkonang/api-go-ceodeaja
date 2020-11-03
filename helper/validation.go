package helper

import (
	"database/sql"
	"encoding/json"
	"regexp"
	"time"

	"github.com/danangkonang/ceodeaja-go/config"
)

type Registrasion struct {
	Id         int
	Username   string
	Email      string
	Password   string
	Active     bool
	Created_at time.Time
	Updated_at time.Time
}

func IsEmail(email string) ([]byte, bool) {
	emailRegexString1 := "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	// emailRegexString2 := "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	re := regexp.MustCompile(emailRegexString1)
	isEmail := re.MatchString(email)
	if isEmail {
		res := map[string]interface{}{
			"success": true,
			"error": map[string]interface{}{
				"message": "valid email",
			},
		}
		r, _ := json.Marshal(res)
		return r, true
	} else {
		res := map[string]interface{}{
			"success": false,
			"error": map[string]interface{}{
				"message": "invalid email",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	}
}

func EmailRegistered(email string) ([]byte, bool) {
	db := config.Connect()
	sqlStatement := `SELECT email FROM users WHERE email=$1;`
	var user Registrasion
	row := db.QueryRow(sqlStatement, email)
	err := row.Scan(&user.Email)
	if err != nil {
		res := map[string]interface{}{
			"success": false,
			"error": map[string]interface{}{
				"message": "email not registered",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	} else {
		res := map[string]interface{}{
			"success": true,
			"error": map[string]interface{}{
				"message": "email already registered",
			},
		}
		r, _ := json.Marshal(res)
		return r, true
	}
}

func Required(value, name string) ([]byte, bool) {
	if len(value) == 0 {
		res := map[string]interface{}{
			"success": false,
			"error": map[string]interface{}{
				"message": name + " is required",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	} else {
		res := map[string]interface{}{
			"success": true,
			"error": map[string]interface{}{
				"message": "ok",
			},
		}
		r, _ := json.Marshal(res)
		return r, true
	}
}

func Length(text string, min, max int) bool {
	if len(text) < min {
		return false
	} else if len(text) > max {
		return false
	} else {
		return true
	}
}

func IsUser(username string) bool {
	db := config.Connect()
	sqlStatement := `SELECT username FROM users WHERE username=$1;`
	var user Registrasion
	row := db.QueryRow(sqlStatement, username)
	err := row.Scan(&user.Username)
	switch err {
	case sql.ErrNoRows:
		return true
	case nil:
		return false
	default:
		panic(err)
	}
}

func MinMax(value, name string, min, max int) ([]byte, bool) {
	if len(value) < min {
		res := map[string]interface{}{
			"success": false,
			"error": map[string]interface{}{
				"message": name + " min length 8",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	} else if len(value) > max {
		res := map[string]interface{}{
			"success": false,
			"error": map[string]interface{}{
				"message": name + " max length 32",
			},
		}
		r, _ := json.Marshal(res)
		return r, false
	} else {
		res := map[string]interface{}{
			"success": true,
			"error": map[string]interface{}{
				"message": name + "ok",
			},
		}
		r, _ := json.Marshal(res)
		return r, true
	}
}
func ValidForm() {

}
func IsDomoin() {

}
