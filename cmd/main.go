package main

import (
	// Add my Packages
	"github.com/amin4193/go-boilerplate/services"
	"github.com/amin4193/go-boilerplate/handlers"
	"github.com/amin4193/go-boilerplate/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize MongoDB connection
	services.ConnectDB()

	// Initiate Echo Application
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.GET("/ping", handlers.Ping)
	app.POST("/admin/login", handlers.LoginAdmin)
	app.POST("/login", handlers.Login)
	app.GET("/user/:name", handlers.GetUser)

	authorized := app.Group("/admin")
	authorized.Use(middlewares.Auth)
	authorized.POST("/user", handlers.CreateUser)

	// Start server
	app.Logger.Fatal(app.Start(":4000"))
}
