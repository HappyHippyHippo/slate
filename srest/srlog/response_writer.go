package srlog

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type responseWriter interface {
	gin.ResponseWriter
	Body() []byte
}

type writer struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var _ gin.ResponseWriter = &writer{}

func newResponseWriter(w gin.ResponseWriter) (responseWriter, error) {
	if w == nil {
		return nil, errNilPointer("writer")
	}

	return &writer{
		ResponseWriter: w,
		body:           &bytes.Buffer{},
	}, nil
}

// Write executes the writing the desired bytes into the underlying writer
// and storing them in the internal buffer.
func (w writer) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Body will retrieve the stored bytes given on the previous calls
// to the Write method.
func (w writer) Body() []byte {
	return w.body.Bytes()
}
