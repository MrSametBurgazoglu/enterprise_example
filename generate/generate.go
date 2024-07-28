package main

import (
	"example/db_models"
	"github.com/MrSametBurgazoglu/enterprise/generate"
)

func main() {
	generate.Models(
		db_models.Deneme(),
		db_models.Test(),
		db_models.Account(),
		db_models.Group(),
	)
}
