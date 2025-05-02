package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bonjourrog/jb/entity/job"
	"github.com/bonjourrog/jb/service"
	"github.com/bonjourrog/jb/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type JobController interface {
	NewJob(c *gin.Context)
	GetJobs(c *gin.Context)
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
	userId, ok := c.Get("user_id")
	if userId == "" || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	job.CompanyID = userId.(string)
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
func (*jobController) GetJobs(c *gin.Context) {
	var (
		filter = bson.M{}
		jobs   []job.Post
	)
	query := c.Request.URL.Query()
	if search := query.Get("search"); search != "" {
		orFilter := bson.A{
			bson.M{"title": bson.M{"$regex": search, "$options": "i"}},
			bson.M{"short_description": bson.M{"$regex": search, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": search, "$options": "i"}},
		}
		filter["$or"] = orFilter
	}
	if schedule := query.Get("schedule"); schedule != "" {
		if err := util.VerifySchedule(schedule); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		filter["schedule"] = schedule
	}
	if contract := query.Get("contract"); contract != "" {
		if err := util.VerifyContractType(contract); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		filter["contract_type"] = contract
	}
	if industry := query.Get("industry"); industry != "" {
		filter["industry"] = industry
	}
	jobs, err := _jobService.GetJobs(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successful request",
		"data":    jobs,
	})
}
