package log

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewRequestReaderDecoratorJSON(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		if _, err := NewRequestReaderDecoratorJSON(nil, nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := func(_ *gin.Context) (map[string]interface{}, error) {
			return nil, nil
		}
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, err := decorator(nil)
		switch {
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
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
		reader := func(_ *gin.Context) (map[string]interface{}, error) {
			return nil, expected
		}
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, err := decorator(ctx)
		switch {
		case err == nil:
			t.Error("didn't returned the expected error")
		case !reflect.DeepEqual(err, expected):
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		case result != nil:
			t.Errorf("returned the unexpeted context data : %v", result)
		}
	})

	t.Run("empty content-type does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := map[string]interface{}{"body": `{"field":"value"}`}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		reader := func(_ *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
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

		data := map[string]interface{}{"body": `{"field":"value"}`}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML)
		reader := func(_ *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
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

		data := map[string]interface{}{"body": "field"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEJSON)
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML)
		reader := func(_ *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
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

		data := map[string]interface{}{"body": `{"field":"value"}`}
		expected := map[string]interface{}{"field": "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEJSON)
		reader := func(_ *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorJSON(reader, nil)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
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
