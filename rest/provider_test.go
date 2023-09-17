package rest

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/fs"
	"github.com/happyhippyhippo/slate/log"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(EngineID):
			t.Errorf("no REST engine instance : %v", sut)
		case !container.Has(ProcessID):
			t.Errorf("no watchdog process instance : %v", sut)
		}
	})

	t.Run("retrieving REST engine", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		sut, e := container.Get(EngineID)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a reference to the REST engine")
		default:
			switch sut.(type) {
			case Engine:
			default:
				t.Error("didn't returned the REST engine")
			}
		}
	})

	t.Run("error retrieving config when retrieving the watchdog process", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(config.ID, func() (*config.Config, error) {
			return nil, expected
		})

		if _, e := container.Get(ProcessID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving logger when retrieving the watchdog process", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(log.ID, func() (*log.Log, error) {
			return nil, expected
		})

		if _, e := container.Get(ProcessID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving engine when retrieving the watchdog process", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(EngineID, func() (Engine, error) {
			return nil, expected
		})

		if _, e := container.Get(ProcessID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = (Provider{}).Register(container)

		sut, e := container.Get(ProcessID)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a reference to the watchdog process")
		default:
			switch sut.(type) {
			case *Process:
			default:
				t.Error("didn't returned the watchdog process")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error retrieving engine", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)
		_ = container.Service(EngineID, func() (Engine, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid engine reference", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)
		_ = container.Service(EngineID, func() string {
			return "string"
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving register", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)
		_ = container.Service("id", func() (EndpointRegister, error) {
			return nil, expected
		}, EndpointRegisterTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("successful boot", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("error when register", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)

		expected := fmt.Errorf("error message")
		engine := sut.getEngine(container)
		register := NewMockEndpointRegister(ctrl)
		register.EXPECT().Reg(engine).Return(expected).Times(1)
		_ = container.Service("id1", func() EndpointRegister {
			return register
		}, EndpointRegisterTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("successful boot with single register", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)

		engine := sut.getEngine(container)
		register := NewMockEndpointRegister(ctrl)
		register.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id1", func() EndpointRegister {
			return register
		}, EndpointRegisterTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("successful boot with multiple register", func(t *testing.T) {
		type Register1 struct{ *MockEndpointRegister }
		type Register2 struct{ *MockEndpointRegister }
		type Register3 struct{ *MockEndpointRegister }

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)

		engine := sut.getEngine(container)
		register1 := NewMockEndpointRegister(ctrl)
		register1.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id1", func() *Register1 {
			return &Register1{MockEndpointRegister: register1}
		}, EndpointRegisterTag)
		register2 := NewMockEndpointRegister(ctrl)
		register2.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id2", func() *Register2 {
			return &Register2{MockEndpointRegister: register2}
		}, EndpointRegisterTag)
		register3 := NewMockEndpointRegister(ctrl)
		register3.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id3", func() *Register3 {
			return &Register3{MockEndpointRegister: register3}
		}, EndpointRegisterTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}
