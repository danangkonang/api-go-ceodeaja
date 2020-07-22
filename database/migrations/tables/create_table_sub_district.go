package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/rest-api/config"
)

func SubDistricts() {
	db := config.Connect()
	db.Exec(`DROP TABLE sub_districts`)
	_, err := db.Exec(`CREATE TABLE sub_districts(
		id serial PRIMARY KEY,
		subDistrict VARCHAR (225) NOT NULL,
		regency_id INTEGER NOT NULL,
		province_id INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table sub_districts")
}
