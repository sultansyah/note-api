package user

import "time"

type UserWithToken struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserFormatter(user User, token string) UserWithToken {
	return UserWithToken{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Token:     token,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
