package rdb

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/pkg/errors"
	"testing"
)

func Test_DialectFactory_Register(t *testing.T) {
	t.Run("missing strategy", func(t *testing.T) {
		if e := (&DialectFactory{}).Register(nil); e == nil {
			t.Error("didn't return the expected error")
		} else if !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("store strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockDialectStrategy(ctrl)
		sut := &DialectFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("return the unexpected error : %v", e)
		} else if (*sut)[0] != strategy {
			t.Error("didn't stored the requested strategy")
		}
	})
}

func Test_DialectFactory_Get(t *testing.T) {
	t.Run("missing cfg", func(t *testing.T) {
		sut := &DialectFactory{}

		dialect, e := sut.Get(nil)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("requested connection cfg is not a partial", func(t *testing.T) {
		cfg := &config.Partial{"dialect": 123}
		sut := &DialectFactory{}

		dialect, e := sut.Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrConversion)
		}
	})

	t.Run("unsupported dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &config.Partial{"dialect": "unsupported"}
		strategy := NewMockDialectStrategy(ctrl)
		strategy.EXPECT().Accept("unsupported").Return(false).Times(1)

		sut := &DialectFactory{}
		_ = sut.Register(strategy)

		dialect, e := sut.Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.ErrUnknownDatabaseDialect):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrUnknownDatabaseDialect)
		}
	})

	t.Run("return strategy error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := &config.Partial{"dialect": "unsupported"}
		strategy := NewMockDialectStrategy(ctrl)
		strategy.EXPECT().Accept("unsupported").Return(true).Times(1)
		strategy.EXPECT().Get(cfg).Return(nil, expected).Times(1)

		sut := &DialectFactory{}
		_ = sut.Register(strategy)

		if _, e := sut.Get(cfg); e == nil {
			t.Error("didn't return the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("return strategy provided dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &config.Partial{"dialect": "unsupported"}
		dialect := NewMockDialect(ctrl)
		strategy := NewMockDialectStrategy(ctrl)
		strategy.EXPECT().Accept("unsupported").Return(true).Times(1)
		strategy.EXPECT().Get(cfg).Return(dialect, nil).Times(1)

		sut := &DialectFactory{}
		_ = sut.Register(strategy)

		if check, e := sut.Get(cfg); e != nil {
			t.Errorf("return the unexpected error (%v)", e)
		} else if check != dialect {
			t.Error("didn't returned the strategy provided dialect")
		}
	})
}
