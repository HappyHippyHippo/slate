package logmw

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewRequestReaderDecoratorJSON(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		if _, e := NewRequestReaderDecoratorJSON(nil, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := func(_ *gin.Context) (log.Context, error) { return nil, nil }
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, e := decorator(nil)
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
		reader := func(_ *gin.Context) (log.Context, error) { return nil, expected }
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !reflect.DeepEqual(e, expected):
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		case result != nil:
			t.Errorf("returned the unexpeted context data : %v", result)
		}
	})

	t.Run("empty content-type does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": `{"field":"value"}`}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		reader := func(_ *gin.Context) (log.Context, error) { return data, nil }
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx)
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

	t.Run("non-json content-type does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": `{"field":"value"}`}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML)
		reader := func(_ *gin.Context) (log.Context, error) { return data, nil }
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx)
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

		data := log.Context{"body": "field"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEJSON)
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML)
		reader := func(_ *gin.Context) (log.Context, error) { return data, nil }
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx)
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

	t.Run("correctly add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := log.Context{"body": `{"field":"value"}`}
		expected := map[string]interface{}{"field": "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEJSON)
		reader := func(_ *gin.Context) (log.Context, error) { return data, nil }
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, e := decorator(ctx)
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
