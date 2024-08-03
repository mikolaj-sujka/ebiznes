package controllers

import (
	"go_app/database"
	"go_app/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
        return err
    }

    u.Password = string(hashedPassword)
    u.ID = uuid.New().String()

    uc.Store.AddUser(u)

    return c.JSON(http.StatusCreated, u)
}

func (uc *UserController) GetUser(c echo.Context) error {
    id := c.Param("id")
    user, exists := uc.Store.GetUser(id)
    if !exists {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
    }
    return c.JSON(http.StatusOK, user)
}