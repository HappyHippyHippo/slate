package srdb

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/pkg/errors"
	"testing"
)

func Test_DialectFactory_Register(t *testing.T) {
	t.Run("missing strategy", func(t *testing.T) {
		dFactory := &DialectFactory{}

		if err := dFactory.Register(nil); err == nil {
			t.Error("didn't return an expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("store strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockDialectStrategy(ctrl)
		dFactory := &DialectFactory{}

		if err := dFactory.Register(strategy); err != nil {
			t.Errorf("return the unexpected error : %v", err)
		} else if (*dFactory)[0] != strategy {
			t.Error("didn't stored the requested strategy")
		}
	})
}

func Test_DialectFactory_Get(t *testing.T) {
	t.Run("missing config", func(t *testing.T) {
		dFactory := &DialectFactory{}

		dialect, err := dFactory.Get(nil)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("requested connection config is not a partial", func(t *testing.T) {
		cfg := &sconfig.Partial{"dialect": 123}
		dFactory := &DialectFactory{}

		dialect, err := dFactory.Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("unsupported dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &sconfig.Partial{"dialect": "unsupported"}
		strategy := NewMockDialectStrategy(ctrl)
		strategy.EXPECT().Accept("unsupported").Return(false).Times(1)
		dFactory := &DialectFactory{}
		_ = dFactory.Register(strategy)

		dialect, err := dFactory.Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, serror.ErrUnknownDatabaseDialect):
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrUnknownDatabaseDialect)
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
		dFactory := &DialectFactory{}
		_ = dFactory.Register(strategy)

		if _, err := dFactory.Get(cfg); err == nil {
			t.Error("didn't return an expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
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
		dFactory := &DialectFactory{}
		_ = dFactory.Register(strategy)

		if check, err := dFactory.Get(cfg); err != nil {
			t.Errorf("return the unexpected error (%v)", err)
		} else if check != dialect {
			t.Error("didn't returned the strategy provided dialect")
		}
	})
}
