package log

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

// NewRequestReaderDecoratorJSON will instantiate a new request
// event context reader JSON decorator used to parse the request body as a JSON
// and add the parsed content into the logging data.
func NewRequestReaderDecoratorJSON(reader RequestReader, model interface{}) (RequestReader, error) {
	if reader == nil {
		return nil, errNilPointer("reader")
	}

	return func(ctx *gin.Context) (map[string]interface{}, error) {
		if ctx == nil {
			return nil, errNilPointer("ctx")
		}

		data, err := reader(ctx)
		if err != nil {
			return nil, err
		}

		contentType := strings.ToLower(ctx.Request.Header.Get("Content-Type"))
		if strings.HasPrefix(contentType, gin.MIMEJSON) {
			if err = json.Unmarshal([]byte(data["body"].(string)), &model); err == nil {
				data["bodyJson"] = model
			}
		}

		return data, nil
	}, nil
}
