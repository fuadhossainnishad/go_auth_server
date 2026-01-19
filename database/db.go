package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)
var DB *gorm.DB

func Connect(){
	var err error
	DB, err = gorm.Open(sqlite.Open("auth.db", &gorm.Config{}))
	if err !=nil{
		log.Fatal("Failed to connect database",err)
	}

}
