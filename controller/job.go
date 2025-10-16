package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/bonjourrog/jb/entity/job"
	"github.com/bonjourrog/jb/service"
	"github.com/bonjourrog/jb/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type JobController interface {
	NewJob(c *gin.Context)
	GetJobs(c *gin.Context)
	UpdateJob(c *gin.Context)
	DeleteJob(c *gin.Context)
	GetJob(c *gin.Context)
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
		cxt = c.Request.Context()
	)

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userId, ok := c.Get("user_id")
	if userId == "" || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	user_id, err := bson.ObjectIDFromHex(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	job.CompanyID = user_id

	jobInserted, err := _jobService.NewJob(job, cxt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "job created succesfully",
		"data":    jobInserted,
	})

}
func (*jobController) GetJobs(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		filter = bson.M{}
		jobs   []job.PostWithCompany
	)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	query := c.Request.URL.Query()
	if user_id := query.Get("user_id"); user_id != "" {
		userId, err := bson.ObjectIDFromHex(user_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		filter["user_id"] = userId
	}

	filter["published"] = true
	if company_id := query.Get("company_id"); company_id != "" {
		companyId, err := bson.ObjectIDFromHex(company_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		filter["company_id"] = companyId
		delete(filter, "published")
	}
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
	jobs, total, err := _jobService.GetJobs(filter, page, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	totalPages := int(math.Ceil(float64(total) / float64(12)))
	c.JSON(http.StatusOK, gin.H{
		"message":     "successful request",
		"data":        jobs,
		"page":        page,
		"page_zise":   12,
		"total":       total,
		"total_pages": totalPages,
	})
}
func (jobController) UpdateJob(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		job job.Post
	)

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := _jobService.UpdateJob(job, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "job updated succesfully",
	})
}
func (*jobController) DeleteJob(c *gin.Context) {
	ctx := c.Request.Context()
	userId, ok := c.Get("user_id")

	if userId == "" || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid user ID",
		})
		return
	}
	user_id, err := bson.ObjectIDFromHex(userId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	job_id := c.Param("id")

	if job_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no id provided",
		})
		return
	}
	jobId, err := bson.ObjectIDFromHex(job_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = _jobService.DeleteJob(jobId, user_id, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Job deleted successfully",
	})
}
func (*jobController) GetJob(c *gin.Context) {
	var (
		companyName, slug string
		ctx               = c.Request.Context()
		job               *job.PostWithCompany
	)
	if companyName = c.Param("company_name"); companyName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "company name is missing",
		})
		return
	}
	if slug = c.Param("slug"); slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"data":    nil,
			"message": "slug is missing",
		})
		return
	}
	job, err := _jobService.GetBySlug(companyName, slug, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   true,
		"data":    job,
		"message": "succesful request",
	})
}
