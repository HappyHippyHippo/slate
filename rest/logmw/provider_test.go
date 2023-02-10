package logmw

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Register(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ID):
			t.Errorf("didn't registered the generator : %v", sut)
		case !container.Has(RequestReaderID):
			t.Errorf("didn't registered the default request reader : %v", sut)
		case !container.Has(ResponseReaderID):
			t.Errorf("didn't registered the default response reader : %v", sut)
		}
	})

	t.Run("retrieving default request reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prevJSOM := DecorateJSON
		prevXML := DecorateXML
		DecorateJSON = true
		DecorateXML = true
		defer func() {
			DecorateJSON = prevJSOM
			DecorateXML = prevXML
		}()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		sut, e := container.Get(RequestReaderID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the generator")
		default:
			switch sut.(type) {
			case RequestReader:
			default:
				t.Error("didn't returned a request reader reference")
			}
		}
	})

	t.Run("retrieving default response reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prevJSOM := DecorateJSON
		prevXML := DecorateXML
		DecorateJSON = true
		DecorateXML = true
		defer func() {
			DecorateJSON = prevJSOM
			DecorateXML = prevXML
		}()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		sut, e := container.Get(ResponseReaderID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the generator")
		default:
			switch sut.(type) {
			case ResponseReader:
			default:
				t.Error("didn't returned a response reader reference")
			}
		}
	})

	t.Run("error retrieving logger when retrieving the generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(log.ID, func() (log.ILog, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving default request reader when retrieving the generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		sut := &Provider{}
		_ = sut.Register(container)
		_ = (&log.Provider{}).Register(container)
		_ = container.Service(RequestReaderID, func() (RequestReader, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving default response reader when retrieving the generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		sut := &Provider{}
		_ = sut.Register(container)
		_ = (&log.Provider{}).Register(container)
		_ = container.Service(ResponseReaderID, func() (ResponseReader, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving the middleware generator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)

		sut, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the generator")
		default:
			switch sut.(type) {
			case MiddlewareGenerator:
			default:
				t.Error("didn't returned a middleware generator reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Boot(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("successful boot", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Provide(Provider{})

		if e := app.Boot(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
