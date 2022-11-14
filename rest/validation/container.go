package validation

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/happyhippyhippo/slate"
)

// GetUniversalTranslator will try to retrieve the registered universal translator
// instance from the application service container.
func GetUniversalTranslator(c slate.ServiceContainer) (*ut.UniversalTranslator, error) {
	instance, err := c.Get(ContainerUniversalTranslatorID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*ut.UniversalTranslator)
	if !ok {
		return nil, errConversion(instance, "*ut.UniversalTranslator")
	}
	return i, nil
}

// GetTranslator will try to retrieve the registered translator
// instance from the application service container.
func GetTranslator(c slate.ServiceContainer) (ut.Translator, error) {
	instance, err := c.Get(ContainerTranslatorID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(ut.Translator)
	if !ok {
		return nil, errConversion(instance, "ut.Translator")
	}
	return i, nil
}

// GetParser will try to retrieve the registered error parser
// instance from the application service container.
func GetParser(c slate.ServiceContainer) (Parser, error) {
	instance, err := c.Get(ContainerParserID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(Parser)
	if !ok {
		return nil, errConversion(instance, "validation.Parser")
	}
	return i, nil
}

// Get will try to retrieve the registered validator
// instance from the application service container.
func Get(c slate.ServiceContainer) (Validator, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(Validator)
	if !ok {
		return nil, errConversion(instance, "validation.Validator")
	}
	return i, nil
}
