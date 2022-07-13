package slog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/happyhippyhippo/slate/sfs"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		p := &Provider{}

		if err := p.Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		provider := &Provider{}

		err := provider.Register(container)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case !container.Has(ContainerFormatterStrategyJSONID):
			t.Error("didn't registered the log formatter strategy json", err)
		case !container.Has(ContainerFormatterFactoryID):
			t.Error("didn't registered the log formatter factory", err)
		case !container.Has(ContainerStreamStrategyConsoleID):
			t.Error("didn't registered the log console stream strategy", err)
		case !container.Has(ContainerStreamStrategyFileID):
			t.Error("didn't registered the log file stream strategy", err)
		case !container.Has(ContainerStreamStrategRotatingFileID):
			t.Error("didn't registered the log rotate file stream strategy", err)
		case !container.Has(ContainerStreamFactoryID):
			t.Error("didn't registered the log stream factory", err)
		case !container.Has(ContainerID):
			t.Error("didn't registered the logger", err)
		case !container.Has(ContainerLoaderID):
			t.Error("didn't registered the log loader", err)
		}
	})

	t.Run("retrieving log formatter factory strategy json", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		strategy, err := container.Get(ContainerFormatterStrategyJSONID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving log formatter factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		factory, err := container.Get(ContainerFormatterFactoryID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
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

		if _, err := container.Get(ContainerStreamStrategyConsoleID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid formatter factory on retrieving the stream factory strategy console", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerStreamStrategyConsoleID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving log stream strategy console", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, err := container.Get(ContainerStreamStrategyConsoleID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
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

		if _, err := container.Get(ContainerStreamStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid file system on retrieving the stream factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerStreamStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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

		if _, err := container.Get(ContainerStreamStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid formatter factory on retrieving the stream factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerStreamStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving log stream strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, err := container.Get(ContainerStreamStrategyFileID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
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

		if _, err := container.Get(ContainerStreamStrategRotatingFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid file system on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		expected := errConversion("string", "afero.Fs")
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerStreamStrategRotatingFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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

		if _, err := container.Get(ContainerStreamStrategRotatingFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid formatter factory on retrieving the stream factory strategy rotate file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerStreamStrategRotatingFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving log stream strategy rotate file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		strategy, err := container.Get(ContainerStreamStrategRotatingFileID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving log stream factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		strategy, err := container.Get(ContainerStreamFactoryID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving logger", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)

		log, err := container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case log == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving config on retrieving logger loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (Provider{}).Register(container)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid config on retrieving logger loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving log loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		load, err := container.Get(ContainerLoaderID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case load == nil:
			t.Error("didn't returned a valid reference")
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (&Provider{}).Boot(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error retrieving formatter factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid formatter factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error retrieving formatter factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerFormatterStrategyTag)

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid formatter factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerFormatterStrategyTag)

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error retrieving stream factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerStreamFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid stream factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerStreamFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error retrieving stream factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerStreamStrategyTag)

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid stream factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerStreamStrategyTag)

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("don't run loader if globally configured so", func(t *testing.T) {
		LoaderActive = false
		defer func() { LoaderActive = true }()

		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			panic(fmt.Errorf("error message"))
		})

		if err := provider.Boot(container); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("error retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return nil, expected
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return "string", nil
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("invalid log entry config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (sfs.Provider{}).Register(container)
		_ = (sconfig.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(sconfig.Partial{"log": sconfig.Partial{"streams": "string"}}, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		if err := provider.Boot(container); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(sconfig.Partial{"path": []interface{}{}}, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		if err := provider.Boot(container); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})
}

func Test_GetFormatterFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetFormatterFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non formatter factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerFormatterFactoryID, func() (any, error) {
			return "string", nil
		})

		s, err := GetFormatterFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid formatter factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetFormatterFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetFormatterStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		e := fmt.Errorf("dummy message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, e
		}, ContainerFormatterStrategyTag)

		s, err := GetFormatterStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, e):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non formatter strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerFormatterStrategyTag)

		s, err := GetFormatterStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid formatter strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetFormatterStrategies(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetStreamFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetStreamFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non stream factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerStreamFactoryID, func() (any, error) {
			return "string", nil
		})

		s, err := GetStreamFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid stream factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetStreamFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetStreamStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		e := fmt.Errorf("dummy message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, e
		}, ContainerStreamStrategyTag)

		s, err := GetStreamStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, e):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non stream strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerStreamStrategyTag)

		s, err := GetStreamStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid stream strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, err := GetStreamStrategies(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetLogger(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetLogger(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non logger instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		s, err := GetLogger(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid logger instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetLogger(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetLoader(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetLoader(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non loader instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerLoaderID, func() (any, error) {
			return "string", nil
		})

		s, err := GetLoader(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid loader instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(c)
		_ = (&sconfig.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, err := GetLoader(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
