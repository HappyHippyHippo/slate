package envelope

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	sconfig "github.com/happyhippyhippo/slate/config"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewMiddlewareGenerator(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		generator, err := NewMiddlewareGenerator(nil)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error getting the service id from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(
			sconfig.Partial{"slate": sconfig.Partial{"service": sconfig.Partial{"id": "invalid"}}},
			nil,
		).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)

		generator, err := NewMiddlewareGenerator(cfg)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error getting the service encode accept list from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(
			sconfig.Partial{
				"slate": sconfig.Partial{
					"service": sconfig.Partial{"id": 1},
					"rest":    sconfig.Partial{"accept": "invalid"},
				},
			},
			nil,
		).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)

		generator, err := NewMiddlewareGenerator(cfg)
		switch {
		case generator != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("running will return a middleware function generator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(
			sconfig.Partial{
				"slate": sconfig.Partial{
					"service": sconfig.Partial{"id": 1},
					"rest":    sconfig.Partial{"accept": []interface{}{"application/json"}},
				},
			},
			nil,
		).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)

		if generator, err := NewMiddlewareGenerator(cfg); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if generator == nil {
			t.Error("didn't returned the expected generator function")
		}
	})

	t.Run("error while retrieving endpoint path when generating middleware", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": "invalid",
					},
				},
			},
		}
		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(cfgData, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		generator, _ := NewMiddlewareGenerator(cfg)

		mw, err := generator("index")
		switch {
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		case mw != nil:
			t.Error("returned an unexpected valid reference to a middleware")
		}
	})

	t.Run("calling the generated handler calls the given original handler", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(cfgData, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		generator, _ := NewMiddlewareGenerator(cfg)

		calls := 0
		mw, _ := generator("index")
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

		cfgData := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(cfgData, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		generator, _ := NewMiddlewareGenerator(cfg)
		mw, _ := generator("index")

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", NewEnvelope(200, []string{"data1", "data2"}, nil))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":true,"errors":[]},"data":["data1","data2"]}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("parse error stored in the response field of context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(cfgData, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		generator, _ := NewMiddlewareGenerator(cfg)
		mw, _ := generator("index")

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"errors":[{"code":"s:1.e:2.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("parse invalid stored in the response field of context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(cfgData, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		generator, _ := NewMiddlewareGenerator(cfg)
		mw, _ := generator("index")

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", "string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"errors":[{"code":"s:1.e:2.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("parse panic error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(cfgData, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		generator, _ := NewMiddlewareGenerator(cfg)
		mw, _ := generator("index")

		handler := mw(func(ctx *gin.Context) {
			panic(fmt.Errorf("error message"))
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"errors":[{"code":"s:1.e:2.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("panic non-error value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source := NewMockConfigManager(ctrl)
		source.EXPECT().Get("").Return(cfgData, nil).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		generator, _ := NewMiddlewareGenerator(cfg)
		mw, _ := generator("index")

		handler := mw(func(ctx *gin.Context) {
			panic("string message")
		})

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"errors":[{"code":"s:1.e:2.c:0","message":"internal server error"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("registered observer update the service id value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData1 := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source1 := NewMockConfigManager(ctrl)
		source1.EXPECT().Get("").Return(cfgData1, nil).MinTimes(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id1", 0, source1)
		generator, _ := NewMiddlewareGenerator(cfg)
		mw, _ := generator("index")

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		cfgData2 := sconfig.Partial{"slate": sconfig.Partial{"service": sconfig.Partial{"id": 2}}}
		source2 := NewMockConfigManager(ctrl)
		source2.EXPECT().Get("").Return(cfgData2, nil).MinTimes(1)
		_ = cfg.AddSource("id2", 10, source2)

		time.Sleep(time.Millisecond * 100)

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `{"status":{"success":false,"errors":[{"code":"s:2.e:2.c:0","message":"error message"}]}}`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})

	t.Run("registered observer update the accepted mime type value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgData1 := sconfig.Partial{
			"slate": sconfig.Partial{
				"service": sconfig.Partial{
					"id": 1,
				},
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/json"},
				},
				"endpoints": sconfig.Partial{
					"index": sconfig.Partial{
						"id": 2,
					},
				},
			},
		}
		source1 := NewMockConfigManager(ctrl)
		source1.EXPECT().Get("").Return(cfgData1, nil).MinTimes(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id1", 0, source1)
		generator, _ := NewMiddlewareGenerator(cfg)
		mw, _ := generator("index")

		handler := mw(func(ctx *gin.Context) {
			ctx.Set("response", fmt.Errorf("error message"))
		})

		cfgData2 := sconfig.Partial{
			"slate": sconfig.Partial{
				"rest": sconfig.Partial{
					"accept": []interface{}{"application/xml"},
				},
			},
		}
		source2 := NewMockConfigManager(ctrl)
		source2.EXPECT().Get("").Return(cfgData2, nil).MinTimes(1)
		_ = cfg.AddSource("id2", 10, source2)

		time.Sleep(time.Millisecond * 100)

		gin.SetMode(gin.ReleaseMode)
		writer := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(writer)
		ctx.Request = &http.Request{}
		handler(ctx)

		expected := `<envelope><status><success>false</success><errors><error code="s:1.e:2.c:0" message="error message"></error></errors></status></envelope>`

		if check := writer.Body.String(); check != expected {
			t.Errorf("parsed (%v) response data when expecting : %v", check, expected)
		}
	})
}
