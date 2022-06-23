package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/happyhippyhippo/slate/sfs"
	"io"
	"os"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		p := &Provider{}
		_ = p.Register(nil)

		if err := p.Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}

		err := p.Register(container)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case !container.Has(ContainerDecoderStrategyYAMLID):
			t.Errorf("didn't registered the config decoder strategy yaml : %v", p)
		case !container.Has(ContainerDecoderStrategyJSONID):
			t.Errorf("didn't registered the config decoder strategy json : %v", p)
		case !container.Has(ContainerDecoderFactoryID):
			t.Errorf("didn't registered the config decoder factory : %v", p)
		case !container.Has(ContainerSourceStrategyFileID):
			t.Errorf("didn't registered the config file source strategy : %v", p)
		case !container.Has(ContainerSourceStrategyFileObservableID):
			t.Errorf("didn't registered the config observable file source strategy : %v", p)
		case !container.Has(ContainerSourceStrategyDirID):
			t.Errorf("didn't registered the config dir source strategy : %v", p)
		case !container.Has(ContainerSourceStrategyRemoteID):
			t.Errorf("didn't registered the config remote source strategy : %v", p)
		case !container.Has(ContainerSourceStrategyRemoteObservableID):
			t.Errorf("didn't registered the config observable remote source strategy : %v", p)
		case !container.Has(ContainerSourceStrategyEnvID):
			t.Errorf("didn't registered the config environment source strategy : %v", p)
		case !container.Has(ContainerSourceFactoryID):
			t.Errorf("didn't registered the config source factory : %v", p)
		case !container.Has(ContainerID):
			t.Errorf("didn't registered the config : %v", p)
		case !container.Has(ContainerLoaderID):
			t.Errorf("didn't registered the config loader : %v", p)
		}
	})

	t.Run("retrieving config yaml decoder factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerDecoderStrategyYAMLID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *DecoderStrategyYAML:
			default:
				t.Error("didn't returned a yaml decoder factory strategy reference")
			}
		}
	})

	t.Run("retrieving config json decoder factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerDecoderStrategyJSONID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *DecoderStrategyJSON:
			default:
				t.Error("didn't returned a json decoder factory strategy reference")
			}
		}
	})

	t.Run("retrieving config decoder factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		factory, err := container.Get(ContainerDecoderFactoryID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch factory.(type) {
			case *DecoderFactory:
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

		if _, err := container.Get(ContainerSourceStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid file system on retrieving the source factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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

		if _, err := container.Get(ContainerSourceStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyFileID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerSourceStrategyFileID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving file system on retrieving the source factory strategy observable file", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerSourceStrategyFileObservableID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid file system on retrieving the source factory strategy observable file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyFileObservableID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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

		if _, err := container.Get(ContainerSourceStrategyFileObservableID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy observable file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyFileObservableID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy observable file", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerSourceStrategyFileObservableID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving file system on retrieving the source factory strategy dir", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerSourceStrategyDirID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid file system on retrieving the source factory strategy dir", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sfs.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyDirID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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

		if _, err := container.Get(ContainerSourceStrategyDirID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy dir", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyDirID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy dir", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerSourceStrategyDirID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving decoder factory on retrieving the source factory strategy remote", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerSourceStrategyRemoteID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy remote", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyRemoteID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy remote", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerSourceStrategyRemoteID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving decoder factory on retrieving the source factory strategy observable remote", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerSourceStrategyRemoteObservableID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid decoder factory on retrieving the source factory strategy observable remote", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerSourceStrategyRemoteObservableID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving the source factory strategy observable remote", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerSourceStrategyRemoteObservableID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving the source factory strategy environment", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		strategy, err := container.Get(ContainerSourceStrategyEnvID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving config source factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		factory, err := container.Get(ContainerSourceFactoryID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving config", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

		cfg, err := container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case cfg == nil:
			t.Error("didn't returned a valid reference")
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

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid config on retrieving loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid config source factory on retrieving loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerSourceFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerLoaderID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving config loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		_ = (&Provider{}).Register(container)

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
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error retrieving config decoder factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("retrieving invalid config decoder factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error retrieving config decoder strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerDecoderStrategyTag)

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("retrieving invalid config decoder factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerDecoderStrategyTag)

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error retrieving config source factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerSourceFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("retrieving invalid config source factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerSourceFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error retrieving config source factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerSourceStrategyTag)

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("retrieving invalid config source factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerSourceStrategyTag)

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("no entry source active", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		LoaderActive = false
		defer func() { LoaderActive = true }()

		expected := fmt.Errorf("error message")
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return nil, expected
		})

		if err := p.Boot(container); err != nil {
			t.Errorf("returned the unexpected error (%v)", err)
		}
	})

	t.Run("error retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return nil, expected
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid loader", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&(sfs.Provider{})).Register(container)
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerLoaderID, func() (interface{}, error) {
			return "string", nil
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
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
		p := &Provider{}
		_ = p.Register(container)

		if err := p.Boot(container); err != nil {
			t.Errorf("returned the unexpected error (%v)", err)
		}
	})
}
