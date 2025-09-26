package owner

import (
	"fmt"
	"go-auth/database"

	// "net/http"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	// "go-auth/utils"
)

var DB *gorm.DB

func Connect(db *gorm.DB) {
	DB = db
}

func AddAdmin(c *gin.Context) {

	var adminDetails struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		ContactNumber int    `json:"contact_number"` // Use snake case
		Role          string `json:"role"`
		LibId         int    `json:"lib_id"` // Use snake case
		Password      string `json:"password"`
		Owner         string `json:"owner"`
	}

	var owner database.User
	var admin database.User

	if err := c.Bind(&adminDetails); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	} else {

		adminDetails.Email = strings.Trim(adminDetails.Email, " ")
		adminDetails.Email = strings.ToLower(adminDetails.Email)

		adminDetails.Password = strings.Trim(adminDetails.Password, " ")
		adminDetails.Name = strings.Trim(adminDetails.Name, " ")

		// if !(utils.IsValidEmail(adminDetails.Email)) {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, "invalid Email Format")
		// 	return
		// }
		// if !(utils.IsContactNumberValid(adminDetails.ContactNumber)) {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, "invalid Contact number format enter without +91 and must be 10 number only")
		// 	return
		// }
		// if !(utils.IsNameValid(adminDetails.Name)) {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, "invalid Name Format")
		// 	return
		// }
		// if !(utils.IsPasswordValid(adminDetails.Password)) {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, "password must be between 8 to 15 characters included capital letters smaller letters symbols and numbers ")
		// 	return
		// }

		ownerRes := DB.Where("email = ?", adminDetails.Owner).First(&owner)
		if ownerRes.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "error finding owner"})
			return
		}

		var existingAdmin database.User
		result := DB.Where("lib_id = ? AND role = ? AND email = ?", owner.LibId, "admin").First(&existingAdmin)
		fmt.Println(result)
		if result.RowsAffected == 0 {
			admin.Name = adminDetails.Name
			admin.ContactNumber = adminDetails.ContactNumber
			admin.Email = adminDetails.Email
			admin.Role = "admin"
			admin.LibId = owner.LibId

			bs, err := bcrypt.GenerateFromPassword([]byte(adminDetails.Password), bcrypt.DefaultCost)
			if err != nil {
				panic(err)

			}
			admin.Password = string(bs)
			tx := DB.Begin()
			if err := tx.Create(&admin).Error; err != nil {
				tx.Rollback()
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"msg": "user already exists"})
				return
			}
			if err := tx.Commit().Error; err == nil {
				c.JSON(http.StatusCreated, gin.H{"admin": admin})

			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"msg": "admin already exists"})
			return
		}
	}

}
