package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/danangkonang/ceodeaja-go/config"
	"github.com/danangkonang/ceodeaja-go/helper"
	"golang.org/x/crypto/bcrypt"
)

type Registra struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Registrasi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decode := json.NewDecoder(r.Body)
	var req Registra
	err := decode.Decode(&req)
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

	msgRequiredEmail, requiredEmail := helper.Required(req.Email, "email")
	if !requiredEmail {
		helper.MakeResponJson(msgRequiredEmail, w)
		return
	}

	msgIsEmail, isEmail := helper.IsEmail(req.Email)
	if !isEmail {
		helper.MakeResponJson(msgIsEmail, w)
		return
	}

	msgEmailRegistered, EmailRegistered := helper.EmailRegistered(req.Email)
	if EmailRegistered {
		helper.MakeResponJson(msgEmailRegistered, w)
		return
	}

	msgRequiredPassword, requiredPassword := helper.Required(req.Password, "password")
	if !requiredPassword {
		helper.MakeResponJson(msgRequiredPassword, w)
		return
	}

	msgMinMaxPassword, minMaxPassword := helper.MinMax(req.Password, "password", 8, 32)
	if !minMaxPassword {
		helper.MakeResponJson(msgMinMaxPassword, w)
		return
	}

	trimPassSpace := strings.Join(strings.Fields(req.Password), "")
	trimPassUnix := strings.TrimFunc(trimPassSpace, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	hashed, _ := HashPassword(trimPassUnix)
	password := string(hashed)
	re := saveRegistrasi(req.Email, password)
	if re {
		res := map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"message": "success registrasi",
			},
		}
		userJson, _ := json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		w.Write(userJson)
	}
}

func storeUser(email, password string) {
	// model.Store(email, password)
}

func saveRegistrasi(email, password string) bool {
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO users(user_id,email,password,user_role,is_active,is_verify,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(helper.UnixRandomString(32), email, password, 1, false, false, waktu, waktu)
	if err != nil {
		panic(err.Error())
	}
	return true
}

// func cekDobleName(name string) bool {
// 	db := config.Connect()
// 	sqlStatement := `SELECT * FROM users WHERE username=$1;`
// 	var user Registra
// 	row := db.QueryRow(sqlStatement, name)
// 	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Active, &user.CreatedAt, &user.UpdatedAt)
// 	switch err {
// 	case sql.ErrNoRows:
// 		return true
// 	case nil:
// 		return false
// 	default:
// 		panic(err)
// 	}
// }

func HashPassword(password string) (string, error) {
	pass := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func ActivasiAccunt(w http.ResponseWriter, r *http.Request) {
	fmt.Println("true")
	// var url string = "http://localhost:900"
	// eml := "dngrifai21@gmail.com"
	// tkn := "rkdmeozhrwtkclso"
	// e := email.NewEmail()
	// e.From = "ceodeaja <admin@ceodeaja.com>"
	// e.To = []string{"dngrifai21@gmail.com"}
	// e.Subject = "Hello"
	// e.HTML = []byte("silahkan klik link untuk aktifasi akun<a href=' " + url + "/registrasi/aktifasi?email=" + eml + "&token=" + tkn + " '> " + url + "/registrasi/aktifasi?email=" + eml + "&token=" + tkn + " </a> <br>link akan kadaluarsa dalam waktu 24 jam")
	// err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "ceodeaja@gmail.com", "rkdmeozhrwtkclso", "smtp.gmail.com"))
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("Mail sent!")
}
