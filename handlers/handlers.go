package handlers

import (
	"fmt"

	"github.com/amin4193/go-boilerplate/configs"
	"github.com/amin4193/go-boilerplate/models"
	srv "github.com/amin4193/go-boilerplate/services"

	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

var (
	adminUsername = "admin"
	adminPassword = "1234"
)

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginAdmin(ctx echo.Context) error {
	// Bind the request body to the AdminLoginRequest struct
	req := new(AdminLoginRequest)
	if err := ctx.Bind(req); err != nil {
		// return srv.SendResponse(ctx, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return srv.SendResponse(ctx, 400, "Invalid request body")
	}

	// Validate credentials (this is a simple example, use your authentication logic here)
	if req.Username == adminUsername && req.Password == adminPassword {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = adminUsername
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

		// Convert string secret key to byte slice
		secretKeyBytes := []byte(configs.SECRET_KEY)

		tokenString, err := token.SignedString(secretKeyBytes)
		if err != nil {
			fmt.Println(">>>> Error Generating Token: ", err)
			return srv.SendResponse(ctx, 500, "Error Generating Token", err)
		}

		return srv.SendResponse(ctx, 200, "success", map[string]interface{}{"token": tokenString})
	}

	return srv.SendResponse(ctx, 401, "invalid credentials")
}

func Login(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return srv.SendResponse(ctx, 400, "body validation error", err)
	}

	// Get user from DB
	storedUser, err := models.GetByUsername(user.UserName)
	if err != nil {
		fmt.Println(">>>> User not found: ", err)
		return srv.SendResponse(ctx, 401, "User not found.")
	}

	// Validate credentials (this is a simple example, use your authentication logic here)
	if storedUser.Password == user.Password {
		return srv.SendResponse(ctx, 401, "Invalid Credentials.")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	tokenString, err := token.SignedString(configs.SECRET_KEY)
	if err != nil {
		fmt.Println(">>>> Error on generating token: ", err)
		return srv.SendResponse(ctx, 500, "Error on generating token.", err)
	}

	return srv.SendResponse(ctx, 200, "success", map[string]interface{}{"token": tokenString})
}

func GetUser(ctx echo.Context) error {
	username := ctx.Param("name")
	user, err := models.GetByUsername(username)
	if err != nil {
		return srv.SendResponse(ctx, 401, "User not found.")
	}
	return srv.SendResponse(ctx, 200, "success", map[string]interface{}{"user": user})
}

func CreateUser(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return srv.SendResponse(ctx, 400, "Body Validation Error", err)
	}

	// Check if the username already exists in the db
	exists, _ := models.GetByUsername(user.UserName)
	if exists != nil {
		return srv.SendResponse(ctx, 409, "Username already exists.")
	}

	// adds user in db
	newUser, err := models.User.Create(user)
	if err != nil {
		fmt.Println(">>>> Could not create user: ", err)
		return srv.SendResponse(ctx, 401, "Could not create user.", err)
	}
	return srv.SendResponse(ctx, 200, "success", map[string]interface{}{"status": "ok", "user": newUser})
}
