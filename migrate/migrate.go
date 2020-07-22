package migrate

import (
	"github.com/danangkonang/rest-api/migrate/tables"
)

func Migrate() {
	tables.Tables()
}
