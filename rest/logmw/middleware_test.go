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
		reqReader := func(ctx *gin.Context) (log.Context, error) { return nil, nil }
		resReader := func(ctx *gin.Context, writer Writer, statusCode int) (log.Context, error) { return nil, nil }

		generator, e := NewMiddlewareGenerator(nil, reqReader, resReader)
		switch {
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		case generator != nil:
			t.Error("unexpected valid middleware generator reference")
		}
	})

	t.Run("nil request reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		resReader := func(ctx *gin.Context, writer Writer, statusCode int) (log.Context, error) { return nil, nil }

		generator, e := NewMiddlewareGenerator(log.NewLog(), nil, resReader)
		switch {
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		case generator != nil:
			t.Error("unexpected valid middleware generator reference")
		}
	})

	t.Run("nil response reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reqReader := func(ctx *gin.Context) (log.Context, error) { return nil, nil }

		generator, e := NewMiddlewareGenerator(log.NewLog(), reqReader, nil)
		switch {
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		case generator != nil:
			t.Error("unexpected valid middleware generator reference")
		}
	})

	t.Run("valid middleware generator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reqReader := func(ctx *gin.Context) (log.Context, error) { return nil, nil }
		resReader := func(ctx *gin.Context, writer Writer, statusCode int) (log.Context, error) { return nil, nil }

		generator, e := NewMiddlewareGenerator(log.NewLog(), reqReader, resReader)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
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
		req := log.Context{"type": "request"}
		res := log.Context{"type": "response"}
		logStream := NewMockLogStream(ctrl)
		gomock.InOrder(
			logStream.
				EXPECT().
				Signal(
					RequestChannel,
					RequestLevel,
					RequestMessage,
					log.Context{"request": req},
				),
			logStream.
				EXPECT().
				Signal(
					ResponseChannel,
					ResponseLevel,
					ResponseMessage,
					log.Context{"request": req, "response": res, "duration": int64(0)},
				),
		)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		requestReader := func(context *gin.Context) (log.Context, error) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
			}
			return req, nil
		}
		responseReader := func(context *gin.Context, _ Writer, sc int) (log.Context, error) {
			if context != ctx {
				t.Errorf("handler called with unexpected context instance")
			}
			if sc != statusCode {
				t.Errorf("handler called with unexpected status code")
			}
			return res, nil
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
