package dto

import "github.com/go-playground/validator/v10"





// Validate runs validator.Struct on any input struct
func Validate(s interface{}) error {
	var validate = validator.New()
	return validate.Struct(s)
}