package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBConn var
// var DBConn *gorm.DB

// InitDB func
func InitDB() {
	var err error
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&Book{})
	fmt.Println("Database Migrated")
}
