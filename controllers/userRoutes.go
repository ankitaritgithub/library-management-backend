package controllers

import (
	middleware "go-auth/Middlewares"
	"go-auth/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {

	r.Use(middleware.RequireAuth("user"))

	r.POST("/raise-issue", user.CreateIssueRequestHandler)
	// r.POST("/raise-issue", user.CreateIssue)

	r.GET("/books", user.SearchBooks)
	// r.GET("/issue-registries", user.GetIssueRegistryHandler)
	// r.POST("/request", user.CreateIssue)
	// r.POST("/return/:isbn", user.ReturnBook)

}
