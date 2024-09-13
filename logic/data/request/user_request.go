package request

type UserLoginRequest struct {
	Account  string `validate:"required,max=64,min=1"  json:"account"`
	Password string `validate:"required,max=64,min=1" json:"password"`
}

type UpdateUserRequest struct {
	UserId   uint   `json:"userid"`
	Avatar   string `json:"avatar"`
	UserName string `validate:"required,max=64,min=1" json:"username"`
	Account  string `validate:"required,max=64,min=1"  json:"account"`
	Email    string `validate:"required,max=64,min=1" json:"email"`
	Role     string `validate:"required,max=20,min=1" json:"role"`
}

type UserChangePasswordRequest struct {
	UserId      uint   `validate:"max=64,min=1"  json:"user_id"`
	OldPassword string `validate:"required,max=64,min=1" json:"old_password"`
	NewPassword string `validate:"required,max=64,min=1" json:"new_password"`
}
