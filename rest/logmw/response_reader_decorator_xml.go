package logmw

import (
	"encoding/xml"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/log"
)

// NewResponseReaderDecoratorXML will instantiate a new response
// event context reader XML decorator used to parse the response body as an XML
// and add the parsed content into the logging data.
func NewResponseReaderDecoratorXML(
	reader ResponseReader,
	model interface{},
) (ResponseReader, error) {
	// check the reader argument reference
	if reader == nil {
		return nil, errNilPointer("reader")
	}
	// return the decorated response reader method
	return func(
		ctx *gin.Context,
		writer responseWriter,
		statusCode int,
	) (log.Context, error) {
		// check the context argument reference
		if ctx == nil {
			return nil, errNilPointer("ctx")
		}
		// check the writer argument reference
		if writer == nil {
			return nil, errNilPointer("writer")
		}
		// read the logging response data from the context
		data, err := reader(ctx, writer, statusCode)
		if err != nil {
			return nil, err
		}
		// check if there is content in the response body logging data
		// and try to unmarshall it if the response is in XML to be logged
		// in the bodyXml field
		if body, ok := data["body"]; ok == true {
			accept := strings.ToLower(ctx.Request.Header.Get("Accept"))
			if strings.Contains(accept, gin.MIMEXML) || strings.Contains(accept, gin.MIMEXML2) {
				if err = xml.Unmarshal([]byte(body.(string)), &model); err == nil {
					data["bodyXml"] = model
				}
			}
		}
		// return the response information
		return data, nil
	}, nil
}
