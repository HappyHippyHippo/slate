package srlog

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/slog"
	"net/http"
)

// Provider defines the gapp-rest-log module service provider to be used on
// the application initialization to register the logging middleware service.
type Provider struct{}

var _ slate.ServiceProvider = &Provider{}

// Register will register the log middleware package instances in the
// application container
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	createReaders := func(statusCode int) (request RequestReader, response ResponseReader) {
		request = RequestReaderDefault
		response = NewResponseReaderDefault(statusCode)

		if DecorateJSON {
			request, _ = NewRequestReaderDecoratorJSON(request, nil)
			response, _ = NewResponseReaderDecoratorJSON(response, nil)
		}

		if DecorateXML {
			request, _ = NewRequestReaderDecoratorXML(request, nil)
			response, _ = NewResponseReaderDecoratorXML(response, nil)
		}

		return request, response
	}

	_ = c.Factory(ContainerOkID, func() (interface{}, error) {
		logger, err := slog.GetLogger(c)
		if err != nil {
			return nil, err
		}

		request, response := createReaders(http.StatusOK)

		return NewMiddleware(logger, request, response)
	})

	_ = c.Factory(ContainerCreatedID, func() (interface{}, error) {
		logger, err := slog.GetLogger(c)
		if err != nil {
			return nil, err
		}

		request, response := createReaders(http.StatusCreated)

		return NewMiddleware(logger, request, response)
	})

	_ = c.Factory(ContainerNoContentID, func() (interface{}, error) {
		logger, err := slog.GetLogger(c)
		if err != nil {
			return nil, err
		}

		request, response := createReaders(http.StatusNoContent)

		return NewMiddleware(logger, request, response)
	})

	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	return nil
}
