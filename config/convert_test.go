package config

import (
	"reflect"
	"testing"
)

func Test_Convert(t *testing.T) {
	t.Run("Convert float32 into int", func(t *testing.T) {
		data := float32(123)
		expected := 123

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert float64 into int", func(t *testing.T) {
		data := float64(123)
		expected := 123

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map", func(t *testing.T) {
		data := map[string]interface{}{"node": "value"}
		expected := Partial{"node": "value"}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial", func(t *testing.T) {
		data := Partial{"node": "value"}
		expected := Partial{"node": "value"}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert list", func(t *testing.T) {
		data := []interface{}{1, 2, 3}
		expected := []interface{}{1, 2, 3}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a float32 into a map with a int", func(t *testing.T) {
		data := map[string]interface{}{"node": float32(123)}
		expected := Partial{"node": 123}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a float64 into a map with a int", func(t *testing.T) {
		data := map[string]interface{}{"node": float64(123)}
		expected := Partial{"node": 123}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with another map", func(t *testing.T) {
		data := map[string]interface{}{"node": map[string]interface{}{"node2": "value"}}
		expected := Partial{"node": Partial{"node2": "value"}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a list", func(t *testing.T) {
		data := map[string]interface{}{"node": []interface{}{1, 2, 3}}
		expected := Partial{"node": []interface{}{1, 2, 3}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial with a map", func(t *testing.T) {
		data := Partial{"node": map[string]interface{}{"node2": "value"}}
		expected := Partial{"node": Partial{"node2": "value"}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a list of maps", func(t *testing.T) {
		data := map[string]interface{}{"node": []interface{}{map[string]interface{}{"node2": "value"}}}
		expected := Partial{"node": []interface{}{Partial{"node2": "value"}}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with a list of configs", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{Partial{"node2": "value"}}}
		expected := Partial{"node": []interface{}{Partial{"node2": "value"}}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial with numeric fields", func(t *testing.T) {
		data := Partial{1: map[string]interface{}{"node2": "value"}}
		expected := Partial{1: Partial{"node2": "value"}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial with a uppercase fields", func(t *testing.T) {
		data := Partial{"NoDE": map[string]interface{}{"nODE2": "value"}}
		expected := Partial{"node": Partial{"node2": "value"}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map uppercase fields", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{map[string]interface{}{"NOde2": "value"}}}
		expected := Partial{"node": []interface{}{Partial{"node2": "value"}}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert partial uppercase fields", func(t *testing.T) {
		data := map[string]interface{}{"NoDe": []interface{}{Partial{"NOde2": "value"}}}
		expected := Partial{"node": []interface{}{Partial{"node2": "value"}}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with numeric keys", func(t *testing.T) {
		data := map[interface{}]interface{}{1: []interface{}{Partial{2: "value"}}}
		expected := Partial{1: []interface{}{Partial{2: "value"}}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})

	t.Run("Convert map with string keys", func(t *testing.T) {
		data := map[interface{}]interface{}{"NoDE1": []interface{}{Partial{2: "value"}}}
		expected := Partial{"node1": []interface{}{Partial{2: "value"}}}

		if check := Convert(data); !reflect.DeepEqual(check, expected) {
			t.Errorf("resulted in (%v) when converting (%v), expecting (%v)", check, data, expected)
		}
	})
}
