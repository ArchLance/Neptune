package request

type CreateManagerRequest struct {
	Level    int    `json:"level"`
	Name     string `validate:"required,max=255,min=1" json:"name"`
	Account  string `validate:"required,max=255,min=1" json:"account"`
	Password string `validate:"required,max=255,min=1" json:"password"`
}
type UpdateManagerRequest struct {
	Id       int    `json:"id"`
	Level    int    `json:"level"`
	Name     string `validate:"required,max=255,min=1" json:"name"`
	Account  string `validate:"required,max=255,min=1" json:"account"`
	Password string `validate:"required,max=255,min=1" json:"password"`
}
