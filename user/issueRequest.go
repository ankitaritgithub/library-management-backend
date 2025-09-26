package user

import (
	"go-auth/database"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(db *gorm.DB) {
	DB = db
}

func CreateIssueRequestHandler(c *gin.Context) {

	var data struct{
		ReqId 	uint    `json:"reqId"`
		BookId 	uint 	`json:"bookId"`
		ReaderId uint 	`json:"readerId"`
		RequestType string 	`json:"requestType"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	request := database.RequestEvents{BookID: int(data.BookId), ReaderID: int(data.ReaderId), RequestType: data.RequestType, RequestDate: time.Now(), Status: "pending", ReqID: int(data.ReqId)}

	res := DB.Create(&request)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "request created"})


}
