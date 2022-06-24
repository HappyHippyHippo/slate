package srvalidation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/happyhippyhippo/slate"
)

// Provider @todo doc
type Provider struct{}

var _ slate.ServiceProvider = &Provider{}

// Register will register the validation package instances in the
// application container
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Service(ContainerUniversalTranslatorID, func() (interface{}, error) {
		lang := en.New()
		return ut.New(lang, lang), nil
	})

	_ = c.Service(ContainerTranslatorID, func() (interface{}, error) {
		universalTranslator, err := GetUniversalTranslator(c)
		if err != nil {
			return nil, err
		}

		translator, found := universalTranslator.GetTranslator(Locale)
		if found == false {
			return nil, errTranslatorNotFound(Locale)
		}

		return translator, nil
	})

	_ = c.Service(ContainerParserID, func() (interface{}, error) {
		translator, err := GetTranslator(c)
		if err != nil {
			return nil, err
		}

		return NewParser(translator)
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		translator, err := GetTranslator(c)
		if err != nil {
			return nil, err
		}

		parser, err := GetParser(c)
		if err != nil {
			return nil, err
		}

		validate := validator.New()
		_ = translations.RegisterDefaultTranslations(validate, translator)

		return NewValidator(validate, parser)
	})

	return nil
}

// Boot will start the validation package
func (Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	return nil
}
