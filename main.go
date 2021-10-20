package main

import (
	"github.com/arifluthfi16/game-matchmaking/services/db"
	"gorm.io/gorm"
)

type Env struct {
	db *gorm.DB
}

func main(){
	DBConfig := db.DBConfig{
		DBHost: "localhost",
		DBUser: "postgres",
		DBName: "game_matchmaking",
		DBPort: "5432",
		DbPass: "universalPassword2020",
	}
	env := Env{db: db.LoadDB(DBConfig)}
	db.DBMigrate(env.db)
}