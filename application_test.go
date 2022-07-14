package slate

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"testing"
)

func assertPanic(t *testing.T, expected interface{}) {
	if r := recover(); r == nil {
		t.Error("did not panic")
	} else if !reflect.DeepEqual(expected, r) {
		t.Errorf("panic with the (%v) when expecting (%v)", r, expected)
	}
}

func Test_NewApplication(t *testing.T) {
	t.Run("instantiate a application container", func(t *testing.T) {
		if NewApplication().Container == nil {
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

func Test_Application_Add(t *testing.T) {
	t.Run("nil provider", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if err := NewApplication().Add(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error registering provider", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error")
		sut := NewApplication()
		provider := NewMockServiceProvider(ctrl)
		provider.EXPECT().Register(sut.Container).Return(expected).Times(1)

		err := sut.Add(provider)
		switch {
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		case len(sut.providers) != 0:
			t.Error("stored the failing provider")
		}
	})

	t.Run("adding a valid provider", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewApplication()
		provider := NewMockServiceProvider(ctrl)
		provider.EXPECT().Register(sut.Container).Return(nil).Times(1)

		if err := sut.Add(provider); err != nil {
			t.Errorf("returned the (%v) error", err)
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
		provider := NewMockServiceProvider(ctrl)
		provider.EXPECT().Register(sut.Container).Return(nil).Times(1)
		provider.EXPECT().Boot(sut.Container).DoAndReturn(func(ServiceContainer) error {
			panic(expected)
		}).Times(1)
		_ = sut.Add(provider)

		if err := sut.Boot(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("panic something not an error on boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := "error message"
		sut := NewApplication()
		provider := NewMockServiceProvider(ctrl)
		provider.EXPECT().Register(sut.Container).Return(nil).Times(1)
		provider.EXPECT().Boot(sut.Container).DoAndReturn(func(ServiceContainer) error {
			panic(expected)
		}).Times(1)
		_ = sut.Add(provider)

		defer assertPanic(t, expected)
		_ = sut.Boot()
	})

	t.Run("error on boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := "error"
		sut := NewApplication()
		provider := NewMockServiceProvider(ctrl)
		provider.EXPECT().Register(sut.Container).Return(nil).Times(1)
		provider.EXPECT().Boot(sut.Container).Return(fmt.Errorf("%s", expected)).Times(1)
		_ = sut.Add(provider)

		if err := sut.Boot(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected {
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("boot all providers only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewApplication()
		provider := NewMockServiceProvider(ctrl)
		provider.EXPECT().Register(sut.Container).Times(1)
		provider.EXPECT().Boot(sut.Container).Times(1)
		_ = sut.Add(provider)
		_ = sut.Boot()
		_ = sut.Boot()

		if !sut.isBoot {
			t.Error("didn't flagged the application as booted")
		}
	})
}
