package slate

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"reflect"
	"testing"
)

func Test_ServiceContainer_Close(t *testing.T) {
	t.Run("dont try to remove non-instantiated entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Times(0)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})
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
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Return(expected).Times(1)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})
		_, _ = sut.Get(id)

		if e := sut.Close(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("remove all entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Return(nil).Times(1)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})
		_, _ = sut.Get(id)
		_ = sut.Close()

		if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})
}

func Test_ServiceContainer_Has(t *testing.T) {
	t.Run("validate service existence", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})

		if !sut.Has(id) {
			t.Error("didn't found the entry")
		}
	})

	t.Run("checking a non-existent service should return false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})

		if sut.Has(id + "salt") {
			t.Error("unexpectedly found a valid entry")
		}
	})

	t.Run("validate factory existence with a true value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Factory(id, func() (interface{}, error) {
			return entry, nil
		})

		if !sut.Has(id) {
			t.Error("didn't found the entry")
		}
	})

	t.Run("checking a non-existent factory should return false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Factory(id, func() (interface{}, error) {
			return entry, nil
		})

		if sut.Has(id + "salt") {
			t.Error("unexpectedly found a valid entry")
		}
	})
}

func Test_ServiceContainer_Remove(t *testing.T) {
	t.Run("removing a non-registered service/factory should not error", func(t *testing.T) {
		id := "id"
		if e := (ServiceContainer{}).Remove(id); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("removing a non-loaded service should not error", func(t *testing.T) {
		id := "id"
		sut := ServiceContainer{}
		_ = sut.Service(id, func() (interface{}, error) {
			return "value", nil
		})

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
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Times(1)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})
		_, _ = sut.Get(id)

		if e := sut.Remove(id); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})

	t.Run("removing a non-loaded factory should remove the entry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Times(0)
		_ = sut.Factory(id, func() (interface{}, error) {
			return entry, nil
		})

		if e := sut.Remove(id); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})

	t.Run("removing a loaded factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Times(0)
		_ = sut.Factory(id, func() (interface{}, error) {
			return entry, nil
		})
		_, _ = sut.Get(id)
		_ = sut.Remove(id)

		if e := sut.Remove(id); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})
}

func Test_ServiceContainer_Clear(t *testing.T) {
	t.Run("dont try to remove non-instantiated entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Times(0)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})

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
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Return(expected).Times(1)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})
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
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		entry.EXPECT().Close().Return(nil).Times(1)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry, nil
		})
		_, _ = sut.Get(id)

		if e := sut.Clear(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if sut.Has(id) {
			t.Error("didn't removed the entry")
		}
	})
}

func Test_ServiceContainer_Service(t *testing.T) {
	t.Run("nil factory", func(t *testing.T) {
		id := "id"

		if e := (ServiceContainer{}).Service(id, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrNilPointer) {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("adding a service", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)

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
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry1.EXPECT().Close().Return(expected).Times(1)
		entry2 := NewMockCloser(ctrl)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry1, nil
		})
		_, _ = sut.Get(id)

		if e := sut.Service(id, func() (interface{}, error) {
			return entry2, nil
		}); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("overriding a loaded service", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry1.EXPECT().Close().Times(1)
		entry2 := NewMockCloser(ctrl)

		if e := sut.Service(id, func() (interface{}, error) {
			return entry1, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry1) {
			t.Error("didn't stored the requested first entry")
		} else if e := sut.Service(id, func() (interface{}, error) {
			return entry2, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry2) {
			t.Error("didn't stored the requested second entry")
		}
	})

	t.Run("overriding a factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry2 := NewMockCloser(ctrl)

		if e := sut.Factory(id, func() (interface{}, error) {
			return entry1, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry1) {
			t.Error("didn't stored the requested first entry")
		} else if e := sut.Service(id, func() (interface{}, error) {
			return entry2, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry2) {
			t.Error("didn't stored the requested second entry")
		}
	})
}

func Test_ServiceContainer_Factory(t *testing.T) {
	t.Run("nil factory", func(t *testing.T) {
		id := "id"

		if e := (ServiceContainer{}).Factory(id, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrNilPointer) {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("adding a factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)

		if e := sut.Factory(id, func() (interface{}, error) {
			return entry, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry) {
			t.Error("didn't stored the requested entry")
		}
	})

	t.Run("error while removing a overriding loaded factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		expected := fmt.Errorf("error message")
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry1.EXPECT().Close().Return(expected).Times(1)
		entry2 := NewMockCloser(ctrl)
		_ = sut.Service(id, func() (interface{}, error) {
			return entry1, nil
		})
		_, _ = sut.Get(id)

		if e := sut.Factory(id, func() (interface{}, error) {
			return entry2, nil
		}); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("overriding a loaded factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry2 := NewMockCloser(ctrl)

		if e := sut.Factory(id, func() (interface{}, error) {
			return entry1, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry1) {
			t.Error("didn't stored the requested first entry")
		} else if e := sut.Factory(id, func() (interface{}, error) {
			return entry1, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry2) {
			t.Error("didn't stored the requested second entry")
		}
	})

	t.Run("overriding a loaded service", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry1.EXPECT().Close().Times(1)
		entry2 := NewMockCloser(ctrl)

		if e := sut.Service(id, func() (interface{}, error) {
			return entry1, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry1) {
			t.Error("didn't stored the requested first entry")
		} else if e := sut.Factory(id, func() (interface{}, error) {
			return entry1, nil
		}); e != nil {
			t.Errorf("returned the (%s) error", e)
		} else if !sut.Has(id) {
			t.Error("didn't found the added entry")
		} else if check, _ := sut.Get(id); !reflect.DeepEqual(check, entry2) {
			t.Error("didn't stored the requested second entry")
		}
	})
}

func Test_ServiceContainer_Get(t *testing.T) {
	t.Run("retrieving a non-registered service/factory", func(t *testing.T) {
		id := "invalid_id"

		check, e := (ServiceContainer{}).Get(id)
		switch {
		case check != nil:
			t.Error("returned an unexpected valid instance reference")
		case e == nil:
			t.Error("didn't returned the expected error instance")
		case !errors.Is(e, serr.ErrServiceNotFound):
			t.Errorf("returned the error (%v) when was expecting (%v)", e, serr.ErrServiceNotFound)
		}
	})

	t.Run("error while calling the service", func(t *testing.T) {
		id := "id"
		expected := fmt.Errorf("error message")
		sut := ServiceContainer{}
		_ = sut.Service(id, func() (interface{}, error) {
			return nil, expected
		})

		check, e := sut.Get(id)
		switch {
		case check != nil:
			t.Error("returned an unexpected valid instance reference")
		case e == nil:
			t.Error("didn't returned the expected error instance")
		case e.Error() != expected.Error():
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("retrieving a non-loaded service", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		count := 0
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Service(id, func() (interface{}, error) {
			count++
			return entry, nil
		})

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
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Service(id, func() (interface{}, error) {
			count++
			return entry, nil
		})

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

	t.Run("error while calling the factory", func(t *testing.T) {
		sut := ServiceContainer{}
		id := "id"
		expected := fmt.Errorf("error message")
		_ = sut.Factory(id, func() (interface{}, error) {
			return nil, expected
		})

		check, e := sut.Get(id)
		switch {
		case check != nil:
			t.Error("returned an unexpected valid instance reference")
		case e == nil:
			t.Error("didn't returned the expected error instance")
		case e.Error() != expected.Error():
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("retrieving a non-loaded factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		count := 0
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Factory(id, func() (interface{}, error) {
			count++
			return entry, nil
		})

		if check, e := sut.Get(id); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if check == nil {
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving a loaded factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		count := 0
		sut := ServiceContainer{}
		entry := NewMockCloser(ctrl)
		_ = sut.Factory(id, func() (interface{}, error) {
			count++
			return entry, nil
		})

		runs := 1
		for runs < 3 {
			check, e := sut.Get(id)
			switch {
			case e != nil:
				t.Errorf("returned the unexpected error (%v)", e)
			case check == nil:
				t.Error("didn't returned a valid reference")
			case count != runs:
				t.Error("called the factory more than once")
			}
			runs++
		}
	})
}

func Test_ServiceContainer_Tagged(t *testing.T) {
	t.Run("error while instantiating retrieved service", func(t *testing.T) {
		sut := ServiceContainer{}
		id := "id"
		expected := fmt.Errorf("error message")
		factory := func() (interface{}, error) { return nil, expected }
		_ = sut.Service(id, factory, "tag1")

		check, e := sut.Tagged("tag1")
		switch {
		case check != nil:
			t.Error("returned an unexpected valid instance reference")
		case e == nil:
			t.Error("didn't returned the expected error instance")
		case e.Error() != expected.Error():
			t.Errorf("returned the error (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("retrieving a non-assigned tag list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "id1"
		id2 := "id2"
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry2 := NewMockCloser(ctrl)
		_ = sut.Service(id1, func() (interface{}, error) {
			return entry1, nil
		}, "tag1")
		_ = sut.Service(id2, func() (interface{}, error) {
			return entry2, nil
		}, "tag1", "tag2")

		if list, e := sut.Tagged("tag3"); e != nil {
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
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry2 := NewMockCloser(ctrl)
		_ = sut.Service(id1, func() (interface{}, error) {
			return entry1, nil
		}, "tag1")
		_ = sut.Service(id2, func() (interface{}, error) {
			return entry2, nil
		}, "tag1", "tag2")

		if list, e := sut.Tagged("tag2"); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if len(list) != 1 {
			t.Errorf("returned the unexpected (%v) list", list)
		} else if check := list[0]; !reflect.DeepEqual(check, entry2) {
			t.Errorf("returned the unexpected (%v) entry", entry2)
		}
	})

	t.Run("retrieving a tagged entries", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "id1"
		id2 := "id2"
		sut := ServiceContainer{}
		entry1 := NewMockCloser(ctrl)
		entry2 := NewMockCloser(ctrl)
		_ = sut.Service(id1, func() (interface{}, error) {
			return entry1, nil
		}, "tag1")
		_ = sut.Service(id2, func() (interface{}, error) {
			return entry2, nil
		}, "tag1", "tag2")

		if list, e := sut.Tagged("tag1"); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if len(list) != 2 {
			t.Errorf("returned the unexpected (%v) list", list)
		} else if check := list[0]; !reflect.DeepEqual(check, entry1) {
			t.Errorf("returned the unexpected (%v) entry", entry1)
		} else if check := list[1]; !reflect.DeepEqual(check, entry2) {
			t.Errorf("returned the unexpected (%v) entry", entry2)
		}
	})
}
