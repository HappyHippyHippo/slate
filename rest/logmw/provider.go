package logmw

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rest"
)

const (
	// ID defines the id to be used as the container
	// registration id of a logging middleware instance factory function.
	ID = rest.ID + ".logmw"

	// RequestReaderID @todo doc
	RequestReaderID = ID + ".reader.request"

	// ResponseReaderID @todo doc
	ResponseReaderID = ID + ".reader.response"
)

// Provider defines the slate.rest.log module service provider to be used on
// the application initialization to register the logging middleware service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the log middleware package instances in the
// application container
func (Provider) Register(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// register the default request reader method
	_ = container[0].Service(RequestReaderID, func() RequestReader {
		requestReader := RequestReaderDefault
		if DecorateJSON {
			requestReader, _ = NewRequestReaderDecoratorJSON(requestReader, nil)
		}
		if DecorateXML {
			requestReader, _ = NewRequestReaderDecoratorXML(requestReader, nil)
		}
		return requestReader
	})
	// register the default response reader method
	_ = container[0].Service(ResponseReaderID, func() ResponseReader {
		responseReader := ResponseReaderDefault
		if DecorateJSON {
			responseReader, _ = NewResponseReaderDecoratorJSON(responseReader, nil)
		}
		if DecorateXML {
			responseReader, _ = NewResponseReaderDecoratorXML(responseReader, nil)
		}
		return responseReader
	})
	// register the logging middleware generator
	_ = container[0].Service(ID, NewMiddlewareGenerator)
	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	return nil
}
