package validation

import (
	senv "github.com/happyhippyhippo/slate/env"
	srest "github.com/happyhippyhippo/slate/rest"
)

const (
	// ContainerID defines the id to be used
	// as the container registration id of a validation.
	ContainerID = srest.ContainerID + ".validation"

	// ContainerUniversalTranslatorID defines the id to be used
	// as the container registration id of a universal translator.
	ContainerUniversalTranslatorID = ContainerID + ".universal_translator"

	// ContainerTranslatorID defines the id to be used
	// as the container registration id of a translator.
	ContainerTranslatorID = ContainerID + ".translator"

	// ContainerParserID defines the id to be used
	// as the container registration id of an error parser instance.
	ContainerParserID = ContainerID + ".parser"
)

const (
	// EnvID defines the slaterest.validation package base environment variable name.
	EnvID = srest.EnvID + "_VALIDATION"
)

var (
	// Locale defines the default locale string to be used when
	// instantiating the translator.
	Locale = senv.String(EnvID+"_LOCALE", "en")
)
