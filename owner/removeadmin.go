package owner

import (
	"go-auth/database"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func RemoveAdminByID(c *gin.Context) {

	// Parse the admin ID from the URL parameter

	adminID := c.Param("id")

	// Convert the admin ID to integer

	adminIDInt, err := strconv.Atoi(adminID)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})

		return

	}

	// Check if the admin exists in the database

	var existingAdmin database.User

	if err := DB.Where("id = ?", adminIDInt).First(&existingAdmin).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})

		return

	}

	// Delete the admin from the database

	if err := DB.Delete(&existingAdmin).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove admin"})

		return

	}

	// Return success response

	c.JSON(http.StatusOK, gin.H{"message": "Admin removed successfully"})

}
