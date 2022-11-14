package log

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewRequestReaderDecoratorXML(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		if _, err := NewRequestReaderDecoratorXML(nil, nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := func(ctx *gin.Context) (map[string]interface{}, error) {
			return nil, nil
		}
		decorator, _ := NewRequestReaderDecoratorXML(reader, nil)

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
		reader := func(ctx *gin.Context) (map[string]interface{}, error) {
			return nil, expected
		}
		decorator, _ := NewRequestReaderDecoratorXML(reader, nil)

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

		model := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{}
		data := map[string]interface{}{"body": "<message><field>value</field></message>"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		reader := func(ctx *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorXML(reader, &model)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyXml"]; ok {
				t.Error("added the bodyXml field")
			}
		}
	})

	t.Run("non-xml content-type does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{}
		data := map[string]interface{}{"body": "<message><field>value</field></message>"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEJSON)
		reader := func(ctx *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorXML(reader, &model)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyXml"]; ok {
				t.Error("added the bodyXml field")
			}
		}
	})

	t.Run("invalid xml content does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{}
		data := map[string]interface{}{"body": "<message field value /field /message>"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML)
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML)
		reader := func(ctx *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorXML(reader, &model)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyXml"]; ok {
				t.Error("added the bodyXml field")
			}
		}
	})

	t.Run("correctly add decorated field for application/xml", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{}
		data := map[string]interface{}{"body": "<message><field>value</field></message>"}
		expected := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{XMLName: xml.Name{Local: "message"}, Field: "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML)
		reader := func(ctx *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorXML(reader, &model)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if body, ok := result["bodyXml"]; !ok {
				t.Error("didn't added the bodyXml field")
			} else if !reflect.DeepEqual(body, &expected) {
				t.Errorf("added the (%v) content when expecting : %v", body, &expected)
			}
		}
	})

	t.Run("correctly add decorated field for text/xml", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{}
		data := map[string]interface{}{"body": "<message><field>value</field></message>"}
		expected := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{XMLName: xml.Name{Local: "message"}, Field: "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Content-Type", gin.MIMEXML2)
		reader := func(ctx *gin.Context) (map[string]interface{}, error) {
			return data, nil
		}
		decorator, _ := NewRequestReaderDecoratorXML(reader, &model)

		result, err := decorator(ctx)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if body, ok := result["bodyXml"]; !ok {
				t.Error("didn't added the bodyXml field")
			} else if !reflect.DeepEqual(body, &expected) {
				t.Errorf("added the (%v) content when expecting : %v", body, &expected)
			}
		}
	})
}
