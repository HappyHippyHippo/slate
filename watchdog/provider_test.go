package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !container.Has(LogFormatterFactoryID):
			t.Errorf("no log formatter creator : %v", sut)
		case !container.Has(WatchdogFactoryID):
			t.Errorf("no watchdog creator : %v", sut)
		case !container.Has(ID):
			t.Errorf("no kannel : %v", sut)
		}
	})

	t.Run("retrieving log formatter creator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&log.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		sut, e := container.Get(LogFormatterFactoryID)
		switch {
		case e != nil:
			t.Errorf("unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the log formatter creator")
		default:
			switch sut.(type) {
			case *LogFormatterFactory:
			default:
				t.Error("didn't returned the log formatter creator")
			}
		}
	})

	t.Run("retrieving watchdog creator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		sut, e := container.Get(WatchdogFactoryID)
		switch {
		case e != nil:
			t.Errorf("unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the watchdog creator")
		default:
			switch sut.(type) {
			case *Factory:
			default:
				t.Error("didn't returned the watchdog creator")
			}
		}
	})

	t.Run("retrieving kennel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		sut, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the kennel")
		default:
			switch sut.(type) {
			case *Kennel:
			default:
				t.Error("didn't returned the kennel")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error retrieving the log formatter creator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LogFormatterFactoryID, func() (*LogFormatterFactory, error) {
			return nil, fmt.Errorf("error message")
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid log formatter creator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(LogFormatterFactoryID, func() string {
			return "string"
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving a log formatter strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (LogFormatterStrategy, error) {
			return nil, fmt.Errorf("error message")
		}, LogFormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid log formatter strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() interface{} {
			return "invalid"
		}, LogFormatterStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving the kennel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ID, func() (*Kennel, error) {
			return nil, fmt.Errorf("error message")
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid kennel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ID, func() string {
			return "string"
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving a watchdog process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (Processor, error) {
			return nil, fmt.Errorf("error message")
		}, ProcessTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid watchdog process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() interface{} {
			return "invalid"
		}, ProcessTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("duplicate watchdog process service name", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		type Process1 struct{ Process }
		type Process2 struct{ Process }

		container := slate.NewContainer()
		formatter := NewMockLogFormatter(ctrl)
		formatterStrategy := NewMockLogFormatterStrategy(ctrl)
		formatterStrategy.EXPECT().Accept(&config.Partial{"type": "simple"}).Return(true)
		formatterStrategy.EXPECT().Create(&config.Partial{"type": "simple"}).Return(formatter, nil)
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("log.formatter", func() LogFormatterStrategy {
			return formatterStrategy
		}, LogFormatterStrategyTag)
		_ = container.Service("id1", func() (*Process1, error) {
			p, e := NewProcess("service", func() error { return nil })
			return &Process1{Process: *p}, e
		}, ProcessTag)
		_ = container.Service("id2", func() (*Process2, error) {
			p, e := NewProcess("service", func() error { return nil })
			return &Process2{Process: *p}, e
		}, ProcessTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrDuplicateService) {
			t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateService)
		}
	})

	t.Run("valid boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		formatter := NewMockLogFormatter(ctrl)
		formatterStrategy := NewMockLogFormatterStrategy(ctrl)
		formatterStrategy.EXPECT().Accept(&config.Partial{"type": "simple"}).Return(true)
		formatterStrategy.EXPECT().Create(&config.Partial{"type": "simple"}).Return(formatter, nil)
		_ = (&config.Provider{}).Register(container)
		_ = (&log.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("log.formatter", func() LogFormatterStrategy {
			return formatterStrategy
		}, LogFormatterStrategyTag)
		_ = container.Service("id", func() (Processor, error) {
			return NewProcess("service", func() error { return nil })
		}, ProcessTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}
