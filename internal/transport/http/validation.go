package http

import (
	"encoding/json"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ValidationErrorField struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type ValidationErrors struct {
	Status string                 `json:"status"`
	Reason string                 `json:"reason"`
	Fields []ValidationErrorField `json:"fields"`
}

func (v *ValidationErrors) Error() string {
	message, _ := json.Marshal(v)

	return string(message)
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		Status: "error",
		Reason: "some fields failed validation",
	}
}

func validate(input AddProductReviewRequest) error {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New(validator.WithRequiredStructEnabled())
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(input)
	if err != nil {
		validationErr := NewValidationErrors()
		for _, err := range err.(validator.ValidationErrors) {
			validationErr.Fields = append(validationErr.Fields, ValidationErrorField{
				Name:   err.Field(),
				Reason: err.Translate(trans),
			})
		}

		return validationErr
	}

	return nil
}
