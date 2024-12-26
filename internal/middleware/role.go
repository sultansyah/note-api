package middleware

import (
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/sultansyah/note-api/internal/helper"
)

func RoleMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.MustGet("userRole").(string)

		if !slices.Contains(roles, userRole) {
			helper.HandleErrorResponse(c, helper.ErrUnauthorized)
			c.Abort()
		}

		c.Next()
	}
}
