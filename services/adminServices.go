package services

import "gorm.io/gorm"

var DB *gorm.DB

func Connect(db *gorm.DB) {
	DB = db
}

