package kel30

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Kel39() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 39")
}
