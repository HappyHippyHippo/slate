package logmw

import (
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/log"
)

// ResponseReaderDefault @todo doc.
func ResponseReaderDefault(
	_ *gin.Context,
	writer responseWriter,
	statusCode int,
) (log.Context, error) {
	// check the writer argument reference
	if writer == nil {
		return nil, errNilPointer("writer")
	}
	// obtain the response status code
	status := writer.Status()
	// store the default logging information
	data := log.Context{
		"status":  status,
		"headers": responseHeaders(writer),
	}
	// add the response body to the logging information if the
	// response status code differs from the expected
	if status != statusCode {
		data["body"] = string(writer.Body())
	}
	// return the response logging information
	return data, nil
}

func responseHeaders(
	response responseWriter,
) log.Context {
	// try to flat single entry header fields
	headers := log.Context{}
	for index, header := range response.Header() {
		if len(header) == 1 {
			headers[index] = header[0]
		} else {
			headers[index] = header
		}
	}
	return headers
}
