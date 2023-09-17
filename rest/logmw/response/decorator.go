package response

import (
	"github.com/happyhippyhippo/slate/rest/logmw"
)

// Decorator defines a function used to decorate a response
// reader output.
type Decorator func(reader logmw.ResponseReader, model interface{}) (logmw.ResponseReader, error)
