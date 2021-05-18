package database

import (
	"go-api/models"
	"go-api/util"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	util.Env_load()

	database, err := gorm.Open(mysql.Open(os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/"+os.Getenv("DB_DATABASE")), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{})
}
