package user

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FindUserRequest struct {
	Id string `uri:"id"`
}

type EditNameUserRequest struct {
	Name string `json:"name"`
}

type EditPasswordUserRequest struct {
	Password string `json:"password"`
}

type EditEmailUserRequest struct {
	Email string `json:"email"`
}
