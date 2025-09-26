package admin

import (
	"go-auth/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBook(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}

	// Your code for handling DELETE requests goes here
	id := c.Param("id")

	if err := DB.Delete(&database.Book{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
