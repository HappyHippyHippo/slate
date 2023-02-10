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
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Register(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
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
		case !container.Has(EngineID):
			t.Errorf("didn't registered the REST engine instance : %v", sut)
		case !container.Has(ProcessID):
			t.Errorf("didn't registered the watchdog process instance : %v", sut)
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
			t.Errorf("returned the unexpected error (%v)", e)
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
		_ = container.Service(config.ID, func() (config.IManager, error) { return nil, expected })

		if _, e := container.Get(ProcessID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving logger when retrieving the watchdog process", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(log.ID, func() (log.ILog, error) { return nil, expected })

		if _, e := container.Get(ProcessID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving engine when retrieving the watchdog process", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = (Provider{}).Register(container)
		_ = container.Service(EngineID, func() (Engine, error) { return nil, expected })

		if _, e := container.Get(ProcessID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving REST engine", func(t *testing.T) {
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
			t.Errorf("returned the unexpected error (%v)", e)
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
	t.Run("no argument", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
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
		_ = container.Service(EngineID, func() (Engine, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid engine reference", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)
		_ = container.Service(EngineID, func() string { return "string" })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
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
		_ = container.Service("id", func() (IEndpointRegister, error) { return nil, expected }, EndpointRegisterTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
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
			t.Errorf("returned the unexpected error : %v", e)
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
		engine, _ := sut.getEngine(container)
		register := NewMockRegister(ctrl)
		register.EXPECT().Reg(engine).Return(expected).Times(1)
		_ = container.Service("id1", func() (IEndpointRegister, error) { return register, nil }, EndpointRegisterTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
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

		engine, _ := sut.getEngine(container)
		register := NewMockRegister(ctrl)
		register.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id1", func() (IEndpointRegister, error) { return register, nil }, EndpointRegisterTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("successful boot with multiple register", func(t *testing.T) {
		type Register1 struct{ *MockRegister }
		type Register2 struct{ *MockRegister }
		type Register3 struct{ *MockRegister }

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}

		_ = (fs.Provider{}).Register(container)
		_ = (config.Provider{}).Register(container)
		_ = (log.Provider{}).Register(container)
		_ = sut.Register(container)

		engine, _ := sut.getEngine(container)
		register1 := NewMockRegister(ctrl)
		register1.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id1", func() (*Register1, error) { return &Register1{MockRegister: register1}, nil }, EndpointRegisterTag)
		register2 := NewMockRegister(ctrl)
		register2.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id2", func() (*Register2, error) { return &Register2{MockRegister: register2}, nil }, EndpointRegisterTag)
		register3 := NewMockRegister(ctrl)
		register3.EXPECT().Reg(engine).Return(nil).Times(1)
		_ = container.Service("id3", func() (*Register3, error) { return &Register3{MockRegister: register3}, nil }, EndpointRegisterTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})
}
