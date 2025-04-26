package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bonjourrog/jb/entity/job"
	"github.com/bonjourrog/jb/service"
	"github.com/gin-gonic/gin"
)

type JobController interface {
	NewJob(c *gin.Context)
}

type jobController struct{}

var (
	_jobService service.JobService
)

func NewJobController(jobService service.JobService) JobController {
	_jobService = jobService
	return &jobController{}
}

func (*jobController) NewJob(c *gin.Context) {
	var (
		job job.Post
	)

	if err := json.NewDecoder(c.Request.Body).Decode(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := _jobService.NewJob(job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "job created succesfully",
	})

}
