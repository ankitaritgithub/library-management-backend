package controllers

import (
	middleware "go-auth/Middlewares"
	"go-auth/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {
	r.Use(middleware.RequireAuth("admin"))
	r.POST("/add-book", admin.AddBook)

	r.GET("/book", admin.GetAllBooks)

	r.GET("/book/:id", admin.GetBookByID)

	r.PUT("/update-book/:isbn", admin.UpdateBook)
	r.DELETE("/book/:id", admin.DeleteBook)

	r.GET("/issue-requests", admin.ListRequestEventsHandler) // Define route for listing issue requests

	r.PUT("/issue-requests/approve", admin.ApproveRequestEventHandler) // Corrected    // Define route for approving an issue request

	r.PUT("/issue-requests/reject", admin.RejectRequestEventHandler)

}
