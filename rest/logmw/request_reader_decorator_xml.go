package logmw

import (
	"encoding/xml"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/log"
)

// NewRequestReaderDecoratorXML will instantiate a new request
// event context reader XML decorator used to parse the request body as an XML
// and add the parsed content into the logging data.
func NewRequestReaderDecoratorXML(
	reader RequestReader,
	model interface{},
) (RequestReader, error) {
	// check the reader argument reference
	if reader == nil {
		return nil, errNilPointer("reader")
	}
	// return the decorated request reader method
	return func(
		ctx *gin.Context,
	) (log.Context, error) {
		// check the context argument reference
		if ctx == nil {
			return nil, errNilPointer("ctx")
		}
		// read the logging request data from the context
		data, err := reader(ctx)
		if err != nil {
			return nil, err
		}
		// try to unmarshall the request body content if the request
		// is in XML format, and store it in the data map on the
		// bodyXml field
		contentType := strings.ToLower(ctx.Request.Header.Get("Content-Type"))
		if strings.HasPrefix(contentType, gin.MIMEXML) || strings.HasPrefix(contentType, gin.MIMEXML2) {
			if err = xml.Unmarshal([]byte(data["body"].(string)), &model); err == nil {
				data["bodyXml"] = model
			}
		}
		// return the request information
		return data, nil
	}, nil
}
