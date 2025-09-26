package main

import (
	db "go-auth/DB"
	middleware "go-auth/Middlewares"
	admin "go-auth/admin"
	"go-auth/controllers"
	"os"

	"go-auth/owner"

	"go-auth/user"

	"go-auth/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = db.InitDB()
	admin.Connect(DB)
	user.Connect(DB)
	owner.Connect(DB)
	middleware.Connect(DB)
	services.Connect(DB)

}
func main() {
	router := gin.Default()
	controllers.Routes(router)

	// Get port from environment variable or default to 8080
	port := getPort()

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
