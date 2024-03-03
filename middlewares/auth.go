package middlewares

import (
	"fmt"

	"github.com/amin-abbasi/go-boilerplate/configs"
	srv "github.com/amin-abbasi/go-boilerplate/services"

	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		tokenString := ctx.Request().Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		if tokenString == "" {
			return srv.SendResponse(ctx, 401, "Token is missing.")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.GetEnvVariable("JWT_SECRET_KEY")), nil
		})
		fmt.Println(">>>> Auth - token: ", token)

		if err != nil || !token.Valid {
			fmt.Println(">>>> Invalid Token: ", err)
			return srv.SendResponse(ctx, 401, "Invalid Token.")
		}

		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(">>>> Auth - claims: ", claims)
		ctx.Set("claims", claims)
		return next(ctx)
	}
}
