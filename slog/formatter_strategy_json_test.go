package slog

import (
	"testing"
)

func Test_FormatterStrategyJSON_Accept(t *testing.T) {
	t.Run("accept only json format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test json format
				format:   FormatJSON,
				expected: true,
			},
			{ // _test non-json format
				format:   FormatUnknown,
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			if check := (&formatterStrategyJSON{}).Accept(scenario.format); check != scenario.expected {
				t.Errorf("returned (%v) for the (%s) format", check, scenario.format)
			}
		}
	})
}

func Test_FormatterStrategyJSON_Create(t *testing.T) {
	t.Run("create json formatter", func(t *testing.T) {
		sut, e := (&formatterStrategyJSON{}).Create()
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch sut.(type) {
			case *formatterJSON:
			default:
				t.Errorf("didn't returned a new json formatter")
			}
		}
	})
}
