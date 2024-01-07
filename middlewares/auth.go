package middlewares

import (
	"github.com/amin4193/go-boilerplate/models"

	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenString := ctx.Request().Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == "" {
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "missing token"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return models.SECRET_KEY, nil
		})

		if err != nil || !token.Valid {
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)
		return next(ctx)
	}
}
