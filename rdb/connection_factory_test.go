package rdb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_NewConnectionFactory(t *testing.T) {
	t.Run("missing dialect factory", func(t *testing.T) {
		sut, e := NewConnectionFactory(nil)
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("valid creation", func(t *testing.T) {
		if sut, e := NewConnectionFactory(&DialectFactory{}); sut == nil {
			t.Error("didn't returned the expected valid connection factory instance")
		} else if e != nil {
			t.Errorf("return the unexpected error : %v", e)
		}
	})
}

func Test_ConnectionFactory_Create(t *testing.T) {
	t.Run("error instantiating dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := &config.Config{"dialect": "invalid", "host": ":memory:"}
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(cfg).Return(nil, expected).Times(1)

		sut, _ := NewConnectionFactory(dialectFactory)
		conn, e := sut.Create(cfg, &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("error instantiating connector", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := &config.Config{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(expected).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(cfg).Return(dialect, nil).Times(1)

		sut, _ := NewConnectionFactory(dialectFactory)

		conn, e := sut.Create(cfg, &gorm.Config{Logger: logger.Discard})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &config.Config{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(cfg).Return(dialect, nil).Times(1)

		sut, _ := NewConnectionFactory(dialectFactory)

		if check, e := sut.Create(cfg, &gorm.Config{Logger: logger.Discard}); check == nil {
			t.Error("didn't return the expected connection instance")
		} else if e != nil {
			t.Errorf("return the unexpected error : (%v)", e)
		}
	})
}
