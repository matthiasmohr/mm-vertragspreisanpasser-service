package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const numberOfSubstrings = 2

// Validator represents custom validator that validates structs based on tags.
type Validator struct {
	validateStruct      *validator.Validate
	universalTranslator *ut.UniversalTranslator
	translator          ut.Translator
	translations        map[string]string
}

// Error represents validation error.
type Error struct {
	Errors []string `json:"validation_errors"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("validation error: (%s)", strings.Join(e.Errors, ", "))
}

// NewValidator returns instance of validator.
func NewValidator() (*Validator, error) {
	en := en.New()
	validator := &Validator{
		validateStruct:      validator.New(),
		universalTranslator: ut.New(en, en),
	}
	validator.setTranslations()

	if err := validator.registerTranslations("en"); err != nil {
		return nil, err
	}

	return validator, nil
}

// Validate validates the struct and returns translated errors.
func (v *Validator) Validate(i interface{}) error {
	if err := v.validateStruct.Struct(i); err != nil {
		var vErrs validator.ValidationErrors
		if errors.As(err, &vErrs) {
			errs := v.translateErrors(vErrs)

			return &Error{Errors: errs}
		}
	}

	return nil
}

func (v *Validator) setTranslations() {
	v.translations = map[string]string{
		"required": "{0} is a required field",
		"gte":      "{0} must be equal or greater than {1} chars",
		"lte":      "{0} must be equal or less than {1} chars",
		"email":    "{0} invalid email address",
	}
}

// setupTranslations sets up the translations for certain validation rules.
func (v *Validator) registerTranslations(language string) error {
	v.translator, _ = v.universalTranslator.GetTranslator(language)

	for condition, translation := range v.translations {
		condition := condition
		translation := translation

		if err := v.validateStruct.RegisterTranslation(
			condition,
			v.translator,
			func(ut ut.Translator) error {
				if err := ut.Add(condition, translation, true); err != nil {
					return fmt.Errorf("translation Add func returned an error: %w", err)
				}

				return nil
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				pathAndName := strings.SplitN(fe.Namespace(), ".", numberOfSubstrings)[1]
				t, _ := ut.T(fe.Tag(), pathAndName, fe.Param())

				return t
			},
		); err != nil {
			return fmt.Errorf("register validator translation: %w", err)
		}
	}

	v.validateStruct.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", numberOfSubstrings)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return nil
}

func (v *Validator) translateErrors(errs validator.ValidationErrors) []string {
	translatedErrsStr := []string{}
	for _, fieldErr := range errs {
		translatedErrsStr = append(translatedErrsStr, fieldErr.Translate(v.translator))
	}

	return translatedErrsStr
}
