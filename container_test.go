package slate

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_Container_Close(t *testing.T) {
	t.Run("dont try to remove non-instantiated entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Times(0)

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })
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

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })
		_, _ = sut.Get(id)

		if e := sut.Close(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("remove all entries, even if instantiated", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Return(nil).Times(1)

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })
		_, _ = sut.Get(id)
		_ = sut.Close()

		if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})
}

func Test_Container_Has(t *testing.T) {
	t.Run("validate service existence", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry := NewMockCloser(ctrl)

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })

		if !sut.Has(id) {
			t.Error("didn't found the entry")
		}
	})

	t.Run("checking a non-existent service should return false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry := NewMockCloser(ctrl)

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })

		if sut.Has(id + "salt") {
			t.Error("unexpectedly found a valid entry")
		}
	})
}

func Test_Container_Service(t *testing.T) {
	t.Run("nil factory", func(t *testing.T) {
		id := "id"

		sut := NewContainer()
		if e := sut.Service(id, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrNilPointer) {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, ErrNilPointer)
		}
	})

	t.Run("non-function factory", func(t *testing.T) {
		id := "id"

		sut := NewContainer()
		if e := sut.Service(id, "string"); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrNonFunctionFactory) {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, ErrNonFunctionFactory)
		}
	})

	t.Run("factory function not returning a service", func(t *testing.T) {
		id := "id"

		sut := NewContainer()
		if e := sut.Service(id, func() {}); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrFactoryWithoutResult) {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, ErrFactoryWithoutResult)
		}
	})

	t.Run("adding a service", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry := NewMockCloser(ctrl)

		sut := NewContainer()
		if e := sut.Service(id, func() (interface{}, error) {
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

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry1 })
		_, _ = sut.Get(id)

		if e := sut.Service(id, func() interface{} { return entry2 }); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("overriding a loaded service", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry1 := NewMockCloser(ctrl)
		entry1.EXPECT().Close().Times(1)
		entry2 := NewMockCloser(ctrl)

		sut := NewContainer()
		if e := sut.Service(id, func() interface{} { return entry1 }); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry1) {
			t.Error("didn't stored the requested first entry")
		} else if e := sut.Service(id, func() interface{} { return entry2 }); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry2) {
			t.Error("didn't stored the requested second entry")
		}
	})
}

func Test_Container_Get(t *testing.T) {
	t.Run("retrieving a non-registered service/factory", func(t *testing.T) {
		id := "invalid_id"

		sut := NewContainer()
		check, e := sut.Get(id)

		switch {
		case check != nil:
			t.Error("returned an unexpected valid instance reference")
		case e == nil:
			t.Error("didn't returned the expected error instance")
		case !errors.Is(e, ErrServiceNotFound):
			t.Errorf("returned the error (%v) when was expecting (%v)", e, ErrServiceNotFound)
		}
	})

	t.Run("error creating the requested service", func(t *testing.T) {
		id := "id"

		sut := NewContainer()
		_ = sut.Service(id, func(dep io.Closer) int { return 4 })
		check, e := sut.Get(id)

		switch {
		case check != nil:
			t.Error("returned an unexpected valid instance reference")
		case e == nil:
			t.Error("didn't returned the expected error instance")
		case !errors.Is(e, ErrContainer):
			t.Errorf("returned the error (%v) when was expecting (%v)", e, ErrContainer)
		}
	})

	t.Run("retrieving a non-loaded service", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		count := 0
		entry := NewMockCloser(ctrl)

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { count++; return entry })

		if check, e := sut.Get(id); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
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

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { count++; return entry })

		runs := 2
		for runs > 0 {
			check, e := sut.Get(id)
			switch {
			case e != nil:
				t.Errorf("returned the unexpected error (%v)", e)
			case check == nil:
				t.Error("didn't returned a valid reference")
			case count != 1:
				t.Error("called the factory more than once")
			}
			runs--
		}
	})
}

func Test_Container_Tag(t *testing.T) {
	type A struct{}
	type B struct{}

	t.Run("error creating the requested tagged service", func(t *testing.T) {
		id := "id"

		sut := NewContainer()
		_ = sut.Service(id, func(dep io.Closer) int { return 4 }, "tag1")
		check, e := sut.Tag("tag1")

		switch {
		case check != nil:
			t.Error("returned an unexpected valid instance reference")
		case e == nil:
			t.Error("didn't returned the expected error instance")
		case !errors.Is(e, ErrContainer):
			t.Errorf("returned the error (%v) when was expecting (%v)", e, ErrContainer)
		}
	})

	t.Run("retrieving a non-assigned tag list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "id1"
		id2 := "id2"

		sut := NewContainer()
		_ = sut.Service(id1, func() *A { return &A{} }, "tag1")
		_ = sut.Service(id2, func() *B { return &B{} }, "tag1", "tag2")

		if list, e := sut.Tag("tag3"); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if len(list) != 0 {
			t.Errorf("returned the unexpected non-empty (%v) list", list)
		}
	})

	t.Run("retrieving a single tagged entry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "id1"
		id2 := "id2"

		sut := NewContainer()
		_ = sut.Service(id1, func() *A { return &A{} }, "tag1")
		_ = sut.Service(id2, func() *B { return &B{} }, "tag1", "tag2")

		if list, e := sut.Tag("tag2"); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if len(list) != 1 {
			t.Errorf("returned the unexpected (%v) list", list)
		}
	})

	t.Run("retrieving a tagged entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "id1"
		id2 := "id2"

		sut := NewContainer()
		_ = sut.Service(id1, func() *A { return &A{} }, "tag1")
		_ = sut.Service(id2, func() *B { return &B{} }, "tag1", "tag2")

		if list, e := sut.Tag("tag1"); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if len(list) != 2 {
			t.Errorf("returned the unexpected (%v) list", list)
		}
	})
}

func Test_Container_Remove(t *testing.T) {
	t.Run("removing a non-registered service/factory should not error", func(t *testing.T) {
		id := "id"
		sut := NewContainer()
		if e := sut.Remove(id); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("removing a non-loaded service should not error", func(t *testing.T) {
		id := "id"
		sut := NewContainer()
		_ = sut.Service(id, func() string { return "value" })

		if e := sut.Remove(id); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
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

		sut := NewContainer()
		_ = sut.Service(id, func() io.Closer { return entry })
		_, _ = sut.Get(id)

		if e := sut.Remove(id); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})
}

func Test_Container_Clear(t *testing.T) {
	t.Run("dont try to remove non-instantiated entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Times(0)

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })

		if e := sut.Clear(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
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

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })
		_, _ = sut.Get(id)

		if e := sut.Clear(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("remove all entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Return(nil).Times(1)

		sut := NewContainer()
		_ = sut.Service(id, func() interface{} { return entry })
		_, _ = sut.Get(id)

		if e := sut.Clear(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})
}
