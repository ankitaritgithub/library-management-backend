package db

import (
	"go-auth/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	DB, err := gorm.Open(sqlite.Open("Library.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate the provided structs
	DB.AutoMigrate(&database.IssueRegistery{})
	DB.AutoMigrate(&database.User{})
	DB.AutoMigrate(&database.RequestEvents{})
	DB.AutoMigrate(&database.Book{})
	DB.AutoMigrate(&database.LibraryAdmin{})
	return DB
}
