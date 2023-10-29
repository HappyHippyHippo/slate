package slate

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_config_err(t *testing.T) {
	t.Run("errInvalidEmptyConfigPath", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : invalid empty config path"

		t.Run("creation without context", func(t *testing.T) {
			if e := errInvalidEmptyConfigPath(arg); !errors.Is(e, ErrInvalidEmptyConfigPath) {
				t.Errorf("error not a instance of ErrInvalidEmptyPath")
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
			if e := errInvalidEmptyConfigPath(arg, context); !errors.Is(e, ErrInvalidEmptyConfigPath) {
				t.Errorf("error not a instance of ErrInvalidEmptyPath")
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

	t.Run("errConfigPathNotFound", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : config path not found"

		t.Run("creation without context", func(t *testing.T) {
			if e := errConfigPathNotFound(arg); !errors.Is(e, ErrConfigPathNotFound) {
				t.Errorf("error not a instance of ErrPathNotFound")
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
			if e := errConfigPathNotFound(arg, context); !errors.Is(e, ErrConfigPathNotFound) {
				t.Errorf("error not a instance of ErrPathNotFound")
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

	t.Run("errInvalidConfigFormat", func(t *testing.T) {
		arg := ConfigFormatJSON
		context := map[string]interface{}{"field": "value"}
		message := "json : invalid config format"

		t.Run("creation without context", func(t *testing.T) {
			if e := errInvalidConfigFormat(arg); !errors.Is(e, ErrInvalidConfigFormat) {
				t.Errorf("error not a instance of ErrInvalidFormat")
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
			if e := errInvalidConfigFormat(arg, context); !errors.Is(e, ErrInvalidConfigFormat) {
				t.Errorf("error not a instance of ErrInvalidFormat")
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

	t.Run("errInvalidConfigSupplier", func(t *testing.T) {
		arg := ConfigPartial{"field": "value"}
		context := map[string]interface{}{"field": "value"}
		message := "map[field:value] : invalid config supplier"

		t.Run("creation without context", func(t *testing.T) {
			if e := errInvalidConfigSupplier(arg); !errors.Is(e, ErrInvalidConfigSupplier) {
				t.Errorf("error not a instance of ErrInvalidSource")
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
			if e := errInvalidConfigSupplier(arg, context); !errors.Is(e, ErrInvalidConfigSupplier) {
				t.Errorf("error not a instance of ErrInvalidSource")
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

	t.Run("errConfigSupplierNotFound", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : config supplier not found"

		t.Run("creation without context", func(t *testing.T) {
			if e := errConfigSupplierNotFound(arg); !errors.Is(e, ErrConfigSupplierNotFound) {
				t.Errorf("error not a instance of ErrSourceNotFound")
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
			if e := errConfigSupplierNotFound(arg, context); !errors.Is(e, ErrConfigSupplierNotFound) {
				t.Errorf("error not a instance of ErrSourceNotFound")
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

	t.Run("errDuplicateConfigSupplier", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : config supplier already registered"

		t.Run("creation without context", func(t *testing.T) {
			if e := errDuplicateConfigSupplier(arg); !errors.Is(e, ErrDuplicateConfigSupplier) {
				t.Errorf("error not a instance of ErrDuplicateSource")
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
			if e := errDuplicateConfigSupplier(arg, context); !errors.Is(e, ErrDuplicateConfigSupplier) {
				t.Errorf("error not a instance of ErrDuplicateSource")
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

func Test_ConfigPartial(t *testing.T) {
	t.Run("Clone", func(t *testing.T) {
		t.Run("clone empty partial", func(t *testing.T) {
			sut := ConfigPartial{}
			c := sut.Clone()
			sut["extra"] = "value"

			if c == nil {
				t.Error("didn't returned a valid reference")
			} else if len(c) != 0 {
				t.Errorf("cloned partial is not empty : %v", c)
			}
		})

		t.Run("clone non-empty partial", func(t *testing.T) {
			sut := ConfigPartial{"field": "value"}
			expected := ConfigPartial{"field": "value"}
			c := sut.Clone()
			sut["extra"] = "value"

			if c == nil {
				t.Error("didn't returned a valid reference")
			} else if !reflect.DeepEqual(c, expected) {
				t.Errorf("cloned (%v) is not the expected : %v", c, expected)
			}
		})

		t.Run("recursive cloning", func(t *testing.T) {
			sut := ConfigPartial{"field": ConfigPartial{"field": "value"}}
			expected := ConfigPartial{"field": ConfigPartial{"field": "value"}}
			c := sut.Clone()
			sut["extra"] = "value"
			sut["field"].(ConfigPartial)["extra"] = "value"

			if c == nil {
				t.Error("didn't returned a valid reference")
			} else if !reflect.DeepEqual(c, expected) {
				t.Errorf("cloned (%v) is not the expected : %v", c, expected)
			}
		})

		t.Run("recursive cloning with lists", func(t *testing.T) {
			sut := ConfigPartial{"field": []interface{}{ConfigPartial{"field": "value"}}}
			expected := ConfigPartial{"field": []interface{}{ConfigPartial{"field": "value"}}}
			c := sut.Clone()
			sut["extra"] = "value"
			sut["field"].([]interface{})[0].(ConfigPartial)["extra"] = "value"

			if c == nil {
				t.Error("didn't returned a valid reference")
			} else if !reflect.DeepEqual(c, expected) {
				t.Errorf("cloned (%v) is not the expected : %v", c, expected)
			}
		})

		t.Run("recursive cloning with multi-level lists", func(t *testing.T) {
			sut := ConfigPartial{"field": []interface{}{[]interface{}{ConfigPartial{"field": "value"}}}}
			expected := ConfigPartial{"field": []interface{}{[]interface{}{ConfigPartial{"field": "value"}}}}
			c := sut.Clone()
			sut["extra"] = "value"
			sut["field"].([]interface{})[0].([]interface{})[0].(ConfigPartial)["extra"] = "value"

			if c == nil {
				t.Error("didn't returned a valid reference")
			} else if !reflect.DeepEqual(c, expected) {
				t.Errorf("cloned (%v) is not the expected : %v", c, expected)
			}
		})
	})

	t.Run("Entries", func(t *testing.T) {
		t.Run("empty partial", func(t *testing.T) {
			if (&ConfigPartial{}).Entries() != nil {
				t.Errorf("didn't returned the expected empty list")
			}
		})

		t.Run("single entry partial", func(t *testing.T) {
			if !reflect.DeepEqual((&ConfigPartial{"field": "value"}).Entries(), []string{"field"}) {
				t.Errorf("didn't returned the expected single entry list")
			}
		})

		t.Run("multi entry partial", func(t *testing.T) {
			check := (&ConfigPartial{
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
	})

	t.Run("Has", func(t *testing.T) {
		t.Run("check if a valid path exists", func(t *testing.T) {
			scenarios := []struct {
				partial ConfigPartial
				search  string
			}{
				{ // _test empty ConfigPartial, search for everything
					partial: ConfigPartial{},
					search:  "",
				},
				{ // _test single node, search for root node
					partial: ConfigPartial{"node": "value"},
					search:  "",
				},
				{ // _test single node search
					partial: ConfigPartial{"node": "value"},
					search:  "node",
				},
				{ // _test multiple node, search for root node
					partial: ConfigPartial{"node1": "value", "node2": "value"},
					search:  "",
				},
				{ // _test multiple node search for first
					partial: ConfigPartial{"node1": "value", "node2": "value"},
					search:  "node1",
				},
				{ // _test multiple node search for non-first
					partial: ConfigPartial{"node1": "value", "node2": "value"},
					search:  "node2",
				},
				{ // _test tree, search for root node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "",
				},
				{ // _test tree, search for root level node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node1",
				},
				{ // _test tree, search for sub node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node1.node2",
				},
			}

			for _, s := range scenarios {
				if check := s.partial.Has(s.search); !check {
					t.Errorf("didn't found the (%s) path in (%v)", s.search, s.partial)
				}
			}
		})

		t.Run("check if a invalid path do not exists", func(t *testing.T) {
			scenarios := []struct {
				partial ConfigPartial
				search  string
			}{
				{ // _test single node search (invalid)
					partial: ConfigPartial{"node": "value"},
					search:  "node2",
				},
				{ // _test multiple node search for invalid node
					partial: ConfigPartial{"node1": "value", "node2": "value"},
					search:  "node3",
				},
				{ // _test tree search for invalid root node
					partial: ConfigPartial{"node": ConfigPartial{"node": "value"}},
					search:  "node1",
				},
				{ // _test tree search for invalid sub node
					partial: ConfigPartial{"node": ConfigPartial{"node": "value"}},
					search:  "node.node1",
				},
				{ // _test tree search for invalid sub-sub-node
					partial: ConfigPartial{"node": ConfigPartial{"node": "value"}},
					search:  "node.node.node",
				},
			}

			for _, s := range scenarios {
				if check := s.partial.Has(s.search); check {
					t.Errorf("founded the (%s) path in (%v)", s.search, s.partial)
				}
			}
		})
	})

	t.Run("Set", func(t *testing.T) {
		t.Run("error on empty path", func(t *testing.T) {
			if chk, e := (&ConfigPartial{}).Set("", 123); chk != nil {
				t.Error("unexpected valid partial reference")
			} else if !errors.Is(e, ErrInvalidEmptyConfigPath) {
				t.Errorf("returned (%v) when expecting (%e)", e, ErrInvalidEmptyConfigPath)
			}
		})

		t.Run("set", func(t *testing.T) {
			scenarios := []struct {
				partial  ConfigPartial
				path     string
				value    interface{}
				expected ConfigPartial
			}{
				{ // set to existing root partial
					partial:  ConfigPartial{"test": 456},
					path:     "test",
					value:    123,
					expected: ConfigPartial{"test": 123},
				},
				{ // set to non-existing root partial
					partial:  ConfigPartial{},
					path:     "test",
					value:    123,
					expected: ConfigPartial{"test": 123},
				},
				{ // set to existing sub partial and existing value
					partial:  ConfigPartial{"test": ConfigPartial{"node": 456}},
					path:     "test.node",
					value:    123,
					expected: ConfigPartial{"test": ConfigPartial{"node": 123}},
				},
				{ // set to existing sub partial and non-existing value
					partial:  ConfigPartial{"test": ConfigPartial{}},
					path:     "test.node",
					value:    123,
					expected: ConfigPartial{"test": ConfigPartial{"node": 123}},
				},
				{ // set to non-existing sub partial
					partial:  ConfigPartial{},
					path:     "test.node",
					value:    123,
					expected: ConfigPartial{"test": ConfigPartial{"node": 123}},
				},
				{ // set to existing value on path
					partial:  ConfigPartial{"test": 123},
					path:     "test.node",
					value:    123,
					expected: ConfigPartial{"test": ConfigPartial{"node": 123}},
				},
				{ // set to non-existing sub partial with nil path part
					partial:  ConfigPartial{},
					path:     "test...node",
					value:    123,
					expected: ConfigPartial{"test": ConfigPartial{"node": 123}},
				},
			}

			for _, s := range scenarios {
				check, e := s.partial.Set(s.path, s.value)
				switch {
				case e != nil:
					t.Errorf("(%v) error when assigning (%v) to (%v) path", e, s.value, s.path)
				case check != &s.partial:
					t.Error("didn't returned the reference ot the assigned partial")
				case !reflect.DeepEqual(s.partial, s.expected):
					t.Errorf("didn't updated (%v) partial, expected (%v)", s.partial, s.expected)
				}
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("retrieve a value of a existent path", func(t *testing.T) {
			scenarios := []struct {
				partial  ConfigPartial
				search   string
				expected interface{}
			}{
				{ // _test empty ConfigPartial, search for everything
					partial:  ConfigPartial{},
					search:   "",
					expected: ConfigPartial{},
				},
				{ // _test single node, search for root node
					partial:  ConfigPartial{"node": "value"},
					search:   "",
					expected: ConfigPartial{"node": "value"},
				},
				{ // _test single node search
					partial:  ConfigPartial{"node": "value"},
					search:   "node",
					expected: "value",
				},
				{ // _test multiple node, search for root node
					partial:  ConfigPartial{"node1": "value1", "node2": "value2"},
					search:   "",
					expected: ConfigPartial{"node1": "value1", "node2": "value2"},
				},
				{ // _test multiple node search for first
					partial:  ConfigPartial{"node1": "value1", "node2": "value2"},
					search:   "node1",
					expected: "value1",
				},
				{ // _test multiple node search for non-first
					partial:  ConfigPartial{"node1": "value1", "node2": "value2"},
					search:   "node2",
					expected: "value2",
				},
				{ // _test tree, search for root node
					partial:  ConfigPartial{"node": ConfigPartial{"node": "value"}},
					search:   "",
					expected: ConfigPartial{"node": ConfigPartial{"node": "value"}},
				},
				{ // _test tree, search for root level node
					partial:  ConfigPartial{"node": ConfigPartial{"node": "value"}},
					search:   "node",
					expected: ConfigPartial{"node": "value"},
				},
				{ // _test tree, search for sub node
					partial:  ConfigPartial{"node": ConfigPartial{"node": "value"}},
					search:   "node.node",
					expected: "value",
				},
			}

			for _, s := range scenarios {
				if check, e := s.partial.Get(s.search); e != nil {
					t.Errorf("(%v) error when retrieving (%v)", e, s.search)
				} else if !reflect.DeepEqual(check, s.expected) {
					t.Errorf("returned (%v) when retrieving (%v), expected (%v)", check, s.search, s.expected)
				}
			}
		})

		t.Run("return nil if a path don't exists", func(t *testing.T) {
			scenarios := []struct {
				partial ConfigPartial
				search  string
			}{
				{ // _test empty ConfigPartial search for non-existent node
					partial: ConfigPartial{},
					search:  "node",
				},
				{ // _test single node search for non-existent node
					partial: ConfigPartial{"node": "value"},
					search:  "node2",
				},
				{ // _test multiple node search for non-existent node
					partial: ConfigPartial{"node1": "value1", "node2": "value2"},
					search:  "node3",
				},
				{ // _test tree search for non-existent root node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node2",
				},
				{ // _test tree search for non-existent sub node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node1.node1",
				},
				{ // _test tree search for non-existent sub-sub-node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node1.node2.node3",
				},
			}

			for _, s := range scenarios {
				check, e := s.partial.Get(s.search)
				switch {
				case e == nil:
					t.Error("didn't returned the expected error")
				case !errors.Is(e, ErrConfigPathNotFound):
					t.Errorf("not a partial path not found error : %v", e)
				case check != nil:
					t.Error("valid reference to a stored partial value")
				}
			}
		})

		t.Run("return nil if the node actually stores nil", func(t *testing.T) {
			sut := ConfigPartial{"node1": nil, "node2": "value2"}

			if check, e := sut.Get("node1", "default_value"); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != nil {
				t.Errorf("(%v) returned", check)
			}
		})

		t.Run("return the simple value if a path don't exists", func(t *testing.T) {
			scenarios := []struct {
				partial ConfigPartial
				search  string
			}{
				{ // _test empty ConfigPartial search for non-existent node
					partial: ConfigPartial{},
					search:  "node",
				},
				{ // _test single node search for non-existent node
					partial: ConfigPartial{"node": "value"},
					search:  "node2",
				},
				{ // _test multiple node search for non-existent node
					partial: ConfigPartial{"node1": "value1", "node2": "value2"},
					search:  "node3",
				},
				{ // _test tree search for non-existent root node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node2",
				},
				{ // _test tree search for non-existent sub node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node1.node1",
				},
				{ // _test tree search for non-existent sub-sub-node
					partial: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					search:  "node1.node2.node3",
				},
			}

			def := "default_value"
			for _, s := range scenarios {
				if check, e := s.partial.Get(s.search, def); e != nil {
					t.Errorf("unexpected (%v) error", e)
				} else if check != def {
					t.Errorf("returned (%v) when retrieving (%v)", check, s.search)
				}
			}
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("return valid stored value", func(t *testing.T) {
			path := "node"
			sut := ConfigPartial{path: true}

			if check, e := sut.Bool(path, false); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !check {
				t.Errorf("(%v) when expecting : %v", check, true)
			}
		})

		t.Run("return conversion error if not storing a bool", func(t *testing.T) {
			path := "node"
			value := "123"
			sut := ConfigPartial{path: value}

			check, e := sut.Bool(path, true)
			switch {
			case check:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("not convertion error : %v", e)
			}
		})

		t.Run("return path not found error if no simple value is given", func(t *testing.T) {
			sut := ConfigPartial{}

			check, e := sut.Bool("node")
			switch {
			case check:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("not path not found error : %v", e)
			}
		})

		t.Run("return simple value if the path don't exists", func(t *testing.T) {
			sut := ConfigPartial{}

			check, e := sut.Bool("node", true)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !check:
				t.Errorf("(%v) value when expecting : %v", check, true)
			}
		})
	})

	t.Run("Int", func(t *testing.T) {
		t.Run("return valid stored value", func(t *testing.T) {
			path := "node"
			value := 123
			sut := ConfigPartial{path: value}

			if check, e := sut.Int(path, 456); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != value {
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})

		t.Run("return conversion error if not storing an int", func(t *testing.T) {
			path := "node"
			value := "123"
			sut := ConfigPartial{path: value}

			check, e := sut.Int(path, 456)
			switch {
			case check != 0:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("not convertion error : %v", e)
			}
		})

		t.Run("return path not found error if no simple value is given", func(t *testing.T) {
			sut := ConfigPartial{}

			check, e := sut.Int("node")
			switch {
			case check != 0:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("not path not found error : %v", e)
			}
		})

		t.Run("return simple value if the path don't exists", func(t *testing.T) {
			value := 123
			sut := ConfigPartial{}

			check, e := sut.Int("node", value)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case check != value:
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})
	})

	t.Run("Float", func(t *testing.T) {
		t.Run("return valid stored value", func(t *testing.T) {
			path := "node"
			value := 123.456
			sut := ConfigPartial{path: value}

			if check, e := sut.Float(path, 456.789); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != value {
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})

		t.Run("return conversion error if not storing an float", func(t *testing.T) {
			path := "node"
			value := "123.456"
			sut := ConfigPartial{path: value}

			check, e := sut.Float(path, 456)
			switch {
			case check != 0:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("not convertion error : %v", e)
			}
		})

		t.Run("return path not found e if no simple value is given", func(t *testing.T) {
			sut := ConfigPartial{}

			check, e := sut.Float("node")
			switch {
			case check != 0:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("not path not found error : %v", e)
			}
		})

		t.Run("return simple value if the path don't exists", func(t *testing.T) {
			value := 123.456
			sut := ConfigPartial{}

			check, e := sut.Float("node", value)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case check != value:
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("return valid stored value", func(t *testing.T) {
			path := "node"
			value := "value"
			sut := ConfigPartial{path: value}

			if check, e := sut.String(path, "simple value"); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != value {
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})

		t.Run("return conversion error if not storing an string", func(t *testing.T) {
			path := "node"
			value := 123
			sut := ConfigPartial{path: value}

			check, e := sut.String(path, "simple value")
			switch {
			case check != "":
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("not convertion error : %v", e)
			}
		})

		t.Run("return path not found e if no simple value is given", func(t *testing.T) {
			sut := ConfigPartial{}

			check, e := sut.String("node")
			switch {
			case check != "":
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("not path not found error : %v", e)
			}
		})

		t.Run("return simple value if the path don't exists", func(t *testing.T) {
			value := "simple value"
			sut := ConfigPartial{}

			check, e := sut.String("node", value)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case check != value:
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("return valid stored value", func(t *testing.T) {
			path := "node"
			value := []interface{}{"value"}
			sut := ConfigPartial{path: value}

			if check, e := sut.List(path, []interface{}{"simple value"}); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(check, value) {
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})

		t.Run("return conversion error if not storing an list", func(t *testing.T) {
			path := "node"
			value := 123
			sut := ConfigPartial{path: value}

			check, e := sut.List(path, []interface{}{"simple value"})
			switch {
			case check != nil:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("not convertion error : %v", e)
			}
		})

		t.Run("return path not found e if no simple value is given", func(t *testing.T) {
			sut := ConfigPartial{}

			check, e := sut.List("node")
			switch {
			case check != nil:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("not path not found error : %v", e)
			}
		})

		t.Run("return simple value if the path don't exists", func(t *testing.T) {
			value := []interface{}{"simple value"}
			sut := ConfigPartial{}

			check, e := sut.List("node", value)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(check, value):
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})
	})

	t.Run("Partial", func(t *testing.T) {
		t.Run("return valid copied stored value", func(t *testing.T) {
			path := "node"
			value := ConfigPartial{"id": "value"}
			sut := ConfigPartial{path: value}

			check, e := sut.Partial(path, ConfigPartial{"id": "simple value"})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(check, value):
				t.Errorf("(%v) value when expecting : %v", check, value)
			default:
				check["id"] = "other value"
				changeCheck, _ := sut.Partial(path, ConfigPartial{"id": "simple value"})
				if !reflect.DeepEqual(changeCheck, value) {
					t.Errorf("(%v) value when expecting : %v", check, value)
				}
			}
		})

		t.Run("return conversion error if not storing an ConfigPartial", func(t *testing.T) {
			path := "node"
			value := 123
			sut := ConfigPartial{path: value}

			check, e := sut.Partial(path, ConfigPartial{"id": "simple value"})
			switch {
			case check != nil:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("not convertion error : %v", e)
			}
		})

		t.Run("return path not found e if no simple value is given", func(t *testing.T) {
			sut := ConfigPartial{}

			check, e := sut.Partial("node")
			switch {
			case check != nil:
				t.Errorf("unexpected value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("not path not found error : %v", e)
			}
		})

		t.Run("return default value if the path don't exists", func(t *testing.T) {
			value := ConfigPartial{"id": "simple value"}
			sut := ConfigPartial{}

			check, e := sut.Partial("node", value)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(check, value):
				t.Errorf("(%v) value when expecting : %v", check, value)
			}
		})
	})

	t.Run("Merge", func(t *testing.T) {
		t.Run("merges two partials", func(t *testing.T) {
			scenarios := []struct {
				partial1 ConfigPartial
				partial2 ConfigPartial
				expected ConfigPartial
			}{
				{ // _test merging nil ConfigPartial source
					partial1: ConfigPartial{},
					partial2: nil,
					expected: ConfigPartial{},
				},
				{ // _test merging empty ConfigPartial
					partial1: ConfigPartial{},
					partial2: ConfigPartial{},
					expected: ConfigPartial{},
				},
				{ // _test merging empty ConfigPartial with a non-empty ConfigPartial
					partial1: ConfigPartial{"node1": "value1"},
					partial2: ConfigPartial{},
					expected: ConfigPartial{"node1": "value1"},
				},
				{ // _test merging ConfigPartial into empty ConfigPartial
					partial1: ConfigPartial{},
					partial2: ConfigPartial{"node1": "value1"},
					expected: ConfigPartial{"node1": "value1"},
				},
				{ // _test merging override source value
					partial1: ConfigPartial{"node1": "value1"},
					partial2: ConfigPartial{"node1": "value2"},
					expected: ConfigPartial{"node1": "value2"},
				},
				{ // _test merging does not override non-present value in merged ConfigPartial (create)
					partial1: ConfigPartial{"node1": "value1"},
					partial2: ConfigPartial{"node2": "value2"},
					expected: ConfigPartial{"node1": "value1", "node2": "value2"},
				},
				{ // _test merging does not override non-present value in merged ConfigPartial (override)
					partial1: ConfigPartial{"node1": "value1", "node2": "value2"},
					partial2: ConfigPartial{"node2": "value3"},
					expected: ConfigPartial{"node1": "value1", "node2": "value3"},
				},
				{ // _test merging override source value to a subtree
					partial1: ConfigPartial{"node1": "value1"},
					partial2: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					expected: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
				},
				{ // _test merging override source value in a subtree (to a value)
					partial1: ConfigPartial{"node1": ConfigPartial{"node2": "value1"}},
					partial2: ConfigPartial{"node1": ConfigPartial{"node2": "value2"}},
					expected: ConfigPartial{"node1": ConfigPartial{"node2": "value2"}},
				},
				{ // _test merging override source value in a subtree (to a subtree)
					partial1: ConfigPartial{"node1": ConfigPartial{"node2": "value"}},
					partial2: ConfigPartial{"node1": ConfigPartial{"node2": ConfigPartial{"node3": "value"}}},
					expected: ConfigPartial{"node1": ConfigPartial{"node2": ConfigPartial{"node3": "value"}}},
				},
			}

			for _, s := range scenarios {
				check := s.partial1
				check.Merge(s.partial2)

				if !reflect.DeepEqual(check, s.expected) {
					t.Errorf("(%s) when merging (%v) and (%v), expecting (%v)", check, s.partial1, s.partial2, s.expected)
				}
			}
		})

		t.Run("merging works with copies", func(t *testing.T) {
			data1 := ConfigPartial{"node1": ConfigPartial{"node2": "value 2"}}
			data2 := ConfigPartial{"node1": ConfigPartial{"node3": ConfigPartial{"node4": "value 4"}}}
			expected := ConfigPartial{"node1": ConfigPartial{"node2": "value 2", "node3": ConfigPartial{"node4": "value 4"}}}

			check := ConfigPartial{}
			check.Merge(data1)
			check.Merge(data2)

			if !reflect.DeepEqual(check, expected) {
				t.Errorf("(%s) when merging (%v) and (%v), expecting (%v)", check, data1, data2, expected)
			}
		})
	})

	t.Run("Populate", func(t *testing.T) {
		t.Run("error if path not found", func(t *testing.T) {
			data := ConfigPartial{"field1": 123, "field2": 456}
			path := "field3"
			target := 0

			v, e := data.Populate(path, target)
			switch {
			case v != nil:
				t.Error("valid reference to a data")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error on populating an invalid type", func(t *testing.T) {
			data := ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}}
			path := "field1"
			target := struct{ Field1 string }{}

			v, e := data.Populate(path, target)
			switch {
			case v != nil:
				t.Error("valid reference to a data")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error on populating an invalid type struct field", func(t *testing.T) {
			data := ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}}
			path := "field2"
			target := struct{ Field1 string }{}

			v, e := data.Populate(path, target)
			switch {
			case v != nil:
				t.Error("valid reference to a data")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error on populating an inner invalid type struct field", func(t *testing.T) {
			data := ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}}
			path := ""
			target := struct{ Field2 struct{ Field1 string } }{}

			v, e := data.Populate(path, target)
			switch {
			case v != nil:
				t.Error("valid reference to a data")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("no-op if field is not found in partial", func(t *testing.T) {
			data := ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}}
			path := ""
			target := struct{ Field3 int }{Field3: 123}
			expValue := struct{ Field3 int }{Field3: 0}

			v, e := data.Populate(path, target)
			switch {
			case !reflect.DeepEqual(v, expValue):
				t.Errorf("(%v) value when expecting : %v", v, expValue)
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("no-op if field is not found in inner partial", func(t *testing.T) {
			data := ConfigPartial{"field1": 123}
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
				t.Errorf("(%v) value when expecting : %v", v, expValue)
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("populate scalar values", func(t *testing.T) {
			scenarios := []struct {
				data     ConfigPartial
				path     string
				target   interface{}
				expValue interface{}
			}{
				{ // populate an integer
					data:     ConfigPartial{"field1": 123, "field2": 456},
					path:     "field2",
					target:   0,
					expValue: 456,
				},
				{ // populate an integer from inner field
					data:     ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}},
					path:     "field2.field2",
					target:   0,
					expValue: 456,
				},
				{ // populate a string from inner field
					data:     ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": "test string"}},
					path:     "field2.field2",
					target:   "",
					expValue: "test string",
				},
			}

			for _, s := range scenarios {
				v, e := s.data.Populate(s.path, s.target)
				switch {
				case !reflect.DeepEqual(v, s.expValue):
					t.Errorf("(%v) value when expecting : %v", v, s.expValue)
				case e != nil:
					t.Errorf("unexpected (%v) error", e)
				}
			}
		})

		t.Run("populate structure references", func(t *testing.T) {
			scenarios := []struct {
				data     ConfigPartial
				path     string
				target   interface{}
				expValue interface{}
			}{
				{ // populate a single exposed field structure
					data:     ConfigPartial{"field1": 123, "field2": 456},
					path:     "",
					target:   struct{ field1, Field2 int }{},
					expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
				},
				{ // populate a single exposed field structure from inner field
					data:     ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}},
					path:     "field2",
					target:   struct{ field1, Field2 int }{},
					expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
				},
				{ // populate a multiple exposed field structure
					data:     ConfigPartial{"field1": 123, "field2": 456},
					path:     "",
					target:   struct{ Field1, Field2 int }{},
					expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
				},
				{ // populate a multiple exposed field structure from inner field
					data:     ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}},
					path:     "Field2",
					target:   struct{ Field1, Field2 int }{},
					expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
				},
				{ // populate a multiple level structure
					data: ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}},
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

			for _, s := range scenarios {
				v, e := s.data.Populate(s.path, s.target)
				switch {
				case !reflect.DeepEqual(v, s.expValue):
					t.Errorf("(%v) value when expecting : %v", v, s.expValue)
				case e != nil:
					t.Errorf("unexpected (%v) error", e)
				}
			}
		})

		t.Run("populate structure pointers", func(t *testing.T) {
			scenarios := []struct {
				data     ConfigPartial
				path     string
				target   interface{}
				expValue interface{}
			}{
				{ // populate a single exposed field structure
					data:     ConfigPartial{"field1": 123, "field2": 456},
					path:     "",
					target:   &struct{ field1, Field2 int }{},
					expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
				},
				{ // no-op if field is not in partial
					data:     ConfigPartial{"field1": 123, "field2": 456},
					path:     "",
					target:   &struct{ field1, Field2, Field3 int }{Field3: 789},
					expValue: struct{ field1, Field2, Field3 int }{field1: 0, Field2: 456, Field3: 789},
				},
				{ // populate a single exposed field structure from inner field
					data:     ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}},
					path:     "field2",
					target:   &struct{ field1, Field2 int }{},
					expValue: struct{ field1, Field2 int }{field1: 0, Field2: 456},
				},
				{ // populate a multiple exposed field structure
					data:     ConfigPartial{"field1": 123, "field2": 456},
					path:     "",
					target:   &struct{ Field1, Field2 int }{},
					expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
				},
				{ // populate a multiple exposed field structure from inner field
					data:     ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}},
					path:     "Field2",
					target:   &struct{ Field1, Field2 int }{},
					expValue: struct{ Field1, Field2 int }{Field1: 123, Field2: 456},
				},
				{ // populate a multiple level structure
					data: ConfigPartial{"field1": 123, "field2": ConfigPartial{"field1": 123, "field2": 456}},
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

			for _, s := range scenarios {
				v, e := s.data.Populate(s.path, s.target)
				switch {
				case !reflect.DeepEqual(v, s.expValue):
					t.Errorf("(%v) value when expecting : %v", v, s.expValue)
				case e != nil:
					t.Errorf("unexpected (%v) error", e)
				}
			}
		})
	})
}

func Test_ConfigConvert(t *testing.T) {
	t.Run("Convert float32 into int", func(t *testing.T) {
		data := float32(123)
		expected := 123

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert float64 into int", func(t *testing.T) {
		data := float64(123)
		expected := 123

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map", func(t *testing.T) {
		data := map[string]interface{}{"node": "value"}
		expected := ConfigPartial{"node": "value"}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial", func(t *testing.T) {
		data := ConfigPartial{"node": "value"}
		expected := ConfigPartial{"node": "value"}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert list", func(t *testing.T) {
		data := []interface{}{1, 2, 3}
		expected := []interface{}{1, 2, 3}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a float32 into a map with a int", func(t *testing.T) {
		data := map[string]interface{}{"node": float32(123)}
		expected := ConfigPartial{"node": 123}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a float64 into a map with a int", func(t *testing.T) {
		data := map[string]interface{}{"node": float64(123)}
		expected := ConfigPartial{"node": 123}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with another map", func(t *testing.T) {
		data := map[string]interface{}{"node": map[string]interface{}{"node2": "value"}}
		expected := ConfigPartial{"node": ConfigPartial{"node2": "value"}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a list", func(t *testing.T) {
		data := map[string]interface{}{"node": []interface{}{1, 2, 3}}
		expected := ConfigPartial{"node": []interface{}{1, 2, 3}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial with a map", func(t *testing.T) {
		data := ConfigPartial{"node": map[string]interface{}{"node2": "value"}}
		expected := ConfigPartial{"node": ConfigPartial{"node2": "value"}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a list of maps", func(t *testing.T) {
		data := map[string]interface{}{"node": []interface{}{map[string]interface{}{"node2": "value"}}}
		expected := ConfigPartial{"node": []interface{}{ConfigPartial{"node2": "value"}}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a list of configs", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{ConfigPartial{"node2": "value"}}}
		expected := ConfigPartial{"node": []interface{}{ConfigPartial{"node2": "value"}}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial with numeric fields", func(t *testing.T) {
		data := ConfigPartial{1: map[string]interface{}{"node2": "value"}}
		expected := ConfigPartial{1: ConfigPartial{"node2": "value"}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial with a uppercase fields", func(t *testing.T) {
		data := ConfigPartial{"NoDE": map[string]interface{}{"nODE2": "value"}}
		expected := ConfigPartial{"node": ConfigPartial{"node2": "value"}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map uppercase fields", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{map[string]interface{}{"NOde2": "value"}}}
		expected := ConfigPartial{"node": []interface{}{ConfigPartial{"node2": "value"}}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial uppercase fields", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{ConfigPartial{"NOde2": "value"}}}
		expected := ConfigPartial{"node": []interface{}{ConfigPartial{"node2": "value"}}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with numeric keys", func(t *testing.T) {
		data := map[interface{}]interface{}{1: []interface{}{ConfigPartial{2: "value"}}}
		expected := ConfigPartial{1: []interface{}{ConfigPartial{2: "value"}}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with string keys", func(t *testing.T) {
		data := map[interface{}]interface{}{"NoDE1": []interface{}{ConfigPartial{2: "value"}}}
		expected := ConfigPartial{"node1": []interface{}{ConfigPartial{2: "value"}}}

		if check := ConfigConvert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})
}

func Test_ConfigParserFactory(t *testing.T) {
	t.Run("NewConfigParserFactory", func(t *testing.T) {
		t.Run("creation with empty creator list", func(t *testing.T) {
			sut := NewConfigParserFactory(nil)
			if sut == nil {
				t.Error("didn't returned the expected reference")
			}
		})

		t.Run("creation with creator list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			creator := NewMockConfigParserCreator(ctrl)

			sut := NewConfigParserFactory([]ConfigParserCreator{creator})
			if sut == nil {
				t.Error("didn't returned the expected reference")
			} else if (*sut)[0] != creator {
				t.Error("didn't stored the passed creator")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error if the format is unrecognized", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			format := ConfigFormatJSON
			reader := NewMockReader(ctrl)
			creator := NewMockConfigParserCreator(ctrl)
			creator.EXPECT().Accept(format).Return(false).Times(1)
			sut := NewConfigParserFactory([]ConfigParserCreator{creator})

			check, e := sut.Create(format, reader)
			switch {
			case check != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigFormat)
			}
		})

		t.Run("should create the requested parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			format := ConfigFormatJSON
			reader := NewMockReader(ctrl)
			parser := NewMockConfigParser(ctrl)
			creator := NewMockConfigParserCreator(ctrl)
			creator.EXPECT().Accept(format).Return(true).Times(1)
			creator.EXPECT().Create(reader).Return(parser, nil).Times(1)
			sut := NewConfigParserFactory([]ConfigParserCreator{creator})

			if check, e := sut.Create(format, reader); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(check, parser) {
				t.Error("didn't returned the created parser")
			}
		})
	})
}

func Test_ConfigDecoder_Close(t *testing.T) {
	t.Run("error while closing the jsonReader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Return(expected).Times(1)
		sut := ConfigDecoder{Reader: reader}

		if e := sut.Close(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("call close method on jsonReader only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut := ConfigDecoder{Reader: reader}

		_ = sut.Close()
		_ = sut.Close()
	})
}

func Test_ConfigDecoder(t *testing.T) {
	t.Run("Close", func(t *testing.T) {
		t.Run("error while closing the decoder", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Return(expected).Times(1)
			sut := ConfigDecoder{Reader: reader}

			if e := sut.Close(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("call close method on decoder only once", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut := ConfigDecoder{Reader: reader}

			_ = sut.Close()
			_ = sut.Close()
		})
	})

	t.Run("Decode", func(t *testing.T) {
		t.Run("return decode error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut := ConfigDecoder{Reader: reader}
			defer func() { _ = sut.Close() }()
			baseDecoder := NewMockConfigUnderlyingDecoder(ctrl)
			baseDecoder.
				EXPECT().
				Decode(&map[string]interface{}{}).
				DoAndReturn(func(p *map[string]interface{}) error {
					return expected
				}).Times(1)
			sut.UnderlyingDecoder = baseDecoder

			check, e := sut.Parse()
			switch {
			case check != nil:
				t.Error("returned an reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("redirect to the underlying decoder", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data := ConfigPartial{"node": "data"}
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut := ConfigDecoder{Reader: reader}
			defer func() { _ = sut.Close() }()
			baseDecoder := NewMockConfigUnderlyingDecoder(ctrl)
			baseDecoder.
				EXPECT().
				Decode(&map[string]interface{}{}).
				DoAndReturn(func(p *map[string]interface{}) error {
					(*p)["node"] = data["node"]
					return nil
				}).Times(1)
			sut.UnderlyingDecoder = baseDecoder

			check, e := sut.Parse()
			switch {
			case check == nil:
				t.Error("returned a nil data")
			case !reflect.DeepEqual(*check, data):
				t.Errorf("returned (%v)", check)
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})
}

func Test_ConfigJSONDecoder(t *testing.T) {
	t.Run("NewConfigJSONDecoder", func(t *testing.T) {
		t.Run("nil reader", func(t *testing.T) {
			sut, e := NewConfigJSONDecoder(nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new json decoder adapter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)

			if sut, e := NewConfigJSONDecoder(reader); sut == nil {
				t.Errorf("didn't returned a valid reference")
			} else {
				defer func() { _ = sut.Close() }()
				if e != nil {
					t.Errorf("unexpected (%v) error", e)
				} else if sut.Reader != reader {
					t.Error("didn't store the jsonReader reference")
				}
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("error while closing the jsonReader", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Return(expected).Times(1)
			sut, _ := NewConfigJSONDecoder(reader)

			if e := sut.Close(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("call close method on jsonReader only once", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut, _ := NewConfigJSONDecoder(reader)

			_ = sut.Close()
			_ = sut.Close()
		})
	})

	t.Run("Parse", func(t *testing.T) {
		t.Run("return parse error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut, _ := NewConfigJSONDecoder(reader)
			defer func() { _ = sut.Close() }()
			underlyingDecoder := NewMockConfigUnderlyingDecoder(ctrl)
			underlyingDecoder.
				EXPECT().
				Decode(&map[string]interface{}{}).
				DoAndReturn(func(p *map[string]interface{}) error {
					return expected
				}).Times(1)
			sut.UnderlyingDecoder = underlyingDecoder

			check, e := sut.Parse()
			switch {
			case check != nil:
				t.Error("returned an reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("redirect to the underlying decoder", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data := ConfigPartial{"node": "data"}
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut, _ := NewConfigJSONDecoder(reader)
			defer func() { _ = sut.Close() }()
			underlyingDecoder := NewMockConfigUnderlyingDecoder(ctrl)
			underlyingDecoder.
				EXPECT().
				Decode(&map[string]interface{}{}).
				DoAndReturn(func(p *map[string]interface{}) error {
					(*p)["node"] = data["node"]
					return nil
				}).Times(1)
			sut.UnderlyingDecoder = underlyingDecoder

			check, e := sut.Parse()
			switch {
			case check == nil:
				t.Error("returned a nil data")
			case !reflect.DeepEqual(*check, data):
				t.Errorf("returned (%v)", check)
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("parse json string", func(t *testing.T) {
			json := `{"node": {"sub_node": "data"}}`
			expected := ConfigPartial{"node": ConfigPartial{"sub_node": "data"}}
			reader := strings.NewReader(json)
			sut, _ := NewConfigJSONDecoder(reader)
			defer func() { _ = sut.Close() }()

			check, e := sut.Parse()
			switch {
			case check == nil:
				t.Error("returned a nil data")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(*check, expected):
				t.Errorf("(%v) when expecting (%v)", *check, expected)
			}
		})
	})
}

func Test_ConfigJSONDecoderCreator(t *testing.T) {
	t.Run("NewConfigJSONDecoderCreator", func(t *testing.T) {
		t.Run("creation", func(t *testing.T) {
			if NewConfigJSONDecoderCreator() == nil {
				t.Error("didn't returned the expected reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		t.Run("accept only json format", func(t *testing.T) {
			scenarios := []struct {
				format   string
				expected bool
			}{
				{ // _test json format
					format:   ConfigFormatJSON,
					expected: true,
				},
				{ // _test non-json format
					format:   ConfigFormatYAML,
					expected: false,
				},
			}

			for _, s := range scenarios {
				test := func() {
					if check := NewConfigJSONDecoderCreator().Accept(s.format); check != s.expected {
						t.Errorf("returned (%v) when checking (%v)", check, s.format)
					}
				}
				test()
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("nil reader", func(t *testing.T) {
			if decoder, e := NewConfigJSONDecoderCreator().Create(); decoder != nil {
				t.Error("returned an unexpected valid decoder instance")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("invalid reader instance", func(t *testing.T) {
			if decoder, e := NewConfigJSONDecoderCreator().Create("string"); decoder != nil {
				t.Error("returned an unexpected valid decoder instance")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the decoder", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			decoder, e := NewConfigJSONDecoderCreator().Create(NewMockReader(ctrl))
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case decoder == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch decoder.(type) {
				case *ConfigJSONDecoder:
				default:
					t.Error("didn't returned a JSON decoder")
				}
			}
		})
	})
}

func Test_ConfigYAMLDecoder(t *testing.T) {
	t.Run("NewConfigYAMLDecoder", func(t *testing.T) {
		t.Run("nil reader", func(t *testing.T) {
			sut, e := NewConfigYAMLDecoder(nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new yaml decoder adapter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)

			if sut, e := NewConfigYAMLDecoder(reader); sut == nil {
				t.Errorf("didn't returned a valid reference")
			} else {
				defer func() { _ = sut.Close() }()
				if e != nil {
					t.Errorf("unexpected (%v) error", e)
				} else if sut.Reader != reader {
					t.Error("didn't store the reader reference")
				}
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("error while closing the reader", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Return(expected).Times(1)
			sut, _ := NewConfigYAMLDecoder(reader)

			if e := sut.Close(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("call close method on reader only once", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut, _ := NewConfigYAMLDecoder(reader)

			_ = sut.Close()
			_ = sut.Close()
		})
	})

	t.Run("Parse", func(t *testing.T) {
		t.Run("return parse error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut, _ := NewConfigYAMLDecoder(reader)
			defer func() { _ = sut.Close() }()
			underlyingDecoder := NewMockConfigUnderlyingDecoder(ctrl)
			underlyingDecoder.
				EXPECT().
				Decode(&ConfigPartial{}).
				DoAndReturn(func(p *ConfigPartial) error {
					return expected
				}).Times(1)
			sut.UnderlyingDecoder = underlyingDecoder

			check, e := sut.Parse()
			switch {
			case check != nil:
				t.Error("returned an reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("redirect to the underlying decoder", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data := ConfigPartial{"node": "data"}
			reader := NewMockReader(ctrl)
			reader.EXPECT().Close().Times(1)
			sut, _ := NewConfigYAMLDecoder(reader)
			defer func() { _ = sut.Close() }()
			underlyingDecoder := NewMockConfigUnderlyingDecoder(ctrl)
			underlyingDecoder.
				EXPECT().
				Decode(&ConfigPartial{}).
				DoAndReturn(func(p *ConfigPartial) error {
					(*p)["node"] = data["node"]
					return nil
				}).Times(1)
			sut.UnderlyingDecoder = underlyingDecoder

			check, e := sut.Parse()
			switch {
			case check == nil:
				t.Error("returned a nil data")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(*check, data):
				t.Errorf("returned (%v)", check)
			}
		})

		t.Run("parse yaml string", func(t *testing.T) {
			yaml := "node:\n  sub_node: data"
			expected := ConfigPartial{"node": ConfigPartial{"sub_node": "data"}}
			reader := strings.NewReader(yaml)
			sut, _ := NewConfigYAMLDecoder(reader)
			defer func() { _ = sut.Close() }()

			check, e := sut.Parse()
			switch {
			case check == nil:
				t.Error("returned a nil data")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(*check, expected):
				t.Errorf("(%v) when expecting (%v)", *check, expected)
			}
		})
	})
}

func Test_ConfigYAMLDecoderCreator(t *testing.T) {
	t.Run("NewConfigYAMLDecoderCreator", func(t *testing.T) {
		t.Run("creation", func(t *testing.T) {
			if NewConfigYAMLDecoderCreator() == nil {
				t.Error("didn't returned the expected reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		t.Run("accept only yaml format", func(t *testing.T) {
			scenarios := []struct {
				format   string
				expected bool
			}{
				{ // _test yaml format
					format:   ConfigFormatYAML,
					expected: true,
				},
				{ // _test non-yaml format
					format:   ConfigFormatJSON,
					expected: false,
				},
			}

			for _, s := range scenarios {
				test := func() {
					if check := NewConfigYAMLDecoderCreator().Accept(s.format); check != s.expected {
						t.Errorf("returned (%v) when checking (%v)", check, s.format)
					}
				}
				test()
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("nil reader", func(t *testing.T) {
			if decoder, e := NewConfigYAMLDecoderCreator().Create(); decoder != nil {
				t.Error("returned an unexpected valid decoder instance")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("invalid reader instance", func(t *testing.T) {
			if decoder, e := NewConfigYAMLDecoderCreator().Create("string"); decoder != nil {
				t.Error("returned an unexpected valid decoder instance")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the decoder", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			decoder, e := NewConfigYAMLDecoderCreator().Create(NewMockReader(ctrl))
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case decoder == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch decoder.(type) {
				case *ConfigYAMLDecoder:
				default:
					t.Error("didn't returned a YAML decoder")
				}
			}
		})
	})
}

func Test_ConfigSupplierFactory(t *testing.T) {
	t.Run("NewConfigSupplierFactory", func(t *testing.T) {
		t.Run("creation with empty creator list", func(t *testing.T) {
			sut := NewConfigSupplierFactory(nil)
			if sut == nil {
				t.Error("didn't returned the expected reference")
			}
		})

		t.Run("creation with creator list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			creator := NewMockConfigSupplierCreator(ctrl)
			sut := NewConfigSupplierFactory([]ConfigSupplierCreator{creator})
			if sut == nil {
				t.Error("didn't returned the expected reference")
			} else if (*sut)[0] != creator {
				t.Error("didn't stored the passed creator")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("nil partial", func(t *testing.T) {
			sut := NewConfigSupplierFactory(nil)
			src, e := sut.Create(nil)
			switch {
			case src != nil:
				t.Error("returned an unexpected supplier")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error on unrecognized format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sourceType := "type"
			path := "path"
			format := "format"
			partial := ConfigPartial{
				"type":   sourceType,
				"path":   path,
				"format": format,
			}
			creator := NewMockConfigSupplierCreator(ctrl)
			creator.EXPECT().Accept(&partial).Return(false).Times(1)
			sut := NewConfigSupplierFactory([]ConfigSupplierCreator{creator})

			src, e := sut.Create(&partial)
			switch {
			case src != nil:
				t.Error("returned an unexpected supplier")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("create the requested supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sourceType := "type"
			path := "path"
			format := "format"
			partial := ConfigPartial{
				"type":   sourceType,
				"path":   path,
				"format": format,
			}
			supplier := NewMockConfigSupplier(ctrl)
			creator := NewMockConfigSupplierCreator(ctrl)
			creator.EXPECT().Accept(&partial).Return(true).Times(1)
			creator.EXPECT().Create(&partial).Return(supplier, nil).Times(1)
			sut := NewConfigSupplierFactory([]ConfigSupplierCreator{creator})

			if check, e := sut.Create(&partial); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(check, supplier) {
				t.Error("didn't returned the created source")
			}
		})
	})
}

func Test_ConfigSource(t *testing.T) {
	t.Run("has", func(t *testing.T) {
		t.Run("lock and check the stored config value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			search := "path"
			data := ConfigPartial{search: "value"}
			locker := NewMockLocker(ctrl)
			locker.EXPECT().Lock().Times(1)
			locker.EXPECT().Unlock().Times(1)

			sut := &ConfigSource{Mutex: locker, Partial: data}

			if value := sut.Has(search); value != true {
				t.Errorf("returned the (%v) value", value)
			}
		})
	})

	t.Run("get", func(t *testing.T) {
		t.Run("lock and redirect to the stored config value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			search := "path"
			expected := "value"
			data := ConfigPartial{search: expected}
			locker := NewMockLocker(ctrl)
			locker.EXPECT().Lock().Times(1)
			locker.EXPECT().Unlock().Times(1)

			sut := &ConfigSource{Mutex: locker, Partial: data}

			if value, e := sut.Get(search); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if value != expected {
				t.Errorf("returned the (%v) value", value)
			}
		})
	})
}

func Test_ConfigAggregateSource(t *testing.T) {
	t.Run("NewNewConfigAggregateSource", func(t *testing.T) {
		t.Run("nil list of configs", func(t *testing.T) {
			sut, e := NewConfigAggregateSource(nil)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned the expected valid reference")
			case !reflect.DeepEqual(sut.Partial, ConfigPartial{}):
				t.Errorf("(%v) when expecting (%v)", sut.Partial, ConfigPartial{})
			}
		})

		t.Run("error while retrieving config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")

			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("", ConfigPartial{}).Return(nil, expected).Times(1)

			sut, e := NewConfigAggregateSource([]ConfigSupplier{supplier})
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("valid single config load", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"id": "value"}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("", ConfigPartial{}).Return(expected, nil).Times(1)

			sut, e := NewConfigAggregateSource([]ConfigSupplier{supplier})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned the expected valid reference")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("(%v) when expecting (%v)", sut.Partial, expected)
			}
		})

		t.Run("valid multiple partials load", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"id 1": "value 1", "id 2": "value 2"}
			partial1 := ConfigPartial{"id 1": "value 1"}
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("", ConfigPartial{}).Return(partial1, nil).Times(1)
			partial2 := ConfigPartial{"id 2": "value 2"}
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("", ConfigPartial{}).Return(partial2, nil).Times(1)

			sut, e := NewConfigAggregateSource([]ConfigSupplier{supplier1, supplier2})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned the expected valid reference")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("(%v) when expecting (%v)", sut.Partial, expected)
			}
		})
	})
}

func Test_ConfigAggregateSourceCreator(t *testing.T) {
	t.Run("NewConfigAggregateSourceCreator", func(t *testing.T) {
		t.Run("nil suppliers list", func(t *testing.T) {
			if sut, e := NewConfigAggregateSourceCreator(nil); e != nil {
				t.Errorf("return the unexpected error : %v", e)
			} else if sut == nil {
				t.Error("didn't return the expected valid reference")
			}
		})

		t.Run("non-empty suppliers list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			supplier := NewMockConfigSupplier(ctrl)

			if sut, e := NewConfigAggregateSourceCreator([]ConfigSupplier{supplier}); e != nil {
				t.Errorf("return the unexpected error : %v", e)
			} else if sut == nil {
				t.Error("didn't return the expected valid reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		supplier := NewMockConfigSupplier(ctrl)
		sut, _ := NewConfigAggregateSourceCreator([]ConfigSupplier{supplier})

		t.Run("don't accept on invalid config pointer", func(t *testing.T) {
			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is missing", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not a string", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not aggregate", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": ConfigTypeDir}) {
				t.Error("returned true")
			}
		})

		t.Run("accept if type is aggregate", func(t *testing.T) {
			if !sut.Accept(&ConfigPartial{"type": ConfigTypeAggregate}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("accept nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := ConfigPartial{"key": "value"}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("", ConfigPartial{}).Return(value, nil).Times(1)
			sut, _ := NewConfigAggregateSourceCreator([]ConfigSupplier{supplier})

			src, e := sut.Create(nil)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch src.(type) {
				case *ConfigAggregateSource:
				default:
					t.Error("didn't returned a new env src")
				}
			}
		})

		t.Run("create the source with a single config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := ConfigPartial{"key": "value"}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("", ConfigPartial{}).Return(value, nil).Times(1)
			sut, _ := NewConfigAggregateSourceCreator([]ConfigSupplier{supplier})

			src, e := sut.Create(nil)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigAggregateSource:
					if !reflect.DeepEqual(s.Partial, value) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new env src")
				}
			}
		})

		t.Run("create the source with multiple config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value1 := ConfigPartial{"key1": "value 1"}
			value2 := ConfigPartial{"key2": "value 2"}

			expected := ConfigPartial{"key1": "value 1", "key2": "value 2"}
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("", ConfigPartial{}).Return(value1, nil).Times(1)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("", ConfigPartial{}).Return(value2, nil).Times(1)
			sut, _ := NewConfigAggregateSourceCreator([]ConfigSupplier{supplier1, supplier2})

			src, e := sut.Create(nil)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigAggregateSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new env src")
				}
			}
		})
	})
}

func Test_ConfigEnvSource(t *testing.T) {
	t.Run("NewConfigEnvSource", func(t *testing.T) {
		t.Run("with empty mappings", func(t *testing.T) {
			sut, e := NewConfigEnvSource(map[string]string{})
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(sut.Partial, ConfigPartial{}):
					t.Error("didn't loaded the content correctly")
				}
			}
		})

		t.Run("with empty environment", func(t *testing.T) {
			env := "env"

			expected := ConfigPartial{}

			sut, e := NewConfigEnvSource(map[string]string{env: "id"})
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(sut.Partial, expected):
					t.Errorf("(%v) when expecting (%v)", sut.Partial, expected)
				}
			}
		})

		t.Run("with root mappings", func(t *testing.T) {
			env := "env"
			value := "value"
			_ = os.Setenv(env, value)
			defer func() { _ = os.Setenv(env, "") }()

			expected := ConfigPartial{"id": value}

			sut, e := NewConfigEnvSource(map[string]string{env: "id"})
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(sut.Partial, expected):
					t.Errorf("(%v) when expecting (%v)", sut.Partial, expected)
				}
			}
		})

		t.Run("with multi-level mappings", func(t *testing.T) {
			env := "env"
			value := "value"
			_ = os.Setenv(env, value)
			defer func() { _ = os.Setenv(env, "") }()

			expected := ConfigPartial{}
			_, _ = expected.Set("root.node", value)

			sut, e := NewConfigEnvSource(map[string]string{env: "root.node"})
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(sut.Partial, expected):
					t.Errorf("(%v) when expecting (%v)", sut.Partial, expected)
				}
			}
		})

		t.Run("with multi-level mapping", func(t *testing.T) {
			env := "env1"
			value := "value"
			_ = os.Setenv(env, value)
			defer func() { _ = os.Setenv(env, "") }()

			expected := ConfigPartial{}
			_, _ = expected.Set("root.node", "value")

			sut, e := NewConfigEnvSource(map[string]string{"env1": "root.node"})
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(sut.Partial, expected):
					t.Errorf("(%v) when expecting (%v)", sut.Partial, expected)
				}
			}
		})

		t.Run("with multi-level mapping (deeper)", func(t *testing.T) {
			env := "env1"
			value := "value"
			_ = os.Setenv(env, value)
			defer func() { _ = os.Setenv(env, "") }()

			expected := ConfigPartial{}
			_, _ = expected.Set("root.node1.node2", "value")

			sut, e := NewConfigEnvSource(map[string]string{"env1": "root.node1.node2"})
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(sut.Partial, expected):
					t.Errorf(" (%v) when expecting (%v)", sut.Partial, expected)
				}
			}
		})
	})
}

func Test_ConfigEnvSourceCreator(t *testing.T) {
	t.Run("NewConfigEnvSourceCreator", func(t *testing.T) {
		t.Run("creation", func(t *testing.T) {
			if NewConfigEnvSourceCreator() == nil {
				t.Error("didn't returned the expected reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		t.Run("don't accept on invalid config pointer", func(t *testing.T) {
			if NewConfigEnvSourceCreator().Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is missing", func(t *testing.T) {
			if NewConfigEnvSourceCreator().Accept(&ConfigPartial{}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not a string", func(t *testing.T) {
			if NewConfigEnvSourceCreator().Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not env", func(t *testing.T) {
			if NewConfigEnvSourceCreator().Accept(&ConfigPartial{"type": ConfigTypeAggregate}) {
				t.Error("returned true")
			}
		})

		t.Run("accept if type is env", func(t *testing.T) {
			if !NewConfigEnvSourceCreator().Accept(&ConfigPartial{"type": ConfigTypeEnv}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			src, e := NewConfigEnvSourceCreator().Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("non-map mappings", func(t *testing.T) {
			src, e := NewConfigEnvSourceCreator().Create(&ConfigPartial{
				"mappings": 123,
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string key map mappings", func(t *testing.T) {
			src, e := NewConfigEnvSourceCreator().Create(&ConfigPartial{
				"mappings": ConfigPartial{
					1: "value",
				}})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string value map mappings", func(t *testing.T) {
			src, e := NewConfigEnvSourceCreator().Create(&ConfigPartial{
				"mappings": ConfigPartial{
					"key": 1,
				}})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the source", func(t *testing.T) {
			env := "env"
			value := "value"
			_ = os.Setenv(env, value)
			defer func() { _ = os.Setenv(env, "") }()

			path := "root"
			expected := ConfigPartial{path: value}

			src, e := NewConfigEnvSourceCreator().Create(&ConfigPartial{
				"mappings": ConfigPartial{env: path},
			})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigEnvSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new env src")
				}
			}
		})

		t.Run("no mappings on config", func(t *testing.T) {
			expected := ConfigPartial{}

			src, e := NewConfigEnvSourceCreator().Create(&ConfigPartial{})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigEnvSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new env src")
				}
			}
		})
	})
}

func Test_ConfigFileSource(t *testing.T) {
	t.Run("NewConfigFileSource", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigFileSource("path", "format", nil, NewConfigParserFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigFileSource("path", "format", NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error that may be raised when opening the file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := fmt.Errorf("error message")
			fs := NewMockFs(ctrl)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)

			sut, e := NewConfigFileSource(path, "format", fs, NewConfigParserFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error that may be raised when creating the parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept(ConfigFormatJSON).Return(false).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigFileSource(path, ConfigFormatJSON, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigFormat)
			}
		})

		t.Run("error that may be raised when running the parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := fmt.Errorf("error message")
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(nil, expected).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigFileSource(path, "format", fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("creates the config file source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{"field": "value"}
			path := "path"
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&partial, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigFileSource(path, "format", fs, parserFactory)
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case sut.path != path:
					t.Error("didn't stored the file path")
				case sut.format != "format":
					t.Error("didn't stored the file content format")
				case sut.fileSystem != fs:
					t.Error("didn't stored the file system adapter reference")
				case sut.parserFactory != parserFactory:
					t.Error("didn't stored the parser factory reference")
				case !reflect.DeepEqual(sut.Partial, partial):
					t.Error("didn't stored the parser information")
				}
			}
		})
	})
}

func Test_ConfigFileSourceCreator(t *testing.T) {
	t.Run("NewConfigFileSourceCreator", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			sut, e := NewConfigFileSourceCreator(nil, NewConfigParserFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigFileSourceCreator(NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new file source factory creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fs := NewMockFs(ctrl)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigFileSourceCreator(fs, parserFactory)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case sut.fileSystem != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.parserFactory != parserFactory:
				t.Error("didn't stored the parser factory reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		t.Run("don't accept on invalid config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is missing", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			if sut.Accept(&ConfigPartial{}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not a string", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if invalid type", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			if sut.Accept(&ConfigPartial{"type": ConfigTypeAggregate}) {
				t.Error("returned true")
			}
		})

		t.Run("accept config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			if !sut.Accept(&ConfigPartial{"type": ConfigTypeFile}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("missing path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{"format": "format"})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("non-string path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{"path": 123, "format": "format"})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{"path": "path", "format": 123})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the file source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&expected, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigFileSourceCreator(fs, parserFactory)

			src, e := sut.Create(&ConfigPartial{"path": path, "format": "format"})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigFileSource:
					p, _ := s.Get("")
					if !reflect.DeepEqual(p, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new file src")
				}
			}
		})

		t.Run("create the file source defaulting format if not present in config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&expected, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigFileSourceCreator(fs, parserFactory)

			src, e := sut.Create(&ConfigPartial{"path": path})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigFileSource:
					p, _ := s.Get("")
					if !reflect.DeepEqual(p, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new file src")
				}
			}
		})
	})
}

func Test_ConfigObsFileSource(t *testing.T) {
	t.Run("NewConfigObsFileSource", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			sut, e := NewConfigObsFileSource("path", "format", nil, NewConfigParserFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigObsFileSource("path", "format", NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error that may be raised when retrieving the file info", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := fmt.Errorf("error message")
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(nil, expected).Times(1)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigObsFileSource(path, "format", fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error that may be raised when opening the file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := fmt.Errorf("error message")
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigObsFileSource(path, "format", fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error that may be raised when creating the parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Times(1)
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept(ConfigFormatJSON).Return(false).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigObsFileSource(path, ConfigFormatJSON, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("error that may be raised when running the parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			path := "path"
			file := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(nil, expected).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigObsFileSource(path, "format", fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("create the config observable file source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			file := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&expected, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigObsFileSource(path, "format", fs, parserFactory)
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.path != path:
					t.Error("didn't stored the file path")
				case sut.format != "format":
					t.Error("didn't stored the file content format")
				case sut.fileSystem != fs:
					t.Error("didn't stored the file system adapter reference")
				case sut.parserFactory != parserFactory:
					t.Error("didn't stored the parser factory reference")
				}
			}
		})
	})

	t.Run("Reload", func(t *testing.T) {
		t.Run("error if fail to retrieving the file info", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			expected := fmt.Errorf("error message")
			file := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
			fs := NewMockFs(ctrl)
			gomock.InOrder(
				fs.EXPECT().Stat(path).Return(fileInfo, nil),
				fs.EXPECT().Stat(path).Return(nil, expected),
			)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{field: value}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsFileSource(path, "format", fs, parserFactory)

			reloaded, e := sut.Reload()
			switch {
			case reloaded:
				t.Error("flagged that was reloaded")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error if fails to load the file content", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			expected := fmt.Errorf("error message")
			file := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			gomock.InOrder(
				fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)),
				fileInfo.EXPECT().ModTime().Return(time.Unix(0, 2)),
			)
			fs := NewMockFs(ctrl)
			gomock.InOrder(
				fs.EXPECT().Stat(path).Return(fileInfo, nil),
				fs.EXPECT().Stat(path).Return(fileInfo, nil),
			)
			gomock.InOrder(
				fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil),
				fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected),
			)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{field: value}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsFileSource(path, "format", fs, parserFactory)

			reloaded, e := sut.Reload()
			switch {
			case reloaded:
				t.Error("flagged that was reloaded")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("prevent reload of a unchanged source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			file := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(2)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(2)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{field: value}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsFileSource(path, "format", fs, parserFactory)

			if reloaded, e := sut.Reload(); reloaded {
				t.Error("flagged that was reloaded")
			} else if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("should reload a changed source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value1 := "value1"
			value2 := "value2"
			expected := ConfigPartial{field: value2}
			file1 := NewMockFile(ctrl)
			file2 := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			gomock.InOrder(
				fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)),
				fileInfo.EXPECT().ModTime().Return(time.Unix(0, 2)),
			)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(2)
			gomock.InOrder(
				fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil),
				fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil),
			)

			parser1 := NewMockConfigParser(ctrl)
			parser1.EXPECT().Parse().Return(&ConfigPartial{field: value1}, nil).Times(1)
			parser1.EXPECT().Close().Return(nil).Times(1)
			parser2 := NewMockConfigParser(ctrl)
			parser2.EXPECT().Parse().Return(&ConfigPartial{field: value2}, nil).Times(1)
			parser2.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			gomock.InOrder(
				parserCreator.EXPECT().Accept("format").Return(true),
				parserCreator.EXPECT().Accept("format").Return(true),
			)
			gomock.InOrder(
				parserCreator.EXPECT().Create(file1).Return(parser1, nil),
				parserCreator.EXPECT().Create(file2).Return(parser2, nil),
			)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsFileSource(path, "format", fs, parserFactory)

			reloaded, e := sut.Reload()
			p, _ := sut.Get("")

			switch {
			case !reloaded:
				t.Error("flagged that was not reloaded")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(expected, p):
				t.Error("didn't stored the check configuration")
			}
		})
	})
}

func Test_ConfigObsFileSourceCreator(t *testing.T) {
	t.Run("NewConfigObsFileSourceCreator", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			sut, e := NewConfigObsFileSourceCreator(nil, NewConfigParserFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigObsFileSourceCreator(NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new file source factory creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fs := NewMockFs(ctrl)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigObsFileSourceCreator(fs, parserFactory)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case sut.fileSystem != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.parserFactory != parserFactory:
				t.Error("didn't stored the parser factory reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewConfigObsFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

		t.Run("don't accept on invalid config pointer", func(t *testing.T) {
			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is missing", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not a string", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if invalid type", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": ConfigTypeFile}) {
				t.Error("returned true")
			}
		})

		t.Run("accept config", func(t *testing.T) {
			if !sut.Accept(&ConfigPartial{"type": ConfigTypeObsFile}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("missing path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{"format": "format"})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("non-string path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{"path": 123, "format": "format"})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsFileSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{"path": "path", "format": 123})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the observable file source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			file := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&expected, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsFileSourceCreator(fs, parserFactory)

			src, e := sut.Create(&ConfigPartial{"path": path, "format": "format"})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigObsFileSource:
					p, _ := s.Get("")
					if !reflect.DeepEqual(p, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new file src")
				}
			}
		})

		t.Run("create the observable file source defaulting format if not present in config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			file := NewMockFile(ctrl)
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&expected, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsFileSourceCreator(fs, parserFactory)

			src, e := sut.Create(&ConfigPartial{"path": path})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigObsFileSource:
					p, _ := s.Get("")
					if !reflect.DeepEqual(p, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new file src")
				}
			}
		})
	})
}

func Test_ConfigDirSource(t *testing.T) {
	t.Run("NewConfigDirSource", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigDirSource("path", "format", true, nil, NewConfigParserFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigDirSource("path", "format", true, NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error that may be raised when opening the dir", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := fmt.Errorf("error message")
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(nil, expected).Times(1)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigDirSource(path, "format", true, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error that may be raised when reading the dir", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := fmt.Errorf("error message")
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return(nil, expected).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigDirSource(path, "format", true, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("empty dir", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigDirSource(path, "format", true, fs, parserFactory)
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case sut.path != path:
					t.Error("didn't stored the file path")
				case sut.format != "format":
					t.Error("didn't stored the file content format")
				case sut.fileSystem != fs:
					t.Error("didn't stored the file system adapter reference")
				case sut.parserFactory != parserFactory:
					t.Error("didn't stored the parser factory reference")
				}
			}
		})

		t.Run("error opening the config file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := fmt.Errorf("error message")
			fileInfoName := "file.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(nil, expected).
				Times(1)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigDirSource(path, ConfigFormatJSON, true, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error retrieving the proper parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			fileInfoName := "file.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Return(nil).Times(1)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept(ConfigFormatJSON).Return(false).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigDirSource(path, ConfigFormatJSON, true, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("error decoding file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			msg := "error message"
			expected := fmt.Errorf("yaml: input error: %s", msg)
			fileInfoName := "file.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(nil, expected).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigDirSource(path, "format", true, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("correctly load single file on directory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := &ConfigPartial{"field": "value"}
			path := "path"
			fileInfoName := "file.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(partial, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigDirSource(path, "format", true, fs, parserFactory)
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case sut.path != path:
					t.Error("didn't stored the file path")
				case sut.format != "format":
					t.Error("didn't stored the file content format")
				case sut.fileSystem != fs:
					t.Error("didn't stored the file system adapter reference")
				case sut.parserFactory != parserFactory:
					t.Error("didn't stored the parser factory reference")
				case !reflect.DeepEqual(sut.Partial, *partial):
					t.Errorf("(%v) when expecting (%v)", sut.Partial, *partial)
				}
			}
		})

		t.Run("don't follow sub dirs if not recursive", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := &ConfigPartial{"field": "value"}
			path := "path"
			fileInfoName := "file.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			subDirInfo := NewMockFileInfo(ctrl)
			subDirInfo.EXPECT().IsDir().Return(true).Times(1)
			subDirInfo.EXPECT().Name().Times(0)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo, subDirInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(partial, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigDirSource(path, "format", false, fs, parserFactory)
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case sut.path != path:
					t.Error("didn't stored the file path")
				case sut.format != "format":
					t.Error("didn't stored the file content format")
				case sut.fileSystem != fs:
					t.Error("didn't stored the file system adapter reference")
				case sut.parserFactory != parserFactory:
					t.Error("didn't stored the parser factory reference")
				case !reflect.DeepEqual(sut.Partial, *partial):
					t.Errorf("(%v) when expecting (%v)", sut.Partial, *partial)
				}
			}
		})

		t.Run("error while loading sub dir", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial1 := &ConfigPartial{"field": "value"}
			path := "path"
			expected := fmt.Errorf("error message")
			fileInfoName := "file1.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			subDirInfoName := "sub_dir"
			subDirInfo := NewMockFileInfo(ctrl)
			subDirInfo.EXPECT().IsDir().Return(true).Times(1)
			subDirInfo.EXPECT().Name().Return(subDirInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo, subDirInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			gomock.InOrder(
				fs.EXPECT().Open(path).Return(dir, nil).Times(1),
				fs.EXPECT().Open(path+"/"+subDirInfoName).Return(nil, expected).Times(1),
			)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(partial1, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigDirSource(path, "format", true, fs, parserFactory)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("follow sub dirs if recursive", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial1 := &ConfigPartial{"field1": "value"}
			partial2 := &ConfigPartial{"field2": "value"}
			expected := ConfigPartial{"field1": "value", "field2": "value"}
			path := "path"
			fileInfoName := "file1.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			subFileInfoName := "file2.yaml"
			subFileInfo := NewMockFileInfo(ctrl)
			subFileInfo.EXPECT().IsDir().Return(false).Times(1)
			subFileInfo.EXPECT().Name().Return(subFileInfoName).Times(1)
			subDirInfoName := "sub_dir"
			subDirInfo := NewMockFileInfo(ctrl)
			subDirInfo.EXPECT().IsDir().Return(true).Times(1)
			subDirInfo.EXPECT().Name().Return(subDirInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo, subDirInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			subDir := NewMockFile(ctrl)
			subDir.EXPECT().Readdir(0).Return([]os.FileInfo{subFileInfo}, nil).Times(1)
			subDir.EXPECT().Close().Return(nil).Times(1)
			file1 := NewMockFile(ctrl)
			file2 := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			gomock.InOrder(
				fs.EXPECT().Open(path).Return(dir, nil),
				fs.EXPECT().Open(path+"/"+subDirInfoName).Return(subDir, nil),
			)
			gomock.InOrder(
				fs.
					EXPECT().
					OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
					Return(file1, nil),
				fs.
					EXPECT().
					OpenFile(path+"/"+subDirInfoName+"/"+subFileInfoName, os.O_RDONLY, os.FileMode(0o644)).
					Return(file2, nil),
			)
			parser1 := NewMockConfigParser(ctrl)
			parser1.EXPECT().Parse().Return(partial1, nil).Times(1)
			parser1.EXPECT().Close().Return(nil).Times(1)
			parser2 := NewMockConfigParser(ctrl)
			parser2.EXPECT().Parse().Return(partial2, nil).Times(1)
			parser2.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			gomock.InOrder(
				parserCreator.EXPECT().Accept("format").Return(true),
				parserCreator.EXPECT().Accept("format").Return(true),
			)
			gomock.InOrder(
				parserCreator.EXPECT().Create(file1).Return(parser1, nil),
				parserCreator.EXPECT().Create(file2).Return(parser2, nil),
			)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigDirSource(path, "format", true, fs, parserFactory)
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			default:
				switch {
				case sut.Mutex == nil:
					t.Error("didn't created the access mutex")
				case sut.path != path:
					t.Error("didn't stored the file path")
				case sut.format != "format":
					t.Error("didn't stored the file content format")
				case sut.fileSystem != fs:
					t.Error("didn't stored the file system adapter reference")
				case sut.parserFactory != parserFactory:
					t.Error("didn't stored the parser factory reference")
				case !reflect.DeepEqual(sut.Partial, expected):
					t.Errorf("(%v) when expecting (%v)", sut.Partial, expected)
				}
			}
		})
	})
}

func Test_ConfigDirSourceCreator(t *testing.T) {
	t.Run("NewConfigDirSourceCreator", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigDirSourceCreator(nil, NewConfigParserFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigDirSourceCreator(NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new file source factory creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fs := NewMockFs(ctrl)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigDirSourceCreator(fs, parserFactory)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case sut.fileSystem != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.parserFactory != parserFactory:
				t.Error("didn't stored the parser factory reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewConfigDirSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

		t.Run("don't accept on invalid config pointer", func(t *testing.T) {
			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is missing", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not a string", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if invalid type", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": ConfigTypeFile}) {
				t.Error("returned true")
			}
		})

		t.Run("accept config", func(t *testing.T) {
			if !sut.Accept(&ConfigPartial{"type": ConfigTypeDir}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigDirSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("missing path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigDirSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"format":    "format",
				"recursive": true,
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("non-string path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigDirSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"path":      123,
				"format":    "format",
				"recursive": true,
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigDirSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"path":      "path",
				"format":    123,
				"recursive": true,
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-bool recursive flag", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigDirSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"path":      "path",
				"format":    123,
				"recursive": "true",
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-bool recursive flag", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigDirSourceCreator(NewMockFs(ctrl), NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"path":      "path",
				"format":    "format",
				"recursive": "true",
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the dir source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := ConfigPartial{"field": "value"}
			fileInfoName := "file.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&expected, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigDirSourceCreator(fs, parserFactory)

			src, e := sut.Create(&ConfigPartial{
				"path":      path,
				"format":    "format",
				"recursive": true,
			})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigDirSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new file src")
				}
			}
		})

		t.Run("create the dir source defaulting format if not present in config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path"
			expected := ConfigPartial{"field": "value"}
			fileInfoName := "file.yaml"
			fileInfo := NewMockFileInfo(ctrl)
			fileInfo.EXPECT().IsDir().Return(false).Times(1)
			fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
			dir := NewMockFile(ctrl)
			dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
			dir.EXPECT().Close().Return(nil).Times(1)
			file := NewMockFile(ctrl)
			fs := NewMockFs(ctrl)
			fs.EXPECT().Open(path).Return(dir, nil).Times(1)
			fs.
				EXPECT().
				OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).
				Return(file, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&expected, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept(ConfigFormatYAML).Return(true).Times(1)
			parserCreator.EXPECT().Create(file).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigDirSourceCreator(fs, parserFactory)

			src, e := sut.Create(&ConfigPartial{
				"path":      path,
				"recursive": true,
			})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigDirSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new file src")
				}
			}
		})
	})
}

func Test_ConfigRestSource(t *testing.T) {
	t.Run("NewConfigRestSource", func(t *testing.T) {
		t.Run("nil client", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigRestSource(nil, "uri", ConfigFormatJSON, parserFactory, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			client := NewMockConfigRestRequester(ctrl)

			sut, e := NewConfigRestSource(client, "uri", ConfigFormatJSON, nil, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error while creating the request object", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf(`parse "\n": net/url: invalid control character in URL`)
			client := NewMockConfigRestRequester(ctrl)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigRestSource(client, "\n", ConfigFormatJSON, parserFactory, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error executing the http request", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf(`test exception`)
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(nil, expected).Times(1)
			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigRestSource(client, "uri", ConfigFormatJSON, parserFactory, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("unable to get a format parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path"`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept(ConfigFormatJSON).Return(false).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigRestSource(client, "uri", ConfigFormatJSON, parserFactory, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigFormat)
			}
		})

		t.Run("invalid json body", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf(`error message`)
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path"`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(nil, expected).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigRestSource(client, "uri", "format", parserFactory, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("response path not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"other_path": 123}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigRestSource(client, "uri", "format", parserFactory, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("response invalid path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path": 123}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{"path": 123}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigRestSource(client, "uri", "format", parserFactory, "path.node")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("response path not pointing to a config Partial", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path": 123}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{"path": 123}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigRestSource(client, "uri", "format", parserFactory, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("correctly load", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"field": "data"}
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "data"}}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{"path": expected}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigRestSource(client, "uri", "format", parserFactory, "path")
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Error("didn't correctly stored the parsed partial")
			}
		})

		t.Run("correctly load complex path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"field": "data"}
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"node": {"inner_node": {"field": "data"}}}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{
				"node": ConfigPartial{
					"inner_node": expected,
				},
			}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, e := NewConfigRestSource(client, "uri", "format", parserFactory, "node..inner_node")
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Error("didn't correctly stored the parsed partial")
			}
		})
	})
}

func Test_ConfigRestSourceCreator(t *testing.T) {
	t.Run("NewConfigRestSourceCreator", func(t *testing.T) {
		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigRestSourceCreator(nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new rest source factory creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigRestSourceCreator(parserFactory)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case sut.parserFactory != parserFactory:
				t.Error("didn't stored the parser factory reference")
			default:
				client := sut.clientFactory()
				switch client.(type) {
				case *http.Client:
				default:
					t.Error("didn't stored a valid http client")
				}
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		t.Run("don't accept on invalid config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is missing", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			if sut.Accept(&ConfigPartial{}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not a string", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if invalid type", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			if sut.Accept(&ConfigPartial{"type": ConfigTypeFile}) {
				t.Error("returned true")
			}
		})

		t.Run("accept config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			if !sut.Accept(&ConfigPartial{"type": ConfigTypeRest}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("missing uri", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"format":     "format",
				"configPath": "path",
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("missing config path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    "path",
				"format": "format",
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("non-string uri", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":        123,
				"format":     "format",
				"configPath": "path",
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":        "uri",
				"format":     123,
				"configPath": "path",
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    "uri",
				"format": "format",
				"path": ConfigPartial{
					"config": 123,
				},
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the rest source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uri := "uri"
			format := "format"
			path := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			respData := ConfigPartial{"path": ConfigPartial{"field": "value"}}
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&respData, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigRestSourceCreator(parserFactory)
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			sut.clientFactory = func() configRestRequester { return client }

			src, e := sut.Create(&ConfigPartial{
				"uri":    uri,
				"format": format,
				"path": ConfigPartial{
					"config": path,
				},
			})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigRestSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new rest source")
				}
			}
		})

		t.Run("create the rest source defaulting format if not present in config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uri := "uri"
			path := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			respData := ConfigPartial{"path": ConfigPartial{"field": "value"}}
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&respData, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("json").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigRestSourceCreator(parserFactory)
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			sut.clientFactory = func() configRestRequester { return client }

			src, e := sut.Create(&ConfigPartial{
				"uri": uri,
				"path": ConfigPartial{
					"config": path,
				},
			})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigRestSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new rest source")
				}
			}
		})
	})
}

func Test_ConfigObsRestSource(t *testing.T) {
	t.Run("NewConfigObsRestSource", func(t *testing.T) {
		t.Run("nil client", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigObsRestSource(
				nil,
				"uri",
				"format",
				NewConfigParserFactory(nil),
				"timestampPath",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigObsRestSource(
				NewMockConfigRestRequester(ctrl),
				"uri",
				"format",
				nil,
				"timestampPath",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error while creating the request object", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf(`parse "\n": net/url: invalid control character in URL`)

			sut, e := NewConfigObsRestSource(
				NewMockConfigRestRequester(ctrl),
				"\n",
				"format",
				NewConfigParserFactory(nil),
				"timestampPath",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error executing the http request", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf(`test exception`)
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(nil, expected).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"format",
				NewConfigParserFactory(nil),
				"timestampPath",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})
		t.Run("unable to get a format parser", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path"`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"format",
				NewConfigParserFactory(nil),
				"timestampPath",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigFormat)
			}
		})

		t.Run("error decoding body", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf(`error message`)
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"path"`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(nil, expected).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestampPath",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("response timestamp path not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"other_path": 123}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestampPath",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("invalid timestamp value type", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"timestamp": 123}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{"timestamp": 123}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid timestamp value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := "parsing time \"abc\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"abc\" as \"2006\""
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "abc"}`))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{"timestamp": "abc"}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected:
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("response config path not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(
				`{"timestamp": "2000-01-01T00:00:00Z", other_path": 123}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
			}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"configPath",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("response invalid path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(
				`{"timestamp": "2000-01-01T00:00:00Z", "path": 123}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
				"path":      123,
			}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"path.node",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("response path not pointing to a config Partial", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(
				`{"timestamp": "2000-01-01T00:00:00Z", "path": 123}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
				"path":      123,
			}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"path",
			)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("correctly load", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"field": "data"}
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(
				`{"timestamp": "2000-01-01T00:00:00Z", "path": {"field": "data"}}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
				"path":      expected,
			}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"path",
			)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Error("didn't correctly stored the parsed partial")
			}
		})

		t.Run("correctly load complex path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"field": "data"}
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(
				`{"timestamp": "2000-01-01T00:00:00Z", "node": {"inner_node": {"field": "data"}}}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
				"node": ConfigPartial{
					"inner_node": expected,
				},
			}, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("yaml").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)

			sut, e := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"node..inner_node",
			)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Error("didn't correctly stored the parsed partial")
			}
		})
	})

	t.Run("Reload", func(t *testing.T) {
		t.Run("dont reload on same timestamp", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"field": "data 1"}
			response1 := http.Response{}
			response1.Body = io.NopCloser(strings.NewReader(
				`{"node": {"field": "data 1"}, "timestamp": "2021-12-15T21:07:48.239Z"}`,
			))
			response2 := http.Response{}
			response2.Body = io.NopCloser(strings.NewReader(
				`{"node": {"field": "data 2"}, "timestamp": "2021-12-15T21:07:48.239Z"}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			gomock.InOrder(
				client.EXPECT().Do(gomock.Any()).Return(&response1, nil),
				client.EXPECT().Do(gomock.Any()).Return(&response2, nil),
			)
			parser1 := NewMockConfigParser(ctrl)
			parser1.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
				"node":      expected,
			}, nil).Times(1)
			parser1.EXPECT().Close().Return(nil).Times(1)
			parser2 := NewMockConfigParser(ctrl)
			parser2.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
				"node": ConfigPartial{
					"field": "data 2",
				},
			}, nil).Times(1)
			parser2.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			gomock.InOrder(
				parserCreator.EXPECT().Accept("yaml").Return(true),
				parserCreator.EXPECT().Accept("yaml").Return(true),
			)
			gomock.InOrder(
				parserCreator.EXPECT().Create(gomock.Any()).Return(parser1, nil),
				parserCreator.EXPECT().Create(gomock.Any()).Return(parser2, nil),
			)

			sut, _ := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"node",
			)

			loaded, e := sut.Reload()
			switch {
			case loaded != false:
				t.Error("unexpectedly reload the source config")
			case e != nil:
				t.Errorf("returned the eunexpected e : %v", e)
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Error("didn't correctly stored the parsed partial")
			}
		})

		t.Run("correctly reload config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := ConfigPartial{"field": "data 2"}
			response1 := http.Response{}
			response1.Body = io.NopCloser(strings.NewReader(
				`{"node": {"field": "data 1"}, "timestamp": "2021-12-15T21:07:48.239Z"}`,
			))
			response2 := http.Response{}
			response2.Body = io.NopCloser(strings.NewReader(
				`{"node": {"field": "data 2"}, "timestamp": "2021-12-15T21:07:48.240Z"}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			gomock.InOrder(
				client.EXPECT().Do(gomock.Any()).Return(&response1, nil),
				client.EXPECT().Do(gomock.Any()).Return(&response2, nil),
			)
			parser1 := NewMockConfigParser(ctrl)
			parser1.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:00Z",
				"node": ConfigPartial{
					"field": "data 1",
				},
			}, nil).Times(1)
			parser1.EXPECT().Close().Return(nil).Times(1)
			parser2 := NewMockConfigParser(ctrl)
			parser2.EXPECT().Parse().Return(&ConfigPartial{
				"timestamp": "2000-01-01T00:00:01Z",
				"node":      expected,
			}, nil).Times(1)
			parser2.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			gomock.InOrder(
				parserCreator.EXPECT().Accept("yaml").Return(true),
				parserCreator.EXPECT().Accept("yaml").Return(true),
			)
			gomock.InOrder(
				parserCreator.EXPECT().Create(gomock.Any()).Return(parser1, nil),
				parserCreator.EXPECT().Create(gomock.Any()).Return(parser2, nil),
			)

			sut, _ := NewConfigObsRestSource(
				client,
				"uri",
				"yaml",
				NewConfigParserFactory([]ConfigParserCreator{parserCreator}),
				"timestamp",
				"node",
			)

			loaded, e := sut.Reload()
			switch {
			case loaded != true:
				t.Error("didn't reload the source config")
			case e != nil:
				t.Errorf("returned the eunexpected e : %v", e)
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Error("didn't correctly stored the parsed partial")
			}
		})
	})
}

func Test_ConfigObsRestSourceCreator(t *testing.T) {
	t.Run("NewConfigObsRestSourceCreator", func(t *testing.T) {
		t.Run("nil parser factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewConfigObsRestSourceCreator(nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new observable rest source factory creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			parserFactory := NewConfigParserFactory(nil)

			sut, e := NewConfigObsRestSourceCreator(parserFactory)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case sut.parserFactory != parserFactory:
				t.Error("didn't stored the parser factory reference")
			default:
				client := sut.clientFactory()
				switch client.(type) {
				case *http.Client:
				default:
					t.Error("didn't stored a valid http client factory")
				}
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

		t.Run("don't accept on invalid config pointer", func(t *testing.T) {
			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is missing", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if type is not a string", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept if invalid type", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": ConfigTypeRest}) {
				t.Error("returned true")
			}
		})

		t.Run("accept config", func(t *testing.T) {
			if !sut.Accept(&ConfigPartial{"type": ConfigTypeObsRest}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("missing uri", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"format":        "format",
				"timestampPath": "path",
				"configPath":    "path",
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("missing timestamp path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    "uri",
				"format": "format",
				"path": ConfigPartial{
					"config": "path",
				},
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("missing config path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    "uri",
				"format": "format",
				"path": ConfigPartial{
					"timestamp": "path",
				},
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidConfigSupplier):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidConfigSupplier)
			}
		})

		t.Run("non-string uri", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    123,
				"format": "format",
				"path": ConfigPartial{
					"config":    "path",
					"timestamp": "path",
				},
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    "uri",
				"format": 123,
				"path": ConfigPartial{
					"config":    "path",
					"timestamp": "path",
				},
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string timestamp path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    "uri",
				"format": "format",
				"path": ConfigPartial{
					"config":    "path",
					"timestamp": 123,
				},
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-string config path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewConfigObsRestSourceCreator(NewConfigParserFactory(nil))

			src, e := sut.Create(&ConfigPartial{
				"uri":    "uri",
				"format": "format",
				"path": ConfigPartial{
					"config":    123,
					"timestamp": "path",
				},
			})
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("create the rest source", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uri := "uri"
			format := "format"
			timestampPath := "timestamp"
			configPath := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			respData := ConfigPartial{
				"path":      ConfigPartial{"field": "value"},
				"timestamp": "2000-01-01T00:00:00.000Z",
			}
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&respData, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("format").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsRestSourceCreator(parserFactory)
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(
				`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			sut.clientFactory = func() configRestRequester { return client }

			src, e := sut.Create(&ConfigPartial{
				"uri":    uri,
				"format": format,
				"path": ConfigPartial{
					"config":    configPath,
					"timestamp": timestampPath,
				},
			})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigObsRestSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new rest src")
				}
			}
		})

		t.Run("create the rest source defaulting format if not present in config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uri := "uri"
			timestampPath := "timestamp"
			configPath := "path"
			field := "field"
			value := "value"
			expected := ConfigPartial{field: value}
			respData := ConfigPartial{
				"path":      ConfigPartial{"field": "value"},
				"timestamp": "2000-01-01T00:00:00.000Z",
			}
			parser := NewMockConfigParser(ctrl)
			parser.EXPECT().Parse().Return(&respData, nil).Times(1)
			parser.EXPECT().Close().Return(nil).Times(1)
			parserCreator := NewMockConfigParserCreator(ctrl)
			parserCreator.EXPECT().Accept("json").Return(true).Times(1)
			parserCreator.EXPECT().Create(gomock.Any()).Return(parser, nil).Times(1)
			parserFactory := NewConfigParserFactory([]ConfigParserCreator{parserCreator})

			sut, _ := NewConfigObsRestSourceCreator(parserFactory)
			response := http.Response{}
			response.Body = io.NopCloser(strings.NewReader(
				`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`,
			))
			client := NewMockConfigRestRequester(ctrl)
			client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
			sut.clientFactory = func() configRestRequester { return client }

			src, e := sut.Create(&ConfigPartial{
				"uri": uri,
				"path": ConfigPartial{
					"config":    configPath,
					"timestamp": timestampPath,
				},
			})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case src == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := src.(type) {
				case *ConfigObsRestSource:
					if !reflect.DeepEqual(s.Partial, expected) {
						t.Error("didn't loaded the content correctly")
					}
				default:
					t.Error("didn't returned a new rest src")
				}
			}
		})
	})
}

func Test_Config(t *testing.T) {
	t.Run("NewConfig", func(t *testing.T) {
		t.Run("new config without reload", func(t *testing.T) {
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			switch {
			case sut.mutex == nil:
				t.Error("didn't instantiate the access mutex")
			case sut.suppliers == nil:
				t.Error("didn't instantiate the suppliers storing array")
			case sut.observers == nil:
				t.Error("didn't instantiate the observers storing array")
			case sut.observer != nil:
				t.Error("instantiated the suppliers reload trigger")
			}
		})

		t.Run("new config with reload", func(t *testing.T) {
			ConfigObserveFrequency = 10
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			switch {
			case sut.mutex == nil:
				t.Error("didn't instantiate the access mutex")
			case sut.suppliers == nil:
				t.Error("didn't instantiate the suppliers storing array")
			case sut.observers == nil:
				t.Error("didn't instantiate the observers storing array")
			case sut.observer == nil:
				t.Error("didn't instantiate the suppliers reload trigger")
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("error while closing supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")

			ConfigObserveFrequency = 0
			sut := NewConfig()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).AnyTimes()
			supplier.EXPECT().Close().Return(expected).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if e := sut.Close(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})
		t.Run("error while closing observer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			ConfigObserveFrequency = 0
			sut := NewConfig()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).AnyTimes()
			supplier.EXPECT().Close().Return(nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)
			observer := NewMockTrigger(ctrl)
			observer.EXPECT().Close().Return(expected).Times(1)
			sut.observer = observer

			if e := sut.Close(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("propagate close to suppliers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			id1 := "supplier.1"
			id2 := "supplier.2"
			priority1 := 0
			priority2 := 1
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(ConfigPartial{}, nil).AnyTimes()
			supplier1.EXPECT().Close().Times(1)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(ConfigPartial{}, nil).AnyTimes()
			supplier2.EXPECT().Close().Times(1)
			_ = sut.AddSupplier(id1, priority1, supplier1)
			_ = sut.AddSupplier(id2, priority2, supplier2)

			_ = sut.Close()
		})
	})

	t.Run("Entries", func(t *testing.T) {
		t.Run("return partial entries", func(t *testing.T) {
			scenarios := []struct {
				config   ConfigPartial
				expected []string
			}{
				{ // _test the empty partial
					config:   ConfigPartial{},
					expected: nil,
				},
				{ // _test the single entry partial
					config:   ConfigPartial{"field": "value"},
					expected: []string{"field"},
				},
				{ // _test the multi entry partial
					config:   ConfigPartial{"field1": "value 1", "field2": "value 2"},
					expected: []string{"field1", "field2"},
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)

					ConfigObserveFrequency = 0
					sut := NewConfig()
					supplier := NewMockConfigSupplier(ctrl)
					supplier.EXPECT().Close().Times(1)
					supplier.EXPECT().Get("").Return(s.config, nil).Times(1)
					_ = sut.AddSupplier("supplier", 0, supplier)

					defer func() { _ = sut.Close(); ctrl.Finish() }()

					check := sut.Entries()

					sort.Strings(s.expected)
					sort.Strings(check)
					if !reflect.DeepEqual(check, s.expected) {
						t.Errorf("(%v) when expecting (%v)", check, s.expected)
					}
				}
				test()
			}
		})
	})

	t.Run("Has", func(t *testing.T) {
		t.Run("return the existence of the path", func(t *testing.T) {
			scenarios := []struct {
				config   ConfigPartial
				search   string
				expected bool
			}{
				{ // _test the existence of a present path
					config:   ConfigPartial{"node": "value"},
					search:   "node",
					expected: true,
				},
				{ // _test the non-existence of a missing path
					config:   ConfigPartial{"node": "value"},
					search:   "invalid-node",
					expected: false,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)

					ConfigObserveFrequency = 0
					sut := NewConfig()
					supplier := NewMockConfigSupplier(ctrl)
					supplier.EXPECT().Close().Times(1)
					supplier.EXPECT().Get("").Return(s.config, nil).Times(1)
					_ = sut.AddSupplier("supplier", 0, supplier)

					defer func() { _ = sut.Close(); ctrl.Finish() }()

					if check := sut.Has(s.search); check != s.expected {
						t.Errorf("(%v) when expecting (%v)", check, s.expected)
					}
				}
				test()
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("return path value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			search := "node"
			expected := "value"
			config := ConfigPartial{search: expected}

			ConfigObserveFrequency = 0
			sut := NewConfig()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(config, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			defer func() { _ = sut.Close() }()

			if check, e := sut.Get(search); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != expected {
				t.Errorf("(%v) when expecting (%v)", check, expected)
			}
		})

		t.Run("return internal partial get error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data := ConfigPartial{"node1": ConfigPartial{"node2": 101}}
			path := "node3"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(data, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			check, e := sut.Get(path)
			switch {
			case check != nil:
				t.Errorf("unexpected valid value : %v", check)
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("return simple if path was not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			data := ConfigPartial{"node1": ConfigPartial{"node2": 101}}
			path := "node3"
			val := 3
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(data, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.Get(path, val); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != val {
				t.Errorf("(%v) when expecting (%v)", check, val)
			}
		})
	})

	t.Run("Bool", func(t *testing.T) {
		t.Run("return the stored boolean value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "node"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{path: true}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.Bool(path); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !check {
				t.Errorf("returned (%v)", check)
			}
		})
	})

	t.Run("Int", func(t *testing.T) {
		t.Run("return the stored integer value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := 123
			path := "node"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{path: value}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.Int(path); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != value {
				t.Errorf("returned (%v) when expecting : %v", check, value)
			}
		})
	})

	t.Run("Float", func(t *testing.T) {
		t.Run("return the stored integer value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := 123.4
			path := "node"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{path: value}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.Float(path); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != value {
				t.Errorf("returned (%v) when expecting : %v", check, value)
			}
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("return the stored integer value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := "value"
			path := "node"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{path: value}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.String(path); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if check != value {
				t.Errorf("returned (%v) when expecting : %v", check, value)
			}
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("return the stored integer value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := []interface{}{1, 2, 3}
			path := "node"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{path: value}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.List(path); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(check, value) {
				t.Errorf("returned (%v) when expecting : %v", check, value)
			}
		})
	})

	t.Run("Partial", func(t *testing.T) {
		t.Run("return the stored partial value", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := ConfigPartial{"field": "value"}
			path := "node"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{path: value}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.Partial(path); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(check, value) {
				t.Errorf("returned (%v) when expecting : %v", check, value)
			}
		})
	})

	t.Run("Populate", func(t *testing.T) {
		t.Run("populate the given structure", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			value := ConfigPartial{"field": ConfigPartial{"field": "value"}}
			target := struct{ Field string }{}
			expected := struct{ Field string }{Field: "value"}
			path := "node"
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{path: value}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if check, e := sut.Populate(path+"."+"field", target); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(check, expected) {
				t.Errorf("returned (%v) when expecting : %v", check, expected)
			}
		})
	})

	t.Run("HasSupplier", func(t *testing.T) {
		t.Run("validate if the supplier is registered", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if !sut.HasSupplier("supplier") {
				t.Error("returned false")
			}
		})

		t.Run("invalidate if the supplier is not registered", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if sut.HasSupplier("invalid supplier id") {
				t.Error("returned true")
			}
		})
	})

	t.Run("AddSupplier", func(t *testing.T) {
		t.Run("nil supplier", func(t *testing.T) {
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			if e := sut.AddSupplier("supplier", 0, nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register a new supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)

			if e := sut.AddSupplier("supplier", 0, supplier); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !sut.HasSupplier("supplier") {
				t.Error("didn't stored the supplier")
			}
		})

		t.Run("duplicate id", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			if e := sut.AddSupplier("supplier", 0, supplier); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrDuplicateConfigSupplier) {
				t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateConfigSupplier)
			}
		})

		t.Run("override path if the insert have higher priority", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().Times(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().Times(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 1, supplier1)
			_ = sut.AddSupplier("supplier.2", 2, supplier2)

			if check, _ := sut.Get("node"); check != "value.2" {
				t.Errorf("returned the (%v) value when expecting (value.2)", check)
			}
		})

		t.Run("do not override path if the insert have lower priority", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().Times(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().Times(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 2, supplier1)
			_ = sut.AddSupplier("supplier.2", 1, supplier2)

			if check, _ := sut.Get("node"); check != "value.1" {
				t.Errorf("returned the (%v) value when expecting (value.1)", check)
			}
		})

		t.Run("still be able to get not overridden paths of a inserted lower priority", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().Times(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().Times(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2", "extendedNode": "extendedValue"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 2, supplier1)
			_ = sut.AddSupplier("supplier.2", 1, supplier2)

			if check, _ := sut.Get("extendedNode"); check != "extendedValue" {
				t.Errorf("returned the (%v) value when expecting (extendedValue)", check)
			}
		})
	})

	t.Run("RemoveSupplier", func(t *testing.T) {
		t.Run("unregister a non-registered supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			if e := sut.RemoveSupplier("supplier"); e != nil {
				t.Errorf("unexpected error (%v)", e)
			}
		})

		t.Run("error unregister a previously registered supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Return(expected).Times(2)
			supplier.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier", 0, supplier)

			if e := sut.RemoveSupplier("supplier"); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("unregister a previously registered supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().Times(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().Times(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2"}, nil).AnyTimes()
			supplier3 := NewMockConfigSupplier(ctrl)
			supplier3.EXPECT().Close().Times(1)
			supplier3.EXPECT().Get("").Return(ConfigPartial{"node": "value.3"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 0, supplier1)
			_ = sut.AddSupplier("supplier.2", 0, supplier2)
			_ = sut.AddSupplier("supplier.3", 0, supplier3)
			_ = sut.RemoveSupplier("supplier.2")

			if sut.HasSupplier("supplier.2") {
				t.Error("didn't remove the supplier")
			}
		})

		t.Run("recover path overridden by the removed supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().Times(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().Times(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2"}, nil).AnyTimes()
			supplier3 := NewMockConfigSupplier(ctrl)
			supplier3.EXPECT().Close().Times(1)
			supplier3.EXPECT().Get("").Return(ConfigPartial{"node": "value.3"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 0, supplier1)
			_ = sut.AddSupplier("supplier.2", 1, supplier2)
			_ = sut.AddSupplier("supplier.3", 2, supplier3)
			_ = sut.RemoveSupplier("supplier.3")

			if check, _ := sut.Get("node"); check != "value.2" {
				t.Errorf("returned (%check) value when expecting (value.2)", check)
			}
		})
	})

	t.Run("RemoveAllSuppliers", func(t *testing.T) {
		t.Run("remove all the suppliers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			expected := fmt.Errorf("error string")
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().MinTimes(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().MinTimes(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2"}, nil).AnyTimes()
			supplier3 := NewMockConfigSupplier(ctrl)
			supplier3.EXPECT().Close().Return(expected).MinTimes(1)
			supplier3.EXPECT().Get("").Return(ConfigPartial{"node": "value.3"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 0, supplier1)
			_ = sut.AddSupplier("supplier.2", 1, supplier2)
			_ = sut.AddSupplier("supplier.3", 2, supplier3)

			if e := sut.RemoveAllSuppliers(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("remove all the suppliers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().Times(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().Times(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2"}, nil).AnyTimes()
			supplier3 := NewMockConfigSupplier(ctrl)
			supplier3.EXPECT().Close().Times(1)
			supplier3.EXPECT().Get("").Return(ConfigPartial{"node": "value.3"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 0, supplier1)
			_ = sut.AddSupplier("supplier.2", 1, supplier2)
			_ = sut.AddSupplier("supplier.3", 2, supplier3)
			_ = sut.RemoveAllSuppliers()

			if len(sut.suppliers) != 0 {
				t.Error("didn't removed all the registered suppliers")
			}
		})
	})

	t.Run("Supplier", func(t *testing.T) {
		t.Run("error if the supplier don't exists", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			check, e := sut.Supplier("invalid id")
			switch {
			case check != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConfigSupplierNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigSupplierNotFound)
			}
		})

		t.Run("return the registered supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			check, e := sut.Supplier("supplier")
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case check == nil:
				t.Error("returned nil")
			case !reflect.DeepEqual(check, supplier):
				t.Errorf("(%v) when expecting (%v)", check, supplier)
			}
		})
	})

	t.Run("SupplierPriority", func(t *testing.T) {
		t.Run("error if the supplier was not found", func(t *testing.T) {
			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			if e := sut.SupplierPriority("invalid id", 0); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrConfigSupplierNotFound) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigSupplierNotFound)
			}
		})

		t.Run("update the priority of the supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 0
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Close().Times(1)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": "value.1"}, nil).AnyTimes()
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Close().Times(1)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value.2"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier.1", 1, supplier1)
			_ = sut.AddSupplier("supplier.2", 2, supplier2)

			if check, _ := sut.Get("node"); check != "value.2" {
				t.Errorf("returned the (%v) value prior the change, when expecting (value.2)", check)
			}
			if e := sut.SupplierPriority("supplier.2", 0); e != nil {
				t.Errorf("returned the unexpeced error : (%v)", e)
			}
			if check, _ := sut.Get("node"); check != "value.1" {
				t.Errorf("returned the (%v) value after the change, when expecting (value.1)", check)
			}
		})
	})

	t.Run("HasObserver", func(t *testing.T) {
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

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer ctrl.Finish()

					ConfigObserveFrequency = 0
					sut := NewConfig()
					defer func() { _ = sut.Close() }()
					supplier := NewMockConfigSupplier(ctrl)
					supplier.EXPECT().Close().Times(1)
					supplier.EXPECT().Get("").Return(ConfigPartial{
						"node1": "value1",
						"node2": "value2",
						"node3": "value3",
					}, nil).Times(1)
					_ = sut.AddSupplier("config", 0, supplier)

					for _, observer := range s.observers {
						_ = sut.AddObserver(observer, func(old, new interface{}) {})
					}

					if check := sut.HasObserver(s.search); check != s.exp {
						t.Errorf("(%v) when expecting (%v)", check, s.exp)
					}
				}
				test()
			}
		})
	})

	t.Run("AddObserver", func(t *testing.T) {
		t.Run("nil callback", func(t *testing.T) {
			ConfigObserveFrequency = 60
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			if e := sut.AddObserver("path", nil); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error if path not present", func(t *testing.T) {
			ConfigObserveFrequency = 60
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			if e := sut.AddObserver("path", func(interface{}, interface{}) {
			}); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrConfigPathNotFound) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("valid callback", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 60
			partial := ConfigPartial{"path": "value"}
			sut := NewConfig()
			defer func() { _ = sut.Close() }()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			_ = sut.AddSupplier("config", 0, supplier)

			if e := sut.AddObserver("path", func(interface{}, interface{}) {
			}); e != nil {
				t.Errorf("unexpected error, %v", e)
			} else if len(sut.observers) != 1 {
				t.Error("didn't stored the requested observer")
			}
		})
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		t.Run("remove a registered observer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 60
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			partial := ConfigPartial{
				"node": ConfigPartial{
					"1": "value1",
					"2": "value2",
					"3": "value3",
				},
			}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			_ = sut.AddSupplier("config", 0, supplier)

			_ = sut.AddObserver("node.1", func(old, new interface{}) {})
			_ = sut.AddObserver("node.2", func(old, new interface{}) {})
			_ = sut.AddObserver("node.3", func(old, new interface{}) {})
			sut.RemoveObserver("node.2")

			if sut.HasObserver("node.2") {
				t.Errorf("didn't removed the observer")
			}
		})
	})

	t.Run("running", func(t *testing.T) {
		t.Run("reload on observable suppliers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 10
			sut := NewConfig()
			defer func() { _ = sut.Close() }()

			partial := ConfigPartial{"node": "value"}
			supplier := NewMockConfigObsSupplier(ctrl)
			supplier.EXPECT().Close().Times(1)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			supplier.EXPECT().Reload().Return(false, nil).MinTimes(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			time.Sleep(100 * time.Millisecond)
		})

		t.Run("rebuild if the observable supplier notify changes", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigObserveFrequency = 10
			sut := NewConfig()

			partial := ConfigPartial{"node": "value"}
			supplier := NewMockConfigObsSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).MinTimes(2)
			supplier.EXPECT().Reload().Return(true, nil).MinTimes(1)
			_ = sut.AddSupplier("supplier", 0, supplier)

			time.Sleep(200 * time.Millisecond)

			if check, _ := sut.Get("node"); check != "value" {
				t.Errorf("returned (%v) when expecting (value)", check)
			}
		})

		t.Run("should call observer callback function on partial changes", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			check := false
			ConfigObserveFrequency = 10
			sut := NewConfig()

			partial := ConfigPartial{"node": "value1"}
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(partial, nil).AnyTimes()
			_ = sut.AddSupplier("supplier1", 0, supplier1)

			_ = sut.AddObserver("node", func(old, new interface{}) {
				check = true

				if old != "value1" {
					t.Errorf("callback called with (%v) as old value", old)
				} else if new != "value2" {
					t.Errorf("callback called with (%v) as new value", new)
				}
			})

			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": "value2"}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier2", 1, supplier2)

			if !check {
				t.Errorf("didn't actually called the callback")
			} else if check := sut.observers[0].current; check != "value2" {
				t.Errorf("stored {%v} instead of {%v}", check, "value2")
			}
		})

		t.Run("should call observer callback function on partial changes on a list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			check := false
			ConfigObserveFrequency = 10
			sut := NewConfig()
			initial := []interface{}{ConfigPartial{"sub_node": "value1"}}
			expected := []interface{}{ConfigPartial{"sub_node": "value2"}}

			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": initial}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier1", 0, supplier1)

			_ = sut.AddObserver("node", func(old, new interface{}) {
				check = true

				if old.([]interface{})[0].(ConfigPartial)["sub_node"] != initial[0].(ConfigPartial)["sub_node"] {
					t.Errorf("callback called with (%v) as old value", old)
				} else if new.([]interface{})[0].(ConfigPartial)["sub_node"] != expected[0].(ConfigPartial)["sub_node"] {
					t.Errorf("callback called with (%v) as new value", new)
				}
			})

			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": expected}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier2", 1, supplier2)

			if !check {
				t.Errorf("didn't actually called the callback")
			} else if check := sut.observers[0].current; check.([]interface{})[0].(ConfigPartial)["sub_node"] != expected[0].(ConfigPartial)["sub_node"] {
				t.Errorf("stored {%v} instead of {%v}", check, expected)
			}
		})

		t.Run("should call observer callback function on partial changes on a partial", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			check := false
			ConfigObserveFrequency = 10
			sut := NewConfig()
			initial := ConfigPartial{"sub_node": "value1"}
			expected := ConfigPartial{"sub_node": "value2"}

			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(ConfigPartial{"node": initial}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier1", 0, supplier1)

			_ = sut.AddObserver("node", func(old, new interface{}) {
				check = true

				if reflect.DeepEqual(old, initial) {
					t.Errorf("callback called with (%v) as old value", old)
				} else if reflect.DeepEqual(old, expected) {
					t.Errorf("callback called with (%v) as new value", new)
				}
			})

			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(ConfigPartial{"node": expected}, nil).AnyTimes()
			_ = sut.AddSupplier("supplier2", 1, supplier2)

			if !check {
				t.Errorf("didn't actually called the callback")
			} else if check := sut.observers[0].current; check.(ConfigPartial)["sub_node"] != expected["sub_node"] {
				t.Errorf("stored {%v} instead of {%v}", check, expected)
			}
		})
	})
}

func Test_ConfigLoader(t *testing.T) {
	t.Run("NewConfigLoader", func(t *testing.T) {
		t.Run("nil config", func(t *testing.T) {
			sut, e := NewConfigLoader(nil, NewConfigSupplierFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil supplier factory", func(t *testing.T) {
			sut, e := NewConfigLoader(NewConfig(), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new observer", func(t *testing.T) {
			if sut, e := NewConfigLoader(NewConfig(), NewConfigSupplierFactory(nil)); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("Load", func(t *testing.T) {
		ConfigLoaderSupplierID = "base_supplier_id"
		ConfigLoaderFileSupplierPath = "base_supplier_path"
		ConfigLoaderSupplierFormat = "format"
		defer func() {
			ConfigLoaderSupplierID = "main"
			ConfigLoaderFileSupplierPath = "partial/suppliers.yaml"
			ConfigLoaderSupplierFormat = "format"
		}()
		baseSupplierPartial := ConfigPartial{
			"type":   "file",
			"path":   "base_supplier_path",
			"format": "format",
		}

		t.Run("error getting the base supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true).Times(1)
			supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(nil, expected).Times(1)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})

			sut, _ := NewConfigLoader(NewConfig(), supplierFactory)

			if e := sut.Load(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error storing the base supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true).Times(1)
			supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(supplier, nil).Times(1)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})
			config := NewConfig()
			_ = config.AddSupplier(ConfigLoaderSupplierID, 0, supplier)

			sut, _ := NewConfigLoader(config, supplierFactory)

			if e := sut.Load(); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrDuplicateConfigSupplier) {
				t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateConfigSupplier)
			}
		})

		t.Run("add base supplier into the partial", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true).Times(1)
			supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(supplier, nil).Times(1)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})
			config := NewConfig()

			sut, _ := NewConfigLoader(config, supplierFactory)

			if e := sut.Load(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if len(config.suppliers) == 0 {
				t.Error("didn't stored the base supplier")
			}
		})

		t.Run("invalid list of suppliers results in an empty suppliers list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.config.suppliers", "string")
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true).Times(1)
			supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(supplier, nil).Times(1)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})

			sut, _ := NewConfigLoader(NewConfig(), supplierFactory)

			if e := sut.Load(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error parsing supplier entry", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			entry := ConfigPartial{"priority": "string"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.config.suppliers", ConfigPartial{"supplier": entry})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true).Times(1)
			supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(supplier, nil).Times(1)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})

			sut, _ := NewConfigLoader(NewConfig(), supplierFactory)

			if e := sut.Load(); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("error creating the supplier entry", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			entry := ConfigPartial{"type": "my type"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.config.suppliers", ConfigPartial{"supplier": entry})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			gomock.InOrder(
				supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true),
				supplierCreator.EXPECT().Accept(&entry).Return(true),
			)
			gomock.InOrder(
				supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(supplier, nil),
				supplierCreator.EXPECT().Create(&entry).Return(nil, expected),
			)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})

			sut, _ := NewConfigLoader(NewConfig(), supplierFactory)

			if e := sut.Load(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error on supplier registration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			supplierEntry := ConfigPartial{"type": "my type"}
			suppliers := ConfigPartial{}
			_, _ = suppliers.Set("slate.config.suppliers", ConfigPartial{"supplier": supplierEntry})
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(suppliers, nil).Times(1)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(supplierEntry, nil).Times(2)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			gomock.InOrder(
				supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true),
				supplierCreator.EXPECT().Accept(&supplierEntry).Return(true),
			)
			gomock.InOrder(
				supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(supplier1, nil),
				supplierCreator.EXPECT().Create(&supplierEntry).Return(supplier2, nil),
			)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier2)

			sut, _ := NewConfigLoader(config, supplierFactory)

			if e := sut.Load(); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrDuplicateConfigSupplier) {
				t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateConfigSupplier)
			}
		})

		t.Run("register the loaded supplier", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			supplierEntry := ConfigPartial{"type": "my type"}
			suppliers := ConfigPartial{}
			_, _ = suppliers.Set("slate.config.suppliers", ConfigPartial{"supplier": supplierEntry})
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(suppliers, nil).Times(2)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(supplierEntry, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			gomock.InOrder(
				supplierCreator.EXPECT().Accept(&baseSupplierPartial).Return(true),
				supplierCreator.EXPECT().Accept(&supplierEntry).Return(true),
			)
			gomock.InOrder(
				supplierCreator.EXPECT().Create(&baseSupplierPartial).Return(supplier1, nil),
				supplierCreator.EXPECT().Create(&supplierEntry).Return(supplier2, nil),
			)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})
			config := NewConfig()

			sut, _ := NewConfigLoader(config, supplierFactory)

			if e := sut.Load(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("load from defined supplier path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := ConfigLoaderFileSupplierPath
			ConfigLoaderFileSupplierPath = "config.yaml"
			defer func() { ConfigLoaderFileSupplierPath = prev }()

			supplierEntry := ConfigPartial{"type": "my type"}
			suppliers := ConfigPartial{}
			baseSourcePartial := ConfigPartial{
				"type":   "file",
				"path":   ConfigLoaderFileSupplierPath,
				"format": "format",
			}
			_, _ = suppliers.Set("slate.config.suppliers", ConfigPartial{"supplier": supplierEntry})
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(suppliers, nil).Times(2)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(supplierEntry, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			gomock.InOrder(
				supplierCreator.EXPECT().Accept(&baseSourcePartial).Return(true),
				supplierCreator.EXPECT().Accept(&supplierEntry).Return(true),
			)
			gomock.InOrder(
				supplierCreator.EXPECT().Create(&baseSourcePartial).Return(supplier1, nil),
				supplierCreator.EXPECT().Create(&supplierEntry).Return(supplier2, nil),
			)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})
			config := NewConfig()

			sut, _ := NewConfigLoader(config, supplierFactory)

			if e := sut.Load(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("load from defined format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := ConfigLoaderSupplierFormat
			ConfigLoaderSupplierFormat = "json"
			defer func() { ConfigLoaderSupplierFormat = prev }()

			supplierEntry := ConfigPartial{"type": "my type"}
			suppliers := ConfigPartial{}
			baseSourcePartial := ConfigPartial{
				"type":   "file",
				"path":   "base_supplier_path",
				"format": ConfigLoaderSupplierFormat,
			}
			_, _ = suppliers.Set("slate.config.suppliers", ConfigPartial{"supplier": supplierEntry})
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(suppliers, nil).Times(2)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(supplierEntry, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			gomock.InOrder(
				supplierCreator.EXPECT().Accept(&baseSourcePartial).Return(true),
				supplierCreator.EXPECT().Accept(&supplierEntry).Return(true),
			)
			gomock.InOrder(
				supplierCreator.EXPECT().Create(&baseSourcePartial).Return(supplier1, nil),
				supplierCreator.EXPECT().Create(&supplierEntry).Return(supplier2, nil),
			)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})
			config := NewConfig()

			sut, _ := NewConfigLoader(config, supplierFactory)

			if e := sut.Load(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("load from defined supplier list path", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := ConfigLoaderSupplierListPath
			ConfigLoaderSupplierListPath = "config_list"
			defer func() { ConfigLoaderSupplierListPath = prev }()

			supplierEntry := ConfigPartial{"type": "my type"}
			suppliers := ConfigPartial{}
			baseSourcePartial := ConfigPartial{
				"type":   "file",
				"path":   "base_supplier_path",
				"format": "format",
			}
			_, _ = suppliers.Set(ConfigLoaderSupplierListPath, ConfigPartial{"supplier": supplierEntry})
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(suppliers, nil).Times(2)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(supplierEntry, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			gomock.InOrder(
				supplierCreator.EXPECT().Accept(&baseSourcePartial).Return(true),
				supplierCreator.EXPECT().Accept(&supplierEntry).Return(true),
			)
			gomock.InOrder(
				supplierCreator.EXPECT().Create(&baseSourcePartial).Return(supplier1, nil),
				supplierCreator.EXPECT().Create(&supplierEntry).Return(supplier2, nil),
			)
			supplierFactory := NewConfigSupplierFactory([]ConfigSupplierCreator{supplierCreator})
			config := NewConfig()

			sut, _ := NewConfigLoader(config, supplierFactory)

			if e := sut.Load(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})
}

func Test_ConfigServiceRegister(t *testing.T) {
	t.Run("NewConfigServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewConfigServiceRegister() == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewConfigServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			sut := NewConfigServiceRegister(nil)

			if e := sut.Provide(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register components", func(t *testing.T) {
			container := NewServiceContainer()
			sut := NewConfigServiceRegister(nil)

			e := sut.Provide(container)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !container.Has(ConfigYAMLDecoderCreatorContainerID):
				t.Errorf("no YAML parser creator : %v", sut)
			case !container.Has(ConfigJSONDecoderCreatorContainerID):
				t.Errorf("no JSON parser creator : %v", sut)
			case !container.Has(ConfigAllParserCreatorsContainerID):
				t.Errorf("no parser creators aggregator : %v", sut)
			case !container.Has(ConfigParserFactoryContainerID):
				t.Errorf("no parser factory : %v", sut)
			case !container.Has(ConfigAggregateSourceCreatorContainerID):
				t.Errorf("no aggregate source creator : %v", sut)
			case !container.Has(ConfigEnvSourceCreatorContainerID):
				t.Errorf("no env source creator : %v", sut)
			case !container.Has(ConfigFileSourceCreatorContainerID):
				t.Errorf("no file source creator : %v", sut)
			case !container.Has(ConfigObsFileSourceCreatorContainerID):
				t.Errorf("no observable file source creator : %v", sut)
			case !container.Has(ConfigDirSourceCreatorContainerID):
				t.Errorf("no dir source creator : %v", sut)
			case !container.Has(ConfigRestSourceCreatorContainerID):
				t.Errorf("no rest source creator : %v", sut)
			case !container.Has(ConfigObsRestSourceCreatorContainerID):
				t.Errorf("no observable rest source creator : %v", sut)
			case !container.Has(ConfigAllSupplierCreatorsContainerID):
				t.Errorf("no supplier creators aggregator : %v", sut)
			case !container.Has(ConfigSupplierFactoryContainerID):
				t.Errorf("no supplier factory : %v", sut)
			case !container.Has(ConfigContainerID):
				t.Errorf("no config : %v", sut)
			case !container.Has(ConfigLoaderContainerID):
				t.Errorf("no loader : %v", sut)
			}
		})

		t.Run("retrieving YAML decoder creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigYAMLDecoderCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigYAMLDecoderCreator:
				default:
					t.Error("didn't return a YAML decoder creator reference")
				}
			}
		})

		t.Run("retrieving JSON decoder creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigJSONDecoderCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigJSONDecoderCreator:
				default:
					t.Error("didn't return a JSON decoder creator reference")
				}
			}
		})

		t.Run("retrieving parser creators", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			parserCreator := NewMockConfigParserCreator(ctrl)
			_ = container.Add("parser.id", func() ConfigParserCreator {
				return parserCreator
			}, ConfigParserCreatorTag)

			creators, e := container.Get(ConfigAllParserCreatorsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creators == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := creators.(type) {
				case []ConfigParserCreator:
					found := false
					for _, creator := range c {
						if creator == parserCreator {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a parser creator slice populated with the expected creator instance")
					}
				default:
					t.Error("didn't return a parser creator slice")
				}
			}
		})

		t.Run("retrieving parser factory", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigParserFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigParserFactory:
				default:
					t.Error("didn't return a parser factory reference")
				}
			}
		})

		t.Run("retrieving aggregate source creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigAggregateSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigAggregateSourceCreator:
				default:
					t.Error("didn't return a aggregate source creator reference")
				}
			}
		})

		t.Run("retrieving env source creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigEnvSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigEnvSourceCreator:
				default:
					t.Error("didn't return a env source creator reference")
				}
			}
		})

		t.Run("retrieving file source creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigFileSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigFileSourceCreator:
				default:
					t.Error("didn't return a file source creator reference")
				}
			}
		})

		t.Run("retrieving observable file source creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigObsFileSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigObsFileSourceCreator:
				default:
					t.Error("didn't return a observable file source creator reference")
				}
			}
		})

		t.Run("retrieving dir source creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigDirSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigDirSourceCreator:
				default:
					t.Error("didn't return a dir source creator reference")
				}
			}
		})

		t.Run("retrieving rest source creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigRestSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigRestSourceCreator:
				default:
					t.Error("didn't return a rest source creator reference")
				}
			}
		})

		t.Run("retrieving observable rest source creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)

			factory, e := container.Get(ConfigObsRestSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *ConfigObsRestSourceCreator:
				default:
					t.Error("didn't return a observable rest source creator reference")
				}
			}
		})

		t.Run("retrieving aggregate suppliers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			supplier := NewMockConfigSupplier(ctrl)
			_ = container.Add("supplier.id", func() interface{} {
				return supplier
			}, ConfigAggregateSupplierTag)

			creator, e := container.Get(ConfigAggregateSourceCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creator == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := creator.(type) {
				case *ConfigAggregateSourceCreator:
					found := false
					for _, s := range c.suppliers {
						if s == supplier {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a aggregate supplier with the registered supplier")
					}
				default:
					t.Error("didn't return a aggregate supplier creator")
				}
			}
		})

		t.Run("retrieving supplier creators", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			_ = container.Add("supplier.id", func() ConfigSupplierCreator {
				return supplierCreator
			}, ConfigSupplierCreatorTag)

			creators, e := container.Get(ConfigAllSupplierCreatorsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creators == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := creators.(type) {
				case []ConfigSupplierCreator:
					found := false
					for _, creator := range c {
						if creator == supplierCreator {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a supplier creator slice populated with the expected creator instance")
					}
				default:
					t.Error("didn't return a supplier creator slice")
				}
			}
		})

		t.Run("retrieving the supplier factory", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			creator, e := container.Get(ConfigSupplierFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creator == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch creator.(type) {
				case *ConfigSupplierFactory:
				default:
					t.Error("didn't return a supplier factory reference")
				}
			}
		})

		t.Run("retrieving config", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			config, e := container.Get(ConfigContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case config == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch config.(type) {
				case *Config:
				default:
					t.Error("didn't return a config reference")
				}
			}
		})

		t.Run("error retrieving config on retrieving loader", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = container.Add(ConfigContainerID, func() *Config { return nil })

			if _, e := container.Get(ConfigLoaderContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving supplier factory on retrieving loader", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = container.Add(ConfigSupplierFactoryContainerID, func() *ConfigSupplierFactory { return nil })

			if _, e := container.Get(ConfigLoaderContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("retrieving loader", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)

			l, e := container.Get(ConfigLoaderContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case l == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch l.(type) {
				case *ConfigLoader:
				default:
					t.Error("didn't return a loader reference")
				}
			}
		})
	})

	t.Run("Boot", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewConfigServiceRegister(nil).Boot(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("valid simple boot with creators (no loader)", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ConfigLoaderActive = false
			defer func() { ConfigLoaderActive = true }()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			sut := NewConfigServiceRegister(nil)
			_ = sut.Provide(container)

			_ = container.Add("parser.id", func() ConfigParserCreator {
				return NewMockConfigParserCreator(ctrl)
			}, ConfigParserCreatorTag)
			_ = container.Add("supplier.id", func() ConfigSupplierCreator {
				return NewMockConfigSupplierCreator(ctrl)
			}, ConfigSupplierCreatorTag)

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected error (%v)", e)
			}
		})

		t.Run("no entry supplier active", func(t *testing.T) {
			ConfigLoaderActive = false
			defer func() { ConfigLoaderActive = true }()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			sut := NewConfigServiceRegister(nil)
			_ = sut.Provide(container)
			_ = container.Add(ConfigLoaderContainerID, func() (*ConfigLoader, error) {
				return nil, fmt.Errorf("error message")
			})

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected error (%v)", e)
			}
		})

		t.Run("error retrieving loader", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			sut := NewConfigServiceRegister(nil)
			_ = sut.Provide(container)
			_ = container.Add(ConfigLoaderContainerID, func() (*ConfigLoader, error) {
				return nil, fmt.Errorf("error message")
			})

			if e := sut.Boot(container); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("invalid loader", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			sut := NewConfigServiceRegister(nil)
			_ = sut.Provide(container)
			_ = container.Add(ConfigLoaderContainerID, func() string {
				return "message"
			})

			if e := sut.Boot(container); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("request loader to init config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":   "file",
				"path":   ConfigLoaderFileSupplierPath,
				"format": ConfigLoaderSupplierFormat,
			}
			container := NewServiceContainer()
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(ConfigPartial{}, nil).Times(1)
			supplierCreator := NewMockConfigSupplierCreator(ctrl)
			supplierCreator.EXPECT().Accept(&partial).Return(true).Times(1)
			supplierCreator.EXPECT().Create(&partial).Return(supplier, nil).Times(1)
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			sut := NewConfigServiceRegister(nil)
			_ = sut.Provide(container)
			_ = container.Remove(ConfigFileSourceCreatorContainerID)
			_ = container.Add("supplier", func() ConfigSupplierCreator {
				return supplierCreator
			}, ConfigSupplierCreatorTag)

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected error (%v)", e)
			}
		})
	})
}
