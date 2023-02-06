package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rest"
)

const (
	// ID defines the id to be used
	// as the container registration id of a validation.
	ID = rest.ID + ".validation"

	// UniversalTranslatorID defines the id to be used
	// as the container registration id of a universal translator.
	UniversalTranslatorID = ID + ".universal_translator"

	// TranslatorID defines the id to be used
	// as the container registration id of a translator.
	TranslatorID = ID + ".translator"

	// ParserID defines the id to be used
	// as the container registration id of an error parser instance.
	ParserID = ID + ".parser"
)

// Provider @todo doc
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the validation package instances in the
// application container
func (p Provider) Register(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// register a universal translator
	_ = container[0].Service(UniversalTranslatorID, func() *ut.UniversalTranslator {
		lang := en.New()
		return ut.New(lang, lang)
	})
	// register a translator instance of the defined default locale
	_ = container[0].Service(TranslatorID, func(universalTranslator *ut.UniversalTranslator) (ut.Translator, error) {
		translator, found := universalTranslator.GetTranslator(Locale)
		if found == false {
			return nil, errTranslatorNotFound(Locale)
		}
		return translator, nil
	})
	// register a validation error parser
	_ = container[0].Service(ParserID, NewParser)
	// register a validation method service
	_ = container[0].Service(ID, func(translator ut.Translator, parser IParser) (Validator, error) {
		validate := validator.New()
		_ = translations.RegisterDefaultTranslations(validate, translator)
		return NewValidator(validate, parser)
	})
	return nil
}

// Boot will start the validation package
func (p Provider) Boot(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	return nil
}
