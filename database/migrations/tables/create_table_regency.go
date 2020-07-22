package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/rest-api/config"
)

func Regencies() {
	db := config.Connect()
	db.Exec(`DROP TABLE regencies`)
	_, err := db.Exec(`CREATE TABLE regencies(
		id serial PRIMARY KEY,
		regency VARCHAR (225) NOT NULL,
		province_id INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table regencies")
}
