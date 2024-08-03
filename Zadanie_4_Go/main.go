package main

import (
	"go_app/controllers"
	"go_app/database"
	"go_app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := router.New()

	userStore := database.NewUserStore()
    userController := controllers.NewUserController(userStore)

    e.POST("/register", userController.RegisterUser)
    e.GET("/users/:id", userController.GetUser)

	// Configure CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Adjust this as needed
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	router.Configure(e) // Assuming you have a function to configure routes

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
