package log

import (
	"encoding/xml"
	"strings"

	"github.com/gin-gonic/gin"
)

// NewRequestReaderDecoratorXML will instantiate a new request
// event context reader XML decorator used to parse the request body as an XML
// and add the parsed content into the logging data.
func NewRequestReaderDecoratorXML(reader RequestReader, model interface{}) (RequestReader, error) {
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
		if strings.HasPrefix(contentType, gin.MIMEXML) || strings.HasPrefix(contentType, gin.MIMEXML2) {
			if err = xml.Unmarshal([]byte(data["body"].(string)), &model); err == nil {
				data["bodyXml"] = model
			}
		}

		return data, nil
	}, nil
}
