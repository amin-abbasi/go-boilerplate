package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	db        = make(map[string]User)
	secretKey = []byte("your_secret_key") // Replace with your secret key
)

// User struct to represent a user
type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getUser(ctx echo.Context) error {
	username := ctx.Param("name")
	value, ok := db[username]
	if ok {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"username": username, "value": value})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{"username": username, "status": "no value"})
}

func login(ctx echo.Context) error {
	var user User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	// Get user from DB
	storedUser, ok := db[user.UserName]

	// Validate credentials (this is a simple example, use your authentication logic here)
	if ok && storedUser.Password == user.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.UserName
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "error generating token"})
		}

		// adds user in db
		db[user.UserName] = user

		return ctx.JSON(http.StatusOK, map[string]interface{}{"token": tokenString})
	}

	return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "invalid credentials"})
}

func createUser(ctx echo.Context) error {
	var user User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	// Check if the username already exists in the db
	if _, exists := db[user.UserName]; exists {
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": "username already exists"})
	}

	// adds user in db
	db[user.UserName] = user
	return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "user": user})
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenString := ctx.Request().Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == "" {
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "missing token"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)
		return next(ctx)
	}
}

func main() {
	db["amin"] = User{UserName: "amin", Password: "123"}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "pong")
	})
	e.GET("/user/:name", getUser)
	e.POST("/login", login)

	authorized := e.Group("/admin")
	authorized.Use(authMiddleware)
	authorized.POST("/user", createUser)

	// Start server
	e.Logger.Fatal(e.Start(":4000"))
}
