package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sultansyah/note-api/internal/helper"
	"github.com/sultansyah/note-api/internal/token"
)

func AuthMiddleware(tokenService token.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			helper.HandleErrorResponse(c, helper.ErrUnauthorized)
			c.Abort()
			return
		}

		tokenString := ""

		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := tokenService.ValidateToken(tokenString)
		if err != nil {
			helper.HandleErrorResponse(c, helper.ErrUnauthorized)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helper.HandleErrorResponse(c, helper.ErrUnauthorized)
			c.Abort()
			return
		}

		userId := int(claims["user_id"].(float64))
		userRole := claims["role"]

		c.Set("userId", userId)
		c.Set("userRole", userRole)

		c.Next()
	}
}
