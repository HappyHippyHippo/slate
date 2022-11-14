package log

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestReaderDefault is the default function used to parse the request
// context information.
func RequestReaderDefault(ctx *gin.Context) (map[string]interface{}, error) {
	if ctx == nil {
		return nil, errNilPointer("ctx")
	}

	params := map[string]interface{}{}
	for p, v := range ctx.Request.URL.Query() {
		if len(v) == 1 {
			params[p] = v[0]
		} else {
			params[p] = v
		}
	}

	return map[string]interface{}{
		"headers": requestHeaders(ctx.Request),
		"method":  ctx.Request.Method,
		"path":    ctx.Request.URL.Path,
		"params":  params,
		"body":    requestBody(ctx.Request),
	}, nil
}

func requestHeaders(request *http.Request) map[string]interface{} {
	headers := map[string]interface{}{}
	for index, header := range request.Header {
		if len(header) == 1 {
			headers[index] = header[0]
		} else {
			headers[index] = header
		}
	}
	return headers
}

func requestBody(request *http.Request) string {
	var bodyBytes []byte
	if request.Body != nil {
		bodyBytes, _ = io.ReadAll(request.Body)
	}
	request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}
