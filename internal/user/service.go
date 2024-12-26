package user

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/sultansyah/note-api/internal/helper"
	"github.com/sultansyah/note-api/internal/token"
)

type UserService interface {
	Create(ctx context.Context, input CreateUserRequest) (UserWithToken, error)
	FindById(ctx context.Context, input FindUserRequest) (User, error)
	Login(ctx context.Context, input LoginUserRequest) (UserWithToken, error)
	EditName(ctx context.Context, input EditNameUserRequest, userId int) error
	EditPassword(ctx context.Context, input EditPasswordUserRequest, userId int) error
	EditEmail(ctx context.Context, input EditEmailUserRequest, userId int) error
}

type UserServiceImpl struct {
	UserRepository UserRepository
	DB             *sql.DB
	TokenService   token.TokenService
}

func NewUserService(userRepository UserRepository, DB *sql.DB, tokenService token.TokenService) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, TokenService: tokenService}
}

func (u *UserServiceImpl) Login(ctx context.Context, input LoginUserRequest) (UserWithToken, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return UserWithToken{}, nil
	}

	defer helper.HandleTransaction(tx, &err)

	user, err := u.UserRepository.FindByEmail(ctx, tx, input.Email)
	if err != nil {
		return UserWithToken{}, err
	}

	if user.Id == 0 {
		return UserWithToken{}, helper.ErrNotFound
	}

	isSame, err := CompareHashPassword(user.Password, input.Password)
	if err != nil {
		return UserWithToken{}, err
	}

	if !isSame {
		return UserWithToken{}, helper.ErrUnauthorized
	}

	token, err := u.TokenService.GenerateToken(user.Id, user.Role)
	if err != nil {
		return UserWithToken{}, err
	}

	return UserFormatterWithToken(user, token), nil
}

func (u *UserServiceImpl) Create(ctx context.Context, input CreateUserRequest) (UserWithToken, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return UserWithToken{}, err
	}

	defer helper.HandleTransaction(tx, &err)

	existUser, err := u.UserRepository.FindByEmail(ctx, tx, input.Email)
	if err != nil && err != helper.ErrNotFound {
		return UserWithToken{}, err
	}

	if existUser.Id > 0 {
		return UserWithToken{}, helper.ErrAlreadyExists
	}

	passwordHash, err := HashPassword(input.Password)
	if err != nil {
		return UserWithToken{}, err
	}

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Role:     input.Role,
		Password: passwordHash,
	}

	user, err = u.UserRepository.Create(ctx, tx, user)
	if err != nil {
		return UserWithToken{}, err
	}

	token, err := u.TokenService.GenerateToken(user.Id, user.Role)
	if err != nil {
		return UserWithToken{}, err
	}

	return UserFormatterWithToken(user, token), nil
}

func (u *UserServiceImpl) FindById(ctx context.Context, input FindUserRequest) (User, error) {
	var user User

	tx, err := u.DB.Begin()
	if err != nil {
		return user, err
	}

	defer helper.HandleTransaction(tx, &err)

	id, err := strconv.Atoi(input.Id)
	if err != nil {
		return user, err
	}

	user, err = u.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		return User{}, err
	}

	if user.Id < 0 {
		return User{}, helper.ErrNotFound
	}

	return user, nil
}

func (u *UserServiceImpl) EditEmail(ctx context.Context, input EditEmailUserRequest, userId int) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}

	defer helper.HandleTransaction(tx, &err)

	user, err := u.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return err
	}
	if user.Id < 0 {
		return helper.ErrNotFound
	}

	user = User{
		Id:    userId,
		Email: input.Email,
	}

	err = u.UserRepository.EditEmail(ctx, tx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserServiceImpl) EditName(ctx context.Context, input EditNameUserRequest, userId int) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}

	defer helper.HandleTransaction(tx, &err)

	user, err := u.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return err
	}

	if user.Id < 0 {
		return helper.ErrNotFound
	}

	user = User{
		Id:   userId,
		Name: input.Name,
	}

	err = u.UserRepository.EditName(ctx, tx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserServiceImpl) EditPassword(ctx context.Context, input EditPasswordUserRequest, userId int) error {
	tx, err := u.DB.Begin()
	if err != nil {
		return err
	}

	defer helper.HandleTransaction(tx, &err)

	user, err := u.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return err
	}

	if user.Id < 0 {
		return helper.ErrNotFound
	}

	passwordHash, err := HashPassword(input.Password)
	if err != nil {
		return err
	}

	user = User{
		Id:       userId,
		Password: passwordHash,
	}

	err = u.UserRepository.EditPassword(ctx, tx, user)
	if err != nil {
		return err
	}

	return nil
}
