package xml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/rest/logmw"
)

func Test_NewResponseReaderDecoratorXML(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		if _, e := NewDecorator(nil, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return nil, nil
		}
		decorator, _ := NewDecorator(reader, nil)

		result, e := decorator(nil, writer, 0)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		case result != nil:
			t.Errorf("returned the unexpeted context data : %v", result)
		}
	})

	t.Run("nil writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ctx := &gin.Context{}
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return nil, nil
		}
		decorator, _ := NewDecorator(reader, nil)

		result, e := decorator(ctx, nil, 0)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
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
		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return nil, expected
		}
		decorator, _ := NewDecorator(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !reflect.DeepEqual(e, expected):
			t.Errorf("(%v) when expecting (%v)", e, expected)
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
		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return data, nil
		}
		decorator, _ := NewDecorator(reader, nil)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyXml"]; ok {
				t.Error("added the bodyXml field")
			}
		}
	})

	t.Run("empty accept does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{}
		data := log.Context{"body": "<message><field>value</field></message>"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return data, nil
		}
		decorator, _ := NewDecorator(reader, &model)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if _, ok := result["bodyXml"]; ok {
				t.Error("added the bodyXml field")
			}
		}
	})

	t.Run("non-xml accept does not add decorated field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		model := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{}
		data := log.Context{"body": "<message><field>value</field></message>"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", gin.MIMEJSON)
		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return data, nil
		}
		decorator, _ := NewDecorator(reader, &model)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
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
		data := log.Context{"body": "<message> field value /field /message>"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", gin.MIMEXML)
		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return data, nil
		}
		decorator, _ := NewDecorator(reader, &model)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
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
		data := log.Context{"body": "<message><field>value</field></message>"}
		expected := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{XMLName: xml.Name{Local: "message"}, Field: "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", gin.MIMEXML)
		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return data, nil
		}
		decorator, _ := NewDecorator(reader, &model)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if body, ok := result["bodyXml"]; !ok {
				t.Error("didn't added the bodyXml field")
			} else if !reflect.DeepEqual(body, &expected) {
				t.Errorf("(%v) when expecting (%v)", body, &expected)
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
		data := log.Context{"body": "<message><field>value</field></message>"}
		expected := struct {
			XMLName xml.Name `xml:"message"`
			Field   string   `xml:"field"`
		}{XMLName: xml.Name{Local: "message"}, Field: "value"}
		ctx := &gin.Context{}
		ctx.Request = &http.Request{}
		ctx.Request.Header = http.Header{}
		ctx.Request.Header.Add("Accept", gin.MIMEXML2)
		writer := NewMockResponseWriter(ctrl)
		reader := func(_ *gin.Context, _ logmw.Writer, _ int) (log.Context, error) {
			return data, nil
		}
		decorator, _ := NewDecorator(reader, &model)

		result, e := decorator(ctx, writer, 0)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case result == nil:
			t.Error("didn't returned the expected context data")
		default:
			if body, ok := result["bodyXml"]; !ok {
				t.Error("didn't added the bodyXml field")
			} else if !reflect.DeepEqual(body, &expected) {
				t.Errorf("(%v) when expecting (%v)", body, &expected)
			}
		}
	})
}
