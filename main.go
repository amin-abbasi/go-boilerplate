package main

import (
	// Add my Packages
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amin-abbasi/go-boilerplate/configs"
	"github.com/amin-abbasi/go-boilerplate/handlers"
	"github.com/amin-abbasi/go-boilerplate/middlewares"
	"github.com/amin-abbasi/go-boilerplate/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize MongoDB connection
	services.ConnectDB()
	services.ConnectRedis()
	defer services.DisconnectDB() // Ensure database is closed when main function exits

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

	// Start server in a separate goroutine
	go func() {
		serverPort := configs.GetEnvVariable("SERVER_PORT")
		if err := app.Start(":" + serverPort); err != nil {
			log.Printf("<<< Error starting server >>> : %v", err)
		}
	}()
	// if err := app.Start(":4000"); err != nil {
	// 	app.Logger.Fatal(err)
	// }

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Give some time for ongoing requests to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := app.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v\n", err)
	}

	log.Println("Server shutdown complete.")
}
