package log

import (
	"encoding/xml"
	"strings"

	"github.com/gin-gonic/gin"
)

// NewResponseReaderDecoratorXML will instantiate a new response
// event context reader XML decorator used to parse the response body as an XML
// and add the parsed content into the logging data.
func NewResponseReaderDecoratorXML(reader ResponseReader, model interface{}) (ResponseReader, error) {
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
			if strings.Contains(accept, gin.MIMEXML) || strings.Contains(accept, gin.MIMEXML2) {
				if err = xml.Unmarshal([]byte(body.(string)), &model); err == nil {
					data["bodyXml"] = model
				}
			}
		}

		return data, nil
	}, nil
}
