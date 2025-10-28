package request

type LoginRequest struct {
	Username string `json:"username" label:"用户名" validate:"required"`
	Password string `json:"password" label:"密码" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" label:"用户名" validate:"required,min=3,max=20"`
	Password string `json:"password" label:"密码" validate:"required,min=6,max=20"`
}