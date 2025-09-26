package admin

import (
	"go-auth/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllBooks retrieves all books from the database
func GetAllBooks(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	var books []database.Book

	if err := DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	c.JSON(http.StatusOK, books)
}
