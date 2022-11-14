package log

import "github.com/gin-gonic/gin"

// NewResponseReaderDefault generates a new default response reader with
// the defined expected status code.
func NewResponseReaderDefault(statusCode int) ResponseReader {
	return func(_ *gin.Context, writer responseWriter) (map[string]interface{}, error) {
		if writer == nil {
			return nil, errNilPointer("writer")
		}

		status := writer.Status()
		data := map[string]interface{}{
			"status":  status,
			"headers": responseHeaders(writer),
		}

		if status != statusCode {
			data["body"] = string(writer.Body())
		}

		return data, nil
	}
}

func responseHeaders(response responseWriter) map[string]interface{} {
	headers := map[string]interface{}{}
	for index, header := range response.Header() {
		if len(header) == 1 {
			headers[index] = header[0]
		} else {
			headers[index] = header
		}
	}
	return headers
}
