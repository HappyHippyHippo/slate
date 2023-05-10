package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/fs"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		sut := &Provider{}
		_ = sut.Register(nil)

		if e := sut.Register(nil); e == nil {
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
		case !container.Has(DecoderFactoryID):
			t.Errorf("didn't registered the config decoder factory : %v", sut)
		case !container.Has(SourceFactoryID):
			t.Errorf("didn't registered the config source factory : %v", sut)
		case !container.Has(ID):
			t.Errorf("didn't registered the config : %v", sut)
		case !container.Has(LoaderID):
			t.Errorf("didn't registered the config loader : %v", sut)
		}
	})

	t.Run("retrieving config decoder factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		factory, e := container.Get(DecoderFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
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
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving config on retrieving loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(SourceFactoryID, func() ISourceFactory { return nil })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
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
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error retrieving config decoder factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(DecoderFactoryID, func() (IDecoderFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid config decoder factory", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(DecoderFactoryID, func() IDecoderFactory { return nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving config decoder strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (IDecoderStrategy, error) { return nil, expected }, DecoderStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving invalid config decoder strategy", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() string { return "invalid strategy" }, DecoderStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving config source factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)

		_ = container.Service(SourceFactoryID, func() (ISourceFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid config source factory", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(SourceFactoryID, func() ISourceFactory { return nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving config source strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (ISourceStrategy, error) { return nil, expected }, SourceStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving invalid config source strategy", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() string { return "invalid strategy" }, SourceStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving config source strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (ISourceStrategy, error) { return nil, expected }, SourceStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
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

		decoderStrategy := NewMockDecoderStrategy(ctrl)
		sourceStrategy := NewMockSourceStrategy(ctrl)

		_ = container.Service("decoder.id", func() IDecoderStrategy { return decoderStrategy }, DecoderStrategyTag)
		_ = container.Service("source.id", func() ISourceStrategy { return sourceStrategy }, SourceStrategyTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})

	t.Run("no entry source active", func(t *testing.T) {
		LoaderActive = false
		defer func() { LoaderActive = true }()

		expected := fmt.Errorf("error message")
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
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = fs.Provider{}.Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() (ILoader, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
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
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("request loader to init config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		loader := NewMockLoader(ctrl)
		loader.EXPECT().Load().Times(1)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LoaderID, func() ILoader { return loader })

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})
}
