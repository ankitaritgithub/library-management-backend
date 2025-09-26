package admin

import (
	"go-auth/database"

	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBook(c *gin.Context) {

	var updatedBook database.Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := DB.Where("isbn = ?", updatedBook.ISBN).Updates(&database.Book{Title: updatedBook.Title, Author: updatedBook.Author, Publisher: updatedBook.Publisher, Version: updatedBook.Version, TotalCopies: updatedBook.TotalCopies, AvailableCopies: updatedBook.AvailableCopies}).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})

}
