package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewManager(t *testing.T) {
	t.Run("new config without reload", func(t *testing.T) {
		sut := NewManager(0 * time.Second)
		defer func() { _ = sut.Close() }()

		switch {
		case sut.mutex == nil:
			t.Error("didn't instantiate the access mutex")
		case sut.sources == nil:
			t.Error("didn't instantiate the sources storing array")
		case sut.observers == nil:
			t.Error("didn't instantiate the observers storing array")
		case sut.loader != nil:
			t.Error("instantiated the sources reload trigger")
		}
	})

	t.Run("new config with reload", func(t *testing.T) {
		sut := NewManager(10 * time.Second)
		defer func() { _ = sut.Close() }()

		switch {
		case sut.mutex == nil:
			t.Error("didn't instantiate the access mutex")
		case sut.sources == nil:
			t.Error("didn't instantiate the sources storing array")
		case sut.observers == nil:
			t.Error("didn't instantiate the observers storing array")
		case sut.loader == nil:
			t.Error("didn't instantiate the sources reload trigger")
		}
	})
}

func Test_Manager_Close(t *testing.T) {
	t.Run("error while closing source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		sut := NewManager(60 * time.Second)
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Config{}, nil).AnyTimes()
		src.EXPECT().Close().Return(expected).Times(1)
		_ = sut.AddSource("src", 0, src)

		if e := sut.Close(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
		}
	})
	t.Run("error while closing loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		sut := NewManager(60 * time.Second)
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Config{}, nil).AnyTimes()
		src.EXPECT().Close().Return(nil).Times(1)
		_ = sut.AddSource("src", 0, src)
		trigger := NewMockTrigger(ctrl)
		trigger.EXPECT().Close().Return(expected).Times(1)
		sut.loader = trigger

		if e := sut.Close(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("propagate close to sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		id1 := "src.1"
		id2 := "src.2"
		priority1 := 0
		priority2 := 1
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Config{}, nil).AnyTimes()
		src1.EXPECT().Close().Times(1)
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Config{}, nil).AnyTimes()
		src2.EXPECT().Close().Times(1)
		_ = sut.AddSource(id1, priority1, src1)
		_ = sut.AddSource(id2, priority2, src2)

		_ = sut.Close()
	})
}

func Test_Manager_Entries(t *testing.T) {
	t.Run("return config entries", func(t *testing.T) {
		scenarios := []struct {
			config   Config
			expected []string
		}{
			{ // _test the empty config
				config:   Config{},
				expected: nil,
			},
			{ // _test the single entry config
				config:   Config{"field": "value"},
				expected: []string{"field"},
			},
			{ // _test the multi entry config
				config:   Config{"field1": "value 1", "field2": "value 2"},
				expected: []string{"field1", "field2"},
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				sut := NewManager(60 * time.Second)
				src := NewMockSource(ctrl)
				src.EXPECT().Close().Times(1)
				src.EXPECT().Get("").Return(scenario.config, nil).Times(1)
				_ = sut.AddSource("src", 0, src)

				defer func() { _ = sut.Close(); ctrl.Finish() }()

				if check := sut.Entries(); !reflect.DeepEqual(check, scenario.expected) {
					t.Errorf("returned (%v) when expecting (%v)", check, scenario.expected)
				}
			}
			test()
		}
	})
}

func Test_Manager_Has(t *testing.T) {
	t.Run("return the existence of the path", func(t *testing.T) {
		scenarios := []struct {
			config   Config
			search   string
			expected bool
		}{
			{ // _test the existence of a present path
				config:   Config{"node": "value"},
				search:   "node",
				expected: true,
			},
			{ // _test the non-existence of a missing path
				config:   Config{"node": "value"},
				search:   "invalid-node",
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				sut := NewManager(60 * time.Second)
				src := NewMockSource(ctrl)
				src.EXPECT().Close().Times(1)
				src.EXPECT().Get("").Return(scenario.config, nil).Times(1)
				_ = sut.AddSource("src", 0, src)

				defer func() { _ = sut.Close(); ctrl.Finish() }()

				if check := sut.Has(scenario.search); check != scenario.expected {
					t.Errorf("returned (%v) when expecting (%v)", check, scenario.expected)
				}
			}
			test()
		}
	})
}

func Test_Manager_Get(t *testing.T) {
	t.Run("return path value", func(t *testing.T) {
		search := "node"
		expected := "value"
		config := Config{search: expected}

		ctrl := gomock.NewController(t)

		sut := NewManager(60 * time.Second)
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(config, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		defer func() { _ = sut.Close(); ctrl.Finish() }()

		if check, e := sut.Get(search); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if check != expected {
			t.Errorf("returned (%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("return internal Partial get error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Config{"node1": Config{"node2": 101}}
		path := "node3"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(data, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		check, e := sut.Get(path)
		switch {
		case check != nil:
			t.Errorf("returned the unexpected valid value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("returned (%v) err when expecting (%v)", e, err.ConfigPathNotFound)
		}
	})

	t.Run("return default if path was not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Config{"node1": Config{"node2": 101}}
		path := "node3"
		val := 3
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(data, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.Get(path, val); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if check != val {
			t.Errorf("returned (%v) when expecting (%v)", check, val)
		}
	})
}

func Test_Manager_Bool(t *testing.T) {
	t.Run("return the stored boolean value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "node"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{path: true}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.Bool(path); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if !check {
			t.Errorf("returned (%v)", check)
		}
	})
}

func Test_Manager_Int(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := 123
		path := "node"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{path: value}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.Int(path); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if check != value {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Manager_Float(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := 123.4
		path := "node"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{path: value}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.Float(path); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if check != value {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Manager_String(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := "value"
		path := "node"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{path: value}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.String(path); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if check != value {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Manager_List(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := []interface{}{1, 2, 3}
		path := "node"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{path: value}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.List(path); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if !reflect.DeepEqual(check, value) {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Manager_Config(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := Config{"field": "value"}
		path := "node"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{path: value}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.Config(path); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if !reflect.DeepEqual(*check.(*Config), value) {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Manager_Populate(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := Config{"field": Config{"field": "value"}}
		target := struct{ Field string }{}
		expected := struct{ Field string }{Field: "value"}
		path := "node"
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{path: value}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if check, e := sut.Populate(path+"."+"field", target); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if !reflect.DeepEqual(check, expected) {
			t.Errorf("returned (%v) when expecting : %v", check, expected)
		}
	})
}

func Test_Manager_HasSource(t *testing.T) {
	t.Run("validate if the source is registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if !sut.HasSource("src") {
			t.Error("returned false")
		}
	})

	t.Run("invalidate if the source is not registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if sut.HasSource("invalid source id") {
			t.Error("returned true")
		}
	})
}

func Test_Manager_AddSource(t *testing.T) {
	t.Run("nil source", func(t *testing.T) {
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		if e := sut.AddSource("src", 0, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("register a new source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{}, nil).Times(1)

		if e := sut.AddSource("src", 0, src); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !sut.HasSource("src") {
			t.Error("didn't stored the source")
		}
	})

	t.Run("duplicate id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		if e := sut.AddSource("src", 0, src); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.DuplicateConfigSource) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.DuplicateConfigSource)
		}
	})

	t.Run("override path if the insert have higher priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 1, src1)
		_ = sut.AddSource("src.2", 2, src2)

		if check, _ := sut.Get("node"); check != "value.2" {
			t.Errorf("returned the (%v) value when expecting (value.2)", check)
		}
	})

	t.Run("do not override path if the insert have lower priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 2, src1)
		_ = sut.AddSource("src.2", 1, src2)

		if check, _ := sut.Get("node"); check != "value.1" {
			t.Errorf("returned the (%v) value when expecting (value.1)", check)
		}
	})

	t.Run("still be able to get not overridden paths of a inserted lower priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2", "extendedNode": "extendedValue"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 2, src1)
		_ = sut.AddSource("src.2", 1, src2)

		if check, _ := sut.Get("extendedNode"); check != "extendedValue" {
			t.Errorf("returned the (%v) value when expecting (extendedValue)", check)
		}
	})
}

func Test_Manager_RemoveSource(t *testing.T) {
	t.Run("unregister a non-registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		if e := sut.RemoveSource("src"); e != nil {
			t.Errorf("returned the unexpected err (%v)", e)
		}
	})

	t.Run("error unregister a previously registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Return(expected).Times(2)
		src.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		_ = sut.AddSource("src", 0, src)

		if e := sut.RemoveSource("src"); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the err (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("unregister a previously registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Times(1)
		src3.EXPECT().Get("").Return(Config{"node": "value.3"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 0, src1)
		_ = sut.AddSource("src.2", 0, src2)
		_ = sut.AddSource("src.3", 0, src3)
		_ = sut.RemoveSource("src.2")

		if sut.HasSource("src.2") {
			t.Error("didn't remove the source")
		}
	})

	t.Run("recover path overridden by the removed source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Times(1)
		src3.EXPECT().Get("").Return(Config{"node": "value.3"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 0, src1)
		_ = sut.AddSource("src.2", 1, src2)
		_ = sut.AddSource("src.3", 2, src3)
		_ = sut.RemoveSource("src.3")

		if check, _ := sut.Get("node"); check != "value.2" {
			t.Errorf("returned (%check) value when expecting (value.2)", check)
		}
	})
}

func Test_Manager_RemoveAllSources(t *testing.T) {
	t.Run("remove all the sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		expected := fmt.Errorf("err string")
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().MinTimes(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().MinTimes(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Return(expected).MinTimes(1)
		src3.EXPECT().Get("").Return(Config{"node": "value.3"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 0, src1)
		_ = sut.AddSource("src.2", 1, src2)
		_ = sut.AddSource("src.3", 2, src3)

		if e := sut.RemoveAllSources(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the err (%v) when was expecting (%v)", e, expected)
		}
	})

	t.Run("remove all the sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Times(1)
		src3.EXPECT().Get("").Return(Config{"node": "value.3"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 0, src1)
		_ = sut.AddSource("src.2", 1, src2)
		_ = sut.AddSource("src.3", 2, src3)
		_ = sut.RemoveAllSources()

		if len(sut.sources) != 0 {
			t.Error("didn't removed all the registered sources")
		}
	})
}

func Test_Manager_Source(t *testing.T) {
	t.Run("error if the source don't exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		check, e := sut.Source("invalid id")
		switch {
		case check != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigSourceNotFound):
			t.Errorf("returned (%v) err when expecting (%v)", e, err.ConfigSourceNotFound)
		}
	})

	t.Run("return the registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{}, nil).Times(1)
		_ = sut.AddSource("src", 0, src)

		check, e := sut.Source("src")
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case check == nil:
			t.Error("returned nil")
		case !reflect.DeepEqual(check, src):
			t.Errorf("returned (%v) when expecting (%v)", check, src)
		}
	})
}

func Test_Manager_SourcePriority(t *testing.T) {
	t.Run("error if the source was not found", func(t *testing.T) {
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		if e := sut.SourcePriority("invalid id", 0); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ConfigSourceNotFound) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.ConfigSourceNotFound)
		}
	})

	t.Run("update the priority of the source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Config{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Config{"node": "value.2"}, nil).AnyTimes()
		_ = sut.AddSource("src.1", 1, src1)
		_ = sut.AddSource("src.2", 2, src2)

		if check, _ := sut.Get("node"); check != "value.2" {
			t.Errorf("returned the (%v) value prior the change, when expecting (value.2)", check)
		}
		if e := sut.SourcePriority("src.2", 0); e != nil {
			t.Errorf("returned the unexpeced err : (%v)", e)
		}
		if check, _ := sut.Get("node"); check != "value.1" {
			t.Errorf("returned the (%v) value after the change, when expecting (value.1)", check)
		}
	})
}

func Test_Manager_HasObserver(t *testing.T) {
	t.Run("check the existence of a observer", func(t *testing.T) {
		scenarios := []struct {
			observers []string
			search    string
			exp       bool
		}{
			{ // Search a non-existing path in an empty list of observers
				observers: []string{},
				search:    "node1",
				exp:       false,
			},
			{ // Search a non-existing path in a non-empty list of observers
				observers: []string{"node1", "node2"},
				search:    "node3",
				exp:       false,
			},
			{ // Search an existing path in a list of observers
				observers: []string{"node1", "node2", "node3"},
				search:    "node2",
				exp:       true,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				sut := NewManager(0 * time.Second)
				defer func() { _ = sut.Close() }()
				src := NewMockSource(ctrl)
				src.EXPECT().Close().Times(1)
				src.EXPECT().Get("").Return(Config{"node1": "value1", "node2": "value2", "node3": "value3"}, nil).Times(1)
				_ = sut.AddSource("manager", 0, src)

				for _, observer := range scenario.observers {
					_ = sut.AddObserver(observer, func(old, new interface{}) {})
				}

				if check := sut.HasObserver(scenario.search); check != scenario.exp {
					t.Errorf("returned (%v) when expecting (%v)", check, scenario.exp)
				}
			}
			test()
		}
	})
}

func Test_Manager_AddObserver(t *testing.T) {
	t.Run("nil callback", func(t *testing.T) {
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		if e := sut.AddObserver("path", nil); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("error if path not present", func(t *testing.T) {
		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		if e := sut.AddObserver("path", func(interface{}, interface{}) {}); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, err.ConfigPathNotFound) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.ConfigPathNotFound)
		}
	})

	t.Run("valid callback", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{"path": "value"}, nil).Times(1)
		_ = sut.AddSource("manager", 0, src)

		if e := sut.AddObserver("path", func(interface{}, interface{}) {}); e != nil {
			t.Errorf("returned the unexpected error, %v", e)
		} else if len(sut.observers) != 1 {
			t.Error("didn't stored the requested observer")
		}
	})
}

func Test_Manager_RemoveObserver(t *testing.T) {
	t.Run("remove a registered observer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(60 * time.Second)
		defer func() { _ = sut.Close() }()

		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{"node": Config{"1": "value1", "2": "value2", "3": "value3"}}, nil).Times(1)
		_ = sut.AddSource("manager", 0, src)

		_ = sut.AddObserver("node.1", func(old, new interface{}) {})
		_ = sut.AddObserver("node.2", func(old, new interface{}) {})
		_ = sut.AddObserver("node.3", func(old, new interface{}) {})
		sut.RemoveObserver("node.2")

		if sut.HasObserver("node.2") {
			t.Errorf("didn't removed the observer")
		}
	})
}

func Test_Manager(t *testing.T) {
	t.Run("reload on observable sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(2 * time.Millisecond)
		defer func() { _ = sut.Close() }()

		src := NewMockObservableSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Config{"node": "value"}, nil).Times(1)
		src.EXPECT().Reload().Return(false, nil).MinTimes(1)
		_ = sut.AddSource("src", 0, src)

		time.Sleep(100 * time.Millisecond)
	})

	t.Run("rebuild if the observable source notify changes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := NewManager(20 * time.Millisecond)

		src := NewMockObservableSource(ctrl)
		src.EXPECT().Get("").Return(Config{"node": "value"}, nil).MinTimes(2)
		src.EXPECT().Reload().Return(true, nil).MinTimes(1)
		_ = sut.AddSource("src", 0, src)

		time.Sleep(200 * time.Millisecond)

		if check, _ := sut.Get("node"); check != "value" {
			t.Errorf("returned (%v) when expecting (value)", check)
		}
	})

	t.Run("should call observer callback function on config changes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check := false
		sut := NewManager(20 * time.Millisecond)

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Config{"node": "value1"}, nil).AnyTimes()
		_ = sut.AddSource("src1", 0, src1)

		_ = sut.AddObserver("node", func(old, new interface{}) {
			check = true

			if old != "value1" {
				t.Errorf("callback called with (%v) as old value", old)
			}
			if new != "value2" {
				t.Errorf("callback called with (%v) as new value", new)
			}
		})

		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Config{"node": "value2"}, nil).AnyTimes()
		_ = sut.AddSource("src2", 1, src2)

		if !check {
			t.Errorf("didn't actually called the callback")
		} else if check := sut.observers[0].current; check != "value2" {
			t.Errorf("stored the current value {%v} instead of the expected {%v}", check, "value2")
		}
	})

	t.Run("should call observer callback function on config changes on a list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check := false
		sut := NewManager(20 * time.Millisecond)
		initial := []interface{}{Config{"sub_node": "value1"}}
		expected := []interface{}{Config{"sub_node": "value2"}}

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Config{"node": initial}, nil).AnyTimes()
		_ = sut.AddSource("src1", 0, src1)

		_ = sut.AddObserver("node", func(old, new interface{}) {
			check = true

			if old.([]interface{})[0].(Config)["sub_node"] != initial[0].(Config)["sub_node"] {
				t.Errorf("callback called with (%v) as old value", old)
			}
			if new.([]interface{})[0].(Config)["sub_node"] != expected[0].(Config)["sub_node"] {
				t.Errorf("callback called with (%v) as new value", new)
			}
		})

		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Config{"node": expected}, nil).AnyTimes()
		_ = sut.AddSource("src2", 1, src2)

		if !check {
			t.Errorf("didn't actually called the callback")
		} else if check := sut.observers[0].current; check.([]interface{})[0].(Config)["sub_node"] != expected[0].(Config)["sub_node"] {
			t.Errorf("stored the current value {%v} instead of the expected {%v}", check, expected)
		}
	})

	t.Run("should call observer callback function on config changes on a config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check := false
		sut := NewManager(20 * time.Millisecond)
		initial := Config{"sub_node": "value1"}
		expected := Config{"sub_node": "value2"}

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Config{"node": initial}, nil).AnyTimes()
		_ = sut.AddSource("src1", 0, src1)

		_ = sut.AddObserver("node", func(old, new interface{}) {
			check = true

			if reflect.DeepEqual(old, initial) {
				t.Errorf("callback called with (%v) as old value", old)
			}
			if reflect.DeepEqual(old, expected) {
				t.Errorf("callback called with (%v) as new value", new)
			}
		})

		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Config{"node": expected}, nil).AnyTimes()
		_ = sut.AddSource("src2", 1, src2)

		if !check {
			t.Errorf("didn't actually called the callback")
		} else if check := sut.observers[0].current; check.(Config)["sub_node"] != expected["sub_node"] {
			t.Errorf("stored the current value {%v} instead of the expected {%v}", check, expected)
		}
	})
}
