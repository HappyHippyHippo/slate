package logmw

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/log"
)

// RequestReaderDefault is the default function used to parse the request
// context information.
func RequestReaderDefault(
	ctx *gin.Context,
) (log.Context, error) {
	// check the context argument reference
	if ctx == nil {
		return nil, errNilPointer("ctx")
	}
	// obtain the request parameters
	params := log.Context{}
	for p, v := range ctx.Request.URL.Query() {
		if len(v) == 1 {
			params[p] = v[0]
		} else {
			params[p] = v
		}
	}
	// return the default request information
	return log.Context{
		"headers": requestHeaders(ctx.Request),
		"method":  ctx.Request.Method,
		"path":    ctx.Request.URL.Path,
		"params":  params,
		"body":    requestBody(ctx.Request),
	}, nil
}

func requestHeaders(request *http.Request) log.Context {
	// try to flat single entry header fields
	headers := log.Context{}
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
	// obtain the request body (content destructible action)
	var bodyBytes []byte
	if request.Body != nil {
		bodyBytes, _ = io.ReadAll(request.Body)
	}
	// reassign the request body with a memory buffer
	request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}
