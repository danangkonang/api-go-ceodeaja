package seeder

import (
	"github.com/danangkonang/ceodeaja-go/database/seeder/seeds"
)

func CreateSeeds() {
	seeds.Categories()
	seeds.CreateProvinces()
	seeds.CreateRegencies()
	seeds.SubCategoryItems()
	seeds.SubCategories()
	seeds.CreateKecamatan()
	// seeds.CreateVillage()
}
