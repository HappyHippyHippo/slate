package log

import (
	"regexp"
	"testing"
)

func Test_JSONFormatter_Format(t *testing.T) {
	t.Run("correctly format the message", func(t *testing.T) {
		scenarios := []struct {
			level    Level
			ctx      Context
			message  string
			expected string
		}{
			{ // _test level FATAL
				level:    FATAL,
				ctx:      nil,
				message:  "",
				expected: `"level"\s*\:\s*"FATAL"`,
			},
			{ // _test level ERROR
				level:    ERROR,
				ctx:      nil,
				message:  "",
				expected: `"level"\s*\:\s*"ERROR"`,
			},
			{ // _test level WARNING
				level:    WARNING,
				ctx:      nil,
				message:  "",
				expected: `"level"\s*\:\s*"WARNING"`,
			},
			{ // _test level NOTICE
				level:    NOTICE,
				ctx:      nil,
				message:  "",
				expected: `"level"\s*\:\s*"NOTICE"`,
			},
			{ // _test level INFO
				level:    INFO,
				ctx:      nil,
				message:  "",
				expected: `"level"\s*\:\s*"INFO"`,
			},
			{ // _test level DEBUG
				level:    DEBUG,
				ctx:      nil,
				message:  "",
				expected: `"level"\s*\:\s*"DEBUG"`,
			},
			{ // _test ctx (single value)
				level:    DEBUG,
				ctx:      Context{"field1": "value1"},
				message:  "",
				expected: `"field1"\s*\:\s*"value1"`,
			},
			{ // _test ctx (multiple value)
				level:    DEBUG,
				ctx:      Context{"field1": "value1", "field2": "value2"},
				message:  "",
				expected: `"field1"\s*\:\s*"value1"|"field2"\s*\:\s*"value2"`,
			},
			{ // _test message
				level:    DEBUG,
				ctx:      nil,
				message:  "My_message",
				expected: `"message"\s*\:\s*"My_message"`,
			},
		}

		for _, scenario := range scenarios {
			check := JSONFormatter{}.Format(scenario.level, scenario.message, scenario.ctx)
			match, _ := regexp.Match(scenario.expected, []byte(check))
			if !match {
				t.Errorf("didn't validated (%s) output", check)
			}
		}
	})
}
