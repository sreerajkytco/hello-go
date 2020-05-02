package user

// CreateUserDto for create request
type CreateUserDto struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}