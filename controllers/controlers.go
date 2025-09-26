package controllers

import (
	"go-auth/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	// function name
	// admin := router.Group("/admin")
	// AdminRoutes(admin)

	// owner := router.Group("/owner")
	// OwnersRoutes(owner)

	// user := router.Group("/user")
	// UserRoutes(user)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://*.vercel.app"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "token", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true // Allow all origins in development
		},
		MaxAge: 12 * time.Hour,
	}))

	router.POST("/signup", services.Signup)
	router.POST("/login", services.Login)
	router.GET("/home", services.Home)

	admin := router.Group("/admin")
	AdminRoutes(admin)

	owner := router.Group("/owner")
	OwnersRoutes(owner)

	user := router.Group("/user")
	UserRoutes(user)

	router.GET("/logout", services.Logout)

}
