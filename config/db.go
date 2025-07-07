package config

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "AuraNest/models"
)

var DB *gorm.DB

func ConnectDatabase() {
   dsn := "Katheria_corrupted:Ayushi@1995@tcp(127.0.0.1:3306)/auranest?charset=utf8mb4&parseTime=True&loc=Local"


   database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    DB = database

    // Auto-migrate tables
	err = DB.AutoMigrate(&models.Vibe{}, &models.TimeSlot{}, &models.Song{})
   
	if err != nil {
        log.Fatal("Migration failed: ", err)
    }

    log.Println("Connected to DB & Migrated successfully")
}
