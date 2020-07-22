package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/rest-api/config"
)

func SubCategoriItems() {
	db := config.Connect()
	db.Exec(`DROP TABLE sub_categori_items`)
	_, err := db.Exec(`CREATE TABLE sub_categori_items(
				id serial PRIMARY KEY,
				category_id INTEGER NOT NULL,
				sub_category_id INTEGER NOT NULL,
				sub_item VARCHAR (225) NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			  )`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table sub_categori_items")
}
