package owner

import (
	"go-auth/database"
	"go-auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateLibrary(c *gin.Context) {
	var data struct {
		Name     string `valid:"required"  json:"name"`
		Email    string `valid:"email"      json:"email"`
		Password string `valid:"length(6|20)"  json:"password"`
		LibName  string `valid:"required"    json:"lib_name"`
	}

	var library database.Library
	var user database.User
	if err := c.Bind(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	} else {
		data.Email = strings.Trim(data.Email, " ")
		data.Email = strings.ToLower(data.Email)
		data.Name = strings.Trim(data.Name, " ")
		data.Password = strings.Trim(data.Password, " ")
		data.LibName = strings.Trim(data.LibName, " ")
		if !(utils.IsValidEmail(data.Email)) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid email")
			return
		}

		if !(utils.IsNameValid(data.Name)) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid Name")
			return
		}
		if !(utils.IsNameValid(data.LibName)) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "invalid library name")
			return
		}

		if !(utils.IsPasswordValid(data.Password)) {
			c.AbortWithStatusJSON(http.StatusBadRequest, "password must be between 8 to 15 characters included capital letters smaller letters symbols and numbers ")
			return
		}

		library.Name = data.LibName

		tx := DB.Begin()
		if err := tx.Create(&library).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusConflict, "Library already exists")
			return
		}
		user.Name = data.Name
		user.Email = data.Email
		user.Role = "owner"
		user.LibId = int(library.ID)
		user.Library = library
		bs, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user.Password = string(bs)

		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusConflict, "User already exists")
			return
		}
		if err := tx.Commit().Error; err == nil {
			c.JSON(http.StatusCreated, gin.H{"user": user})
		} else {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, "Not added")

		}

	}

}
