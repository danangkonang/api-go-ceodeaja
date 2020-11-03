package model

import (
	"time"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Store(email, password string) bool {
	t := time.Now()
	waktu := t.Format("2006-01-02 15:04:05")
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO users(email,password,active,created_at,updated_at) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(email, password, false, waktu, waktu)
	if err != nil {
		panic(err.Error())
	}
	return true
}
