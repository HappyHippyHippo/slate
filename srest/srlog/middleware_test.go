package srlog

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/happyhippyhippo/slate/slog"
	"net/http"
	"testing"
)

func Test_NewMiddleware(t *testing.T) {
	t.Run("nil logger", func(t *testing.T) {
		middleware, err := NewMiddleware(nil, RequestReaderDefault, NewResponseReaderDefault(http.StatusOK))
		switch {
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		case middleware != nil:
			t.Error("returned an unexpected valid middleware reference")
		}
	})

	t.Run("nil request reader", func(t *testing.T) {
		middleware, err := NewMiddleware(slog.NewLogger(), nil, NewResponseReaderDefault(http.StatusOK))
		switch {
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		case middleware != nil:
			t.Error("returned an unexpected valid middleware reference")
		}
	})

	t.Run("nil response reader", func(t *testing.T) {
		middleware, err := NewMiddleware(slog.NewLogger(), RequestReaderDefault, nil)
		switch {
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		case middleware != nil:
			t.Error("returned an unexpected valid middleware reference")
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

		writer := NewMockResponseWriter(ctrl)
		ctx := &gin.Context{}
		ctx.Writer = writer
		callCount := 0
		var next gin.HandlerFunc = func(context *gin.Context) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
				return
			}
			callCount++
		}
		request := map[string]interface{}{"type": "request"}
		response := map[string]interface{}{"type": "response"}
		logStream := NewMockStream(ctrl)
		gomock.InOrder(
			logStream.EXPECT().Signal(RequestChannel, RequestLevel, RequestMessage, map[string]interface{}{"request": request}),
			logStream.EXPECT().Signal(ResponseChannel, ResponseLevel, ResponseMessage, map[string]interface{}{"request": request, "response": response, "duration": int64(0)}),
		)
		logger := slog.NewLogger()
		_ = logger.AddStream("id", logStream)
		requestReader := func(_ *gin.Context) (map[string]interface{}, error) {
			return request, nil
		}
		responseReader := func(_ *gin.Context, _ responseWriter) (map[string]interface{}, error) {
			return response, nil
		}

		mw, _ := NewMiddleware(logger, requestReader, responseReader)
		handler := mw(next)
		handler(ctx)

		if callCount != 1 {
			t.Errorf("didn't called the next handler")
		}
	})
}
