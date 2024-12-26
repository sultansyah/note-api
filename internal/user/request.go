package user

type CreateUserRequest struct {
	Name     string `json:"name" example:"example"`
	Email    string `json:"email" example:"example@example.com"`
	Password string `json:"password" example:"secret123"`
	Role     string `json:"role" example:"user"`
}

type LoginUserRequest struct {
	Email    string `json:"email" example:"example@example.com"`
	Password string `json:"password" example:"secret123"`
}

type EditNameUserRequest struct {
	Name string `json:"name" example:"example"`
}

type EditEmailUserRequest struct {
	Email string `json:"email" example:"example@example.com"`
}

type EditPasswordUserRequest struct {
	Password string `json:"password" example:"newpassword123"`
}

type FindUserRequest struct {
	Id string `uri:"id"`
}
