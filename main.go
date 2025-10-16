package main

import (
	"context"
	"log"
	"os"

	"github.com/bonjourrog/jb/config"
	"github.com/bonjourrog/jb/controller"
	"github.com/bonjourrog/jb/db"
	"github.com/bonjourrog/jb/middleware"
	"github.com/bonjourrog/jb/repository/application"
	"github.com/bonjourrog/jb/repository/auth"
	"github.com/bonjourrog/jb/repository/job"
	"github.com/bonjourrog/jb/repository/prospect"
	"github.com/bonjourrog/jb/routes"
	"github.com/bonjourrog/jb/service"
)

func main() {
	var httpRouter routes.Router = routes.NewGinRouter()
	httpRouter.Use(middleware.CorsConfig())

	// Load env variables
	config.Load()
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
		// Repositories
		applicationRepo application.ApplicationRepository = application.NewApplicationRepository(mongoClient)
		jobRepo         job.JobRepository                 = job.NewJobRepository(mongoClient)
		prospectRepo    prospect.ProspectRepo             = prospect.NewProspectRepo(mongoClient)

		// Auth
		authRepo       auth.AuthRepo             = auth.NewAuthRepository(mongoClient)
		authService    service.AuthService       = service.NewAuthService(authRepo)
		authController controller.AuthController = controller.NewAuthController(authService)

		// Porspective
		prospectService    service.ProspectService       = service.NewProspectService(prospectRepo)
		prospectController controller.ProspectController = controller.NewProspectController(prospectService)

		// Application
		applicationService    service.ApplicationService       = service.NewApplicationService(applicationRepo, jobRepo)
		applicationController controller.ApplicationController = controller.NewApplicationController(applicationService)

		// Job
		jobService    service.JobService       = service.NewPostService(jobRepo, applicationRepo)
		jobController controller.JobController = controller.NewJobController(jobService)
	)
	httpRouter.POST("/api/auth/signup", authController.Signup)
	httpRouter.POST("/api/auth/signin", authController.Signin)
	httpRouter.PUT("/api/job", jobController.UpdateJob, middleware.ValidateToken(), middleware.OnlyCompanyAccess())
	httpRouter.POST("/api/job", jobController.NewJob, middleware.ValidateToken(), middleware.OnlyCompanyAccess())
	httpRouter.GET("/api/job", jobController.GetJobs)
	httpRouter.DELETE("/api/job/:id", jobController.DeleteJob, middleware.ValidateToken(), middleware.OnlyCompanyAccess())
	// application routes
	httpRouter.POST("/api/application/:id/apply", applicationController.ApplyToJob, middleware.ValidateToken())
	httpRouter.GET("/api/application/user", applicationController.GetUserApplications, middleware.ValidateToken(), middleware.OnlyUserAccess())
	httpRouter.DELETE("/api/application/:id", applicationController.DeleteApplication, middleware.ValidateToken(), middleware.OnlyUserAccess())
	// prospective routes
	httpRouter.POST("/api/prospective", prospectController.NewProspect)
	httpRouter.Serve(os.Getenv("PORT"))
}
