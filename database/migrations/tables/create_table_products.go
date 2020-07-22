package tables

import (
	"fmt"
	"log"

	"github.com/danangkonang/rest-api/config"
)

func Products() {
	db := config.Connect()
	db.Exec(`DROP TABLE products`)
	_, err := db.Exec(`CREATE TABLE products(
				id_product CHAR (32) NOT NULL,
				user_id CHAR (32) NOT NULL,
				product_name VARCHAR (225) NOT NULL,
				images VARCHAR (225) NOT NULL,
				prince VARCHAR (225) NOT NULL,
				description VARCHAR (225) NOT NULL,
				brand VARCHAR (225) NULL,
				address VARCHAR (225) NOT NULL,
				condition BOOLEAN,
				stock INTEGER NOT NULL,
				category_id INTEGER NOT NULL,
				weight INTEGER NOT NULL,
				weight_unit VARCHAR (225) NOT NULL,
				active BOOLEAN,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			  )`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create table products")
}
