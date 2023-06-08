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
			t.Errorf("didn't registered the decoder factory : %v", sut)
		case !container.Has(SourceFactoryID):
			t.Errorf("didn't registered the source factory : %v", sut)
		case !container.Has(ID):
			t.Errorf("didn't registered the config : %v", sut)
		case !container.Has(LoaderID):
			t.Errorf("didn't registered the loader : %v", sut)
		}
	})

	t.Run("retrieving decoder factory", func(t *testing.T) {
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

	t.Run("retrieving the source factory", func(t *testing.T) {
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
			case *Config:
			default:
				t.Error("didn't return a config reference")
			}
		}
	})

	t.Run("error retrieving config on retrieving loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ID, func() *Config { return nil })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving source factory on retrieving loader", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(SourceFactoryID, func() *SourceFactory { return nil })

		if _, e := container.Get(LoaderID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving loader", func(t *testing.T) {
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
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error retrieving decoder factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(DecoderFactoryID, func() (*DecoderFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid decoder factory", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(DecoderFactoryID, func() (string, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving decoder strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (DecoderStrategy, error) { return nil, expected }, DecoderStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving invalid decoder strategy", func(t *testing.T) {
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

	t.Run("error retrieving source factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)

		_ = container.Service(SourceFactoryID, func() (*SourceFactory, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid source factory", func(t *testing.T) {
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service(SourceFactoryID, func() (string, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving source strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (SourceStrategy, error) { return nil, expected }, SourceStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving invalid source strategy", func(t *testing.T) {
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

	t.Run("error retrieving source strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		sut := &Provider{}
		container := slate.NewContainer()
		_ = sut.Register(container)
		_ = container.Service("id", func() (SourceStrategy, error) { return nil, expected }, SourceStrategyTag)

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

		_ = container.Service("decoder.id", func() DecoderStrategy { return decoderStrategy }, DecoderStrategyTag)
		_ = container.Service("source.id", func() SourceStrategy { return sourceStrategy }, SourceStrategyTag)

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
		_ = container.Service(LoaderID, func() (*Loader, error) { return nil, expected })

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
		_ = container.Service(LoaderID, func() (*Loader, error) { return nil, expected })

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
		_ = container.Service(LoaderID, func() (string, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("request loader to init config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		container := slate.NewContainer()
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(Partial{}, nil).Times(1)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		sourceStrategy.EXPECT().Accept(partial).Return(true).Times(1)
		sourceStrategy.EXPECT().Create(partial).Return(source, nil).Times(1)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("source", func() SourceStrategy { return sourceStrategy }, SourceStrategyTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected e (%v)", e)
		}
	})
}
