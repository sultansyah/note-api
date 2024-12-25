package user

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/sultansyah/note-api/internal/helper"
)

type UserService interface {
	Create(ctx context.Context, input CreateUserRequest) (User, error)
	FindById(ctx context.Context, input FindUserRequest) (User, error)
	Login(ctx context.Context, input LoginUserRequest) (User, error)
	EditName(ctx context.Context, input EditNameUserRequest, userId int) error
	EditPassword(ctx context.Context, input EditPasswordUserRequest, userId int) error
	EditEmail(ctx context.Context, input EditEmailUserRequest, userId int) error
}

type UserServiceImpl struct {
	UserRepository UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB}
}

func (u *UserServiceImpl) Login(ctx context.Context, input LoginUserRequest) (User, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return User{}, nil
	}

	defer helper.HandleTransaction(tx, &err)

	user, err := u.UserRepository.FindByEmail(ctx, tx, input.Email)
	if err != nil {
		return User{}, err
	}

	if user.Id == 0 {
		return User{}, helper.ErrNotFound
	}

	return user, nil
}

func (u *UserServiceImpl) Create(ctx context.Context, input CreateUserRequest) (User, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return User{}, err
	}

	defer helper.HandleTransaction(tx, &err)

	existUser, err := u.UserRepository.FindByEmail(ctx, tx, input.Email)
	if err != nil {
		return existUser, err
	}

	if existUser.Id > 0 {
		return User{}, helper.ErrAlreadyExists
	}

	passwordHash, err := HashPassword(input.Password)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Role:     input.Role,
		Password: passwordHash,
	}

	user, err = u.UserRepository.Create(ctx, tx, user)
	if err != nil {
		return user, err
	}

	return user, nil
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
