package user

type UserHandler interface {
	Register(input CreateUserRequest)
	Login(input LoginUserRequest)
	EditName(input EditNameUserRequest)
	EditPassword(input EditPasswordUserRequest)
	EditEmail(input EditEmailUserRequest)
}

type UserHandlerImpl struct {
	UserService UserService
}
