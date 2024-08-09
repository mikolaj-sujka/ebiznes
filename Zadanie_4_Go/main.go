package main

import (
	"go_app/controllers"
	"go_app/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	userStore := database.NewUserStore()
	userController := controllers.NewUserController(userStore)

	e.POST("/register", userController.RegisterUser)
	e.GET("/users/:id", userController.GetUser)
	e.GET("/login/google", userController.GoogleLogin) 
	e.GET("/auth/callback", userController.GoogleCallback) 
	e.POST("/login", userController.LoginUser)

	// Configure CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, 
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}