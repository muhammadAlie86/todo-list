package main

import (
	"todo-list/shared/config"
	"todo-list/shared/database"
)

func Main() {

	cfg := config.InitConfig()

	db := database.NewDataBase(cfg)

	defer db.CloseDatabase()

}
