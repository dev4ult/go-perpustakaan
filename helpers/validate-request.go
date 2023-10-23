package helpers

import "github.com/go-playground/validator/v10"

func ValidateRequest(data interface{}) error {
	var validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(data)

	return err
}