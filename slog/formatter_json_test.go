package slog

import (
	"regexp"
	"testing"
)

func Test_FormatterJSON_Format(t *testing.T) {
	t.Run("correctly format the message", func(t *testing.T) {
		scenarios := []struct {
			level    Level
			fields   map[string]interface{}
			message  string
			expected string
		}{
			{ // _test level FATAL
				level:    FATAL,
				fields:   nil,
				message:  "",
				expected: `"level"\s*\:\s*"FATAL"`,
			},
			{ // _test level ERROR
				level:    ERROR,
				fields:   nil,
				message:  "",
				expected: `"level"\s*\:\s*"ERROR"`,
			},
			{ // _test level WARNING
				level:    WARNING,
				fields:   nil,
				message:  "",
				expected: `"level"\s*\:\s*"WARNING"`,
			},
			{ // _test level NOTICE
				level:    NOTICE,
				fields:   nil,
				message:  "",
				expected: `"level"\s*\:\s*"NOTICE"`,
			},
			{ // _test level INFO
				level:    INFO,
				fields:   nil,
				message:  "",
				expected: `"level"\s*\:\s*"INFO"`,
			},
			{ // _test level DEBUG
				level:    DEBUG,
				fields:   nil,
				message:  "",
				expected: `"level"\s*\:\s*"DEBUG"`,
			},
			{ // _test fields (single value)
				level:    DEBUG,
				fields:   map[string]interface{}{"field1": "value1"},
				message:  "",
				expected: `"field1"\s*\:\s*"value1"`,
			},
			{ // _test fields (multiple value)
				level:    DEBUG,
				fields:   map[string]interface{}{"field1": "value1", "field2": "value2"},
				message:  "",
				expected: `"field1"\s*\:\s*"value1"|"field2"\s*\:\s*"value2"`,
			},
			{ // _test message
				level:    DEBUG,
				fields:   nil,
				message:  "My_message",
				expected: `"message"\s*\:\s*"My_message"`,
			},
		}

		for _, scenario := range scenarios {
			check := formatterJSON{}.Format(scenario.level, scenario.message, scenario.fields)
			match, _ := regexp.Match(scenario.expected, []byte(check))
			if !match {
				t.Errorf("didn't validated (%s) output", check)
			}
		}
	})
}
