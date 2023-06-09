package json

import (
	"testing"

	"github.com/happyhippyhippo/slate/log"
)

func Test_NewFormatterStrategy(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		if NewFormatterStrategy() == nil {
			t.Error("didn't returned the expected reference")
		}
	})
}

func Test_FormatterStrategy_Accept(t *testing.T) {
	t.Run("accept only json format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test json format
				format:   Format,
				expected: true,
			},
			{ // _test non-json format
				format:   log.UnknownFormatter,
				expected: false,
			},
		}

		for _, s := range scenarios {
			if check := (&FormatterStrategy{}).Accept(s.format); check != s.expected {
				t.Errorf("(%v) for the (%s) format", check, s.format)
			}
		}
	})
}

func Test_FormatterStrategy_Create(t *testing.T) {
	t.Run("create json formatter", func(t *testing.T) {
		sut, e := (&FormatterStrategy{}).Create()
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch sut.(type) {
			case *Formatter:
			default:
				t.Errorf("didn't returned a new json formatter")
			}
		}
	})
}
