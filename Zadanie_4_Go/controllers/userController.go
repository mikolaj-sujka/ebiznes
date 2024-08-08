package controllers

import (
	"context"
	"encoding/json"
	"go_app/database"
	"go_app/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserController struct {
	Store *database.UserStore
}

func NewUserController(store *database.UserStore) *UserController {
	return &UserController{
		Store: store,
	}
}

// RegisterUser handles the user registration endpoint
func (uc *UserController) RegisterUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data provided"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	u.Password = string(hashedPassword)
	u.ID = uuid.New().String()

	uc.Store.AddUser(u)

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully", "user": u})
}

// GetUser handles fetching a user by ID
func (uc *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, exists := uc.Store.GetUser(id)
	if !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"user": user})
}

// GoogleLogin handles the OAuth2 login process with Google
func (uc *UserController) GoogleLogin(c echo.Context) error {
	// Generate a URL to Google's OAuth 2.0 server
	url := oauthConf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback handles the callback from Google with the authorization code
func (uc *UserController) GoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != "state-token" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "State is invalid"})
	}

	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Authorization code not provided"})
	}

	// Exchange the authorization code for an access token
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Failed to exchange code for token"})
	}

	// Retrieve user information from Google
	client := oauthConf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Failed to retrieve user info"})
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to decode user info"})
	}

	// Check if the user already exists in the database
	existingUser, exists := uc.Store.GetUserByEmail(userInfo.Email)
	if exists {
		// User exists, return their data
		return c.JSON(http.StatusOK, echo.Map{
			"message": "User logged in successfully",
			"user":    existingUser,
		})
	}

	// Create a new user if they don't exist
	newUser := &models.User{
		ID:       uuid.New().String(),
		Username: userInfo.Name,
		Email:    userInfo.Email,
		Password: "", // Password not needed for Google login users
	}

	uc.Store.AddUser(newUser)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "New user created and logged in successfully",
		"user":    newUser,
	})
}


var oauthConf = &oauth2.Config{
	ClientID:     "430307597286-tl0u26vvprdh8864ipsloi67gd6cigfq.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/auth/callback", // Replace with your frontend URL
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint:     google.Endpoint,
}