package log

import (
	"testing"

	slog "github.com/happyhippyhippo/slate/log"
)

func Test_envToLogLevel(t *testing.T) {
	t.Run("existing parsing", func(t *testing.T) {
		scenarios := []struct {
			input    string
			def      slog.Level
			expected slog.Level
		}{
			{ // fatal
				input:    "fatal",
				def:      slog.DEBUG,
				expected: slog.FATAL,
			},
			{ // FATAL
				input:    "FATAL",
				def:      slog.DEBUG,
				expected: slog.FATAL,
			},
			{ // error
				input:    "error",
				def:      slog.DEBUG,
				expected: slog.ERROR,
			},
			{ // ERROR
				input:    "ERROR",
				def:      slog.DEBUG,
				expected: slog.ERROR,
			},
			{ // warning
				input:    "warning",
				def:      slog.DEBUG,
				expected: slog.WARNING,
			},
			{ // WARNING
				input:    "WARNING",
				def:      slog.DEBUG,
				expected: slog.WARNING,
			},
			{ // notice
				input:    "notice",
				def:      slog.DEBUG,
				expected: slog.NOTICE,
			},
			{ // NOTICE
				input:    "NOTICE",
				def:      slog.DEBUG,
				expected: slog.NOTICE,
			},
			{ // info
				input:    "info",
				def:      slog.DEBUG,
				expected: slog.INFO,
			},
			{ // INFO
				input:    "INFO",
				def:      slog.DEBUG,
				expected: slog.INFO,
			},
			{ // debug
				input:    "debug",
				def:      slog.DEBUG,
				expected: slog.DEBUG,
			},
			{ // DEBUG
				input:    "DEBUG",
				def:      slog.DEBUG,
				expected: slog.DEBUG,
			},
			{ // unknown -> return default
				input:    "unknown",
				def:      slog.INFO,
				expected: slog.INFO,
			},
		}

		for _, scenario := range scenarios {
			if chk := envToLogLevel(scenario.input, scenario.def); chk != scenario.expected {
				t.Errorf("parsed to  (%v) when expecting (%v)", chk, scenario.expected)
			}
		}
	})
}
