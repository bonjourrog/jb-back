package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bonjourrog/jb/entity"
	"github.com/bonjourrog/jb/service"
	"github.com/bonjourrog/jb/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	var user entity.User
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.Name == "" || user.LastName == "" || user.Account.Email == "" || user.Account.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "some required fields are empty",
		})
		return
	}
	if isRoleValid := util.VerifyRole(user.Role); !isRoleValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid role",
		})
		return
	}
	hashedPassword, err := util.GeneratePassword(user.Account.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.ID = bson.NewObjectID()
	user.Account.Password = hashedPassword
	user.Account.Banned = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	_, err = _authService.Signup(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})
}
func (*authController) Signin(c *gin.Context) {
	var (
		credentials entity.Account
	)
	if err := json.NewDecoder(c.Request.Body).Decode(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := _authService.SignIn(credentials)
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
