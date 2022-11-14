package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	serror "github.com/happyhippyhippo/slate/error"
	sfs "github.com/happyhippyhippo/slate/fs"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		sut := &Provider{}
		_ = sut.Register(nil)

		if e := sut.Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ContainerDecoderStrategyYAMLID):
			t.Errorf("didn't registered the config decoder strategy yaml : %v", sut)
		case !container.Has(ContainerDecoderStrategyJSONID):
			t.Errorf("didn't registered the config decoder strategy json : %v", sut)
		case !container.Has(ContainerDecoderFactoryID):
			t.Errorf("didn't registered the config decoder factory : %v", sut)
		case !container.Has(ContainerSourceStrategyFileID):
			t.Errorf("didn't registered the config file source strategy : %v", sut)
		case !container.Has(ContainerSourceStrategyFileObservableID):
			t.Errorf("didn't registered the config observable file source strategy : %v", sut)
		case !container.Has(ContainerSourceStrategyDirID):
			t.Errorf("didn't registered the config dir source strategy : %v", sut)
		case !container.Has(ContainerSourceStrategyRestID):
			t.Errorf("didn't registered the config rest source strategy : %v", sut)
		case !container.Has(ContainerSourceStrategyRestObservableID):
			t.Errorf("didn't registered the config observable rest source strategy : %v", sut)
		case !container.Has(ContainerSourceStrategyEnvID):
			t.Errorf("didn't registered the config environment source strategy : %v", sut)
		case !container.Has(ContainerSourceStrategyAggregateID):
			t.Errorf("didn't registered the config container loading source strategy : %v", sut)
		case !container.Has(ContainerSourceFactoryID):
			t.Errorf("didn't registered the config source factory : %v", sut)
		case !container.Has(ContainerID):
			t.Errorf("didn't registered the config : %v", sut)
		case !container.Has(ContainerLoaderID):
			t.Errorf("didn't registered the config loader : %v", sut)
		}
	})

	t.Run("retrieving config yaml decoder factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerDecoderStrategyYAMLID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *decoderStrategyYAML:
			default:
				t.Error("didn't returned a yaml decoder factory strategy reference")
			}
		}
	})

	t.Run("retrieving config json decoder factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerDecoderStrategyJSONID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *decoderStrategyJSON:
			default:
				t.Error("didn't returned a json decoder factory strategy reference")
			}
		}
	})

	t.Run("retrieving config decoder factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		factory, e := container.Get(ContainerDecoderFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch factory.(type) {
			case *decoderFactory:
			default:
				t.Error("didn't returned a decoder factory reference")
			}
		}
	})

	t.Run("error retrieving file system on retrieving the source factory strategy file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid file system on retrieving the source factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error retrieving decoder factory on retrieving the source factory strategy file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyFileID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerSourceStrategyFileID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *sourceStrategyFile:
			default:
				t.Error("didn't returned a source file strategy reference")
			}
		}
	})

	t.Run("error retrieving file system on retrieving the source factory strategy observable file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyFileObservableID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid file system on retrieving the source factory strategy observable file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyFileObservableID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error retrieving decoder factory on retrieving the source factory strategy observable file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyFileObservableID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy observable file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyFileObservableID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy observable file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerSourceStrategyFileObservableID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *sourceStrategyObservableFile:
			default:
				t.Error("didn't returned a source observable file strategy reference")
			}
		}
	})

	t.Run("error retrieving file system on retrieving the source factory strategy dir", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyDirID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid file system on retrieving the source factory strategy dir", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyDirID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error retrieving decoder factory on retrieving the source factory strategy dir", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyDirID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy dir", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyDirID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy dir", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerSourceStrategyDirID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *sourceStrategyDir:
			default:
				t.Error("didn't returned a source dir strategy reference")
			}
		}
	})

	t.Run("error retrieving decoder factory on retrieving the source factory strategy rest", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyRestID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy rest", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyRestID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy rest", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerSourceStrategyRestID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *sourceStrategyRest:
			default:
				t.Error("didn't returned a source rest strategy reference")
			}
		}
	})

	t.Run("error retrieving decoder factory on retrieving the source factory strategy observable rest", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerSourceStrategyRestObservableID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy observable rest", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerSourceStrategyRestObservableID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy observable rest", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerSourceStrategyRestObservableID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *sourceStrategyObservableRest:
			default:
				t.Error("didn't returned a source observable rest strategy reference")
			}
		}
	})

	t.Run("retrieving the source factory strategy environment", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerSourceStrategyEnvID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *sourceStrategyEnv:
			default:
				t.Error("didn't returned a source environment strategy reference")
			}
		}
	})

	t.Run("invalid partial on retrieving the source factory strategy container", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerSourceContainerPartialTag)

		if _, e := container.Get(ContainerSourceStrategyAggregateID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy container", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ContainerSourceStrategyAggregateID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *sourceStrategyAggregate:
			default:
				t.Error("didn't returned a source container strategy reference")
			}
		}
	})

	t.Run("retrieving config source factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		factory, e := container.Get(ContainerSourceFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch factory.(type) {
			case *sourceFactory:
			default:
				t.Error("didn't returned a source factory reference")
			}
		}
	})

	t.Run("retrieving config", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		cfg, e := container.Get(ContainerID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case cfg == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch cfg.(type) {
			case *manager:
			default:
				t.Error("didn't returned a config manager reference")
			}
		}
	})

	t.Run("error retrieving config on retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid config on retrieving loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error retrieving config source factory on retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerSourceFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid config source factory on retrieving loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerSourceFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerLoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("retrieving config loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		l, e := container.Get(ContainerLoaderID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case l == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch l.(type) {
			case *loader:
			default:
				t.Error("didn't returned a config loader reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("error retrieving config decoder factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("retrieving invalid config decoder factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error retrieving config decoder strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerDecoderStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("retrieving invalid config decoder factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerDecoderStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error retrieving config source factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerSourceFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("retrieving invalid config source factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerSourceFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error retrieving config source factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerSourceStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("retrieving invalid config source factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerSourceStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("no entry source active", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		LoaderActive = false
		defer func() { LoaderActive = true }()

		expected := fmt.Errorf("error message")
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})

	t.Run("error retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
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
		_ = (&(sfs.Provider{})).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return "string", nil
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("add entry source into the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		content := "field: value"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, content)
			return len(content), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		container := slate.ServiceContainer{}
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return fileSystem, nil
		})
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})
}
