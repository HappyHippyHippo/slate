package envelopemw

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/envelope"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewMiddlewareGenerator(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		generator, e := NewMiddlewareGenerator(nil, log.NewLog())
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil logger", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		generator, e := NewMiddlewareGenerator(config.NewConfig(), nil)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error getting the service id from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.ERROR,
				LogServiceErrorMessage,
				log.Context{"error": errConversion("invalid", "int")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("default to error level logging on invalid log level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = "invalid"
		defer func() { LogLevel = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.ERROR,
				LogServiceErrorMessage,
				log.Context{"error": errConversion("invalid", "int")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log for environment defined channel when getting the service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "new channel"
		defer func() { LogChannel = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.ERROR,
				LogServiceErrorMessage,
				log.Context{"error": errConversion("invalid", "int")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log with environment defined level when getting the service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = "warning"
		defer func() { LogLevel = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.WARNING,
				LogServiceErrorMessage,
				log.Context{"error": errConversion("invalid", "int")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log with environment defined message when getting the service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogServiceErrorMessage
		LogServiceErrorMessage = "test"
		defer func() { LogServiceErrorMessage = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.ERROR,
				"test",
				log.Context{"error": errConversion("invalid", "int")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error getting the service accept list from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.ERROR,
				LogAcceptListErrorMessage,
				log.Context{"error": errConversion("invalid", "[]interface{}")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log for environment defined channel when getting the service accept list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				"test",
				log.ERROR,
				LogAcceptListErrorMessage,
				log.Context{"error": errConversion("invalid", "[]interface{}")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log with environment defined level when getting the service accept list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = "warning"
		defer func() { LogLevel = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.WARNING,
				LogAcceptListErrorMessage,
				log.Context{"error": errConversion("invalid", "[]interface{}")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log with environment defined message when getting the service accept list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogAcceptListErrorMessage
		LogAcceptListErrorMessage = "test"
		defer func() { LogAcceptListErrorMessage = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal(
				LogChannel,
				log.ERROR,
				"test",
				log.Context{"error": errConversion("invalid", "[]interface{}")}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("valid generator instantiation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logger := log.NewLog()

		generator, e := NewMiddlewareGenerator(cfg, logger)
		switch {
		case generator == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("error while retrieving endpoint path when generating middleware", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", "string")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "Invalid endpoint id", gomock.Any()).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)

		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log for environment defined channel when retrieving endpoint path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", "string")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("test", log.ERROR, "Invalid endpoint id", gomock.Any()).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)

		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log for environment level channel when retrieving endpoint path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", "string")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.DEBUG, "Invalid endpoint id", gomock.Any()).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)

		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("log with environment defined message when retrieving endpoint path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogEndpointErrorMessage
		LogEndpointErrorMessage = "test"
		defer func() { LogEndpointErrorMessage = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", "string")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "test", gomock.Any()).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)

		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("valid middleware creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 123)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)

		mw, e := generator(endpoint)
		switch {
		case mw == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("calling the generated handler calls the given original handler", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 123)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		calls := 0
		handler := mw(func(*gin.Context) {
			calls++
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)

		handler(ctx)

		if calls != 1 {
			t.Errorf("didn't called the original underlying handler")
		}
	})

	t.Run("parse data envelope stored in the response field of context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 123)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", envelope.NewEnvelope(200, []string{"data1", "data2"}, nil))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":true,"error":[]},"data":["data1","data2"]}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("parse error stored in the response field of context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("parse invalid stored in the response field of context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", "string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("parse panic error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			panic(fmt.Errorf("error message"))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("panic non-error value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered observer update the service id value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.service.id", 321)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		_ = cfg.AddSource("id2", 1, newSource)

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:321.e:456.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered service id observer log on invalid new service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "Invalid service id", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered service id observer log on invalid new service id with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("test", log.ERROR, "Invalid service id", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered service id observer log on invalid new service id with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.DEBUG, "Invalid service id", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered service id observer log on invalid new service id with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogServiceErrorMessage
		LogServiceErrorMessage = "test"
		defer func() { LogServiceErrorMessage = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.service.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "test", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered observer update the accept formats value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"text/xml"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "Invalid accept list", log.Context{"list": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("test", log.ERROR, "Invalid accept list", log.Context{"list": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.DEBUG, "Invalid accept list", log.Context{"list": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogAcceptListErrorMessage
		LogAcceptListErrorMessage = "test"
		defer func() { LogAcceptListErrorMessage = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "test", log.Context{"list": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list entry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", []interface{}{"application/json", 123})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "Invalid accept list", log.Context{"value": 123}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list entry with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", []interface{}{"application/json", 123})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("test", log.ERROR, "Invalid accept list", log.Context{"value": 123}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list entry with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", []interface{}{"application/json", 123})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.DEBUG, "Invalid accept list", log.Context{"value": 123}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list entry with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogAcceptListErrorMessage
		LogAcceptListErrorMessage = "test"
		defer func() { LogAcceptListErrorMessage = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.accept", []interface{}{"application/json", 123})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "test", log.Context{"value": 123}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered endpoint id observer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.endpoints.index.id", 654)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:654.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered endpoint id observer log on invalid new id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.endpoints.index.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "Invalid endpoint id", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered endpoint id observer log on invalid new id with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.endpoints.index.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("test", log.ERROR, "Invalid endpoint id", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered endpoint id observer log on invalid new id with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.endpoints.index.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.DEBUG, "Invalid endpoint id", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})

	t.Run("registered endpoint id observer log on invalid new id with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogEndpointErrorMessage
		LogEndpointErrorMessage = "test"
		defer func() { LogEndpointErrorMessage = prev }()

		endpoint := "index"
		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.service.id", 123)
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		_, _ = partial.Set("slate.api.rest.endpoints.index.id", 456)
		newPartial := config.Partial{}
		_, _ = newPartial.Set("slate.api.rest.endpoints.index.id", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).AnyTimes()
		newSource := NewMockConfigSource(ctrl)
		newSource.EXPECT().Get("").Return(newPartial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 0, source)
		logStream := NewMockLogStream(ctrl)
		logStream.
			EXPECT().
			Signal("rest", log.ERROR, "test", log.Context{"value": "invalid"}).
			Return(nil).
			Times(1)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		generator, _ := NewMiddlewareGenerator(cfg, logger)
		mw, _ := generator(endpoint)

		_ = cfg.AddSource("id2", 1, newSource)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:123.e:456.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})
}
