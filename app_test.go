package slate

import (
	"errors"
	"fmt"
	"io"
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
	} else {
		e, ok := r.(error)
		if !ok {
			if !reflect.DeepEqual(r, expected) {
				t.Errorf("panic with the (%v) when expecting (%v)", r, expected)
			}
		} else if !errors.As(e, &expected) {
			t.Errorf("panic with the (%v) when expecting (%v)", r, expected)
		}
	}
}

func Test_app_err(t *testing.T) {
	t.Run("errServiceContainer", func(t *testing.T) {
		arg := fmt.Errorf("dummy argument")
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : service di error"

		t.Run("creation without context", func(t *testing.T) {
			if e := errServiceContainer(arg); !errors.Is(e, ErrServiceContainer) {
				t.Errorf("error not a instance of ErrContainer")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errServiceContainer(arg, context); !errors.Is(e, ErrServiceContainer) {
				t.Errorf("error not a instance of ErrContainer")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})

	t.Run("errNonFunctionServiceFactory", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : non-function service factory"

		t.Run("creation without context", func(t *testing.T) {
			if e := errNonFunctionServiceFactory(arg); !errors.Is(e, ErrNonFunctionServiceFactory) {
				t.Errorf("error not a instance of ErrNonFunctionFactory")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errNonFunctionServiceFactory(arg, context); !errors.Is(e, ErrNonFunctionServiceFactory) {
				t.Errorf("error not a instance of ErrNonFunctionFactory")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})

	t.Run("errServiceFactoryWithoutResult", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : service factory without result"

		t.Run("creation without context", func(t *testing.T) {
			if e := errServiceFactoryWithoutResult(arg); !errors.Is(e, ErrServiceFactoryWithoutResult) {
				t.Errorf("error not a instance of ErrFactoryWithoutResult")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errServiceFactoryWithoutResult(arg, context); !errors.Is(e, ErrServiceFactoryWithoutResult) {
				t.Errorf("error not a instance of ErrFactoryWithoutResult")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})

	t.Run("errServiceNotFound", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : service not found"

		t.Run("creation without context", func(t *testing.T) {
			if e := errServiceNotFound(arg); !errors.Is(e, ErrServiceNotFound) {
				t.Errorf("error not a instance of ErrServiceNotFound")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errServiceNotFound(arg, context); !errors.Is(e, ErrServiceNotFound) {
				t.Errorf("error not a instance of ErrServiceNotFound")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})
}

func Test_ServiceContainer(t *testing.T) {
	t.Run("NewServiceContainer", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if sut := NewServiceContainer(); sut == nil {
				t.Error("didn't returned a valid reference")
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("dont try to remove non-instantiated entries", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)
			entry.EXPECT().Close().Times(0)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })
			_ = sut.Close()

			if sut.Has(id) {
				t.Error("didn't removed the entry")
			}
		})

		t.Run("return first entry closing error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			expected := fmt.Errorf("error message")
			entry := NewMockCloser(ctrl)
			entry.EXPECT().Close().Return(expected).Times(1)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })
			_, _ = sut.Get(id)

			if e := sut.Close(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("remove all entries, even if instantiated", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)
			entry.EXPECT().Close().Return(nil).Times(1)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })
			_, _ = sut.Get(id)
			_ = sut.Close()

			if sut.Has(id) {
				t.Error("didn't removed the entry")
			}
		})
	})

	t.Run("Has", func(t *testing.T) {
		t.Run("validate service existence", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })

			if !sut.Has(id) {
				t.Error("didn't found the entry")
			}
		})

		t.Run("checking a non-existent service should return false", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })

			if sut.Has(id + "salt") {
				t.Error("unexpectedly found a valid entry")
			}
		})
	})

	t.Run("Serve", func(t *testing.T) {
		t.Run("nil factory", func(t *testing.T) {
			id := "id"

			sut := NewServiceContainer()
			if e := sut.Add(id, nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("non-function factory", func(t *testing.T) {
			id := "id"

			sut := NewServiceContainer()
			if e := sut.Add(id, "string"); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNonFunctionServiceFactory) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNonFunctionServiceFactory)
			}
		})

		t.Run("factory function not returning a service", func(t *testing.T) {
			id := "id"

			sut := NewServiceContainer()
			if e := sut.Add(id, func() {}); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceFactoryWithoutResult) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceFactoryWithoutResult)
			}
		})

		t.Run("adding a service", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)

			sut := NewServiceContainer()
			if e := sut.Add(id, func() (interface{}, error) {
				return entry, nil
			}); e != nil {
				t.Errorf("returned the (%s) error", e)
			} else if !sut.Has(id) {
				t.Error("didn't found the added entry")
			} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry) {
				t.Error("didn't stored the requested entry")
			}
		})

		t.Run("error while removing a overriding loaded service", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			expected := fmt.Errorf("error message")
			entry1 := NewMockCloser(ctrl)
			entry1.EXPECT().Close().Return(expected).Times(1)
			entry2 := NewMockCloser(ctrl)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry1 })
			_, _ = sut.Get(id)

			if e := sut.Add(id, func() interface{} { return entry2 }); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("overriding a loaded service", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry1 := NewMockCloser(ctrl)
			entry1.EXPECT().Close().Times(1)
			entry2 := NewMockCloser(ctrl)

			sut := NewServiceContainer()
			if e := sut.Add(id, func() interface{} { return entry1 }); e != nil {
				t.Errorf("returned the (%s) error", e)
			} else if !sut.Has(id) {
				t.Error("didn't found the added entry")
			} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry1) {
				t.Error("didn't stored the requested first entry")
			} else if e := sut.Add(id, func() interface{} { return entry2 }); e != nil {
				t.Errorf("returned the (%s) error", e)
			} else if !sut.Has(id) {
				t.Error("didn't found the added entry")
			} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry2) {
				t.Error("didn't stored the requested second entry")
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("retrieving a non-registered service/factory", func(t *testing.T) {
			id := "invalid_id"

			sut := NewServiceContainer()
			check, e := sut.Get(id)

			switch {
			case check != nil:
				t.Error("returned an unexpected valid instance reference")
			case e == nil:
				t.Error("didn't returned the expected error instance")
			case !errors.Is(e, ErrServiceNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceNotFound)
			}
		})

		t.Run("error creating the requested service", func(t *testing.T) {
			id := "id"

			sut := NewServiceContainer()
			_ = sut.Add(id, func(dep io.Closer) int { return 4 })
			check, e := sut.Get(id)

			switch {
			case check != nil:
				t.Error("returned an unexpected valid instance reference")
			case e == nil:
				t.Error("didn't returned the expected error instance")
			case !errors.Is(e, ErrServiceContainer):
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("retrieving a non-loaded service", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			count := 0
			entry := NewMockCloser(ctrl)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { count++; return entry })

			if check, e := sut.Get(id); e != nil {
				t.Errorf("unexpected error (%v)", e)
			} else if check == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("retrieving a loaded service", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			count := 0
			entry := NewMockCloser(ctrl)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { count++; return entry })

			runs := 2
			for runs > 0 {
				check, e := sut.Get(id)
				switch {
				case e != nil:
					t.Errorf("unexpected error (%v)", e)
				case check == nil:
					t.Error("didn't returned a valid reference")
				case count != 1:
					t.Error("called the factory more than once")
				}
				runs--
			}
		})
	})

	t.Run("Tag", func(t *testing.T) {
		type A struct{}
		type B struct{}

		t.Run("error creating the requested tagged service", func(t *testing.T) {
			id := "id"

			sut := NewServiceContainer()
			_ = sut.Add(id, func(dep io.Closer) int { return 4 }, "tag1")
			check, e := sut.Tag("tag1")

			switch {
			case check != nil:
				t.Error("returned an unexpected valid instance reference")
			case e == nil:
				t.Error("didn't returned the expected error instance")
			case !errors.Is(e, ErrServiceContainer):
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("retrieving a non-assigned tag list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "id1"
			id2 := "id2"

			sut := NewServiceContainer()
			_ = sut.Add(id1, func() *A { return &A{} }, "tag1")
			_ = sut.Add(id2, func() *B { return &B{} }, "tag1", "tag2")

			if list, e := sut.Tag("tag3"); e != nil {
				t.Errorf("unexpected error (%v)", e)
			} else if len(list) != 0 {
				t.Errorf("unexpected non-empty (%v) list", list)
			}
		})

		t.Run("retrieving a single tagged entry", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "id1"
			id2 := "id2"

			sut := NewServiceContainer()
			_ = sut.Add(id1, func() *A { return &A{} }, "tag1")
			_ = sut.Add(id2, func() *B { return &B{} }, "tag1", "tag2")

			if list, e := sut.Tag("tag2"); e != nil {
				t.Errorf("unexpected error (%v)", e)
			} else if len(list) != 1 {
				t.Errorf("unexpected (%v) list", list)
			}
		})

		t.Run("retrieving a tagged entries", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "id1"
			id2 := "id2"

			sut := NewServiceContainer()
			_ = sut.Add(id1, func() *A { return &A{} }, "tag1")
			_ = sut.Add(id2, func() *B { return &B{} }, "tag1", "tag2")

			if list, e := sut.Tag("tag1"); e != nil {
				t.Errorf("unexpected error (%v)", e)
			} else if len(list) != 2 {
				t.Errorf("unexpected (%v) list", list)
			}
		})
	})

	t.Run("Remove", func(t *testing.T) {
		t.Run("removing a non-registered service/factory should not error", func(t *testing.T) {
			id := "id"
			sut := NewServiceContainer()
			if e := sut.Remove(id); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("removing a non-loaded service should not error", func(t *testing.T) {
			id := "id"
			sut := NewServiceContainer()
			_ = sut.Add(id, func() string { return "value" })

			if e := sut.Remove(id); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if sut.Has(id) {
				t.Error("didn't removed the entry")
			}
		})

		t.Run("removing a loaded service should close the service", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)
			entry.EXPECT().Close().Times(1)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() io.Closer { return entry })
			_, _ = sut.Get(id)

			if e := sut.Remove(id); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if sut.Has(id) {
				t.Error("didn't removed the entry")
			}
		})
	})

	t.Run("Clear", func(t *testing.T) {
		t.Run("dont try to remove non-instantiated entries", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)
			entry.EXPECT().Close().Times(0)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })

			if e := sut.Clear(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if sut.Has(id) {
				t.Error("didn't removed the entry")
			}
		})

		t.Run("return the first entry closing error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			expected := fmt.Errorf("error message")
			entry := NewMockCloser(ctrl)
			entry.EXPECT().Close().Return(expected).Times(1)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })
			_, _ = sut.Get(id)

			if e := sut.Clear(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("remove all entries", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "id"
			entry := NewMockCloser(ctrl)
			entry.EXPECT().Close().Return(nil).Times(1)

			sut := NewServiceContainer()
			_ = sut.Add(id, func() interface{} { return entry })
			_, _ = sut.Get(id)

			if e := sut.Clear(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if sut.Has(id) {
				t.Error("didn't removed the entry")
			}
		})
	})
}

func Test_ServiceRegister(t *testing.T) {
	t.Run("NewServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if sut := NewServiceRegister(nil); sut == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Reg", func(t *testing.T) {
		t.Run("nil di", func(t *testing.T) {
			app := NewApp()
			if e := NewServiceRegister(app).Provide(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}

			t.Run("successful register", func(t *testing.T) {
				app := NewApp()
				if e := NewServiceRegister(app).Provide(&app.ServiceContainer); e != nil {
					t.Errorf("unexpected (%v) error", e)
				}
			})
		})
	})

	t.Run("Boot", func(t *testing.T) {
		t.Run("nil di", func(t *testing.T) {
			app := NewApp()
			if e := NewServiceRegister(app).Boot(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("successful boot", func(t *testing.T) {
			app := NewApp()
			_ = app.Provide(NewServiceRegister(app))
			if e := app.Boot(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})
}

func Test_App(t *testing.T) {
	t.Run("NewApp", func(t *testing.T) {
		t.Run("instantiate a application entries list", func(t *testing.T) {
			if NewApp().entries == nil {
				t.Error("didn't created the application entries list")
			}
		})

		t.Run("instantiate a application di", func(t *testing.T) {
			if NewApp().di == nil {
				t.Error("didn't created the application di")
			}
		})

		t.Run("instantiate a list of registers", func(t *testing.T) {
			if NewApp().providers == nil {
				t.Error("didn't created the list of registers")
			}
		})

		t.Run("flag the application has not booted", func(t *testing.T) {
			if NewApp().isBoot {
				t.Error("didn't flagged the application as not booted")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil registers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			if e := NewApp().Provide(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error registering the register", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			sut := NewApp()
			provider := NewMockServiceProvider(ctrl)
			provider.EXPECT().Provide(&sut.ServiceContainer).Return(expected).Times(1)

			e := sut.Provide(provider)
			switch {
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			case len(sut.providers) != 0:
				t.Error("stored the failing register")
			}
		})

		t.Run("adding a valid register", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut := NewApp()
			provider := NewMockServiceProvider(ctrl)
			provider.EXPECT().Provide(&sut.ServiceContainer).Return(nil).Times(1)

			if e := sut.Provide(provider); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if len(sut.providers) != 1 || sut.providers[0] != provider {
				t.Error("didn't stored the added register")
			}
		})
	})

	t.Run("Boot", func(t *testing.T) {
		t.Run("panic error on boot", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			sut := NewApp()
			provider := NewMockServiceProvider(ctrl)
			provider.EXPECT().Provide(&sut.ServiceContainer).Return(nil).Times(1)
			provider.
				EXPECT().
				Boot(&sut.ServiceContainer).
				DoAndReturn(func(*ServiceContainer) error { panic(expected) }).
				Times(1)
			_ = sut.Provide(provider)

			defer assertPanic(t, expected)
			_ = sut.Boot()
		})

		t.Run("panic something not an error on boot failure", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := "error message"
			sut := NewApp()
			provider := NewMockServiceProvider(ctrl)
			provider.EXPECT().Provide(&sut.ServiceContainer).Return(nil).Times(1)
			provider.
				EXPECT().
				Boot(&sut.ServiceContainer).
				DoAndReturn(func(*ServiceContainer) error { panic(expected) }).
				Times(1)
			_ = sut.Provide(provider)

			defer assertPanic(t, expected)
			_ = sut.Boot()
		})

		t.Run("error on boot", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := "error message"
			sut := NewApp()
			provider := NewMockServiceProvider(ctrl)
			provider.EXPECT().Provide(&sut.ServiceContainer).Return(nil).Times(1)
			provider.
				EXPECT().
				Boot(&sut.ServiceContainer).
				Return(fmt.Errorf("%s", expected)).
				Times(1)
			_ = sut.Provide(provider)

			if e := sut.Boot(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("boot all registers only once", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut := NewApp()
			provider := NewMockServiceProvider(ctrl)
			provider.EXPECT().Provide(&sut.ServiceContainer).Times(1)
			provider.EXPECT().Boot(&sut.ServiceContainer).Times(1)
			_ = sut.Provide(provider)
			_ = sut.Boot()
			_ = sut.Boot()

			if !sut.isBoot {
				t.Error("didn't flagged the application as booted")
			}
		})
	})
}
