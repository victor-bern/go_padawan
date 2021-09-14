package main

import (
	"go_padawan/src/database"
	"go_padawan/src/database/migrations"
	"go_padawan/src/database/seed"
	"go_padawan/src/server"
)

func main() {
	database.StartDatabase()
	migrations.AutoMigrate(database.GetDatabase())
	seed.SeedCurrency(database.GetDatabase())
	server := server.NewServer()
	server.Run()
}
