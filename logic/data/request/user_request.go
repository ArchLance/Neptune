package request

type UserLoginRequest struct {
	Account  string `validate:"required,max=64,min=1"  json:"account"`
	Password string `validate:"required,max=64,min=1" json:"password"`
}

type UpdateUserRequest struct {
	UserId   int    `json:"userid"`
	UserName string `validate:"required,max=64,min=1" json:"username"`
	Account  string `validate:"required,max=64,min=1"  json:"account"`
	Email    string `validate:"required,max=64,min=1" json:"email"`
	Role     string `validate:"required,max=20,min=1" json:"role"`
}
