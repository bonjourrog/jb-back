package main

import (
	"context"
	"log"
	"os"

	"github.com/bonjourrog/jb/controller"
	"github.com/bonjourrog/jb/db"
	"github.com/bonjourrog/jb/middleware"
	"github.com/bonjourrog/jb/repository/application"
	"github.com/bonjourrog/jb/repository/auth"
	"github.com/bonjourrog/jb/repository/job"
	"github.com/bonjourrog/jb/routes"
	"github.com/bonjourrog/jb/service"
	"github.com/joho/godotenv"
)

func main() {
	var httpRouter routes.Router = routes.NewGinRouter()
	httpRouter.Use(middleware.CorsConfig())

	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not found")
	}
	mongoClient, err := db.NewMongoClient(uri)
	if err != nil {
		log.Fatal("Cannot connect to MongoDB", err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	var (

		// Auth
		authRepo       auth.AuthRepo             = auth.NewAuthRepository(mongoClient)
		authService    service.AuthService       = service.NewAuthService(authRepo)
		authController controller.AuthController = controller.NewAuthController(authService)

		// Job
		jobRepo       job.JobRepository        = job.NewJobRepository(mongoClient)
		jobService    service.JobService       = service.NewPostService(jobRepo)
		jobController controller.JobController = controller.NewJobController(jobService)

		// Application
		applicationRepo       application.ApplicationRepository = application.NewApplicationRepository(mongoClient)
		applicationService    service.ApplicationService        = service.NewApplicationService(applicationRepo)
		applicationController controller.ApplicationController  = controller.NewApplicationController(applicationService)
	)
	httpRouter.POST("/api/auth/signup", authController.Signup)
	httpRouter.POST("/api/auth/signin", authController.Signin)
	httpRouter.PUT("/api/job", jobController.UpdateJob, middleware.ValidateToken(), middleware.OnlyCompanyAccess())
	httpRouter.POST("/api/job", jobController.NewJob, middleware.ValidateToken(), middleware.OnlyCompanyAccess())
	httpRouter.GET("/api/job", jobController.GetJobs)
	httpRouter.DELETE("/api/job/:id", jobController.DeleteJob, middleware.ValidateToken(), middleware.OnlyCompanyAccess())
	httpRouter.POST("/api/job/:id/apply", jobController.ApplyToJob, middleware.ValidateToken())
	// application routes
	httpRouter.GET("/api/application/user", applicationController.GetUserApplications, middleware.ValidateToken(), middleware.OnlyUserAccess())
	httpRouter.Serve(os.Getenv("PORT"))
}
