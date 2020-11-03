package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/ceodeaja-go/config"
)

func SubCategories() {
	db := config.Connect()
	db.Exec(`DROP TABLE sub_categories`)
	_, err := db.Exec(`CREATE TABLE sub_categories(
				id serial PRIMARY KEY,
				category_id INTEGER NOT NULL,
				sub_category VARCHAR (225) NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			  )`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table sub_categories")
}
