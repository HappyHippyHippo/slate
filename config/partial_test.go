package config

import (
	"errors"
	"github.com/happyhippyhippo/slate/err"
	"reflect"
	"testing"
)

func Test_Partial_Clone(t *testing.T) {
	t.Run("clone empty partial", func(t *testing.T) {
		sut := Partial{}
		c := sut.Clone()
		sut["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if len(c) != 0 {
			t.Errorf("cloned partial is not empty : %v", c)
		}
	})

	t.Run("clone non-empty partial", func(t *testing.T) {
		sut := Partial{"field": "value"}
		expected := Partial{"field": "value"}
		c := sut.Clone()
		sut["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned partial (%v) is not the expected : %v", c, expected)
		}
	})

	t.Run("recursive cloning", func(t *testing.T) {
		sut := Partial{"field": Partial{"field": "value"}}
		expected := Partial{"field": Partial{"field": "value"}}
		c := sut.Clone()
		sut["extra"] = "value"
		sut["field"].(Partial)["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned partial (%v) is not the expected : %v", c, expected)
		}
	})

	t.Run("recursive cloning with lists", func(t *testing.T) {
		sut := Partial{"field": []interface{}{Partial{"field": "value"}}}
		expected := Partial{"field": []interface{}{Partial{"field": "value"}}}
		c := sut.Clone()
		sut["extra"] = "value"
		sut["field"].([]interface{})[0].(Partial)["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned partial (%v) is not the expected : %v", c, expected)
		}
	})

	t.Run("recursive cloning with multi-level lists", func(t *testing.T) {
		sut := Partial{"field": []interface{}{[]interface{}{Partial{"field": "value"}}}}
		expected := Partial{"field": []interface{}{[]interface{}{Partial{"field": "value"}}}}
		c := sut.Clone()
		sut["extra"] = "value"
		sut["field"].([]interface{})[0].([]interface{})[0].(Partial)["extra"] = "value"

		if c == nil {
			t.Error("clone call didn't returned a valid reference")
		} else if !reflect.DeepEqual(c, expected) {
			t.Errorf("cloned partial (%v) is not the expected : %v", c, expected)
		}
	})
}

func Test_Partial_Has(t *testing.T) {
	t.Run("check if a valid path exists", func(t *testing.T) {
		scenarios := []struct {
			partial Partial
			search  string
		}{
			{ // _test empty Partial, search for everything
				partial: Partial{},
				search:  "",
			},
			{ // _test single node, search for root node
				partial: Partial{"node": "value"},
				search:  "",
			},
			{ // _test single node search
				partial: Partial{"node": "value"},
				search:  "node",
			},
			{ // _test multiple node, search for root node
				partial: Partial{"node1": "value", "node2": "value"},
				search:  "",
			},
			{ // _test multiple node search for first
				partial: Partial{"node1": "value", "node2": "value"},
				search:  "node1",
			},
			{ // _test multiple node search for non-first
				partial: Partial{"node1": "value", "node2": "value"},
				search:  "node2",
			},
			{ // _test tree, search for root node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "",
			},
			{ // _test tree, search for root level node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "node1",
			},
			{ // _test tree, search for sub node
				partial: Partial{"node1": Partial{"node2": "value"}},
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
			partial Partial
			search  string
		}{
			{ // _test single node search (invalid)
				partial: Partial{"node": "value"},
				search:  "node2",
			},
			{ // _test multiple node search for invalid node
				partial: Partial{"node1": "value", "node2": "value"},
				search:  "node3",
			},
			{ // _test tree search for invalid root node
				partial: Partial{"node": Partial{"node": "value"}},
				search:  "node1",
			},
			{ // _test tree search for invalid sub node
				partial: Partial{"node": Partial{"node": "value"}},
				search:  "node.node1",
			},
			{ // _test tree search for invalid sub-sub-node
				partial: Partial{"node": Partial{"node": "value"}},
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

func Test_Partial_Get(t *testing.T) {
	t.Run("retrieve a value of a existent path", func(t *testing.T) {
		scenarios := []struct {
			partial  Partial
			search   string
			expected interface{}
		}{
			{ // _test empty Partial, search for everything
				partial:  Partial{},
				search:   "",
				expected: Partial{},
			},
			{ // _test single node, search for root node
				partial:  Partial{"node": "value"},
				search:   "",
				expected: Partial{"node": "value"},
			},
			{ // _test single node search
				partial:  Partial{"node": "value"},
				search:   "node",
				expected: "value",
			},
			{ // _test multiple node, search for root node
				partial:  Partial{"node1": "value1", "node2": "value2"},
				search:   "",
				expected: Partial{"node1": "value1", "node2": "value2"},
			},
			{ // _test multiple node search for first
				partial:  Partial{"node1": "value1", "node2": "value2"},
				search:   "node1",
				expected: "value1",
			},
			{ // _test multiple node search for non-first
				partial:  Partial{"node1": "value1", "node2": "value2"},
				search:   "node2",
				expected: "value2",
			},
			{ // _test tree, search for root node
				partial:  Partial{"node": Partial{"node": "value"}},
				search:   "",
				expected: Partial{"node": Partial{"node": "value"}},
			},
			{ // _test tree, search for root level node
				partial:  Partial{"node": Partial{"node": "value"}},
				search:   "node",
				expected: Partial{"node": "value"},
			},
			{ // _test tree, search for sub node
				partial:  Partial{"node": Partial{"node": "value"}},
				search:   "node.node",
				expected: "value",
			},
		}

		for _, scenario := range scenarios {
			if check, e := scenario.partial.Get(scenario.search); e != nil {
				t.Errorf("returned the unexpected error (%v) when retrieving (%v)", e, scenario.search)
			} else if !reflect.DeepEqual(check, scenario.expected) {
				t.Errorf("returned (%v) when retrieving (%v), expected (%v)", check, scenario.search, scenario.expected)
			}
		}
	})

	t.Run("return nil if a path don't exists", func(t *testing.T) {
		scenarios := []struct {
			partial Partial
			search  string
		}{
			{ // _test empty Partial search for non-existent node
				partial: Partial{},
				search:  "node",
			},
			{ // _test single node search for non-existent node
				partial: Partial{"node": "value"},
				search:  "node2",
			},
			{ // _test multiple node search for non-existent node
				partial: Partial{"node1": "value1", "node2": "value2"},
				search:  "node3",
			},
			{ // _test tree search for non-existent root node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "node2",
			},
			{ // _test tree search for non-existent sub node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "node1.node1",
			},
			{ // _test tree search for non-existent sub-sub-node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "node1.node2.node3",
			},
		}

		for _, scenario := range scenarios {
			check, e := scenario.partial.Get(scenario.search)
			switch {
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, err.ErrConfigPathNotFound):
				t.Errorf("returned error was not a config path not found error : %v", e)
			case check != nil:
				t.Error("unexpectedly returned a valid reference to a stored config value")
			}
		}
	})

	t.Run("return nil if the node actually stores nil", func(t *testing.T) {
		sut := Partial{"node1": nil, "node2": "value2"}

		if check, e := sut.Get("node1", "default_value"); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if check != nil {
			t.Errorf("returned the (%v) check", check)
		}
	})

	t.Run("return the default value if a path don't exists", func(t *testing.T) {
		scenarios := []struct {
			partial Partial
			search  string
		}{
			{ // _test empty Partial search for non-existent node
				partial: Partial{},
				search:  "node",
			},
			{ // _test single node search for non-existent node
				partial: Partial{"node": "value"},
				search:  "node2",
			},
			{ // _test multiple node search for non-existent node
				partial: Partial{"node1": "value1", "node2": "value2"},
				search:  "node3",
			},
			{ // _test tree search for non-existent root node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "node2",
			},
			{ // _test tree search for non-existent sub node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "node1.node1",
			},
			{ // _test tree search for non-existent sub-sub-node
				partial: Partial{"node1": Partial{"node2": "value"}},
				search:  "node1.node2.node3",
			},
		}

		def := "default_value"
		for _, scenario := range scenarios {
			if check, e := scenario.partial.Get(scenario.search, def); e != nil {
				t.Errorf("returned the unexpected error : %v", e)
			} else if check != def {
				t.Errorf("returned (%v) when retrieving (%v)", check, scenario.search)
			}
		}
	})
}

func Test_Partial_Bool(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		sut := Partial{path: true}

		if check, e := sut.Bool(path, false); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if !check {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, true)
		}
	})

	t.Run("return conversion error if not storing a bool", func(t *testing.T) {
		path := "node"
		value := "123"
		sut := Partial{path: value}

		check, e := sut.Bool(path, true)
		switch {
		case check:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("the returned error is the expected error convertion error : %v", e)
		}
	})

	t.Run("return path not found error if no default value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Bool("node")
		switch {
		case check:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConfigPathNotFound):
			t.Errorf("the returned error is the expected error convertion error : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Bool("node", true)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		case !check:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, true)
		}
	})
}

func Test_Partial_Int(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Partial{path: value}

		if check, e := sut.Int(path, 456); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if check != value {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion error if not storing an int", func(t *testing.T) {
		path := "node"
		value := "123"
		sut := Partial{path: value}

		check, e := sut.Int(path, 456)
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("the returned error is the expected error convertion error : %v", e)
		}
	})

	t.Run("return path not found error if no default value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Int("node")
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := 123
		sut := Partial{}

		check, e := sut.Int("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case check != value:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Partial_Float(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := 123.456
		sut := Partial{path: value}

		if check, e := sut.Float(path, 456.789); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if check != value {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an float", func(t *testing.T) {
		path := "node"
		value := "123.456"
		sut := Partial{path: value}

		check, e := sut.Float(path, 456)
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Float("node")
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := 123.456
		sut := Partial{}

		check, e := sut.Float("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case check != value:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Partial_String(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := "value"
		sut := Partial{path: value}

		if check, e := sut.String(path, "default value"); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if check != value {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an string", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Partial{path: value}

		check, e := sut.String(path, "default value")
		switch {
		case check != "":
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.String("node")
		switch {
		case check != "":
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := "default value"
		sut := Partial{}

		check, e := sut.String("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case check != value:
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Partial_List(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := []interface{}{"value"}
		sut := Partial{path: value}

		if check, e := sut.List(path, []interface{}{"default value"}); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if !reflect.DeepEqual(check, value) {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an list", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Partial{path: value}

		check, e := sut.List(path, []interface{}{"default value"})
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.List("node")
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := []interface{}{"default value"}
		sut := Partial{}

		check, e := sut.List("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case !reflect.DeepEqual(check, value):
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Partial_Partial(t *testing.T) {
	t.Run("return valid stored value", func(t *testing.T) {
		path := "node"
		value := Partial{"id": "value"}
		sut := Partial{path: value}

		if check, e := sut.Partial(path, Partial{"id": "default value"}); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if !reflect.DeepEqual(check, value) {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an Partial", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Partial{path: value}

		check, e := sut.Partial(path, Partial{"id": "default value"})
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no default value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Partial("node")
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrConfigPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return default value if the path don't exists", func(t *testing.T) {
		value := Partial{"id": "default value"}
		sut := Partial{}

		check, e := sut.Partial("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case !reflect.DeepEqual(check, value):
			t.Errorf("returned the unexpected value (%v) when expecting : %v", check, value)
		}
	})
}

func Test_Partial_Merge(t *testing.T) {
	t.Run("merges two partials", func(t *testing.T) {
		scenarios := []struct {
			partial1 Partial
			partial2 Partial
			expected Partial
		}{
			{ // _test merging nil Partial source
				partial1: Partial{},
				partial2: nil,
				expected: Partial{},
			},
			{ // _test merging empty Partial
				partial1: Partial{},
				partial2: Partial{},
				expected: Partial{},
			},
			{ // _test merging empty Partial with a non-empty Partial
				partial1: Partial{"node1": "value1"},
				partial2: Partial{},
				expected: Partial{"node1": "value1"},
			},
			{ // _test merging Partial into empty Partial
				partial1: Partial{},
				partial2: Partial{"node1": "value1"},
				expected: Partial{"node1": "value1"},
			},
			{ // _test merging override source value
				partial1: Partial{"node1": "value1"},
				partial2: Partial{"node1": "value2"},
				expected: Partial{"node1": "value2"},
			},
			{ // _test merging does not override non-present value in merged Partial (create)
				partial1: Partial{"node1": "value1"},
				partial2: Partial{"node2": "value2"},
				expected: Partial{"node1": "value1", "node2": "value2"},
			},
			{ // _test merging does not override non-present value in merged Partial (override)
				partial1: Partial{"node1": "value1", "node2": "value2"},
				partial2: Partial{"node2": "value3"},
				expected: Partial{"node1": "value1", "node2": "value3"},
			},
			{ // _test merging override source value to a subtree
				partial1: Partial{"node1": "value1"},
				partial2: Partial{"node1": Partial{"node2": "value"}},
				expected: Partial{"node1": Partial{"node2": "value"}},
			},
			{ // _test merging override source value in a subtree (to a value)
				partial1: Partial{"node1": Partial{"node2": "value1"}},
				partial2: Partial{"node1": Partial{"node2": "value2"}},
				expected: Partial{"node1": Partial{"node2": "value2"}},
			},
			{ // _test merging override source value in a subtree (to a subtree)
				partial1: Partial{"node1": Partial{"node2": "value"}},
				partial2: Partial{"node1": Partial{"node2": Partial{"node3": "value"}}},
				expected: Partial{"node1": Partial{"node2": Partial{"node3": "value"}}},
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
		data1 := Partial{"node1": Partial{"node2": "value 2"}}
		data2 := Partial{"node1": Partial{"node3": Partial{"node4": "value 4"}}}
		expected := Partial{"node1": Partial{"node2": "value 2", "node3": Partial{"node4": "value 4"}}}

		check := Partial{}
		check.merge(data1)
		check.merge(data2)

		if !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%s) when merging (%v) and (%v), expecting (%v)", check, data1, data2, expected)
		}
	})
}

func Test_Partial_Convert(t *testing.T) {
	t.Run("convert float32 into int", func(t *testing.T) {
		data := float32(123)
		expected := 123

		if check := (Partial{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert float64 into int", func(t *testing.T) {
		data := float64(123)
		expected := 123

		if check := (Partial{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert map", func(t *testing.T) {
		data := map[string]interface{}{"node": "value"}
		expected := Partial{"node": "value"}

		if check := (Partial{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert recurring float32 into int", func(t *testing.T) {
		data := map[string]interface{}{"node": float32(123)}
		expected := Partial{"node": 123}

		if check := (Partial{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert recurring float64 into int", func(t *testing.T) {
		data := map[string]interface{}{"node": float64(123)}
		expected := Partial{"node": 123}

		if check := (Partial{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("convert recurring map", func(t *testing.T) {
		data := map[string]interface{}{"node": map[string]interface{}{"node2": "value"}}
		expected := Partial{"node": Partial{"node2": "value"}}

		if check := (Partial{}).convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})
}
