package main

import (
	"backend/pkg/app"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func run() error {
	router := gin.Default()
	router.Use(cors.Default())

	router.SetTrustedProxies([]string{"127.0.0.1"})

	app.Routes(router)

	if err := router.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := run(); err != nil {
		log.Fatalf("This is the startup error: %v", err)
		os.Exit(1)
	}
}
