package sdk

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func initValidator() {
	Validate = validator.New()
	// Register custom validation tags if needed
	// Validate.RegisterValidation("customTag", CustomValidationFunc)
}

func GetValidator() *validator.Validate {
	if Validate == nil {
		initValidator()
	}
	return Validate
}
