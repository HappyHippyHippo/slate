package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewAggregateSource(t *testing.T) {
	t.Run("nil list of configs", func(t *testing.T) {
		sut, e := NewAggregateSource(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("error while retrieving config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Config("", Config{}).Return(nil, expected).Times(1)

		sut, e := NewAggregateSource([]IConfig{cfg})
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("valid single config load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Config{"id": "value"}
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Config("", Config{}).Return(&expected, nil).Times(1)

		sut, e := NewAggregateSource([]IConfig{cfg})
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(sut.partial, expected):
			t.Errorf("returned the (%v) config when expecting (%v)", sut.partial, expected)
		}
	})

	t.Run("valid multiple partials load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Config{"id 1": "value 1", "id 2": "value 2"}
		cfg1 := NewMockConfig(ctrl)
		cfg1.EXPECT().Config("", Config{}).Return(&Config{"id 1": "value 1"}, nil).Times(1)
		cfg2 := NewMockConfig(ctrl)
		cfg2.EXPECT().Config("", Config{}).Return(&Config{"id 2": "value 2"}, nil).Times(1)

		sut, e := NewAggregateSource([]IConfig{cfg1, cfg2})
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(sut.partial, expected):
			t.Errorf("returned the (%v) config when expecting (%v)", sut.partial, expected)
		}
	})
}
