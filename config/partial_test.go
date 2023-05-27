package config

import (
	"errors"
	"reflect"
	"sort"
	"testing"

	"github.com/happyhippyhippo/slate"
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

func Test_Partial_Entries(t *testing.T) {
	t.Run("empty partial", func(t *testing.T) {
		if (&Partial{}).Entries() != nil {
			t.Errorf("didn't returned the expected empty list")
		}
	})

	t.Run("single entry partial", func(t *testing.T) {
		if !reflect.DeepEqual((&Partial{"field": "value"}).Entries(), []string{"field"}) {
			t.Errorf("didn't returned the expected single entry list")
		}
	})

	t.Run("multi entry partial", func(t *testing.T) {
		check := (&Partial{
			"field1": "value 1",
			"field2": "value 2",
		}).Entries()
		expected := []string{"field1", "field2"}

		sort.Strings(check)
		sort.Strings(expected)
		if !reflect.DeepEqual(check, expected) {
			t.Errorf("didn't returned the expected single entry list")
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
			case !errors.Is(e, ErrPathNotFound):
				t.Errorf("returned error was not a partial path not found error : %v", e)
			case check != nil:
				t.Error("unexpectedly returned a valid reference to a stored partial value")
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

	t.Run("return the simple value if a path don't exists", func(t *testing.T) {
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
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("the returned error is the expected error convertion error : %v", e)
		}
	})

	t.Run("return path not found error if no simple value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Bool("node")
		switch {
		case check:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrPathNotFound):
			t.Errorf("the returned error is the expected error convertion error : %v", e)
		}
	})

	t.Run("return simple value if the path don't exists", func(t *testing.T) {
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
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("the returned error is the expected error convertion error : %v", e)
		}
	})

	t.Run("return path not found error if no simple value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Int("node")
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return simple value if the path don't exists", func(t *testing.T) {
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
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no simple value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Float("node")
		switch {
		case check != 0:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return simple value if the path don't exists", func(t *testing.T) {
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

		if check, e := sut.String(path, "simple value"); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if check != value {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an string", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Partial{path: value}

		check, e := sut.String(path, "simple value")
		switch {
		case check != "":
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no simple value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.String("node")
		switch {
		case check != "":
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return simple value if the path don't exists", func(t *testing.T) {
		value := "simple value"
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

		if check, e := sut.List(path, []interface{}{"simple value"}); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if !reflect.DeepEqual(check, value) {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an list", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Partial{path: value}

		check, e := sut.List(path, []interface{}{"simple value"})
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no simple value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.List("node")
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return simple value if the path don't exists", func(t *testing.T) {
		value := []interface{}{"simple value"}
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

		if check, e := sut.Partial(path, Partial{"id": "simple value"}); e != nil {
			t.Errorf("returned the unexpected e : %v", e)
		} else if !reflect.DeepEqual(*check, value) {
			t.Errorf("returned the unexpected value of (%v) when expecting : %v", check, value)
		}
	})

	t.Run("return conversion e if not storing an Partial", func(t *testing.T) {
		path := "node"
		value := 123
		sut := Partial{path: value}

		check, e := sut.Partial(path, Partial{"id": "simple value"})
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return path not found e if no simple value is given", func(t *testing.T) {
		sut := Partial{}

		check, e := sut.Partial("node")
		switch {
		case check != nil:
			t.Errorf("returned the unexpected value : %v", check)
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrPathNotFound):
			t.Errorf("the returned e is the expected e convertion e : %v", e)
		}
	})

	t.Run("return simple value if the path don't exists", func(t *testing.T) {
		value := Partial{"id": "simple value"}
		sut := Partial{}

		check, e := sut.Partial("node", value)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case !reflect.DeepEqual(*check, value):
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
			check.Merge(scenario.partial2)

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
		check.Merge(data1)
		check.Merge(data2)

		if !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%s) when merging (%v) and (%v), expecting (%v)", check, data1, data2, expected)
		}
	})
}

func Test_Partial_Populate(t *testing.T) {
	t.Run("error if path not found", func(t *testing.T) {
		data := Partial{"field1": 123, "field2": 456}
		path := "field3"
		target := 0

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrPathNotFound):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on populating an invalid type", func(t *testing.T) {
		data := Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}}
		path := "field1"
		target := struct{ Field1 string }{}

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on populating an invalid type struct field", func(t *testing.T) {
		data := Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}}
		path := "field2"
		target := struct{ Field1 string }{}

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on populating an inner invalid type struct field", func(t *testing.T) {
		data := Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}}
		path := ""
		target := struct{ Field2 struct{ Field1 string } }{}

		v, e := data.Populate(path, target)
		switch {
		case v != nil:
			t.Error("returned an unexpected valid reference to a data")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("no-op if field is not found in partial", func(t *testing.T) {
		data := Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}}
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

	t.Run("no-op if field is not found in inner partial", func(t *testing.T) {
		data := Partial{"field1": 123}
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
			data     Partial
			path     string
			target   interface{}
			expValue interface{}
		}{
			{ // populate an integer
				data:     Partial{"field1": 123, "field2": 456},
				path:     "field2",
				target:   0,
				expValue: 456,
			},
			{ // populate an integer from inner field
				data:     Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}},
				path:     "field2.field2",
				target:   0,
				expValue: 456,
			},
			{ // populate a string from inner field
				data:     Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": "test string"}},
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
			data     Partial
			path     string
			target   interface{}
			expValue interface{}
		}{
			{ // populate a single exposed field structure
				data:     Partial{"field1": 123, "field2": 456},
				path:     "",
				target:   struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // populate a single exposed field structure from inner field
				data:     Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}},
				path:     "field2",
				target:   struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // populate a multiple exposed field structure
				data:     Partial{"field1": 123, "field2": 456},
				path:     "",
				target:   struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple exposed field structure from inner field
				data:     Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}},
				path:     "Field2",
				target:   struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple level structure
				data: Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}},
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
			data     Partial
			path     string
			target   interface{}
			expValue interface{}
		}{
			{ // populate a single exposed field structure
				data:     Partial{"field1": 123, "field2": 456},
				path:     "",
				target:   &struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // no-op if field is not in partial
				data:     Partial{"field1": 123, "field2": 456},
				path:     "",
				target:   &struct{ field1, Field2, Field3 int }{Field3: 789},
				expValue: struct{ field1, Field2, Field3 int }{field1: 0, Field2: 456, Field3: 789},
			},
			{ // populate a single exposed field structure from inner field
				data:     Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}},
				path:     "field2",
				target:   &struct{ field1, Field2 int }{},
				expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
			},
			{ // populate a multiple exposed field structure
				data:     Partial{"field1": 123, "field2": 456},
				path:     "",
				target:   &struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple exposed field structure from inner field
				data:     Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}},
				path:     "Field2",
				target:   &struct{ Field1, Field2 int }{},
				expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
			},
			{ // populate a multiple level structure
				data: Partial{"field1": 123, "field2": Partial{"field1": 123, "field2": 456}},
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
