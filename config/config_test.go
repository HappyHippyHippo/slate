package config

import (
	"errors"
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate/err"
)

func Test_Config_Clone(t *testing.T) {
	t.Run("clone empty config", func(t *testing.T) {
		sut := Config{}
		c := sut.Clone()
		sut["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if len(c) != 0 {
			t.Errorf("cloned config is not empty : %v", c)
		}
	})

	t.Run("clone non-empty config", func(t *testing.T) {
		sut := Config{"field": "value"}
		expected := Config{"field": "value"}
		c := sut.Clone()
		sut["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned config (%v) is not the expected : %v", c, expected)
		}
	})

	t.Run("recursive cloning", func(t *testing.T) {
		sut := Config{"field": Config{"field": "value"}}
		expected := Config{"field": Config{"field": "value"}}
		c := sut.Clone()
		sut["extra"] = "value"
		sut["field"].(Config)["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned config (%v) is not the expected : %v", c, expected)
		}
	})

	t.Run("recursive cloning with lists", func(t *testing.T) {
		sut := Config{"field": []interface{}{Config{"field": "value"}}}
		expected := Config{"field": []interface{}{Config{"field": "value"}}}
		c := sut.Clone()
		sut["extra"] = "value"
		sut["field"].([]interface{})[0].(Config)["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned config (%v) is not the expected : %v", c, expected)
		}
	})

	t.Run("recursive cloning with multi-level lists", func(t *testing.T) {
		sut := Config{"field": []interface{}{[]interface{}{Config{"field": "value"}}}}
		expected := Config{"field": []interface{}{[]interface{}{Config{"field": "value"}}}}
		c := sut.Clone()
		sut["extra"] = "value"
		sut["field"].([]interface{})[0].([]interface{})[0].(Config)["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned config (%v) is not the expected : %v", c, expected)
		}
	})
}

func Test_Config_Entries(t *testing.T) {
	t.Run("empty config", func(t *testing.T) {
		if (&Config{}).Entries() != nil {
			t.Errorf("didn't returned the expected empty list")
		}
	})

	t.Run("single entry config", func(t *testing.T) {
		if !reflect.DeepEqual((&Config{"field": "value"}).Entries(), []string{"field"}) {
			t.Errorf("didn't returned the expected single entry list")
		}
	})

	t.Run("multi entry config", func(t *testing.T) {
		if !reflect.DeepEqual((&Config{
			"field1": "value 1",
			"field2": "value 2",
		}).Entries(), []string{"field1", "field2"}) {
			t.Errorf("didn't returned the expected single entry list")
		}
	})
}

func Test_Config_Has(t *testing.T) {
	t.Run("check if a valid path exists", func(t *testing.T) {
		scenarios := []struct {
			partial Config
			search  string
		}{
			{ // _test empty Config, search for everything
				partial: Config{},
				search:  "",
			},
			{ // _test single node, search for root node
				partial: Config{"node": "value"},
				search:  "",
			},
			{ // _test single node search
				partial: Config{"node": "value"},
				search:  "node",
			},
			{ // _test multiple node, search for root node
				partial: Config{"node1": "value", "node2": "value"},
				search:  "",
			},
			{ // _test multiple node search for first
				partial: Config{"node1": "value", "node2": "value"},
				search:  "node1",
			},
			{ // _test multiple node search for non-first
				partial: Config{"node1": "value", "node2": "value"},
				search:  "node2",
			},
			{ // _test tree, search for root node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "",
			},
			{ // _test tree, search for root level node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node1",
			},
			{ // _test tree, search for sub node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node1.node2",
			},
		}

		for _, scenario := range scenarios {
			if check := scenario.partial.Has(scenario.search); !check {
				t.Errorf("didn't found the (%s) path in (%v)", scenario.search, scenario.partial)
			}
		}
	})

	t.Run("check if a invalid path do not exists", func(t *testing.T) {
		scenarios := []struct {
			partial Config
			search  string
		}{
			{ // _test single node search (invalid)
				partial: Config{"node": "value"},
				search:  "node2",
			},
			{ // _test multiple node search for invalid node
				partial: Config{"node1": "value", "node2": "value"},
				search:  "node3",
			},
			{ // _test tree search for invalid root node
				partial: Config{"node": Config{"node": "value"}},
				search:  "node1",
			},
			{ // _test tree search for invalid sub node
				partial: Config{"node": Config{"node": "value"}},
				search:  "node.node1",
			},
			{ // _test tree search for invalid sub-sub-node
				partial: Config{"node": Config{"node": "value"}},
				search:  "node.node.node",
			},
		}

		for _, scenario := range scenarios {
			if check := scenario.partial.Has(scenario.search); check {
				t.Errorf("founded the (%s) path in (%v)", scenario.search, scenario.partial)
			}
		}
	})
}

func Test_Config_Get(t *testing.T) {
	t.Run("retrieve a value of a existent path", func(t *testing.T) {
		scenarios := []struct {
			partial  Config
			search   string
			expected interface{}
		}{
			{ // _test empty Config, search for everything
				partial:  Config{},
				search:   "",
				expected: Config{},
			},
			{ // _test single node, search for root node
				partial:  Config{"node": "value"},
				search:   "",
				expected: Config{"node": "value"},
			},
			{ // _test single node search
				partial:  Config{"node": "value"},
				search:   "node",
				expected: "value",
			},
			{ // _test multiple node, search for root node
				partial:  Config{"node1": "value1", "node2": "value2"},
				search:   "",
				expected: Config{"node1": "value1", "node2": "value2"},
			},
			{ // _test multiple node search for first
				partial:  Config{"node1": "value1", "node2": "value2"},
				search:   "node1",
				expected: "value1",
			},
			{ // _test multiple node search for non-first
				partial:  Config{"node1": "value1", "node2": "value2"},
				search:   "node2",
				expected: "value2",
			},
			{ // _test tree, search for root node
				partial:  Config{"node": Config{"node": "value"}},
				search:   "",
				expected: Config{"node": Config{"node": "value"}},
			},
			{ // _test tree, search for root level node
				partial:  Config{"node": Config{"node": "value"}},
				search:   "node",
				expected: Config{"node": "value"},
			},
			{ // _test tree, search for sub node
				partial:  Config{"node": Config{"node": "value"}},
				search:   "node.node",
				expected: "value",
			},
		}

		for _, scenario := range scenarios {
			if check, e := scenario.partial.Get(scenario.search); e != nil {
				t.Errorf("returned the unexpected err (%v) when retrieving (%v)", e, scenario.search)
			} else if !reflect.DeepEqual(check, scenario.expected) {
				t.Errorf("returned (%v) when retrieving (%v), expected (%v)", check, scenario.search, scenario.expected)
			}
		}
	})

	t.Run("return nil if a path don't exists", func(t *testing.T) {
		scenarios := []struct {
			partial Config
			search  string
		}{
			{ // _test empty Config search for non-existent node
				partial: Config{},
				search:  "node",
			},
			{ // _test single node search for non-existent node
				partial: Config{"node": "value"},
				search:  "node2",
			},
			{ // _test multiple node search for non-existent node
				partial: Config{"node1": "value1", "node2": "value2"},
				search:  "node3",
			},
			{ // _test tree search for non-existent root node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node2",
			},
			{ // _test tree search for non-existent sub node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node1.node1",
			},
			{ // _test tree search for non-existent sub-sub-node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node1.node2.node3",
			},
		}

		for _, scenario := range scenarios {
			check, e := scenario.partial.Get(scenario.search)
			switch {
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, err.ConfigPathNotFound):
				t.Errorf("returned err was not a config path not found err : %v", e)
			case check != nil:
				t.Error("unexpectedly returned a valid reference to a stored config value")
			}
		}
	})

	t.Run("return nil if the node actually stores nil", func(t *testing.T) {
		sut := Config{"node1": nil, "node2": "value2"}

		if check, e := sut.Get("node1", "default_value"); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if check != nil {
			t.Errorf("returned the (%v) check", check)
		}
	})

	t.Run("return the default value if a path don't exists", func(t *testing.T) {
		scenarios := []struct {
			partial Config
			search  string
		}{
			{ // _test empty Config search for non-existent node
				partial: Config{},
				search:  "node",
			},
			{ // _test single node search for non-existent node
				partial: Config{"node": "value"},
				search:  "node2",
			},
			{ // _test multiple node search for non-existent node
				partial: Config{"node1": "value1", "node2": "value2"},
				search:  "node3",
			},
			{ // _test tree search for non-existent root node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node2",
			},
			{ // _test tree search for non-existent sub node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node1.node1",
			},
			{ // _test tree search for non-existent sub-sub-node
				partial: Config{"node1": Config{"node2": "value"}},
				search:  "node1.node2.node3",
			},
		}

		def := "default_value"
		for _, scenario := range scenarios {
			if check, e := scenario.partial.Get(scenario.search, def); e != nil {
				t.Errorf("returned the unexpected err : %v", e)
			} else if check != def {
				t.Errorf("returned (%v) when retrieving (%v)", check, scenario.search)
			}
		}
	})
}

func Test_Config_Bool(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		sut := Config{path: true}

		if check, e := sut.Bool(path, false); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if !check {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, true)
		}
	})

	t.Run("return conversion err if not storing a bool", func(t *testing.T) {
		path := "node"
		value := "123"
		sut := Config{path: value}

		check, e := sut.Bool(path, true)
		switch {
		case check:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("the returned err is the expected err convertion err : %v", e)
		}
	})

	t.Run("return path not found err if no default value is given", func(t *testing.T) {
		sut := Config{}

		check, e := sut.Bool("node")
		switch {
		case check:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("the returned err is the expected err convertion err : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		sut := Config{}

		check, e := sut.Bool("node", true)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err : %v", e)
		case !check:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, true)
		}
	})
}

func Test_Config_Int(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Config{path: value}

		if check, e := sut.Int(path, 456); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if check != value {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion err if not storing an int", func(t *testing.T) {
		path := "node"
		value := "123"
		sut := Config{path: value}

		check, e := sut.Int(path, 456)
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("the returned err is the expected err convertion err : %v", e)
		}
	})

	t.Run("return path not found err if no default value is given", func(t *testing.T) {
		sut := Config{}

		check, e := sut.Int("node")
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := 123
		sut := Config{}

		check, e := sut.Int("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case check != value:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Config_Float(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := 123.456
		sut := Config{path: value}

		if check, e := sut.Float(path, 456.789); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if check != value {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an float", func(t *testing.T) {
		path := "node"
		value := "123.456"
		sut := Config{path: value}

		check, e := sut.Float(path, 456)
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Config{}

		check, e := sut.Float("node")
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := 123.456
		sut := Config{}

		check, e := sut.Float("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case check != value:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Config_String(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := "value"
		sut := Config{path: value}

		if check, e := sut.String(path, "default value"); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if check != value {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an string", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Config{path: value}

		check, e := sut.String(path, "default value")
		switch {
		case check != "":
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Config{}

		check, e := sut.String("node")
		switch {
		case check != "":
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := "default value"
		sut := Config{}

		check, e := sut.String("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case check != value:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Config_List(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := []interface{}{"value"}
		sut := Config{path: value}

		if check, e := sut.List(path, []interface{}{"default value"}); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if !reflect.DeepEqual(check, value) {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an list", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Config{path: value}

		check, e := sut.List(path, []interface{}{"default value"})
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Config{}

		check, e := sut.List("node")
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := []interface{}{"default value"}
		sut := Config{}

		check, e := sut.List("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case !reflect.DeepEqual(check, value):
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Config_Config(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := Config{"id": "value"}
		sut := Config{path: value}

		if check, e := sut.Config(path, Config{"id": "default value"}); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if !reflect.DeepEqual(*check.(*Config), value) {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an Config", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Config{path: value}

		check, e := sut.Config(path, Config{"id": "default value"})
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Config{}

		check, e := sut.Config("node")
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := Config{"id": "default value"}
		sut := Config{}

		check, e := sut.Config("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case !reflect.DeepEqual(*check.(*Config), value):
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Config_Merge(t *testing.T) {
	t.Run("merges two partials", func(t *testing.T) {
		scenarios := []struct {
			partial1 Config
			partial2 Config
			expected Config
		}{
			{ // _test merging nil Config source
				partial1: Config{},
				partial2: nil,
				expected: Config{},
			},
			{ // _test merging empty Config
				partial1: Config{},
				partial2: Config{},
				expected: Config{},
			},
			{ // _test merging empty Config with a non-empty Config
				partial1: Config{"node1": "value1"},
				partial2: Config{},
				expected: Config{"node1": "value1"},
			},
			{ // _test merging Config into empty Config
				partial1: Config{},
				partial2: Config{"node1": "value1"},
				expected: Config{"node1": "value1"},
			},
			{ // _test merging override source value
				partial1: Config{"node1": "value1"},
				partial2: Config{"node1": "value2"},
				expected: Config{"node1": "value2"},
			},
			{ // _test merging does not override non-present value in merged Config (create)
				partial1: Config{"node1": "value1"},
				partial2: Config{"node2": "value2"},
				expected: Config{"node1": "value1", "node2": "value2"},
			},
			{ // _test merging does not override non-present value in merged Config (override)
				partial1: Config{"node1": "value1", "node2": "value2"},
				partial2: Config{"node2": "value3"},
				expected: Config{"node1": "value1", "node2": "value3"},
			},
			{ // _test merging override source value to a subtree
				partial1: Config{"node1": "value1"},
				partial2: Config{"node1": Config{"node2": "value"}},
				expected: Config{"node1": Config{"node2": "value"}},
			},
			{ // _test merging override source value in a subtree (to a value)
				partial1: Config{"node1": Config{"node2": "value1"}},
				partial2: Config{"node1": Config{"node2": "value2"}},
				expected: Config{"node1": Config{"node2": "value2"}},
			},
			{ // _test merging override source value in a subtree (to a subtree)
				partial1: Config{"node1": Config{"node2": "value"}},
				partial2: Config{"node1": Config{"node2": Config{"node3": "value"}}},
				expected: Config{"node1": Config{"node2": Config{"node3": "value"}}},
			},
		}

		for _, scenario := range scenarios {
			check := scenario.partial1
			check.merge(scenario.partial2)

			if !reflect.DeepEqual(check, scenario.expected) {
				t.Errorf("resulted in (%s) when merging (%v) and (%v), expecting (%v)", check, scenario.partial1, scenario.partial2, scenario.expected)
			}
		}
	})

	t.Run("merging works with copies", func(t *testing.T) {
		data1 := Config{"node1": Config{"node2": "value 2"}}
		data2 := Config{"node1": Config{"node3": Config{"node4": "value 4"}}}
		expected := Config{"node1": Config{"node2": "value 2", "node3": Config{"node4": "value 4"}}}

		check := Config{}
		check.merge(data1)
		check.merge(data2)

		if !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%s) when merging (%v) and (%v), expecting (%v)", check, data1, data2, expected)
		}
	})
}

func Test_Config_Convert(t *testing.T) {
	t.Run("convert float32 into int", func(t *testing.T) {
		data := float32(123)
		expected := 123

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert float64 into int", func(t *testing.T) {
		data := float64(123)
		expected := 123

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map", func(t *testing.T) {
		data := map[string]interface{}{"node": "value"}
		expected := Config{"node": "value"}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert config", func(t *testing.T) {
		data := Config{"node": "value"}
		expected := Config{"node": "value"}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert list", func(t *testing.T) {
		data := []interface{}{1, 2, 3}
		expected := []interface{}{1, 2, 3}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with a float32 into a map with a int", func(t *testing.T) {
		data := map[string]interface{}{"node": float32(123)}
		expected := Config{"node": 123}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with a float64 into a map with a int", func(t *testing.T) {
		data := map[string]interface{}{"node": float64(123)}
		expected := Config{"node": 123}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with another map", func(t *testing.T) {
		data := map[string]interface{}{"node": map[string]interface{}{"node2": "value"}}
		expected := Config{"node": Config{"node2": "value"}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with a list", func(t *testing.T) {
		data := map[string]interface{}{"node": []interface{}{1, 2, 3}}
		expected := Config{"node": []interface{}{1, 2, 3}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert config with a map", func(t *testing.T) {
		data := Config{"node": map[string]interface{}{"node2": "value"}}
		expected := Config{"node": Config{"node2": "value"}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with a list of maps", func(t *testing.T) {
		data := map[string]interface{}{"node": []interface{}{map[string]interface{}{"node2": "value"}}}
		expected := Config{"node": []interface{}{Config{"node2": "value"}}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with a list of configs", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{Config{"node2": "value"}}}
		expected := Config{"node": []interface{}{Config{"node2": "value"}}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert config with numeric fields", func(t *testing.T) {
		data := Config{1: map[string]interface{}{"node2": "value"}}
		expected := Config{1: Config{"node2": "value"}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert config with a uppercase fields", func(t *testing.T) {
		data := Config{"NoDE": map[string]interface{}{"nODE2": "value"}}
		expected := Config{"node": Config{"node2": "value"}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map uppercase fields", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{map[string]interface{}{"NOde2": "value"}}}
		expected := Config{"node": []interface{}{Config{"node2": "value"}}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert config uppercase fields", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{Config{"NOde2": "value"}}}
		expected := Config{"node": []interface{}{Config{"node2": "value"}}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with numeric keys", func(t *testing.T) {
		data := map[interface{}]interface{}{1: []interface{}{Config{2: "value"}}}
		expected := Config{1: []interface{}{Config{2: "value"}}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map with string keys", func(t *testing.T) {
		data := map[interface{}]interface{}{"NoDE1": []interface{}{Config{2: "value"}}}
		expected := Config{"node1": []interface{}{Config{2: "value"}}}

		if check := (Config{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})
}

func Test_Config_Populate(t *testing.T) {
	t.Run("error if path not found", func(t *testing.T) {
		data := Config{"field1": 123, "field2": 456}
		path := "field3"
		target := 0

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on populating an invalid type", func(t *testing.T) {
		data := Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}}
		path := "field1"
		target := struct{ Field1 string }{}

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on populating an invalid type struct field", func(t *testing.T) {
		data := Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}}
		path := "field2"
		target := struct{ Field1 string }{}

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on populating an inner invalid type struct field", func(t *testing.T) {
		data := Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}}
		path := ""
		target := struct{ Field2 struct{ Field1 string } }{}

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("no-op if field is not found in config", func(t *testing.T) {
		data := Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}}
		path := ""
		target := struct{ Field3 int }{Field3: 123}
		expValue := struct{ Field3 int }{Field3: 0}

		v, e := data.Populate(path, target)
		switch {
		case !reflect.DeepEqual(v, expValue):
			t.Errorf("didn't returned the (%v) expected value : %v", expValue, v)
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("no-op if field is not found in inner config", func(t *testing.T) {
		data := Config{"field1": 123}
		path := ""
		target := struct {
			Field1 int
			Field2 struct{ Field3 int }
		}{}
		expValue := struct {
			Field1 int
			Field2 struct{ Field3 int }
		}{Field1: 123}

		v, e := data.Populate(path, target)
		switch {
		case !reflect.DeepEqual(v, expValue):
			t.Errorf("didn't returned the (%v) expected value : %v", expValue, v)
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("populate scalar values", func(t *testing.T) {
		scenarios := []struct {
			data     Config
			path     string
			target   interface{}
			expValue interface{}
		}{
			{ // populate an integer
				data:     Config{"field1": 123, "field2": 456},
				path:     "field2",
				target:   0,
				expValue: 456,
			},
			{ // populate an integer from inner field
				data:     Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}},
				path:     "field2.field2",
				target:   0,
				expValue: 456,
			},
			{ // populate a string from inner field
				data:     Config{"field1": 123, "field2": Config{"field1": 123, "field2": "test string"}},
				path:     "field2.field2",
				target:   "",
				expValue: "test string",
			},
		}

		for _, scenario := range scenarios {
			v, e := scenario.data.Populate(scenario.path, scenario.target)
			switch {
			case !reflect.DeepEqual(v, scenario.expValue):
				t.Errorf("didn't returned the (%v) expected value : %v", scenario.expValue, v)
			case e != nil:
				t.Errorf("returned the unexpected (%v) error", e)
			}
		}
	})

	t.Run("populate structure references", func(t *testing.T) {
		scenarios := []struct {
			data     Config
			path     string
			target   interface{}
			expValue interface{}
		}{
			{ // populate a single exposed field structure
				data:     Config{"field1": 123, "field2": 456},
				path:     "",
				target:   struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // populate a single exposed field structure from inner field
				data:     Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}},
				path:     "field2",
				target:   struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // populate a multiple exposed field structure
				data:     Config{"field1": 123, "field2": 456},
				path:     "",
				target:   struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple exposed field structure from inner field
				data:     Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}},
				path:     "Field2",
				target:   struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple level structure
				data: Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}},
				path: "",
				target: struct {
					Field1 int
					Field2 struct{ Field1, Field2 int }
				}{},
				expValue: struct {
					Field1 int
					Field2 struct{ Field1, Field2 int }
				}{Field1: 123, Field2: struct{ Field1, Field2 int }{Field1: 123, Field2: 456}},
			},
		}

		for _, scenario := range scenarios {
			v, e := scenario.data.Populate(scenario.path, scenario.target)
			switch {
			case !reflect.DeepEqual(v, scenario.expValue):
				t.Errorf("didn't returned the (%v) expected value : %v", v, scenario.expValue)
			case e != nil:
				t.Errorf("returned the unexpected (%v) error", e)
			}
		}
	})

	t.Run("populate structure pointers", func(t *testing.T) {
		scenarios := []struct {
			data     Config
			path     string
			target   interface{}
			expValue interface{}
		}{
			{ // populate a single exposed field structure
				data:     Config{"field1": 123, "field2": 456},
				path:     "",
				target:   &struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // no-op if field is not in config
				data:     Config{"field1": 123, "field2": 456},
				path:     "",
				target:   &struct{ field1, Field2, Field3 int }{Field3: 789},
				expValue: struct{ field1, Field2, Field3 int }{field1: 0, Field2: 456, Field3: 789},
			},
			{ // populate a single exposed field structure from inner field
				data:     Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}},
				path:     "field2",
				target:   &struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // populate a multiple exposed field structure
				data:     Config{"field1": 123, "field2": 456},
				path:     "",
				target:   &struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple exposed field structure from inner field
				data:     Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}},
				path:     "Field2",
				target:   &struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple level structure
				data: Config{"field1": 123, "field2": Config{"field1": 123, "field2": 456}},
				path: "",
				target: &struct {
					Field1 int
					Field2 struct{ Field1, Field2 int }
				}{},
				expValue: struct {
					Field1 int
					Field2 struct{ Field1, Field2 int }
				}{Field1: 123, Field2: struct{ Field1, Field2 int }{Field1: 123, Field2: 456}},
			},
		}

		for _, scenario := range scenarios {
			v, e := scenario.data.Populate(scenario.path, scenario.target)
			switch {
			case !reflect.DeepEqual(v, scenario.expValue):
				t.Errorf("didn't returned the (%v) expected value : %v", scenario.expValue, v)
			case e != nil:
				t.Errorf("returned the unexpected (%v) error", e)
			}
		}
	})
}
