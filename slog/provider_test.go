package slog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serr"
	"github.com/happyhippyhippo/slate/sfs"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		sut := &Provider{}

		if e := sut.Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ContainerFormatterStrategyJSONID):
			t.Error("didn't registered the slog formatter strategy json", e)
		case !container.Has(ContainerFormatterFactoryID):
			t.Error("didn't registered the slog formatter factory", e)
		case !container.Has(ContainerStreamStrategyConsoleID):
			t.Error("didn't registered the slog console stream strategy", e)
		case !container.Has(ContainerStreamStrategyFileID):
			t.Error("didn't registered the slog file stream strategy", e)
		case !container.Has(ContainerStreamStrategyRotatingFileID):
			t.Error("didn't registered the slog rotate file stream strategy", e)
		case !container.Has(ContainerStreamFactoryID):
			t.Error("didn't registered the slog stream factory", e)
		case !container.Has(ContainerID):
			t.Error("didn't registered the logger", e)
		case !container.Has(ContainerLoaderID):
			t.Error("didn't registered the slog loader", e)
		}
	})

	t.Run("retrieving slog formatter factory strategy json", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(ContainerFormatterStrategyJSONID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving slog formatter factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		factory, e := container.Get(ContainerFormatterFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving formatter factory on retrieving the stream factory strategy console", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}

		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerStreamStrategyConsoleID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid formatter factory on retrieving the stream factory strategy console", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerStreamStrategyConsoleID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("retrieving slog stream strategy console", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(ContainerStreamStrategyConsoleID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving file system on retrieving the stream factory strategy file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerStreamStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid file system on retrieving the stream factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerStreamStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error retrieving formatter factory on retrieving the stream factory strategy file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerStreamStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid formatter factory on retrieving the stream factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerStreamStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("retrieving slog stream strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(ContainerStreamStrategyFileID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving file system on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerStreamStrategyRotatingFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid file system on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		expected := errConversion("string", "afero.Fs")
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerStreamStrategyRotatingFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error retrieving formatter factory on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerStreamStrategyRotatingFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid formatter factory on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerStreamStrategyRotatingFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("retrieving slog stream strategy rotate file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(ContainerStreamStrategyRotatingFileID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving slog stream factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		strategy, e := container.Get(ContainerStreamFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving logger", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		log, e := container.Get(ContainerID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case log == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving sconfig on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid sconfig on retrieving logger loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error retrieving logger on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid logger on retrieving logger loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error retrieving stream factory on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerStreamFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid source factory on retrieving logger loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerStreamFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("retrieving slog loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		load, e := container.Get(ContainerLoaderID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case load == nil:
			t.Error("didn't returned a valid reference")
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("error retrieving formatter factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("invalid formatter factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error retrieving formatter factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerFormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("invalid formatter factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerFormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error retrieving stream factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerStreamFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("invalid stream factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerStreamFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error retrieving stream factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerStreamStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("invalid stream factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerStreamStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("don't run loader if globally configured so", func(t *testing.T) {
		LoaderActive = false
		defer func() { LoaderActive = true }()

		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			panic(fmt.Errorf("error message"))
		})

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("invalid loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return "string", nil
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("invalid slog entry sconfig", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().List("slog.streams", []interface{}{}).Return(nil, expected).Times(1)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		if e := sut.Boot(container); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("correct boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().List("slog.streams", []interface{}{}).Return([]interface{}{}, nil).Times(1)
		cfg.EXPECT().AddObserver("slog.streams", gomock.Any()).Return(nil).Times(1)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		if e := provider.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
