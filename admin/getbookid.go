package admin

import (
	"go-auth/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookByID(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Your code for handling GET requests goes here
	var book database.Book
	id := c.Param("id")

	if err := DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}
