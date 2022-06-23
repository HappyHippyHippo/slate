package grdb

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"github.com/pkg/errors"
	"testing"
)

func Test_DialectFactory_Register(t *testing.T) {
	t.Run("missing strategy", func(t *testing.T) {
		factory := &DialectFactory{}

		if err := factory.Register(nil); err == nil {
			t.Error("didn't return an expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("store strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockDialectStrategy(ctrl)
		factory := &DialectFactory{}

		if err := factory.Register(strategy); err != nil {
			t.Errorf("return the unexpected error : %v", err)
		} else if (*factory)[0] != strategy {
			t.Error("didn't stored the requested strategy")
		}
	})
}

func Test_DialectFactory_Get(t *testing.T) {
	t.Run("missing config", func(t *testing.T) {
		factory := &DialectFactory{}

		dialect, err := factory.Get(nil)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("requested connection config is not a partial", func(t *testing.T) {
		cfg := &gconfig.Partial{"dialect": 123}
		factory := &DialectFactory{}

		dialect, err := factory.Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("unsupported dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &gconfig.Partial{"dialect": "unsupported"}
		strategy := NewMockDialectStrategy(ctrl)
		strategy.EXPECT().Accept("unsupported").Return(false).Times(1)
		factory := &DialectFactory{}
		_ = factory.Register(strategy)

		dialect, err := factory.Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrUnknownDatabaseDialect):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrUnknownDatabaseDialect)
		}
	})

	t.Run("return strategy error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := &gconfig.Partial{"dialect": "unsupported"}
		strategy := NewMockDialectStrategy(ctrl)
		strategy.EXPECT().Accept("unsupported").Return(true).Times(1)
		strategy.EXPECT().Get(cfg).Return(nil, expected).Times(1)
		factory := &DialectFactory{}
		_ = factory.Register(strategy)

		if _, err := factory.Get(cfg); err == nil {
			t.Error("didn't return an expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("return strategy provided dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &gconfig.Partial{"dialect": "unsupported"}
		dialect := NewMockDialect(ctrl)
		strategy := NewMockDialectStrategy(ctrl)
		strategy.EXPECT().Accept("unsupported").Return(true).Times(1)
		strategy.EXPECT().Get(cfg).Return(dialect, nil).Times(1)
		factory := &DialectFactory{}
		_ = factory.Register(strategy)

		if check, err := factory.Get(cfg); err != nil {
			t.Errorf("return the unexpected error (%v)", err)
		} else if check != dialect {
			t.Error("didn't returned the strategy provided dialect")
		}
	})
}
