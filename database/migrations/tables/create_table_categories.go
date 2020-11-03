package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Categories() {
	db := config.Connect()
	db.Exec(`DROP TABLE categories`)
	_, err := db.Exec(`CREATE TABLE categories(
				id serial PRIMARY KEY,
				category VARCHAR (225) NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			  )`)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table categories")
}
