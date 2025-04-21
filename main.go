package main

import (
	"log"
	"os"

	"github.com/bonjourrog/jb/controller"
	"github.com/bonjourrog/jb/repository/auth"
	"github.com/bonjourrog/jb/routes"
	"github.com/bonjourrog/jb/service"
	"github.com/joho/godotenv"
)

var (

	// Auth
	authRepo       auth.AuthRepo             = auth.NewAuthRepository()
	authService    service.AuthService       = service.NewAuthService(authRepo)
	authController controller.AuthController = controller.NewAuthController(authService)
	httpRouter     routes.Router             = routes.NewGinRouter()
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading env file")
	}
	httpRouter.POST("/api/auth/signup", authController.Signup)
	httpRouter.Serve(os.Getenv("PORT"))
}
