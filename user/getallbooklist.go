package user

import (
	"go-auth/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIssueRegistryHandler handles GET requests to fetch issue registry data.
func GetIssueRegistryHandler(c *gin.Context) {
	var issueRegistries []database.IssueRegistery

	// Fetch issue registries from the database
	if err := DB.Find(&issueRegistries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch issue registries"})
		return
	}
 
	
	c.JSON(http.StatusOK, issueRegistries)
}
