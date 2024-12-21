package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sultansyah/note-api/internal/helper"
	"github.com/sultansyah/note-api/internal/token"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	EditName(c *gin.Context)
	EditPassword(c *gin.Context)
	EditEmail(c *gin.Context)
}

type UserHandlerImpl struct {
	UserService  UserService
	TokenService token.TokenService
}

func NewUserHandler(userService UserService, tokenService token.TokenService) UserHandler {
	return &UserHandlerImpl{UserService: userService, TokenService: tokenService}
}

func (u *UserHandlerImpl) Register(c *gin.Context) {
	var input CreateUserRequest

	if !helper.BindAndValidateJSON(c, &input) {
		return
	}

	user, err := u.UserService.Create(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	token, err := u.TokenService.GenerateToken(user.Id, user.Role)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "register success",
		Data:    UserFormatter(user, token),
	})
}

func (u *UserHandlerImpl) Login(c *gin.Context) {
	var input LoginUserRequest

	if !helper.BindAndValidateJSON(c, &input) {
		return
	}

	user, err := u.UserService.Login(c.Request.Context(), input)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	token, err := u.TokenService.GenerateToken(user.Id, user.Role)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "login success",
		Data:    UserFormatter(user, token),
	})
}

func (u *UserHandlerImpl) EditEmail(c *gin.Context) {
}

func (u *UserHandlerImpl) EditName(c *gin.Context) {
	panic("unimplemented")
}

func (u *UserHandlerImpl) EditPassword(c *gin.Context) {
	panic("unimplemented")
}
