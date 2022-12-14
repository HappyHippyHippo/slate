package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/fs"
	"github.com/spf13/afero"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		sut := &Provider{}
		_ = sut.Register(nil)

		if e := sut.Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) err when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(YAMLDecoderStrategyID):
			t.Errorf("didn't registered the config decoder strategy yaml : %v", sut)
		case !container.Has(JSONDecoderStrategyID):
			t.Errorf("didn't registered the config decoder strategy json : %v", sut)
		case !container.Has(DecoderFactoryID):
			t.Errorf("didn't registered the config decoder factory : %v", sut)
		case !container.Has(FileSourceStrategyID):
			t.Errorf("didn't registered the config file source strategy : %v", sut)
		case !container.Has(ObservableFileSourceStrategyID):
			t.Errorf("didn't registered the config observable file source strategy : %v", sut)
		case !container.Has(DirSourceStrategyID):
			t.Errorf("didn't registered the config dir source strategy : %v", sut)
		case !container.Has(RestSourceStrategyID):
			t.Errorf("didn't registered the config rest source strategy : %v", sut)
		case !container.Has(ObservableRestSourceStrategyID):
			t.Errorf("didn't registered the config observable rest source strategy : %v", sut)
		case !container.Has(EnvSourceStrategyID):
			t.Errorf("didn't registered the config environment source strategy : %v", sut)
		case !container.Has(AggregateSourceStrategyID):
			t.Errorf("didn't registered the config container loading source strategy : %v", sut)
		case !container.Has(SourceFactoryID):
			t.Errorf("didn't registered the config source factory : %v", sut)
		case !container.Has(ID):
			t.Errorf("didn't registered the config : %v", sut)
		case !container.Has(LoaderID):
			t.Errorf("didn't registered the config loader : %v", sut)
		}
	})

	t.Run("retrieving config yaml decoder factory strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(YAMLDecoderStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *YAMLDecoderStrategy:
			default:
				t.Error("didn't return a yaml decoder factory strategy reference")
			}
		}
	})

	t.Run("retrieving config json decoder factory strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(JSONDecoderStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *JSONDecoderStrategy:
			default:
				t.Error("didn't return a json decoder factory strategy reference")
			}
		}
	})

	t.Run("retrieving config decoder factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		factory, e := container.Get(DecoderFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case factory == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch factory.(type) {
			case *DecoderFactory:
			default:
				t.Error("didn't return a decoder factory reference")
			}
		}
	})

	t.Run("retrieving the source factory strategy env", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(EnvSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *EnvSourceStrategy:
			default:
				t.Error("didn't return a source env strategy reference")
			}
		}
	})

	t.Run("error retrieving the file system when retrieving source factory strategy file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(fs.ID, func() afero.Fs { return nil })

		if _, e := container.Get(FileSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving the decoder factory when retrieving source factory strategy file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DecoderFactoryID, func() IDecoderFactory { return nil })

		if _, e := container.Get(FileSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving the source factory strategy file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(FileSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *FileSourceStrategy:
			default:
				t.Error("didn't return a source file strategy reference")
			}
		}
	})

	t.Run("error retrieving the file system when retrieving source factory strategy observable file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(fs.ID, func() afero.Fs { return nil })

		if _, e := container.Get(ObservableFileSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving the decoder factory when retrieving source factory strategy observable file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DecoderFactoryID, func() IDecoderFactory { return nil })

		if _, e := container.Get(ObservableFileSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving the source factory strategy observable file", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ObservableFileSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *ObservableFileSourceStrategy:
			default:
				t.Error("didn't return a source observable file strategy reference")
			}
		}
	})

	t.Run("error retrieving the file system when retrieving source factory strategy dir", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(fs.ID, func() afero.Fs { return nil })

		if _, e := container.Get(DirSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving the decoder factory when retrieving source factory strategy dir", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DecoderFactoryID, func() IDecoderFactory { return nil })

		if _, e := container.Get(DirSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving the source factory strategy dir", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(DirSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *DirSourceStrategy:
			default:
				t.Error("didn't return a source dir strategy reference")
			}
		}
	})

	t.Run("error retrieving the decoder factory when retrieving source factory strategy rest", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DecoderFactoryID, func() IDecoderFactory { return nil })

		if _, e := container.Get(RestSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving the source factory strategy rest", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(RestSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *RestSourceStrategy:
			default:
				t.Error("didn't return a source rest strategy reference")
			}
		}
	})

	t.Run("error retrieving the decoder factory when retrieving source factory strategy observable rest", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DecoderFactoryID, func() IDecoderFactory { return nil })

		if _, e := container.Get(ObservableRestSourceStrategyID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving the source factory strategy observable rest", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ObservableRestSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *ObservableRestSourceStrategy:
			default:
				t.Error("didn't return a source observable rest strategy reference")
			}
		}
	})

	t.Run("retrieving the source factory strategy aggregate", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(AggregateSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *AggregateSourceStrategy:
			default:
				t.Error("didn't return a source aggregate strategy reference")
			}
		}
	})

	t.Run("invalid config on retrieving the source factory strategy aggregate", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service("id", func() interface{} { return "string" }, AggregateSourceTag)

		strategy, e := container.Get(AggregateSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't return a valid strategy reference")
		case len(strategy.(*AggregateSourceStrategy).partials) != 0:
			t.Error("stored an unexpected instance of a config")
		}
	})

	t.Run("valid config on retrieving the source factory strategy aggregate", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service("id", func() IConfig { return &Config{} }, AggregateSourceTag)

		strategy, e := container.Get(AggregateSourceStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
		case strategy == nil:
			t.Error("didn't return a valid strategy reference")
		case len(strategy.(*AggregateSourceStrategy).partials) != 1:
			t.Error("didn't stored the expected config")
		}
	})

	t.Run("retrieving the config source factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(SourceFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *SourceFactory:
			default:
				t.Error("didn't return a source factory reference")
			}
		}
	})

	t.Run("retrieving config", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		cfg, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case cfg == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch cfg.(type) {
			case *Manager:
			default:
				t.Error("didn't return a config manager reference")
			}
		}
	})

	t.Run("error retrieving config on retrieving loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ID, func() IManager { return nil })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving config on retrieving loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(SourceFactoryID, func() ISourceFactory { return nil })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving config loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		l, e := container.Get(LoaderID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e (%v)", e)
		case l == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch l.(type) {
			case *Loader:
			default:
				t.Error("didn't return a config loader reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) err when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("error retrieving config decoder factory", func(t *testing.T) {
		expected := fmt.Errorf("err message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(DecoderFactoryID, func() (IDecoderFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid config decoder factory", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(DecoderFactoryID, func() IDecoderFactory { return nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("error retrieving config decoder strategy", func(t *testing.T) {
		expected := fmt.Errorf("err message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (IDecoderStrategy, error) { return nil, expected }, DecoderStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving config source factory", func(t *testing.T) {
		expected := fmt.Errorf("err message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)

		_ = container.Service(SourceFactoryID, func() (ISourceFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid config source factory", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(SourceFactoryID, func() ISourceFactory { return nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("error retrieving config source strategy", func(t *testing.T) {
		expected := fmt.Errorf("err message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (ISourceStrategy, error) { return nil, expected }, SourceStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("no entry source active", func(t *testing.T) {
		LoaderActive = false
		defer func() { LoaderActive = true }()

		expected := fmt.Errorf("err message")
		container := slate.NewContainer()
		_ = fs.Provider{}.Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() (ILoader, error) { return nil, expected })

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})

	t.Run("error retrieving loader", func(t *testing.T) {
		expected := fmt.Errorf("err message")
		container := slate.NewContainer()
		_ = fs.Provider{}.Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() (ILoader, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = fs.Provider{}.Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() ILoader { return nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
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
		container := slate.NewContainer()
		_ = container.Service(fs.ID, func() afero.Fs { return fileSystem })
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})
}
