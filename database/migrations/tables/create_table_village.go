package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Villages() {
	db := config.Connect()
	db.Exec(`DROP TABLE villages`)

	_, err := db.Exec(`CREATE TABLE villages(
		id serial PRIMARY KEY,
		village VARCHAR (225) NOT NULL,
		sub_district_id INTEGER NOT NULL,
		regency_id INTEGER NOT NULL,
		province_id INTEGER NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table villages")
}
