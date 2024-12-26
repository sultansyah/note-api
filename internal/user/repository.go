package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sultansyah/note-api/internal/helper"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user User) (User, error)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (User, error)
	EditName(ctx context.Context, tx *sql.Tx, user User) error
	EditPassword(ctx context.Context, tx *sql.Tx, user User) error
	EditEmail(ctx context.Context, tx *sql.Tx, user User) error
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user User) (User, error) {
	sql := "insert into users(name, email, password, role) values(?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}

	user.Id = int(id)

	return user, nil
}

func (u *UserRepositoryImpl) EditName(ctx context.Context, tx *sql.Tx, user User) error {
	sql := "update users set name = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, user.Name, user.Id)
	fmt.Println("r = ", err)
	return err
}

func (u *UserRepositoryImpl) EditEmail(ctx context.Context, tx *sql.Tx, user User) error {
	sql := "update users set email = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, user.Email, user.Id)
	return err
}

func (u *UserRepositoryImpl) EditPassword(ctx context.Context, tx *sql.Tx, user User) error {
	sql := "update users set password = ? where id = ?"

	_, err := tx.ExecContext(ctx, sql, user.Password, user.Id)
	return err
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (User, error) {
	var user User

	sql := `select id, name, email, role, created_at, updated_at
			from users
			where id = ?
			`

	row, err := tx.QueryContext(ctx, sql, userId)
	if err != nil {
		return user, err
	}

	defer row.Close()

	if row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return user, err
		}

		return user, nil
	}

	return user, helper.ErrNotFound
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (User, error) {
	var user User

	sql := `select id, name, email, role, created_at, updated_at
			from users
			where email = ?
			`

	row, err := tx.QueryContext(ctx, sql, email)
	if err != nil {
		return user, err
	}

	defer row.Close()

	if row.Next() {
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return user, err
		}

		return user, nil
	}

	return user, helper.ErrNotFound
}
