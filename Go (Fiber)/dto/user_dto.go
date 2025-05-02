package dto




type NewUserRequest struct {
	Name     string `form:"name" validate:"required,min=3"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8"`
}

type EditUserRequest struct {
	Name  string `form:"name" validate:"required,min=3"`
	Email string `form:"email" validate:"required,email"`
}