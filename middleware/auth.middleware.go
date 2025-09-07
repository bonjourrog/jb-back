package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_token := ctx.GetHeader("Authorization")
		if _token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "token required",
			})
			return
		}

		token, err := jwt.Parse(_token, func(t *jwt.Token) (interface{}, error) {
			//check signing method
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
			}
			return []byte(os.Getenv("SigningKey")), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		if !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "invalid token",
			})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("role", claims["role"])
			ctx.Set("user_id", claims["userId"])
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "invalid token",
		})
	}
}
func OnlyCompanyAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exist := ctx.Get("role")
		if !exist || role != "company" {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "admin role required"})
			return
		}
		ctx.Next()
	}
}
func OnlyUserAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exist := ctx.Get("role")
		if !exist || role != "user" {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "user role required"})
			return
		}
		ctx.Next()
	}
}

func CorsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
		for _, ao := range allowedOrigins {
			if origin == ao {
				c.Header("Access-Control-Allow-Origin", ao)
			}
		}
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
