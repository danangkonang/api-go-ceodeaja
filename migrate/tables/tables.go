package tables

import (
	"github.com/danangkonang/ceodeaja-go/migrate/tables/address"
	"github.com/danangkonang/ceodeaja-go/migrate/tables/itemSubKategories"
	"github.com/danangkonang/ceodeaja-go/migrate/tables/kategories"
	"github.com/danangkonang/ceodeaja-go/migrate/tables/products"
	"github.com/danangkonang/ceodeaja-go/migrate/tables/subKategories"
	"github.com/danangkonang/ceodeaja-go/migrate/tables/users"
)

func Tables() {
	address.Address()
	kategories.Kategori()
	subKategories.SubKategori()
	itemSubKategories.SubItem()
	products.Product()
	users.Users()
}
