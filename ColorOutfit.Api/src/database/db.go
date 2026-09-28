package database

import (
	"log"

	"github.com/adamwahyuh/coloroutfit/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("src/database/db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	models := []interface{}{
		&models.User{},
		&models.Category{},
		&models.Color{},
		&models.TemplateColor{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		panic(err)
	}

	DB = db
}
