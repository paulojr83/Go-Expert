package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/rest_err"
)

var (
	Validate   = validator.New()
	translator ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		enTranslator := ut.New(en, en)
		translator, _ = enTranslator.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, translator)
	}
}

func ValidateErr(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewNotFoundError("Invalid type error")
	} else if errors.Is(validationErr, &jsonValidation) {
		errorCauses := []rest_err.Causes{}

		for _, err := range validationErr.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   err.Field(),
				Message: err.Translate(translator),
			})
		}
		return rest_err.NewBadRequestError("Invalid field values", errorCauses...)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
