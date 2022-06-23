package srlog

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

// NewResponseReaderDecoratorJSON will instantiate a new response
// event context reader JSON decorator used to parse the response   body as
// a JSON and add the parsed content into the logging data.
func NewResponseReaderDecoratorJSON(reader ResponseReader, model interface{}) (ResponseReader, error) {
	if reader == nil {
		return nil, errNilPointer("reader")
	}

	return func(ctx *gin.Context, writer responseWriter) (map[string]interface{}, error) {
		if ctx == nil {
			return nil, errNilPointer("ctx")
		}
		if writer == nil {
			return nil, errNilPointer("writer")
		}

		data, err := reader(ctx, writer)
		if err != nil {
			return nil, err
		}

		if body, ok := data["body"]; ok == true {
			accept := strings.ToLower(ctx.Request.Header.Get("Accept"))
			if accept == "*/*" || strings.Contains(accept, gin.MIMEJSON) {
				if err = json.Unmarshal([]byte(body.(string)), &model); err == nil {
					data["bodyJson"] = model
				}
			}
		}

		return data, nil
	}, nil
}
