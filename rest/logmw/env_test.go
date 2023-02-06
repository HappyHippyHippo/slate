package logmw

import (
	"testing"

	"github.com/happyhippyhippo/slate/log"
)

func Test_envToLogLevel(t *testing.T) {
	t.Run("existing parsing", func(t *testing.T) {
		scenarios := []struct {
			input    string
			def      log.Level
			expected log.Level
		}{
			{ // fatal
				input:    "fatal",
				def:      log.DEBUG,
				expected: log.FATAL,
			},
			{ // FATAL
				input:    "FATAL",
				def:      log.DEBUG,
				expected: log.FATAL,
			},
			{ // error
				input:    "error",
				def:      log.DEBUG,
				expected: log.ERROR,
			},
			{ // ERROR
				input:    "ERROR",
				def:      log.DEBUG,
				expected: log.ERROR,
			},
			{ // warning
				input:    "warning",
				def:      log.DEBUG,
				expected: log.WARNING,
			},
			{ // WARNING
				input:    "WARNING",
				def:      log.DEBUG,
				expected: log.WARNING,
			},
			{ // notice
				input:    "notice",
				def:      log.DEBUG,
				expected: log.NOTICE,
			},
			{ // NOTICE
				input:    "NOTICE",
				def:      log.DEBUG,
				expected: log.NOTICE,
			},
			{ // info
				input:    "info",
				def:      log.DEBUG,
				expected: log.INFO,
			},
			{ // INFO
				input:    "INFO",
				def:      log.DEBUG,
				expected: log.INFO,
			},
			{ // debug
				input:    "debug",
				def:      log.DEBUG,
				expected: log.DEBUG,
			},
			{ // DEBUG
				input:    "DEBUG",
				def:      log.DEBUG,
				expected: log.DEBUG,
			},
			{ // unknown -> return default
				input:    "unknown",
				def:      log.INFO,
				expected: log.INFO,
			},
		}

		for _, scenario := range scenarios {
			if chk := envToLogLevel(scenario.input, scenario.def); chk != scenario.expected {
				t.Errorf("parsed to  (%v) when expecting (%v)", chk, scenario.expected)
			}
		}
	})
}
