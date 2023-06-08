package log

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/fs"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(FormatterFactoryID):
			t.Error("didn't registered the logger Formatter factory", e)
		case !container.Has(StreamFactoryID):
			t.Error("didn't registered the logger stream factory", e)
		case !container.Has(ID):
			t.Error("didn't registered the logger", e)
		case !container.Has(LoaderID):
			t.Error("didn't registered the logger loader", e)
		}
	})

	t.Run("retrieving logger Formatter factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)

		factory, e := container.Get(FormatterFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch factory.(type) {
			case *FormatterFactory:
			default:
				t.Error("didn't return a Formatter factory reference")
			}
		}
	})

	t.Run("retrieving logger stream factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)

		factory, e := container.Get(StreamFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch factory.(type) {
			case *StreamFactory:
			default:
				t.Error("didn't return a stream factory reference")
			}
		}
	})

	t.Run("retrieving logger", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)

		log, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case log == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch log.(type) {
			case *Log:
			default:
				t.Error("didn't return a logger reference")
			}
		}
	})

	t.Run("error retrieving config on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)
		_ = container.Service(config.ID, func() (*config.Config, error) { return nil, expected })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving logger on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ID, func() (*Log, error) { return nil, expected })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving stream factory on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (config.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(StreamFactoryID, func() (*StreamFactory, error) { return nil, expected })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving logger loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		loader, e := container.Get(LoaderID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case loader == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch loader.(type) {
			case *Loader:
			default:
				t.Error("didn't return a loader reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error retrieving Formatter factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(FormatterFactoryID, func() (*FormatterFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid Formatter factory", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(FormatterFactoryID, func() (interface{}, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving Formatter factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return nil, expected }, FormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid Formatter factory strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return "string", nil }, FormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving stream factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(StreamFactoryID, func() (interface{}, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid stream factory", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(StreamFactoryID, func() (interface{}, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving stream factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return nil, expected }, StreamStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid stream factory strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return "string", nil }, StreamStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("valid simple boot with strategies (no loader)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderActive = false
		defer func() { LoaderActive = true }()

		container := slate.NewContainer()
		_ = fs.Provider{}.Register(container)
		sut := &Provider{}
		_ = sut.Register(container)

		formatterStrategy := NewMockFormatterStrategy(ctrl)
		streamStrategy := NewMockStreamStrategy(ctrl)

		_ = container.Service("formatter.id", func() FormatterStrategy { return formatterStrategy }, FormatterStrategyTag)
		_ = container.Service("stream.id", func() StreamStrategy { return streamStrategy }, StreamStrategyTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})

	t.Run("don't run loader if globally configured so", func(t *testing.T) {
		LoaderActive = false
		defer func() { LoaderActive = true }()

		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() (interface{}, error) {
			panic(fmt.Errorf("error message"))
		})

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() (interface{}, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() (interface{}, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("invalid logger entry config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)

		partial := config.Partial{"slate": config.Partial{"logger": config.Partial{"streams": "string"}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 1, source)
		_ = container.Service(config.ID, func() (*config.Config, error) { return cfg, nil })

		if e := sut.Boot(container); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("correct boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		partial := config.Partial{"slate": config.Partial{"logger": config.Partial{"streams": config.Partial{}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 1, source)
		_ = container.Service(config.ID, func() (*config.Config, error) { return cfg, nil })

		if e := provider.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
