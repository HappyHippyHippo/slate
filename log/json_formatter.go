package log

import (
	"encoding/json"
	"strings"
	"time"
)

// JSONFormatter defines an instance used to format a log message into
// a JSON string.
type JSONFormatter struct{}

var _ IFormatter = &JSONFormatter{}

// Format will create the output JSON string message formatted with the
// content of the passed level, message and context
func (f JSONFormatter) Format(
	level Level,
	message string,
	ctx map[string]interface{},
) string {
	// guarantee that the content context is a valid map reference,
	// so it can be used to compose the final formatted message
	if ctx == nil {
		ctx = map[string]interface{}{}
	}
	// store the extra time, level and message in the request context
	ctx["time"] = time.Now().Format("2006-01-02T15:04:05.000-0700")
	ctx["level"] = strings.ToUpper(LevelMapName[level])
	ctx["message"] = message
	// compose the response JSON formatted string with the populated
	// context instance
	bytes, _ := json.Marshal(ctx)
	return string(bytes)
}
