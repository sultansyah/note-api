package user

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type EditUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ChangeEmailUserRequest struct {
	Email string `json:"email"`
}
