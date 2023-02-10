package log

import (
	"testing"
)

func Test_NewJSONFormatterStrategy(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		if NewJSONFormatterStrategy() == nil {
			t.Error("didn't returned the expected reference")
		}
	})
}

func Test_JSONFormatterStrategy_Accept(t *testing.T) {
	t.Run("accept only json format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test json format
				format:   JSONFormatterFormat,
				expected: true,
			},
			{ // _test non-json format
				format:   UnknownFormatterFormat,
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			if check := (&JSONFormatterStrategy{}).Accept(scenario.format); check != scenario.expected {
				t.Errorf("returned (%v) for the (%s) format", check, scenario.format)
			}
		}
	})
}

func Test_JSONFormatterStrategy_Create(t *testing.T) {
	t.Run("create json formatter", func(t *testing.T) {
		sut, e := (&JSONFormatterStrategy{}).Create()
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch sut.(type) {
			case *JSONFormatter:
			default:
				t.Errorf("didn't returned a new json formatter")
			}
		}
	})
}
