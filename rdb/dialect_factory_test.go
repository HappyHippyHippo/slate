package rdb

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	sconfig "github.com/happyhippyhippo/slate/config"
	serror "github.com/happyhippyhippo/slate/error"
	"github.com/pkg/errors"
)

func Test_DialectFactory_Register(t *testing.T) {
	t.Run("missing strategy", func(t *testing.T) {
		if e := (&DialectFactory{}).Register(nil); e == nil {
			t.Error("didn't return the expected error")
		} else if !errors.Is(e, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
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
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("requested connection cfg is not a partial", func(t *testing.T) {
		cfg := &sconfig.Partial{"dialect": 123}
		sut := &DialectFactory{}

		dialect, e := sut.Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("unsupported dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &sconfig.Partial{"dialect": "unsupported"}
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
		case !errors.Is(e, serror.ErrUnknownDatabaseDialect):
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrUnknownDatabaseDialect)
		}
	})

	t.Run("return strategy error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := &sconfig.Partial{"dialect": "unsupported"}
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

		cfg := &sconfig.Partial{"dialect": "unsupported"}
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
