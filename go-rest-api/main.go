package main

import (
	"go-rest-api/config"
	logger "go-rest-api/middlewares"
	"go-rest-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize DB
	config.ConnectDB()
	// config.DB.AutoMigrate(&config.DB)

	// Initialize Echo
	e := echo.New()

	logger.InitLogger(true) // debug = true

	// Middleware
	e.Use(logger.EchoLoggerMiddleware())

	// Routes
	routes.RegisterRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8082"))
}
