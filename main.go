package main

import (
	"log"
	"os"

	"github.com/bonjourrog/jb/controller"
	"github.com/bonjourrog/jb/repository/auth"
	"github.com/bonjourrog/jb/repository/job"
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

	// Job
	jobRepo       job.JobRepository        = job.NewJobRepository()
	jobService    service.JobService       = service.NewPostService(jobRepo)
	jobController controller.JobController = controller.NewJobController(jobService)
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading env file")
	}
	httpRouter.POST("/api/auth/signup", authController.Signup)
	httpRouter.POST("/api/auth/signin", authController.Signin)
	httpRouter.POST("/api/job", jobController.NewJob)
	httpRouter.Serve(os.Getenv("PORT"))
}
