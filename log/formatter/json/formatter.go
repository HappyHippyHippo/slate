package json

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/happyhippyhippo/slate/log"
)

// Formatter defines an instance used to format a log message into
// a JSON string.
type Formatter struct{}

var _ log.IFormatter = &Formatter{}

// Format will create the output JSON string message formatted with the
// content of the passed level, message and context
func (f Formatter) Format(
	level log.Level,
	message string,
	ctx ...log.Context,
) string {
	// guarantee that the content context is a valid map reference,
	// so it can be used to compose the final formatted message
	// for that initialize an empty context map, and merge all the
	// given extra contexts
	data := log.Context{}
	for _, c := range ctx {
		for k, v := range c {
			data[k] = v
		}
	}
	// store the extra time, level and message in the request context
	data["time"] = time.Now().Format("2006-01-02T15:04:05.000-0700")
	data["level"] = strings.ToUpper(log.LevelMapName[level])
	data["message"] = message
	// compose the response JSON formatted string with the populated
	// context instance
	bytes, _ := json.Marshal(data)
	return string(bytes)
}
