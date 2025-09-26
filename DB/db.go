package db

import (
	"go-auth/database"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDBPath() string {
	if path := os.Getenv("DATABASE_PATH"); path != "" {
		return path
	}
	return "Library.db"
}

func InitDB() *gorm.DB {
	var err error
	dbPath := getDBPath()
	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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
