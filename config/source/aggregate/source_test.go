package aggregate

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewSource(t *testing.T) {
	t.Run("nil list of configs", func(t *testing.T) {
		sut, e := NewSource(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error while retrieving config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Config("", config.Config{}).Return(nil, expected).Times(1)

		sut, e := NewSource([]config.IConfig{cfg})
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("valid single config load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := config.Config{"id": "value"}
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Config("", config.Config{}).Return(&expected, nil).Times(1)

		sut, e := NewSource([]config.IConfig{cfg})
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(sut.Config, expected):
			t.Errorf("returned the (%v) config when expecting (%v)", sut.Config, expected)
		}
	})

	t.Run("valid multiple partials load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := config.Config{"id 1": "value 1", "id 2": "value 2"}
		cfg1 := NewMockConfig(ctrl)
		cfg1.EXPECT().Config("", config.Config{}).Return(&config.Config{"id 1": "value 1"}, nil).Times(1)
		cfg2 := NewMockConfig(ctrl)
		cfg2.EXPECT().Config("", config.Config{}).Return(&config.Config{"id 2": "value 2"}, nil).Times(1)

		sut, e := NewSource([]config.IConfig{cfg1, cfg2})
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case !reflect.DeepEqual(sut.Config, expected):
			t.Errorf("returned the (%v) config when expecting (%v)", sut.Config, expected)
		}
	})
}
