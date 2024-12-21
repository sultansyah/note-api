package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAndValidateJSON(c *gin.Context, input any) bool {
	if err := c.ShouldBindJSON(input); err != nil {
		APIResponse(c, WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: "invalid request payload",
			Data:    err.Error(),
		})
		return false
	}
	return true
}

func BindAndValidateURi(c *gin.Context, input any) bool {
	if err := c.ShouldBindUri(input); err != nil {
		errors := FormatValidationErrors(err)

		APIResponse(c, WebResponse{
			Code:    http.StatusUnprocessableEntity,
			Status:  "error",
			Message: "invalid request payload",
			Data:    errors,
		})
		return false
	}
	return true
}
