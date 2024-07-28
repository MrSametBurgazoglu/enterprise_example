package main

import (
	"context"
	"example/db_models"
	"github.com/MrSametBurgazoglu/enterprise/migrate"
	"github.com/MrSametBurgazoglu/enterprise/models"
)

func main() {
	tables := []*models.Table{
		db_models.Deneme(),
		db_models.Test(),
		db_models.Account(),
		db_models.Group(),
	}

	migrate.Migrate(
		context.TODO(),
		"postgresql://testuser:54M3754M37@localhost:5433/testdb?search_path=public",
		"./migrations",
		"new_plan",
		tables,
	)
}
