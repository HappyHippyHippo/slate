package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

// NewTranslator @todo doc
func NewTranslator(
	universalTranslator *ut.UniversalTranslator,
) (ut.Translator, error) {
	translator, found := universalTranslator.GetTranslator(Locale)
	if found == false {
		return nil, errTranslatorNotFound(Locale)
	}
	return translator, nil
}

// NewUniversalTranslator @todo doc
func NewUniversalTranslator() *ut.UniversalTranslator {
	lang := en.New()
	return ut.New(lang, lang)
}
