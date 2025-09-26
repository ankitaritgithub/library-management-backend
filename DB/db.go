package db

import (
	"go-auth/database"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	dbPath := getDBPath()
	DB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// AutoMigrate the provided structs
	DB.AutoMigrate(&database.IssueRegistery{})
	DB.AutoMigrate(&database.User{})
	DB.AutoMigrate(&database.RequestEvents{})
	DB.AutoMigrate(&database.Book{})
	DB.AutoMigrate(&database.LibraryAdmin{})
	return DB
}

func getDBPath() string {
	// Check for environment variable first
	if dbPath := os.Getenv("DATABASE_PATH"); dbPath != "" {
		// Ensure directory exists
		dir := filepath.Dir(dbPath)
		os.MkdirAll(dir, 0755)
		return dbPath
	}

	// Default to Library.db in current directory
	return "Library.db"
}
