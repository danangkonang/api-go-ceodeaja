package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/ceodeaja-go/config"
)

func UserProfiles() {
	db := config.Connect()
	db.Exec(`DROP TABLE user_profiles`)
	_, err := db.Exec(`CREATE TABLE user_profiles(
				user_id CHAR(32) NOT NULL,
				gender VARCHAR (225) NULL,
				telephone VARCHAR (225) NULL,
				avatar VARCHAR (225) NULL,
				address VARCHAR (225) NULL,
				kabupaten INTEGER NULL,
				kecamatan INTEGER NULL,
				kelurahan INTEGER NULL,
				provinsi INTEGER NULL,
				created_at TIMESTAMP NULL,
				updated_at TIMESTAMP NULL
			)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table user_profiles")
}
