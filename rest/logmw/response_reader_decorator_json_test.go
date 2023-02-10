package logmw

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
	"github.com/pkg/errors"
)

func Test_NewResponseReaderDecoratorJSON(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		if _, e := NewResponseReaderDecoratorJSON(nil, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ginWriter := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return nil, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(nil, ginWriter, 0)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		case result != nil:
			t.Errorf("returned the unexpeted context data : %v", result)
		}
	})

	t.Run("nil writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ctx := &gin.Context{}
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return nil, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, nil, 0)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		case result != nil:
			t.Errorf("returned the unexpeted context data : %v", result)
		}
	})

	t.Run("base reader error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ginWriter := NewMockResponseWriter(ctrl)
		writer, _ := newResponseWriter(ginWriter)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return nil, expected }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !reflect.DeepEqual(e, expected):
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		case result != nil:
			t.Errorf("returned the unexpeted context data : %v", result)
		}
	})

	t.Run("missing body does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ginWriter := NewMockResponseWriter(ctrl)
		writer, _ := newResponseWriter(ginWriter)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return data, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyJson"]; ok {
				t.Error("added the bodyJson field")
			}
		}
	})

	t.Run("empty accept does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": `{"field":"value"}`}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ginWriter := NewMockResponseWriter(ctrl)
		writer, _ := newResponseWriter(ginWriter)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return data, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyJson"]; ok {
				t.Error("added the bodyJson field")
			}
		}
	})

	t.Run("non-json accept does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": `{"field":"value"}`}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", gin.MIMEXML)
		ginWriter := NewMockResponseWriter(ctrl)
		writer, _ := newResponseWriter(ginWriter)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return data, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyJson"]; ok {
				t.Error("added the bodyJson field")
			}
		}
	})

	t.Run("invalid json content does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": "{field value}"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", gin.MIMEJSON)
		ginWriter := NewMockResponseWriter(ctrl)
		writer, _ := newResponseWriter(ginWriter)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return data, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyJson"]; ok {
				t.Error("added the bodyJson field")
			}
		}
	})

	t.Run("correctly add decorated field for application/json", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": `{"field":"value"}`}
		expected := map[string]interface{}{"field": "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", gin.MIMEJSON)
		ginWriter := NewMockResponseWriter(ctrl)
		writer, _ := newResponseWriter(ginWriter)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return data, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if body, ok := result["bodyJson"]; !ok {
				t.Error("didn't added the bodyJson field")
			} else if !reflect.DeepEqual(body, expected) {
				t.Errorf("added the (%v) content when expecting : %v", body, expected)
			}
		}
	})

	t.Run("correctly add decorated field for 'any mime type'", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": `{"field":"value"}`}
		expected := map[string]interface{}{"field": "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", "*/*")
		ginWriter := NewMockResponseWriter(ctrl)
		writer, _ := newResponseWriter(ginWriter)
		reader := func(_ *gin.Context, _ responseWriter, _ int) (log.Context, error) { return data, nil }
		decorator, _ := NewResponseReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if body, ok := result["bodyJson"]; !ok {
				t.Error("didn't added the bodyJson field")
			} else if !reflect.DeepEqual(body, expected) {
				t.Errorf("added the (%v) content when expecting : %v", body, expected)
			}
		}
	})
}
