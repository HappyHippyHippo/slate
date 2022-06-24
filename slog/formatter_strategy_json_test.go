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
		formatter, err := (&formatterStrategyJSON{}).Create()
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case formatter == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch formatter.(type) {
			case *FormatterJSON:
			default:
				t.Errorf("didn't returned a new json formatter")
			}
		}
	})
}
