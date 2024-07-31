package validation

import (
	"encoding/json"
	"errors"
	"go-api-jwt/src/configuration/rest_error"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_error.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_error.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_error.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_error.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_error.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_error.NewBadRequestError("Error trying to convert fields")
	}
}
