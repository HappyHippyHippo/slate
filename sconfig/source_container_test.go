package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"testing"
)

func Test_NewSourceContainer(t *testing.T) {
	t.Run("nil list of configs", func(t *testing.T) {
		src, err := newSourceContainer(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error while retrieving config partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Partial("", Partial{}).Return(nil, expected).Times(1)

		src, err := newSourceContainer([]IConfig{cfg})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("valid single partial load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"id": "value"}
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Partial("", Partial{}).Return(expected, nil).Times(1)

		src, err := newSourceContainer([]IConfig{cfg})
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error : %v", err)
		case src == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(src.(*sourceContainer).partial, expected):
			t.Errorf("returned the (%v) partial when expecting (%v)", src.(*sourceContainer).partial, expected)
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

		src, err := newSourceContainer([]IConfig{cfg1, cfg2})
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error : %v", err)
		case src == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(src.(*sourceContainer).partial, expected):
			t.Errorf("returned the (%v) partial when expecting (%v)", src.(*sourceContainer).partial, expected)
		}
	})
}
