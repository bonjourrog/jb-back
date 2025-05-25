package routes

import "github.com/gin-gonic/gin"

type Router interface {
	GET(uri string, f func(c *gin.Context))
	POST(uri string, f func(c *gin.Context), middlewares ...gin.HandlerFunc)
	DELETE(uri string, f func(c *gin.Context))
	PUT(uri string, f func(c *gin.Context))
	Use(middleware ...gin.HandlerFunc)
	Serve(port string)
}
