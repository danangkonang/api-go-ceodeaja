package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Create() {
	db := config.Connect()
	db.Exec(`DROP TABLE firebase_token`)
	_, err := db.Exec(`CREATE TABLE firebase_token(
		id serial PRIMARY KEY,
		user_id CHAR (32) NOT NULL,
		token_fmc VARCHAR (225) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table firebase_token")
}
