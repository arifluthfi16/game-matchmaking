package services

import (
	"fmt"
	"github.com/arifluthfi16/game-matchmaking/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	DBHost string
	DBUser string
	DBName string
	DBPort string
	DbPass string
}

func LoadDB(config DBConfig) *gorm.DB {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUser, config.DbPass, config.DBName, config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("!! Database Connection Loaded !!")
	return db
}

func DBMigrate (db *gorm.DB){
	var err error
	for _,model := range model.RegisterModels(){
		err = db.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("!! Database Migration Succeed !!")
}