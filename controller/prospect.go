package controller

import (
	"net/http"

	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/service"
	"github.com/gin-gonic/gin"
)

type ProspectController interface {
	NewProspect(c *gin.Context)
}
type prospectController struct {
	prospectService service.ProspectService
}

func NewProspectController(prospectService service.ProspectService) ProspectController {
	return &prospectController{
		prospectService: prospectService,
	}
}
func (p *prospectController) NewProspect(c *gin.Context) {
	var (
		prospect entity.Prospect
	)
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&prospect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	if err := p.prospectService.NewProspect(prospect, ctx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "prospect created succesfully",
	})

}
