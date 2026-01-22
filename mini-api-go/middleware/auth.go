package middleware

import (
	"mini-api-go/config"
	"mini-api-go/models"
	"mini-api-go/server"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next server.HandlerFunc) server.HandlerFunc {
	return func(c *server.Context) {
		authHewader := c.Request.Header.Get("Authorization")
		if authHewader == "" {
			models.ResponseError(c, models.NewAppError("No hay token", http.StatusUnauthorized))
			return
		}
		parts := strings.Split(authHewader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			models.ResponseError(c, models.NewAppError("Formato deToken invalido", http.StatusUnauthorized))
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, models.NewAppError("Metodo de firma inesperado", http.StatusUnauthorized)
			}
			return []byte(config.AppConfig.JwtSecret), nil
		})
		if err != nil || !token.Valid {
			models.ResponseError(c, models.NewAppError("Token invalido", http.StatusUnauthorized))
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userID, ok := claims["user_id"].(float64)
		if !ok {
			models.ResponseError(c, models.NewAppError("Error al obtener el id del usuario", http.StatusUnauthorized))
			return
		}
		c.SetUserID(uint(userID))
		next(c)
	}
}
