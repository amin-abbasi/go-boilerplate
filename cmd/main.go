package main

import (
	// Add my Packages
	"github.com/amin4193/go-boilerplate/handlers"
	"github.com/amin4193/go-boilerplate/middlewares"
	"github.com/amin4193/go-boilerplate/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initDB() {
	models.DB["amin"] = models.User{UserName: "amin", Password: "123"}
}

func main() {
	initDB()

	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// Routes
	app.GET("/ping", handlers.Ping)
	app.GET("/user/:name", handlers.GetUser)
	app.POST("/login", handlers.Login)

	authorized := app.Group("/admin")
	authorized.Use(middlewares.Auth)
	authorized.POST("/user", handlers.CreateUser)

	// Start server
	app.Logger.Fatal(app.Start(":4000"))
}
