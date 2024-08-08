package main

import (
	"go_app/controllers"
	"go_app/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	userStore := database.NewUserStore()
	userController := controllers.NewUserController(userStore)

	e.POST("/register", userController.RegisterUser)
	e.GET("/users/:id", userController.GetUser)
	e.GET("/login/google", userController.GoogleLogin) // Initiates login
	e.GET("/auth/callback", userController.GoogleCallback) // Handles Google callback

	// Configure CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Adjust this as needed
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}