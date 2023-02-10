package logmw

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_RequestReaderDefault(t *testing.T) {
	t.Run("nil ctx", func(t *testing.T) {
		if _, e := RequestReaderDefault(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("valid request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		method := "method"
		path := "/resource"
		uri := path + "?param1=value1&param2=value2&param2=value3"
		reqURL, _ := url.Parse("http://domain" + uri)
		headers := map[string][]string{"header1": {"value1", "value2"}, "header2": {"value3"}}
		expHeaders := log.Context{"header1": []string{"value1", "value2"}, "header2": "value3"}
		expParams := log.Context{
			"param1": "value1",
			"param2": []string{"value2", "value3"},
		}
		jsonBody := map[string]interface{}{"field": "value"}
		rawBody, _ := json.Marshal(jsonBody)

		body := NewMockReader(ctrl)
		gomock.InOrder(
			body.
				EXPECT().
				Read(gomock.Any()).
				DoAndReturn(func(p []byte) (int, error) {
					copy(p, rawBody)
					return len(rawBody), nil
				}),
			body.EXPECT().Read(gomock.Any()).Return(0, io.EOF),
		)

		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Method = method
		ctx.Request.URL = reqURL
		ctx.Request.Header = headers
		ctx.Request.Body = body

		data, _ := RequestReaderDefault(ctx)

		t.Run("retrieve the request method", func(t *testing.T) {
			if value := data["method"]; value != method {
				t.Errorf("stored the (%s) method value when expecting (%v)", value, method)
			}
		})

		t.Run("retrieve the request path", func(t *testing.T) {
			if value := data["path"]; value != path {
				t.Errorf("stored the (%s) path value when expecting (%v)", value, path)
			}
		})

		t.Run("retrieve the request params", func(t *testing.T) {
			if value := data["params"]; !reflect.DeepEqual(value, expParams) {
				t.Errorf("stored the (%s) params value when expecting (%v)", value, expParams)
			}
		})

		t.Run("retrieve the request headers", func(t *testing.T) {
			if value := data["headers"]; !reflect.DeepEqual(value, expHeaders) {
				t.Errorf("stored the (%v) headers when expecting (%v)", value, expHeaders)
			}
		})

		t.Run("retrieve the request body", func(t *testing.T) {
			if value := data["body"]; !reflect.DeepEqual(value, string(rawBody)) {
				t.Errorf("stored the (%v) body when expecting (%v)", value, string(rawBody))
			}
		})
	})
}
