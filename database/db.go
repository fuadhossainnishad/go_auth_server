package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)
var DB *gorm.DB

func Connect(){
	vat err error
	DB, err = gorm.Open(sqlite.open("auth.db", &gorm.Config{}))
	if err !=nil{
		log.Fatal("Failed to connect database",err)
	}

}