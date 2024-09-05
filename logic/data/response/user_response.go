package response

type UserLoginResponse struct {
	UserId   int    `json:"userid"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Account  string `json:"account"`
	Role     string `json:"role"`
}

type UserResponse struct {
	UserId   int    `json:"userid"`
	Avatar   string `json:"avatar"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Account  string `json:"account"`
	Role     string `json:"role"`
}
