package slate

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func assertPanic(
	t *testing.T,
	expected interface{},
) {
	if r := recover(); r == nil {
		t.Error("did not panic")
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("panic with the (%v) when expecting (%v)", r, expected)
	}
}

func Test_NewApplication(t *testing.T) {
	t.Run("instantiate a application entries list", func(t *testing.T) {
		if NewApplication().entries == nil {
			t.Error("didn't created the application entries list")
		}
	})

	t.Run("instantiate a application container", func(t *testing.T) {
		if NewApplication().container == nil {
			t.Error("didn't created the application container")
		}
	})

	t.Run("instantiate a list of providers", func(t *testing.T) {
		if NewApplication().providers == nil {
			t.Error("didn't created the list of providers")
		}
	})

	t.Run("flag the application has not booted", func(t *testing.T) {
		if NewApplication().isBoot {
			t.Error("didn't flagged the application as not booted")
		}
	})
}

func Test_Application_Provide(t *testing.T) {
	t.Run("nil provider", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if e := NewApplication().Provide(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, ErrNilPointer)
		}
	})

	t.Run("error registering provider", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		sut := NewApplication()
		provider := NewMockProvider(ctrl)
		provider.EXPECT().Register(sut).Return(expected).Times(1)

		e := sut.Provide(provider)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		case len(sut.providers) != 0:
			t.Error("stored the failing provider")
		}
	})

	t.Run("adding a valid provider", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewApplication()
		provider := NewMockProvider(ctrl)
		provider.EXPECT().Register(sut).Return(nil).Times(1)

		if e := sut.Provide(provider); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if len(sut.providers) != 1 || sut.providers[0] != provider {
			t.Error("didn't stored the added provider")
		}
	})
}

func Test_Application_Boot(t *testing.T) {
	t.Run("panic error on boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		sut := NewApplication()
		provider := NewMockProvider(ctrl)
		provider.EXPECT().Register(sut).Return(nil).Times(1)
		provider.EXPECT().Boot(sut).DoAndReturn(func(IContainer) error { panic(expected) }).Times(1)
		_ = sut.Provide(provider)

		if e := sut.Boot(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("panic something not an error on boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := "error message"
		sut := NewApplication()
		provider := NewMockProvider(ctrl)
		provider.EXPECT().Register(sut).Return(nil).Times(1)
		provider.EXPECT().Boot(sut).DoAndReturn(func(IContainer) error { panic(expected) }).Times(1)
		_ = sut.Provide(provider)

		defer assertPanic(t, expected)
		_ = sut.Boot()
	})

	t.Run("error on boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := "error message"
		sut := NewApplication()
		provider := NewMockProvider(ctrl)
		provider.EXPECT().Register(sut).Return(nil).Times(1)
		provider.EXPECT().Boot(sut).Return(fmt.Errorf("%s", expected)).Times(1)
		_ = sut.Provide(provider)

		if e := sut.Boot(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("boot all providers only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewApplication()
		provider := NewMockProvider(ctrl)
		provider.EXPECT().Register(sut).Times(1)
		provider.EXPECT().Boot(sut).Times(1)
		_ = sut.Provide(provider)
		_ = sut.Boot()
		_ = sut.Boot()

		if !sut.isBoot {
			t.Error("didn't flagged the application as booted")
		}
	})
}
