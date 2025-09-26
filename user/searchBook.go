package user

import (
	"go-auth/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchBooks(c *gin.Context) {

	var books []database.Book

	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	if err := DB.Where("title LIKE ? OR author LIKE ? OR publisher LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books found"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// func SearchBook(c *gin.Context) {
// 	var books []database.Book
// 	title := c.Query("title")
// 	author := c.Query("author")
// 	publisher := c.Query("publisher")

// 	query := DB.Model(&books)

// 	if title != "" {
// 		query = query.Where("title LIKE ?", "%"+title+"%")
// 	}

// 	if author != "" {
// 		query = query.Where("author LIKE ?", "%"+author+"%")
// 	}

// 	if publisher != "" {
// 		query = query.Where("publisher LIKE ?", "%"+publisher+"%")
// 	}

// 	query.Find(&books)

// 	c.JSON(http.StatusOK, books)
// }
