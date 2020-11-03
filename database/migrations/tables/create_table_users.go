package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Users() {
	db := config.Connect()
	db.Exec(`DROP TABLE users`)
	_, err := db.Exec(`CREATE TABLE users(
				id CHAR(32) NOT NULL,
				email VARCHAR (225) NOT NULL,
				password VARCHAR (225) NOT NULL,
				user_token VARCHAR (225) NOT NULL,
				user_role INTEGER NOT NULL,
				is_active BOOLEAN,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table users")
}
