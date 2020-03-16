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
	Base *BaseResponse `json:"base"`
}

type LoginUserRequest struct {
	UserBase *UserBase `json:"user_base"  binding:"required"`
}

type LoginUserResponse struct {
	UserId *int64 `json:"user_id"`
	Base *BaseResponse `json:"base"`
}

type UpdateUserRequest struct {
	Account string `json:"account"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	Base *BaseResponse `json:"base"`
}


type LoadUserMessageRequest struct {
	UserId *int64 `json:"user_id" binding`
}

type LoadUserMessageResponse struct {
	Account string `json:"account" binding`
	UserName string `json:"user_name"  binding:"required"`
	Base *BaseResponse `json:"base"`
}

type CertainAccountRequest struct {
	Account string `json:"account"`
}

type CertainAccountResponse struct {
	Base *BaseResponse `json:"base"`
}