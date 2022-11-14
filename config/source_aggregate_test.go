package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewSourceAggregate(t *testing.T) {
	t.Run("nil list of configs", func(t *testing.T) {
		sut, e := NewSourceAggregate(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("error while retrieving config partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Partial("", Partial{}).Return(nil, expected).Times(1)

		sut, e := NewSourceAggregate([]IConfig{cfg})
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("valid single partial load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"id": "value"}
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Partial("", Partial{}).Return(expected, nil).Times(1)

		sut, e := NewSourceAggregate([]IConfig{cfg})
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(sut.(*sourceAggregate).partial, expected):
			t.Errorf("returned the (%v) partial when expecting (%v)", sut.(*sourceAggregate).partial, expected)
		}
	})

	t.Run("valid multiple partials load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"id 1": "value 1", "id 2": "value 2"}
		cfg1 := NewMockConfig(ctrl)
		cfg1.EXPECT().Partial("", Partial{}).Return(Partial{"id 1": "value 1"}, nil).Times(1)
		cfg2 := NewMockConfig(ctrl)
		cfg2.EXPECT().Partial("", Partial{}).Return(Partial{"id 2": "value 2"}, nil).Times(1)

		sut, e := NewSourceAggregate([]IConfig{cfg1, cfg2})
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(sut.(*sourceAggregate).partial, expected):
			t.Errorf("returned the (%v) partial when expecting (%v)", sut.(*sourceAggregate).partial, expected)
		}
	})
}