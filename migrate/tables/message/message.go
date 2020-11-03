package message

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Message() {
	db := config.Connect()
	db.Exec(`DROP TABLE message`)
	_, err := db.Exec(`CREATE TABLE message(
				id CHAR(32) NOT NULL,
				from CHAR(32) NOT NULL,
				to CHAR(32) NOT NULL,
				message VARCHAR (225) NULL,
				is_read BOOLEAN,,
				created_at TIMESTAMP NULL,
				updated_at TIMESTAMP NULL
			)`)
	if err != nil {
		panic(err)
	}
	fmt.Println("create table message")
}
