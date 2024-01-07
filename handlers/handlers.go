package handlers

import (
	"github.com/amin4193/go-boilerplate/models"

	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func Login(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	// Get user from DB
	storedUser, ok := models.DB[user.UserName]

	// Validate credentials (this is a simple example, use your authentication logic here)
	if ok && storedUser.Password == user.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.UserName
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

		tokenString, err := token.SignedString(models.SECRET_KEY)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "error generating token"})
		}

		// adds user in db
		models.DB[user.UserName] = user

		return ctx.JSON(http.StatusOK, map[string]interface{}{"token": tokenString})
	}

	return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "invalid credentials"})
}

func GetUser(ctx echo.Context) error {
	username := ctx.Param("name")
	value, ok := models.DB[username]
	if ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"username": username, "value": value})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"username": username, "status": "no value"})
}

func CreateUser(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	// Check if the username already exists in the db
	if _, exists := models.DB[user.UserName]; exists {
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": "username already exists"})
	}

	// adds user in db
	models.DB[user.UserName] = user
	return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "user": user})
}
