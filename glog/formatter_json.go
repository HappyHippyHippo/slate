package glog

import (
	"encoding/json"
	"strings"
	"time"
)

// FormatterJSON defines a JSON based logger formatter.
type FormatterJSON struct{}

var _ Formatter = &FormatterJSON{}

// Format will create the output JSON string message formatted with the content
// of the passed level, message and context
func (f FormatterJSON) Format(level Level, message string, ctx map[string]interface{}) string {
	if ctx == nil {
		ctx = map[string]interface{}{}
	}

	ctx["time"] = time.Now().Format("2006-01-02T15:04:05.000-0700")
	ctx["level"] = strings.ToUpper(LevelMapName[level])
	ctx["message"] = message

	bytes, _ := json.Marshal(ctx)
	return string(bytes)
}
