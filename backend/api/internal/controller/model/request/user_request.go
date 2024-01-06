package request

type UserRequest struct {
	Name  string `json:"name" binding:"required,min=4,max=100"`
	Email string `json:"email" binding:"required,email"`
}
