package response

type UserResponse struct {
	UserId   int    `json:"userid"`
	Token    string `json:"token"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Account  string `json:"account"`
	Role     string `json:"role"`
}
