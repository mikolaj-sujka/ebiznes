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

func (uc *UserController) LoginUser(c echo.Context) error {
	credentials := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})

	if err := c.Bind(credentials); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	user, exists := uc.Store.GetUserByEmail(credentials.Email)
	if !exists {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User logged in successfully",
		"user": models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func (uc *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, exists := uc.Store.GetUser(id)
	if !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"user": user})
}

func (uc *UserController) GoogleLogin(c echo.Context) error {
	url := oauthConf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (uc *UserController) GoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != "state-token" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "State is invalid"})
	}

	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Authorization code not provided"})
	}

	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Failed to exchange code for token"})
	}

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

	existingUser, exists := uc.Store.GetUserByEmail(userInfo.Email)
	if exists {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "User logged in successfully",
			"user":    existingUser,
		})
	}

	newUser := &models.User{
		ID:       uuid.New().String(),
		Username: userInfo.Name,
		Email:    userInfo.Email,
		Password: "", 
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
	RedirectURL:  "http://localhost:8080/auth/callback", 
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint:     google.Endpoint,
}