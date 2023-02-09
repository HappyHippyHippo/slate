package log

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/fs"
	"github.com/spf13/afero"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Register(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
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
		case !container.Has(JSONFormatterStrategyID):
			t.Error("didn't registered the log formatter strategy json", e)
		case !container.Has(FormatterFactoryID):
			t.Error("didn't registered the log formatter factory", e)
		case !container.Has(ConsoleStreamStrategyID):
			t.Error("didn't registered the log console stream strategy", e)
		case !container.Has(FileStreamStrategyID):
			t.Error("didn't registered the log file stream strategy", e)
		case !container.Has(RotatingFileStreamStrategyID):
			t.Error("didn't registered the log rotate file stream strategy", e)
		case !container.Has(StreamFactoryID):
			t.Error("didn't registered the log stream factory", e)
		case !container.Has(ID):
			t.Error("didn't registered the logger", e)
		case !container.Has(LoaderID):
			t.Error("didn't registered the log loader", e)
		}
	})

	t.Run("retrieving log formatter factory strategy json", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(JSONFormatterStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *JSONFormatterStrategy:
			default:
				t.Error("didn't return a JSON formatter strategy reference")
			}
		}
	})

	t.Run("retrieving log formatter factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)

		factory, e := container.Get(FormatterFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch factory.(type) {
			case *FormatterFactory:
			default:
				t.Error("didn't return a formatter factory reference")
			}
		}
	})

	t.Run("error retrieving formatter factory on retrieving the stream factory strategy console", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		_ = (fs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(FormatterFactoryID, func() (*FormatterFactory, error) { return nil, expected })

		if _, e := container.Get(ConsoleStreamStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving log stream strategy console", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(ConsoleStreamStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *ConsoleStreamStrategy:
			default:
				t.Error("didn't return a console stream strategy reference")
			}
		}
	})

	t.Run("error retrieving file system on retrieving the stream factory strategy file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)
		_ = container.Service(fs.ID, func() (afero.Fs, error) { return nil, expected })

		if _, e := container.Get(FileStreamStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving formatter factory on retrieving the stream factory strategy file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(FormatterFactoryID, func() (*FormatterFactory, error) { return nil, expected })

		if _, e := container.Get(FileStreamStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving log stream strategy file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(FileStreamStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *FileStreamStrategy:
			default:
				t.Error("didn't return a file stream strategy reference")
			}
		}
	})

	t.Run("error retrieving file system on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)
		_ = container.Service(fs.ID, func() (afero.Fs, error) { return nil, expected })

		if _, e := container.Get(RotatingFileStreamStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving formatter factory on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(FormatterFactoryID, func() (*FormatterFactory, error) { return nil, expected })

		if _, e := container.Get(RotatingFileStreamStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving log stream strategy rotate file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(RotatingFileStreamStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *RotatingFileStreamStrategy:
			default:
				t.Error("didn't return a rotating file stream strategy reference")
			}
		}
	})

	t.Run("retrieving log stream factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)

		factory, e := container.Get(StreamFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
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
			t.Errorf("returned the unexpected err (%v)", e)
		case log == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch log.(type) {
			case *Log:
			default:
				t.Error("didn't return a log reference")
			}
		}
	})

	t.Run("error retrieving config on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (Provider{}).Register(container)
		_ = container.Service(config.ID, func() (config.IManager, error) { return nil, expected })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
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
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
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
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving log loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		loader, e := container.Get(LoaderID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
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
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Boot(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("error retrieving formatter factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(FormatterFactoryID, func() (*FormatterFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid formatter factory", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(FormatterFactoryID, func() (interface{}, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("error retrieving formatter factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return nil, expected }, FormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid formatter factory strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return "string", nil }, FormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
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
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid stream factory", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(StreamFactoryID, func() (interface{}, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
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
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
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
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
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
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
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
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("invalid log entry config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().Config("slate.log.streams", config.Config{}).Return(nil, expected).Times(1)
		_ = container.Service(config.ID, func() (config.IManager, error) { return cfg, nil })

		if e := sut.Boot(container); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("correct boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().Config("slate.log.streams", config.Config{}).Return(&config.Config{}, nil).Times(1)
		cfg.EXPECT().AddObserver("slate.log.streams", gomock.Any()).Return(nil).Times(1)
		_ = container.Service(config.ID, func() (config.IManager, error) { return cfg, nil })

		if e := provider.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
