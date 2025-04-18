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

func (*ginRouter) GET(uri string, f func(c *gin.Context)) {
	ginDispatch.GET(uri, f)
}
func (*ginRouter) POST(uri string, f func(c *gin.Context)) {
	ginDispatch.POST(uri, f)
}
func (*ginRouter) DELETE(uri string, f func(c *gin.Context)) {
	ginDispatch.DELETE(uri, f)
}
func (*ginRouter) PUT(uri string, f func(c *gin.Context)) {
	ginDispatch.PUT(uri, f)
}
func (*ginRouter) Serve(port string) {
	fmt.Printf("Server running in port %v ", port)
	if err := ginDispatch.Run(port); err != nil {
		log.Fatal(err)
	}
}
