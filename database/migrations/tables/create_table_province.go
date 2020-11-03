package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/ceodeaja-go/config"
)

func Provinces() {
	db := config.Connect()
	db.Exec(`DROP TABLE provinces`)

	_, err := db.Exec(`CREATE TABLE provinces(
		id serial PRIMARY KEY,
		province VARCHAR (225) NOT NULL
	)`)
	if err != nil {
		fmt.Println("error profince")
		log.Fatal(err)
	}
	fmt.Println("create table provinces")
}
