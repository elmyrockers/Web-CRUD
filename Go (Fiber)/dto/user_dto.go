package dto




type NewUserRequest struct {
	Name     string `form:"name" validate:"required,min=3"`
	Email    string `form:"email" validate:"required,email"`
	Website  string `form:"website" validate:"required,url"`
}

type EditUserRequest struct {
	Name  string `form:"name" validate:"required,min=3"`
	Email string `form:"email" validate:"required,email"`
}