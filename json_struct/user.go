package json_struct

type UserBase struct {
	Account  string `json:"account"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type RegisterUserRequest struct {
	UserBase *UserBase `json:"user_base"`
	UserName string    `json:"user_name"  binding:"required"`
	ImageId string `json:"image_id"`
}

type RegisterUserResponse struct {
	Base *BaseResponse `json:"base"`
}

type LoginUserRequest struct {
	UserBase *UserBase `json:"user_base"  binding:"required"`
}

type LoginUserResponse struct {
	UserId string        `json:"user_id"`
	Base   *BaseResponse `json:"base"`
}

type UpdateUserRequest struct {
	Account  string `json:"account"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UpdateUserResponse struct {
	Base *BaseResponse `json:"base"`
}

type LoadUserMessageRequest struct {
	UserId string `json:"user_id" binding`
}

type LoadUserMessageResponse struct {
	Account  string        `json:"account" binding`
	UserName string        `json:"user_name"  binding:"required"`
	ImageId string `json:"image_id"`
	Base     *BaseResponse `json:"base"`
}

type CertainAccountRequest struct {
	Account string `json:"account"`
}

type CertainAccountResponse struct {
	Base *BaseResponse `json:"base"`
}

type UploadAvatarRequest struct {
	FileContent []byte `json:"file_content"`
	Name        string `json:"name"`
	FileType string `json:"file_type"`
}
type UploadAvatarResponse struct {
	ImageId string         `json:"image_id"`
	Base    *BaseResponse `json:"base"`
}

type GetAvatarRequest struct {
	ImageId string `json:"image_id"`
}

type GetAvatarResponse struct {
	ImageFile []byte `json:"image_file"`
	Base    *BaseResponse `json:"base"`
}