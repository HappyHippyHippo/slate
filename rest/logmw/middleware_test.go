package logmw

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewMiddlewareGenerator(t *testing.T) {
	t.Run("nil logger", func(t *testing.T) {
		generator, e := NewMiddlewareGenerator(nil, RequestReaderDefault, ResponseReaderDefault)
		switch {
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		case generator != nil:
			t.Error("returned an unexpected valid middleware generator reference")
		}
	})

	t.Run("nil request reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := NewMockLog(ctrl)

		generator, e := NewMiddlewareGenerator(logger, nil, ResponseReaderDefault)
		switch {
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		case generator != nil:
			t.Error("returned an unexpected valid middleware generator reference")
		}
	})

	t.Run("nil response reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := NewMockLog(ctrl)

		generator, e := NewMiddlewareGenerator(logger, RequestReaderDefault, nil)
		switch {
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		case generator != nil:
			t.Error("returned an unexpected valid middleware generator reference")
		}
	})

	t.Run("valid middleware generator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := NewMockLog(ctrl)

		generator, e := NewMiddlewareGenerator(logger, RequestReaderDefault, ResponseReaderDefault)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		case generator == nil:
			t.Error("didn't returned the expected middleware generator reference")
		}
	})

	t.Run("correctly call next handler", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		RequestChannel = "channel.request"
		RequestLevel = log.WARNING
		ResponseChannel = "channel.response"
		ResponseLevel = log.ERROR
		defer func() {
			RequestChannel = "Request"
			RequestLevel = log.DEBUG
			ResponseChannel = "Response"
			ResponseLevel = log.INFO
		}()

		statusCode := 123
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
		request := log.Context{"type": "request"}
		response := log.Context{"type": "response"}
		logger := NewMockLog(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(RequestChannel, RequestLevel, RequestMessage, log.Context{"request": request}),
			logger.EXPECT().Signal(ResponseChannel, ResponseLevel, ResponseMessage, log.Context{"request": request, "response": response, "duration": int64(0)}),
		)
		requestReader := func(context *gin.Context) (log.Context, error) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
			}
			return request, nil
		}
		responseReader := func(context *gin.Context, _ responseWriter, sc int) (log.Context, error) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
			}
			if sc != statusCode {
				t.Errorf("handler called with unexpected status code")
			}
			return response, nil
		}

		generator, _ := NewMiddlewareGenerator(logger, requestReader, responseReader)
		mw := generator(statusCode)
		handler := mw(next)
		handler(ctx)

		if callCount != 1 {
			t.Errorf("didn't called the next handler")
		}
	})
}
