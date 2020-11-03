package tables

import (
	"fmt"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Messages() {
	db := config.Connect()
	db.Exec(`DROP TABLE messages`)
	_, err := db.Exec(`CREATE TABLE messages(
				id CHAR(32) NOT NULL,
				from_user_id CHAR(32) NOT NULL,
				to_user_id CHAR(32) NOT NULL,
				messages VARCHAR (225) NULL,
				is_read BOOLEAN,
				created_at TIMESTAMP NULL,
				updated_at TIMESTAMP NULL
			)`)
	if err != nil {
		panic(err)
	}
	fmt.Println("create table messages")
}
