package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAlreadyExists = errors.New("resource already exists")
	ErrNotFound      = errors.New("resource not found")
	ErrInternal      = errors.New("internal server error")
	ErrUnauthorized  = errors.New("unauthorized")
)

func FormatValidationErrors(err error) []string {
	var errors []string

	for _, v := range err.(validator.ValidationErrors) {
		errors = append(errors, v.Error())
	}

	return errors
}

func HandleErrorResponse(c *gin.Context, err error) {
	webResponse := WebResponse{
		Data: nil,
	}

	switch err {
	case ErrAlreadyExists:
		webResponse.Code = http.StatusConflict
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case ErrNotFound:
		webResponse.Code = http.StatusNotFound
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case ErrInternal:
		webResponse.Code = http.StatusInternalServerError
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	case bcrypt.ErrMismatchedHashAndPassword:
		webResponse.Code = http.StatusUnauthorized
		webResponse.Status = "error"
		webResponse.Message = "email or password is incorrect"
	case ErrUnauthorized:
		webResponse.Code = http.StatusUnauthorized
		webResponse.Status = "error"
		webResponse.Message = "unauthorized"
	default:
		webResponse.Code = http.StatusInternalServerError
		webResponse.Status = "error"
		webResponse.Message = err.Error()
	}

	APIResponse(c, webResponse)
}
