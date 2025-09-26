package admin

import (
	"fmt"
	"go-auth/database"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRequestEventsHandler(c *gin.Context) {
	var requestEvents []database.RequestEvents

	if err := DB.Find(&requestEvents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve requests"})
		return
	}

	c.JSON(http.StatusOK, requestEvents)

}

func ApproveRequestEventHandler(c *gin.Context) {

	var requestEvent database.RequestEvents
	var data struct {
		ReqId   uint `json:"reqId"`
		AdminId uint `json:"adminId"`
	}

	res := c.ShouldBindJSON(&data)
	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
		return
	}
	fmt.Println(&data)

	if err := DB.Where("req_id = ?", data.ReqId).Find(&requestEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve requests"})
		return
	}

	requestEvent.ApproverID = int(data.AdminId)
	requestEvent.ApprovalDate = time.Now()
	requestEvent.Status = "approved"
	DB.Save(requestEvent)

	c.JSON(http.StatusOK, "approved")
}

// Handler function to reject a request event

func RejectRequestEventHandler(c *gin.Context) {

	var requestEvent database.RequestEvents
	var data struct {
		ReqId   uint `json:"reqId"`
		AdminId uint `json:"adminId"`
	}

	res := c.ShouldBindJSON(&data)
	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res})
		return
	}
	fmt.Println(&data)

	if err := DB.Where("req_id = ?", data.ReqId).Find(&requestEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve requests"})
		return
	}

	requestEvent.ApproverID = int(data.AdminId)
	requestEvent.ApprovalDate = time.Now()
	requestEvent.Status = "rejected"
	DB.Save(requestEvent)

	c.JSON(http.StatusOK, "rejected")
}
