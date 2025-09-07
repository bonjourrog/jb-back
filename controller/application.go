package controller

import (
	"github.com/bonjourrog/jb/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ApplicationController interface {
	// Create(c *gin.Context)
	GetUserApplications(c *gin.Context)
	// UpdateStatus(c *gin.Context)
}
type applicationController struct {
	appservice service.ApplicationService
}

func NewApplicationController(applicationService service.ApplicationService) ApplicationController {
	return &applicationController{
		appservice: applicationService,
	}
}

func (a *applicationController) GetUserApplications(c *gin.Context) {
	var (
		cxt = c.Request.Context()
	)
	userId, ok := c.Get("user_id")
	if userId == "" || !ok {
		c.JSON(401, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	user_id, err := bson.ObjectIDFromHex(userId.(string))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	applications, err := a.appservice.GetUserApplications(user_id, cxt)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":      "Applications fetched successfully",
		"applications": applications,
	})
}
