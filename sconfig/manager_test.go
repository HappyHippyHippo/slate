package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"testing"
	"time"
)

func Test_NewConfig(t *testing.T) {
	t.Run("new config without reload", func(t *testing.T) {
		cfg := NewManager(0 * time.Second)
		defer func() { _ = cfg.Close() }()

		switch {
		case cfg.(*manager).mutex == nil:
			t.Error("didn't instantiate the access mutex")
		case cfg.(*manager).sources == nil:
			t.Error("didn't instantiate the sources storing array")
		case cfg.(*manager).observers == nil:
			t.Error("didn't instantiate the observers storing array")
		case cfg.(*manager).loader != nil:
			t.Error("instantiated the sources reload trigger")
		}
	})

	t.Run("new config with reload", func(t *testing.T) {
		cfg := NewManager(10 * time.Second)
		defer func() { _ = cfg.Close() }()

		switch {
		case cfg.(*manager).mutex == nil:
			t.Error("didn't instantiate the access mutex")
		case cfg.(*manager).sources == nil:
			t.Error("didn't instantiate the sources storing array")
		case cfg.(*manager).observers == nil:
			t.Error("didn't instantiate the observers storing array")
		case cfg.(*manager).loader == nil:
			t.Error("didn't instantiate the sources reload trigger")
		}
	})
}

func Test_config_Close(t *testing.T) {
	t.Run("error while closing source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := NewManager(60 * time.Second)
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Partial{}, nil).AnyTimes()
		src.EXPECT().Close().Return(expected).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if err := cfg.Close(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})
	t.Run("error while closing loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := NewManager(60 * time.Second)
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Partial{}, nil).AnyTimes()
		src.EXPECT().Close().Return(nil).Times(1)
		_ = cfg.AddSource("src", 0, src)
		loader := NewMockTrigger(ctrl)
		loader.EXPECT().Close().Return(expected).Times(1)
		cfg.(*manager).loader = loader

		if err := cfg.Close(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("propagate close to sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		id1 := "src.1"
		id2 := "src.2"
		priority1 := 0
		priority2 := 1
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Partial{}, nil).AnyTimes()
		src1.EXPECT().Close().Times(1)
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Partial{}, nil).AnyTimes()
		src2.EXPECT().Close().Times(1)
		_ = cfg.AddSource(id1, priority1, src1)
		_ = cfg.AddSource(id2, priority2, src2)

		_ = cfg.Close()
	})
}

func Test_config_Has(t *testing.T) {
	t.Run("return the existence of the path", func(t *testing.T) {
		scenarios := []struct {
			config   Partial
			search   string
			expected bool
		}{
			{ // _test the existence of a present path
				config:   Partial{"node": "value"},
				search:   "node",
				expected: true,
			},
			{ // _test the non-existence of a missing path
				config:   Partial{"node": "value"},
				search:   "invalid-node",
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				cfg := NewManager(60 * time.Second)
				src := NewMockSource(ctrl)
				src.EXPECT().Close().Times(1)
				src.EXPECT().Get("").Return(scenario.config, nil).Times(1)
				_ = cfg.AddSource("src", 0, src)

				defer func() { _ = cfg.Close(); ctrl.Finish() }()

				if check := cfg.Has(scenario.search); check != scenario.expected {
					t.Errorf("returned (%v) when expecting (%v)", check, scenario.expected)
				}
			}
			test()
		}
	})
}

func Test_config_Get(t *testing.T) {
	t.Run("return path value", func(t *testing.T) {
		search := "node"
		expected := "value"
		config := Partial{search: expected}

		ctrl := gomock.NewController(t)

		cfg := NewManager(60 * time.Second)
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(config, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		defer func() { _ = cfg.Close(); ctrl.Finish() }()

		if check, err := cfg.Get(search); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if check != expected {
			t.Errorf("returned (%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("return internal Partial get error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{"node1": Partial{"node2": 101}}
		path := "node3"
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(data, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		check, err := cfg.Get(path)
		switch {
		case check != nil:
			t.Errorf("returned the unexpected valid value : %v", check)
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigPathNotFound):
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("return default if path was not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{"node1": Partial{"node2": 101}}
		path := "node3"
		val := 3
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(data, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if check, err := cfg.Get(path, val); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if check != val {
			t.Errorf("returned (%v) when expecting (%v)", check, val)
		}
	})
}

func Test_config_Bool(t *testing.T) {
	t.Run("return the stored boolean value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "node"
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{path: true}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if check, err := cfg.Bool(path); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if !check {
			t.Errorf("returned (%v)", check)
		}
	})
}

func Test_config_Int(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := 123
		path := "node"
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{path: value}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if check, err := cfg.Int(path); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if check != value {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_config_Float(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := 123.4
		path := "node"
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{path: value}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if check, err := cfg.Float(path); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if check != value {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_config_String(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := "value"
		path := "node"
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{path: value}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if check, err := cfg.String(path); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if check != value {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_config_List(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := []interface{}{1, 2, 3}
		path := "node"
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{path: value}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if check, err := cfg.List(path); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if !reflect.DeepEqual(check, value) {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_config_Config(t *testing.T) {
	t.Run("return the stored integer value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := Partial{"field": "value"}
		path := "node"
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{path: value}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if check, err := cfg.Partial(path); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if !reflect.DeepEqual(check, value) {
			t.Errorf("returned (%v) when expecting : %v", check, value)
		}
	})
}

func Test_config_HasSource(t *testing.T) {
	t.Run("validate if the source is registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if !cfg.HasSource("src") {
			t.Error("returned false")
		}
	})

	t.Run("invalidate if the source is not registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if cfg.HasSource("invalid source id") {
			t.Error("returned true")
		}
	})
}

func Test_config_AddSource(t *testing.T) {
	t.Run("nil source", func(t *testing.T) {
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		if err := cfg.AddSource("src", 0, nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register a new source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{}, nil).Times(1)

		if err := cfg.AddSource("src", 0, src); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !cfg.HasSource("src") {
			t.Error("didn't stored the source")
		}
	})

	t.Run("duplicate id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		if err := cfg.AddSource("src", 0, src); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrDuplicateConfigSource) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrDuplicateConfigSource)
		}
	})

	t.Run("override path if the insert have higher priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 1, src1)
		_ = cfg.AddSource("src.2", 2, src2)

		if check, _ := cfg.Get("node"); check != "value.2" {
			t.Errorf("returned the (%v) value when expecting (value.2)", check)
		}
	})

	t.Run("do not override path if the insert have lower priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 2, src1)
		_ = cfg.AddSource("src.2", 1, src2)

		if check, _ := cfg.Get("node"); check != "value.1" {
			t.Errorf("returned the (%v) value when expecting (value.1)", check)
		}
	})

	t.Run("still be able to get not overridden paths of a inserted lower priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2", "extendedNode": "extendedValue"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 2, src1)
		_ = cfg.AddSource("src.2", 1, src2)

		if check, _ := cfg.Get("extendedNode"); check != "extendedValue" {
			t.Errorf("returned the (%v) value when expecting (extendedValue)", check)
		}
	})
}

func Test_config_RemoveSource(t *testing.T) {
	t.Run("unregister a non-registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		if err := cfg.RemoveSource("src"); err != nil {
			t.Errorf("returned the unexpected error (%v)", err)
		}
	})

	t.Run("error unregister a previously registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Return(expected).Times(2)
		src.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		_ = cfg.AddSource("src", 0, src)

		if err := cfg.RemoveSource("src"); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", err, expected)
		}
	})

	t.Run("unregister a previously registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Times(1)
		src3.EXPECT().Get("").Return(Partial{"node": "value.3"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 0, src1)
		_ = cfg.AddSource("src.2", 0, src2)
		_ = cfg.AddSource("src.3", 0, src3)
		_ = cfg.RemoveSource("src.2")

		if cfg.HasSource("src.2") {
			t.Error("didn't remove the source")
		}
	})

	t.Run("recover path overridden by the removed source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Times(1)
		src3.EXPECT().Get("").Return(Partial{"node": "value.3"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 0, src1)
		_ = cfg.AddSource("src.2", 1, src2)
		_ = cfg.AddSource("src.3", 2, src3)
		_ = cfg.RemoveSource("src.3")

		if check, _ := cfg.Get("node"); check != "value.2" {
			t.Errorf("returned (%check) value when expecting (value.2)", check)
		}
	})
}

func Test_config_RemoveAllSources(t *testing.T) {
	t.Run("remove all the sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		expected := fmt.Errorf("error string")
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().MinTimes(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().MinTimes(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Return(expected).MinTimes(1)
		src3.EXPECT().Get("").Return(Partial{"node": "value.3"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 0, src1)
		_ = cfg.AddSource("src.2", 1, src2)
		_ = cfg.AddSource("src.3", 2, src3)

		if err := cfg.RemoveAllSources(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the error (%v) when was expecting (%v)", err, expected)
		}
	})

	t.Run("remove all the sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2"}, nil).AnyTimes()
		src3 := NewMockSource(ctrl)
		src3.EXPECT().Close().Times(1)
		src3.EXPECT().Get("").Return(Partial{"node": "value.3"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 0, src1)
		_ = cfg.AddSource("src.2", 1, src2)
		_ = cfg.AddSource("src.3", 2, src3)
		_ = cfg.RemoveAllSources()

		if len(cfg.(*manager).sources) != 0 {
			t.Error("didn't removed all the registered sources")
		}
	})
}

func Test_config_Source(t *testing.T) {
	t.Run("error if the source don't exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		check, err := cfg.Source("invalid id")
		switch {
		case check != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigSourceNotFound):
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConfigSourceNotFound)
		}
	})

	t.Run("return the registered source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{}, nil).Times(1)
		_ = cfg.AddSource("src", 0, src)

		check, err := cfg.Source("src")
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case check == nil:
			t.Error("returned nil")
		case !reflect.DeepEqual(check, src):
			t.Errorf("returned (%v) when expecting (%v)", check, src)
		}
	})
}

func Test_config_SourcePriority(t *testing.T) {
	t.Run("error if the source was not found", func(t *testing.T) {
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		if err := cfg.SourcePriority("invalid id", 0); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConfigSourceNotFound) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConfigSourceNotFound)
		}
	})

	t.Run("update the priority of the source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Close().Times(1)
		src1.EXPECT().Get("").Return(Partial{"node": "value.1"}, nil).AnyTimes()
		src2 := NewMockSource(ctrl)
		src2.EXPECT().Close().Times(1)
		src2.EXPECT().Get("").Return(Partial{"node": "value.2"}, nil).AnyTimes()
		_ = cfg.AddSource("src.1", 1, src1)
		_ = cfg.AddSource("src.2", 2, src2)

		if check, _ := cfg.Get("node"); check != "value.2" {
			t.Errorf("returned the (%v) value prior the change, when expecting (value.2)", check)
		}
		if err := cfg.SourcePriority("src.2", 0); err != nil {
			t.Errorf("returned the unexpeced error : (%v)", err)
		}
		if check, _ := cfg.Get("node"); check != "value.1" {
			t.Errorf("returned the (%v) value after the change, when expecting (value.1)", check)
		}
	})
}

func Test_config_HasObserver(t *testing.T) {
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

				cfg := NewManager(0 * time.Second)
				defer func() { _ = cfg.Close() }()
				src := NewMockSource(ctrl)
				src.EXPECT().Close().Times(1)
				src.EXPECT().Get("").Return(Partial{"node1": "value1", "node2": "value2", "node3": "value3"}, nil).Times(1)
				_ = cfg.AddSource("cfg", 0, src)

				for _, observer := range scenario.observers {
					_ = cfg.AddObserver(observer, func(old, new interface{}) {})
				}

				if check := cfg.HasObserver(scenario.search); check != scenario.exp {
					t.Errorf("returned (%v) when expecting (%v)", check, scenario.exp)
				}
			}
			test()
		}
	})
}

func Test_config_AddObserver(t *testing.T) {
	t.Run("nil callback", func(t *testing.T) {
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		if err := cfg.AddObserver("path", nil); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error if path not present", func(t *testing.T) {
		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		if err := cfg.AddObserver("path", func(interface{}, interface{}) {}); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConfigPathNotFound) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("valid callback", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()
		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{"path": "value"}, nil).Times(1)
		_ = cfg.AddSource("cfg", 0, src)

		if err := cfg.AddObserver("path", func(interface{}, interface{}) {}); err != nil {
			t.Errorf("returned the unexpected error, %v", err)
		} else if len(cfg.(*manager).observers) != 1 {
			t.Error("didn't stored the requested observer")
		}
	})
}

func Test_config_RemoveObserver(t *testing.T) {
	t.Run("remove a registered observer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(60 * time.Second)
		defer func() { _ = cfg.Close() }()

		src := NewMockSource(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{"node": Partial{"1": "value1", "2": "value2", "3": "value3"}}, nil).Times(1)
		_ = cfg.AddSource("cfg", 0, src)

		_ = cfg.AddObserver("node.1", func(old, new interface{}) {})
		_ = cfg.AddObserver("node.2", func(old, new interface{}) {})
		_ = cfg.AddObserver("node.3", func(old, new interface{}) {})
		cfg.RemoveObserver("node.2")

		if cfg.HasObserver("node.2") {
			t.Errorf("didn't removed the observer")
		}
	})
}

func Test_Config(t *testing.T) {
	t.Run("reload on observable sources", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(2 * time.Millisecond)
		defer func() { _ = cfg.Close() }()

		src := NewMockSourceObservable(ctrl)
		src.EXPECT().Close().Times(1)
		src.EXPECT().Get("").Return(Partial{"node": "value"}, nil).Times(1)
		src.EXPECT().Reload().Return(false, nil).MinTimes(1)
		_ = cfg.AddSource("src", 0, src)

		time.Sleep(100 * time.Millisecond)
	})

	t.Run("rebuild if the observable source notify changes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewManager(20 * time.Millisecond)

		src := NewMockSourceObservable(ctrl)
		src.EXPECT().Get("").Return(Partial{"node": "value"}, nil).MinTimes(2)
		src.EXPECT().Reload().Return(true, nil).MinTimes(1)
		_ = cfg.AddSource("src", 0, src)

		time.Sleep(200 * time.Millisecond)

		if check, _ := cfg.Get("node"); check != "value" {
			t.Errorf("returned (%v) when expecting (value)", check)
		}
	})

	t.Run("should call observer callback function on config changes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check := false
		cfg := NewManager(20 * time.Millisecond)

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Partial{"node": "value1"}, nil).AnyTimes()
		_ = cfg.AddSource("src1", 0, src1)

		_ = cfg.AddObserver("node", func(old, new interface{}) {
			check = true

			if old != "value1" {
				t.Errorf("callback called with (%v) as old value", old)
			}
			if new != "value2" {
				t.Errorf("callback called with (%v) as new value", new)
			}
		})

		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Partial{"node": "value2"}, nil).AnyTimes()
		_ = cfg.AddSource("src2", 1, src2)

		if !check {
			t.Errorf("didn't actually called the callback")
		} else if check := cfg.(*manager).observers[0].current; check != "value2" {
			t.Errorf("stored the current value {%v} instead of the expected {%v}", check, "value2")
		}
	})

	t.Run("should call observer callback function on config changes on a list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check := false
		cfg := NewManager(20 * time.Millisecond)
		initial := []interface{}{Partial{"subnode": "value1"}}
		expected := []interface{}{Partial{"subnode": "value2"}}

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Partial{"node": initial}, nil).AnyTimes()
		_ = cfg.AddSource("src1", 0, src1)

		_ = cfg.AddObserver("node", func(old, new interface{}) {
			check = true

			if old.([]interface{})[0].(Partial)["subnode"] != initial[0].(Partial)["subnode"] {
				t.Errorf("callback called with (%v) as old value", old)
			}
			if new.([]interface{})[0].(Partial)["subnode"] != expected[0].(Partial)["subnode"] {
				t.Errorf("callback called with (%v) as new value", new)
			}
		})

		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Partial{"node": expected}, nil).AnyTimes()
		_ = cfg.AddSource("src2", 1, src2)

		if !check {
			t.Errorf("didn't actually called the callback")
		} else if check := cfg.(*manager).observers[0].current; check.([]interface{})[0].(Partial)["subnode"] != expected[0].(Partial)["subnode"] {
			t.Errorf("stored the current value {%v} instead of the expected {%v}", check, expected)
		}
	})

	t.Run("should call observer callback function on config changes on a partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check := false
		cfg := NewManager(20 * time.Millisecond)
		initial := Partial{"subnode": "value1"}
		expected := Partial{"subnode": "value2"}

		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(Partial{"node": initial}, nil).AnyTimes()
		_ = cfg.AddSource("src1", 0, src1)

		_ = cfg.AddObserver("node", func(old, new interface{}) {
			check = true

			if reflect.DeepEqual(old, initial) {
				t.Errorf("callback called with (%v) as old value", old)
			}
			if reflect.DeepEqual(old, expected) {
				t.Errorf("callback called with (%v) as new value", new)
			}
		})

		src2 := NewMockSource(ctrl)
		src2.EXPECT().Get("").Return(Partial{"node": expected}, nil).AnyTimes()
		_ = cfg.AddSource("src2", 1, src2)

		if !check {
			t.Errorf("didn't actually called the callback")
		} else if check := cfg.(*manager).observers[0].current; check.(Partial)["subnode"] != expected["subnode"] {
			t.Errorf("stored the current value {%v} instead of the expected {%v}", check, expected)
		}
	})
}
