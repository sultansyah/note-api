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

// @Summary      Register new user
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body user.CreateUserRequest true "User registration details"
// @Success      200  {object}  helper.WebResponse{data=user.UserWithToken}
// @Router       /auth/register [post]
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
		Data:    UserFormatterWithToken(user, token),
	})
}

// @Summary      User login
// @Description  Authenticate user and get token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body user.LoginUserRequest true "User login credentials"
// @Success      200  {object}  helper.WebResponse{data=user.UserWithToken}
// @Router       /auth/login [post]
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
		Data:    UserFormatterWithToken(user, token),
	})
}

// @Summary      Edit user email
// @Description  Update authenticated user's email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body user.EditEmailUserRequest true "New email"
// @Security     BearerAuth
// @Success      200  {object}  helper.WebResponse
// @Router       /auth/email [post]
func (u *UserHandlerImpl) EditEmail(c *gin.Context) {
	var input EditEmailUserRequest

	if !helper.BindAndValidateJSON(c, &input) {
		return
	}

	userId := c.MustGet("userId").(int)

	err := u.UserService.EditEmail(c.Request.Context(), input, userId)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success change email",
		Data:    nil,
	})
}

// @Summary      Edit user name
// @Description  Update authenticated user's name
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body user.EditNameUserRequest true "New name"
// @Security     BearerAuth
// @Success      200  {object}  helper.WebResponse
// @Router       /auth/name [post]
func (u *UserHandlerImpl) EditName(c *gin.Context) {
	var input EditNameUserRequest

	if !helper.BindAndValidateJSON(c, &input) {
		return
	}

	userId := c.MustGet("userId").(int)

	err := u.UserService.EditName(c.Request.Context(), input, userId)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success change name",
		Data:    nil,
	})
}

// @Summary      Edit user password
// @Description  Update authenticated user's password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body user.EditPasswordUserRequest true "New password"
// @Security     BearerAuth
// @Success      200  {object}  helper.WebResponse
// @Router       /auth/password [post]
func (u *UserHandlerImpl) EditPassword(c *gin.Context) {
	var input EditPasswordUserRequest

	if !helper.BindAndValidateJSON(c, &input) {
		return
	}

	userId := c.MustGet("userId").(int)

	err := u.UserService.EditPassword(c.Request.Context(), input, userId)
	if err != nil {
		helper.HandleErrorResponse(c, err)
		return
	}

	helper.APIResponse(c, helper.WebResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: "success change password",
		Data:    nil,
	})
}
