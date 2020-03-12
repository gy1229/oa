package json_struct

type UserBase struct {
	Account string `json:"account"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type RegisterUserRequest struct {
	UserBase *UserBase `json:"user_base"`
	UserName string `json:"user_name"  binding:"required"`
}

type RegisterUserResponse struct {
	base *BaseResponse `json:"base"`
}

type LoginUserRequest struct {
	UserBase *UserBase `json:"user_base"  binding:"required"`
}

type LoginUserResponse struct {
	base *BaseResponse `json:"base"`
}

type UpdateUserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	base *BaseResponse `json:"base"`
}


type LoadUserMessageRequest struct {
	Account string `json:"account" binding`
}

type LoadUserMessageResponse struct {
	Account string `json:"account" binding`
	UserName string `json:"user_name"  binding:"required"`
}