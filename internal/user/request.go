package user

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required" example:"example"`
	Email    string `json:"email" binding:"required" example:"example@example.com"`
	Password string `json:"password" binding:"required" example:"secret123"`
	Role     string `json:"role" binding:"required" example:"user"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required" example:"example@example.com"`
	Password string `json:"password" binding:"required" example:"secret123"`
}

type EditNameUserRequest struct {
	Name string `json:"name" binding:"required" example:"example"`
}

type EditEmailUserRequest struct {
	Email string `json:"email" binding:"required" example:"example@example.com"`
}

type EditPasswordUserRequest struct {
	Password string `json:"password" binding:"required" example:"newpassword123"`
}

type FindUserRequest struct {
	Id string `uri:"id"`
}
