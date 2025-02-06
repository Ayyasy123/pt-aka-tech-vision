package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, err.Error())
		}
		return errors.New(formatValidationError(err))
	}

	return nil
}

func formatValidationError(err error) string {
	if errors, ok := err.(validator.ValidationErrors); ok {
		var errorMessages []string
		for _, err := range errors {
			errorMessages = append(errorMessages, err.Error())
		}
		return errorMessages[0]
	}
	return err.Error()
}
