package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type ginRouter struct{}

var (
	ginDispatch = gin.Default()
)

func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) GET(uri string, f func(c *gin.Context), middlewares ...gin.HandlerFunc) {
	handlers := append(middlewares, f)
	ginDispatch.GET(uri, handlers...)
}
func (*ginRouter) POST(uri string, f func(c *gin.Context), middlewares ...gin.HandlerFunc) {
	handlers := append(middlewares, f)
	ginDispatch.POST(uri, handlers...)
}
func (*ginRouter) DELETE(uri string, f func(c *gin.Context), middlewares ...gin.HandlerFunc) {
	handlers := append(middlewares, f)
	ginDispatch.DELETE(uri, handlers...)
}
func (*ginRouter) PUT(uri string, f func(c *gin.Context), middlewares ...gin.HandlerFunc) {
	handlers := append(middlewares, f)
	ginDispatch.PUT(uri, handlers...)
}
func (*ginRouter) Serve(port string) {
	fmt.Printf("Server running in port %v ", port)
	if err := ginDispatch.Run(port); err != nil {
		log.Fatal(err)
	}
}
func (*ginRouter) Use(middleware ...gin.HandlerFunc) {
	ginDispatch.Use(middleware...)
}
