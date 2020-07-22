package migrations

import (
	"github.com/danangkonang/rest-api/database/migrations/tables"
	// "github.com/danangkonang/rest-api/database/seeder"
)

func migrations() {
	tables.Categories()
	tables.SubCategories()
	tables.SubCategoriItems()
	tables.Messages()
	tables.Products()
	tables.Provinces()
	tables.Regencies()
	tables.SubDistricts()
	tables.UserProfiles()
	tables.Users()
	tables.Villages()
	// seeder.CreateSeeds()
}
