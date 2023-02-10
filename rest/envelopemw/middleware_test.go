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
	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/rest/envelope"
)

func Test_NewMiddlewareGenerator(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := NewMockLog(ctrl)

		generator, e := NewMiddlewareGenerator(nil, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil logger", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgManager := NewMockConfigManager(ctrl)

		generator, e := NewMiddlewareGenerator(cfgManager, nil)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error getting the service id from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(0, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogServiceErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("default to error level logging on invalid log level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = "invalid"
		defer func() { LogLevel = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(0, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogServiceErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log for environment defined channel when getting the service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(0, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal("test", log.ERROR, LogServiceErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log with environment defined level when getting the service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.WARNING]
		defer func() { LogLevel = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(0, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.WARNING, LogServiceErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log with environment defined message when getting the service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogServiceErrorMessage
		LogServiceErrorMessage = "test"
		defer func() { LogServiceErrorMessage = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(0, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, "test", log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error getting the service accept list from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil).Times(1)
		cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return(nil, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogAcceptListErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log for environment defined channel when getting the service accept list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil).Times(1)
		cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return(nil, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal("test", log.ERROR, LogAcceptListErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log with environment defined level when getting the service accept list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.WARNING]
		defer func() { LogLevel = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil).Times(1)
		cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return(nil, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.WARNING, LogAcceptListErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log with environment defined message when getting the service accept list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogAcceptListErrorMessage
		LogAcceptListErrorMessage = "test"
		defer func() { LogAcceptListErrorMessage = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil).Times(1)
		cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return(nil, expected).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, "test", log.Context{"error": expected}).Return(nil).Times(1)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("valid generator instantiation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil).Times(1)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, e := NewMiddlewareGenerator(cfgManager, logger)
		switch {
		case generator == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error while retrieving endpoint path when generating middleware", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(0, expected),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogEndpointErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("returned an unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log for environment defined channel when retrieving endpoint path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		endpoint := "index"
		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(0, expected),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal("test", log.ERROR, LogEndpointErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("returned an unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log for environment level channel when retrieving endpoint path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		endpoint := "index"
		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(0, expected),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.DEBUG, LogEndpointErrorMessage, log.Context{"error": expected}).Return(nil).Times(1)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("returned an unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("log with environment defined message when retrieving endpoint path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogEndpointErrorMessage
		LogEndpointErrorMessage = "test"
		defer func() { LogEndpointErrorMessage = prev }()

		endpoint := "index"
		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(0, expected),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, "test", log.Context{"error": expected}).Return(nil).Times(1)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, e := generator(endpoint)
		switch {
		case mw != nil:
			t.Error("returned an unexpected valid reference to a middleware")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("valid middleware creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, e := generator(endpoint)
		switch {
		case mw == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("calling the generated handler calls the given original handler", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
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
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
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
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("parse error stored in the response field of context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:1.e:2.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("parse invalid stored in the response field of context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", "string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:1.e:2.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("parse panic error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			panic(fmt.Errorf("error message"))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:1.e:2.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("panic non-error value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:1.e:2.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("registered observer update the service id value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.
				EXPECT().
				AddObserver(ServiceIDConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		callback(1, 2)

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:2.e:2.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("registered service id observer log on invalid new service id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.
				EXPECT().
				AddObserver(ServiceIDConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogServiceErrorMessage, log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(1, newValue)
	})

	t.Run("registered service id observer log on invalid new service id with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.
				EXPECT().
				AddObserver(ServiceIDConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal("test", log.ERROR, LogServiceErrorMessage, log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(1, newValue)
	})

	t.Run("registered service id observer log on invalid new service id with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.
				EXPECT().
				AddObserver(ServiceIDConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.DEBUG, LogServiceErrorMessage, log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(1, newValue)
	})

	t.Run("registered service id observer log on invalid new service id with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogServiceErrorMessage
		LogServiceErrorMessage = "test"
		defer func() { LogServiceErrorMessage = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.
				EXPECT().
				AddObserver(ServiceIDConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, "test", log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(1, newValue)
	})

	t.Run("registered observer update the accept formats value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		callback([]interface{}{gin.MIMEXML}, []interface{}{gin.MIMEJSON})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:1.e:2.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("registered accept format observer log on invalid new format list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogAcceptListErrorMessage, log.Context{"list": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered accept format observer log on invalid new format list with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal("test", log.ERROR, LogAcceptListErrorMessage, log.Context{"list": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered accept format observer log on invalid new format list with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.DEBUG, LogAcceptListErrorMessage, log.Context{"list": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered accept format observer log on invalid new format list with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogAcceptListErrorMessage
		LogAcceptListErrorMessage = "test"
		defer func() { LogAcceptListErrorMessage = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, "test", log.Context{"list": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered accept format observer log on invalid new format list entry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		invalidValue := 2
		newValue := []interface{}{gin.MIMEJSON, invalidValue}
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogAcceptListErrorMessage, log.Context{"value": invalidValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered accept format observer log on invalid new format list entry with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		invalidValue := 2
		newValue := []interface{}{gin.MIMEJSON, invalidValue}
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal("test", log.ERROR, LogAcceptListErrorMessage, log.Context{"value": invalidValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered accept format observer log on invalid new format list entry with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		invalidValue := 2
		newValue := []interface{}{gin.MIMEJSON, invalidValue}
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.DEBUG, LogAcceptListErrorMessage, log.Context{"value": invalidValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered accept format observer log on invalid new format list entry with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogAcceptListErrorMessage
		LogAcceptListErrorMessage = "test"
		defer func() { LogAcceptListErrorMessage = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		invalidValue := 2
		newValue := []interface{}{gin.MIMEJSON, invalidValue}
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEXML}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver(FormatAcceptListConfigPath, gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
			cfgManager.EXPECT().AddObserver("slate.rest.endpoints.index.id", gomock.Any()).Return(nil),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, "test", log.Context{"value": invalidValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback([]interface{}{gin.MIMEXML}, newValue)
	})

	t.Run("registered endpoint id observer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := 10
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver("slate.rest.endpoints.index.id", gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
		)
		logger := NewMockLog(ctrl)

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		mw, _ := generator(endpoint)

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		callback(2, newValue)

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"error":[{"code":"s:1.e:10.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("registered endpoint id observer log on invalid new id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver("slate.rest.endpoints.index.id", gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, LogEndpointErrorMessage, log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(2, newValue)
	})

	t.Run("registered endpoint id observer log on invalid new id with environment defined channel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogChannel
		LogChannel = "test"
		defer func() { LogChannel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver("slate.rest.endpoints.index.id", gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal("test", log.ERROR, LogEndpointErrorMessage, log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(2, newValue)
	})

	t.Run("registered endpoint id observer log on invalid new id with environment defined level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogLevel
		LogLevel = log.LevelMapName[log.DEBUG]
		defer func() { LogLevel = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver("slate.rest.endpoints.index.id", gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.DEBUG, LogEndpointErrorMessage, log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(2, newValue)
	})

	t.Run("registered endpoint id observer log on invalid new id with environment defined message", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LogEndpointErrorMessage
		LogEndpointErrorMessage = "test"
		defer func() { LogEndpointErrorMessage = prev }()

		var callback func(old interface{}, new interface{})

		endpoint := "index"
		newValue := "string"
		cfgManager := NewMockConfigManager(ctrl)
		gomock.InOrder(
			cfgManager.EXPECT().Int(ServiceIDConfigPath, 0).Return(1, nil),
			cfgManager.EXPECT().Int("slate.rest.endpoints.index.id", 0).Return(2, nil),
		)
		cfgManager.EXPECT().List(FormatAcceptListConfigPath).Return([]interface{}{gin.MIMEJSON}, nil).Times(1)
		gomock.InOrder(
			cfgManager.EXPECT().AddObserver(ServiceIDConfigPath, gomock.Any()).Return(nil),
			cfgManager.EXPECT().AddObserver(FormatAcceptListConfigPath, gomock.Any()).Return(nil),
			cfgManager.
				EXPECT().
				AddObserver("slate.rest.endpoints.index.id", gomock.Any()).
				DoAndReturn(func(id string, cb config.IObserver) error {
					callback = cb
					return nil
				}),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().Signal(LogChannel, log.ERROR, "test", log.Context{"value": newValue})

		generator, _ := NewMiddlewareGenerator(cfgManager, logger)
		_, _ = generator(endpoint)

		callback(2, newValue)
	})
}
