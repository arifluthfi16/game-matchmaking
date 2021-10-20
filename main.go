package main

import (
	"github.com/arifluthfi16/game-matchmaking/services"
	"gorm.io/gorm"
)

type Env struct {
	db *gorm.DB
}

func main(){
	DBConfig := services.DBConfig{
		DBHost: "localhost",
		DBUser: "postgres",
		DBName: "game_matchmaking",
		DBPort: "5432",
		DbPass: "universalPassword2020",
	}
	env := Env{db: services.LoadDB(DBConfig)}
	services.DBMigrate(env.db)
}