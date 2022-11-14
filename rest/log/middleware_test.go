package log

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
	slog "github.com/happyhippyhippo/slate/log"
)

func Test_NewMiddlewareGenerator(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := NewMockLogger(ctrl)

		generator, err := NewMiddlewareGenerator(nil, logger)
		switch {
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		case generator != nil:
			t.Error("returned an unexpected valid middleware reference")
		}
	})

	t.Run("nil logger", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfigManager(ctrl)

		generator, err := NewMiddlewareGenerator(cfg, nil)
		switch {
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		case generator != nil:
			t.Error("returned an unexpected valid middleware reference")
		}
	})

	t.Run("valid instance", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfigManager(ctrl)
		logger := NewMockLogger(ctrl)

		generator, err := NewMiddlewareGenerator(cfg, logger)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error : %v", err)
		case generator == nil:
			t.Error("didn't returned the expected valid middleware generator reference")
		}
	})

	t.Run("error retrieving the endpoint status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "id"
		expected := fmt.Errorf("error message")
		writer := NewMockResponseWriter(ctrl)
		ctx := &gin.Context{}
		ctx.Writer = writer
		ctx.Request = &http.Request{}
		ctx.Request.URL = &url.URL{}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().Int("slate.endpoints.id.status").Return(0, expected).Times(1)
		logger := NewMockLogger(ctrl)

		generator, _ := NewMiddlewareGenerator(cfg, logger)
		middleware, e := generator(id)
		switch {
		case middleware != nil:
			t.Error("returned an unexpected instance of the middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting : %v", e, expected)
		}
	})

	t.Run("call next handler", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		RequestChannel = "channel.request"
		RequestLevel = slog.WARNING
		ResponseChannel = "channel.response"
		ResponseLevel = slog.ERROR
		defer func() {
			RequestChannel = "Request"
			RequestLevel = slog.DEBUG
			ResponseChannel = "Response"
			ResponseLevel = slog.INFO
		}()

		id := "id"
		status := 201
		writer := NewMockResponseWriter(ctrl)
		writer.EXPECT().Status().Return(status).Times(1)
		writer.EXPECT().Header().Return(map[string][]string{}).Times(1)
		ctx := &gin.Context{}
		ctx.Writer = writer
		ctx.Request = &http.Request{}
		ctx.Request.URL = &url.URL{}
		callCount := 0
		var next gin.HandlerFunc = func(context *gin.Context) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
				return
			}
			callCount++
		}
		request := map[string]interface{}{
			"body":    "",
			"headers": map[string]interface{}{},
			"method":  "",
			"params":  map[string]interface{}{},
			"path":    "",
		}
		response := map[string]interface{}{
			"body":    "",
			"headers": map[string]interface{}{},
			"status":  status,
		}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().Int("slate.endpoints.id.status").Return(200, nil).Times(1)
		logger := NewMockLogger(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(RequestChannel, RequestLevel, RequestMessage, map[string]interface{}{"request": request}),
			logger.EXPECT().Signal(ResponseChannel, ResponseLevel, ResponseMessage, map[string]interface{}{"request": request, "response": response, "duration": int64(0)}),
		)

		generator, _ := NewMiddlewareGenerator(cfg, logger)
		middleware, _ := generator(id)
		handler := middleware(next)
		handler(ctx)

		if callCount != 1 {
			t.Errorf("didn't called the next handler")
		}
	})

	t.Run("log JSON response body on invalid status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		DecorateJSON = true
		defer func() { DecorateJSON = true }()

		id := "id"
		status := 200
		body := "{\"field\": \"value\"}"
		writer := NewMockResponseWriter(ctrl)
		writer.EXPECT().Status().Return(status + 1).Times(1)
		writer.EXPECT().Header().Return(map[string][]string{"Accept": {"*/*"}}).Times(1)
		writer.EXPECT().WriteString(body).Return(20, nil).Times(1)
		ctx := &gin.Context{}
		ctx.Writer = writer
		ctx.Request = &http.Request{}
		ctx.Request.URL = &url.URL{}
		ctx.Request.Header = map[string][]string{"Accept": {"*/*"}}
		callCount := 0
		var next gin.HandlerFunc = func(context *gin.Context) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
				return
			}
			_, _ = context.Writer.WriteString(body)
			callCount++
		}
		request := map[string]interface{}{
			"body":    "",
			"headers": map[string]interface{}{"Accept": "*/*"},
			"method":  "",
			"params":  map[string]interface{}{},
			"path":    "",
		}
		response := map[string]interface{}{
			"body":     body,
			"bodyJson": map[string]interface{}{"field": "value"},
			"headers":  map[string]interface{}{"Accept": "*/*"},
			"status":   status + 1,
		}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().Int("slate.endpoints.id.status").Return(status, nil).Times(1)
		logger := NewMockLogger(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(RequestChannel, RequestLevel, RequestMessage, map[string]interface{}{"request": request}),
			logger.EXPECT().Signal(ResponseChannel, ResponseLevel, ResponseMessage, map[string]interface{}{"request": request, "response": response, "duration": int64(0)}),
		)

		generator, _ := NewMiddlewareGenerator(cfg, logger)
		middleware, _ := generator(id)
		handler := middleware(next)
		handler(ctx)

		if callCount != 1 {
			t.Errorf("didn't called the next handler")
		}
	})

	t.Run("log XML response body on invalid status code", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		DecorateJSON = false
		DecorateXML = true
		defer func() {
			DecorateJSON = true
			DecorateXML = false
		}()

		id := "id"
		status := 200
		body := "<field>value</field>"
		writer := NewMockResponseWriter(ctrl)
		writer.EXPECT().Status().Return(status + 1).Times(1)
		writer.EXPECT().Header().Return(map[string][]string{"Accept": {gin.MIMEXML}}).Times(1)
		writer.EXPECT().WriteString(body).Return(20, nil).Times(1)
		ctx := &gin.Context{}
		ctx.Writer = writer
		ctx.Request = &http.Request{}
		ctx.Request.URL = &url.URL{}
		ctx.Request.Header = map[string][]string{"Accept": {gin.MIMEXML}}
		callCount := 0
		var next gin.HandlerFunc = func(context *gin.Context) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
				return
			}
			_, _ = context.Writer.WriteString(body)
			callCount++
		}
		request := map[string]interface{}{
			"body":    "",
			"headers": map[string]interface{}{"Accept": gin.MIMEXML},
			"method":  "",
			"params":  map[string]interface{}{},
			"path":    "",
		}
		response := map[string]interface{}{
			"body":    body,
			"bodyXml": nil,
			"headers": map[string]interface{}{"Accept": gin.MIMEXML},
			"status":  status + 1,
		}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().Int("slate.endpoints.id.status").Return(status, nil).Times(1)
		logger := NewMockLogger(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(RequestChannel, RequestLevel, RequestMessage, map[string]interface{}{"request": request}),
			logger.EXPECT().Signal(ResponseChannel, ResponseLevel, ResponseMessage, map[string]interface{}{"request": request, "response": response, "duration": int64(0)}),
		)

		generator, _ := NewMiddlewareGenerator(cfg, logger)
		middleware, _ := generator(id)
		handler := middleware(next)
		handler(ctx)

		if callCount != 1 {
			t.Errorf("didn't called the next handler")
		}
	})
}
