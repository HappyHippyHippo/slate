package srvalidation

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

// GetParser will try to retrieve the registered error perser
// instance from the application service container.
func GetParser(c slate.ServiceContainer) (Parser, error) {
	instance, err := c.Get(ContainerParserID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(Parser)
	if !ok {
		return nil, errConversion(instance, "Parser")
	}
	return i, nil
}

// GetValidator will try to retrieve the registered validator
// instance from the application service container.
func GetValidator(c slate.ServiceContainer) (Validator, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(Validator)
	if !ok {
		return nil, errConversion(instance, "Validator")
	}
	return i, nil
}
