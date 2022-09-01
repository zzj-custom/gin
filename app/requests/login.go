package requests

type UserLoginRequest struct {
	UserName string `form:"userName" binding:"required,gt=0"`
	Password string `form:"password" binding:"required,gt=0"`
}

type RegisterRequest struct {
	UserLoginRequest
}
