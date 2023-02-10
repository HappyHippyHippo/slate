package logmw

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_ResponseReaderDefault(t *testing.T) {
	t.Run("nil writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		statusCode := 200
		if _, e := ResponseReaderDefault(nil, nil, statusCode); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("don't store the body if status code is the expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		statusCode := 200
		headers := map[string][]string{"header1": {"value1", "value2"}, "header2": {"value3"}}
		expHeaders := log.Context{"header1": []string{"value1", "value2"}, "header2": "value3"}
		jsonBody := map[string]interface{}{"field": "value"}
		rawBody, _ := json.Marshal(jsonBody)
		ginWriter := NewMockResponseWriter(ctrl)
		ginWriter.EXPECT().Status().Return(statusCode).Times(1)
		ginWriter.EXPECT().Header().Return(headers).Times(1)
		w, _ := newResponseWriter(ginWriter)
		w.(*writer).body.Write(rawBody)

		if data, e := ResponseReaderDefault(nil, w, statusCode); e != nil {
			t.Errorf("returned the unextected (%v) error", e)
		} else if value := data["status"]; value != statusCode {
			t.Errorf("stored the (%s) status value", value)
		} else if value := data["headers"]; !reflect.DeepEqual(value, expHeaders) {
			t.Errorf("stored the (%v) headers", value)
		} else if value, exists := data["body"]; exists {
			t.Errorf("stored the (%v) body", value)
		}
	})

	t.Run("store the body if status code is different then the expected", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		statusCode := 200
		headers := map[string][]string{"header1": {"value1", "value2"}, "header2": {"value3"}}
		expHeaders := log.Context{"header1": []string{"value1", "value2"}, "header2": "value3"}
		jsonBody := map[string]interface{}{"field": "value"}
		rawBody, _ := json.Marshal(jsonBody)
		ginWriter := NewMockResponseWriter(ctrl)
		ginWriter.EXPECT().Status().Return(statusCode).Times(1)
		ginWriter.EXPECT().Header().Return(headers).Times(1)
		w, _ := newResponseWriter(ginWriter)
		w.(*writer).body.Write(rawBody)

		if data, e := ResponseReaderDefault(nil, w, statusCode+1); e != nil {
			t.Errorf("returned the unextected (%v) error", e)
		} else if value := data["status"]; value != statusCode {
			t.Errorf("stored the (%s) status value", value)
		} else if value := data["headers"]; !reflect.DeepEqual(value, expHeaders) {
			t.Errorf("stored the (%v) headers", value)
		} else if value := data["body"]; !reflect.DeepEqual(value, string(rawBody)) {
			t.Errorf("stored the (%v) body", value)
		}
	})
}
