package admin

import (
	"bytes"
	"encoding/json"
	"go-auth/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(db *gorm.DB) {
	DB = db
}

func AddBook(c *gin.Context) {
	var newBook database.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the book already exists
	var existingBook database.Book
	result := DB.Where("isbn = ?", newBook.ISBN).First(&existingBook)

	if result.Error == nil {
		// Book exists, update its copies
		existingBook.TotalCopies += newBook.TotalCopies
		existingBook.AvailableCopies += newBook.TotalCopies
		DB.Save(&existingBook)
		c.JSON(http.StatusOK, existingBook)
	} else if result.Error == gorm.ErrRecordNotFound {
		// New book, add to database
		DB.Create(&newBook)
		c.JSON(http.StatusOK, newBook)
	} else {
		// Some other error occurred
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, newBook)
	if err := sendTeamsNotification("https://xenonstack1.webhook.office.com/webhookb2/73646a59-81d7-41a4-ad5e-d6883ca9ec7d@7ff914bc-ca07-4c28-8277-73e20a4966c7/IncomingWebhook/1bec37b5c74c4636affdbaeaeb7d7322/50cad2cd-87f4-41ef-bb9e-8e8a88e29076", "Book Added", "A new book has been added", newBook.Title, newBook.Author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification on team"})
		return
	}
	// Send notification email
}

func sendTeamsNotification(webhookURL, title, subtitle, bookTitle, bookAuthor string) error {
	teamsPayload := map[string]interface{}{
		"@type":      "MessageCard",
		"@context":   "http://schema.org/extensions",
		"themeColor": "0076D7",
		"summary":    title,
		"sections": []map[string]string{
			{"activityTitle": title,
				"activitySubtitle": subtitle,
				"text":             "Title:" + bookTitle + "\n Author:" + bookAuthor,
			},
		},
	}
	//converting payload to josn format

	payloadJSON, err := json.Marshal(teamsPayload)
	if err != nil {
		return err
	}

	// final statement to send teams message
	_, err = http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return err
	}
	return nil
}
