package main

import (
	"github.com/fuadhossainnishad/go_auth_server/database"
	"github.com/fuadhossainnishad/go_auth_server/models"
)

func main() {
	print("this is main", "\n")
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
}
