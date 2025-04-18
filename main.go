package main

import (
	"log"
	"os"

	"github.com/bonjourrog/jb/routes"
	"github.com/joho/godotenv"
)

var (
	httpRouter routes.Router = routes.NewGinRouter()
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading env file")
	}

	httpRouter.Serve(os.Getenv("PORT"))
}
