package controllers

import (
	"github.com/gin-gonic/gin"

	middleware "go-auth/Middlewares"
	"go-auth/owner"
)

func OwnersRoutes(r *gin.RouterGroup) {
	r.Use(middleware.RequireAuth("owner"))

	r.POST("/api/admins", owner.AddAdmin)
	r.DELETE("/api/admins/:id", owner.RemoveAdminByID)
	r.POST("/library", owner.CreateLibrary)

}
