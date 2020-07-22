package tables

import (
	"github.com/danangkonang/rest-api/migrate/tables/address"
	"github.com/danangkonang/rest-api/migrate/tables/itemSubKategories"
	"github.com/danangkonang/rest-api/migrate/tables/kategories"
	"github.com/danangkonang/rest-api/migrate/tables/products"
	"github.com/danangkonang/rest-api/migrate/tables/subKategories"
	"github.com/danangkonang/rest-api/migrate/tables/users"
)

func Tables() {
	address.Address()
	kategories.Kategori()
	subKategories.SubKategori()
	itemSubKategories.SubItem()
	products.Product()
	users.Users()
}
