package request

import (
	"github.com/happyhippyhippo/slate/rest/logmw"
)

// Decorator defines a function used to decorate a
// request reader output.
type Decorator func(reader logmw.RequestReader, model interface{}) (logmw.RequestReader, error)
