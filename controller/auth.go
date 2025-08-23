package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}

type authController struct{}

var (
	_authService service.AuthService
)

func NewAuthController(authService service.AuthService) AuthController {
	_authService = authService
	return &authController{}
}

// Signup handles user registration by validating input and creating a new user account.
func (*authController) Signup(c *gin.Context) {
	ctx := c.Request.Context()
	var user entity.User
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	_, err := _authService.Signup(user, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user created",
	})
}
func (*authController) Signin(c *gin.Context) {
	ctx := c.Request.Context()
	var (
		credentials entity.Account
	)
	if err := json.NewDecoder(c.Request.Body).Decode(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := _authService.SignIn(credentials, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}
