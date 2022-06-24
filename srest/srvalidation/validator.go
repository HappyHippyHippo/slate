package srvalidation

import (
	"github.com/go-playground/validator/v10"
	"github.com/happyhippyhippo/slate/srest/srenvelope"
)

// Validator is a function type used to define a calling interface of
// function responsible to validate an instance of a structure and return
// an initialized response envelope with the founded errors
type Validator func(val interface{}) (*srenvelope.Envelope, error)

// NewValidator instantiates a new validation function
func NewValidator(validate *validator.Validate, parser Parser) (Validator, error) {
	if validate == nil {
		return nil, errNilPointer("validate")
	}
	if parser == nil {
		return nil, errNilPointer("parser")
	}

	return func(value interface{}) (*srenvelope.Envelope, error) {
		if value == nil {
			return nil, errNilPointer("value")
		}

		if errs := validate.Struct(value); errs != nil {
			return parser.Parse(value, errs.(validator.ValidationErrors))
		}

		return nil, nil
	}, nil
}
