package response

type ManagerResponse struct {
	Id       int
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
