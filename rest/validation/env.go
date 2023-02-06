package validation

import (
	"github.com/happyhippyhippo/slate/env"
	"github.com/happyhippyhippo/slate/rest"
)

const (
	// EnvID defines the slate.rest.validation package base environment variable name.
	EnvID = rest.EnvID + "_VALIDATION"
)

var (
	// Locale defines the default locale string to be used when
	// instantiating the translator.
	Locale = env.String(EnvID+"_LOCALE", "en")
)
