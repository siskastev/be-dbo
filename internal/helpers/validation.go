package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getParamOrDefault(param string) string {
	if param == "" {
		return "value"
	}
	return param
}

func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required."
	case "email":
		return "Invalid email format."
	case "min":
		return "This field must be at least " + fe.Param() + " characters long."
	case "max":
		return "This field must be at most " + fe.Param() + " characters long."
	default:
		return fe.Error()
	}
}

func parseFieldError(e validator.FieldError) ValidationError {
	fieldPrefix := fmt.Sprintf("The field %s", e.Field())
	tag := strings.Split(e.Tag(), "|")[0]

	customMessages := map[string]string{
		"required": fmt.Sprintf("%s is required", fieldPrefix),
		"email":    fmt.Sprintf("%s must be valid email", fieldPrefix),
		"min":      fmt.Sprintf("%s minimum length is %s", fieldPrefix, getParamOrDefault(e.Param())),
		"max":      fmt.Sprintf("%s maximum length is %s", fieldPrefix, getParamOrDefault(e.Param())),
	}

	if customMessage, found := customMessages[tag]; found {
		return ValidationError{Field: e.Field(), Message: customMessage}
	}

	// If it's a tag for which we don't have a custom message, use the default English translator
	english := en.New()
	translator := ut.New(english, english)
	if translatorInstance, found := translator.GetTranslator("en"); found {
		return ValidationError{Field: e.Field(), Message: e.Translate(translatorInstance)}
	}

	return ValidationError{Field: e.Field(), Message: fmt.Errorf("%v", e).Error()}
}

func HandleValidationErrors(err error) []ValidationError {
	var validationErrors []ValidationError

	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range fieldErrors {
			validationError := ValidationError{
				Field:   fe.Field(),
				Message: getErrorMessage(fe),
			}
			validationErrors = append(validationErrors, validationError)
		}
	}

	return validationErrors
}
