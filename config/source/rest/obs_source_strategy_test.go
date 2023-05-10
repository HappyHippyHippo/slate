package rest

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewObsSourceStrategy(t *testing.T) {
	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewObsSourceStrategy(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new observable rest source factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoderFactory := NewMockDecoderFactory(ctrl)

		sut, e := NewObsSourceStrategy(decoderFactory)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.decoderFactory != decoderFactory:
			t.Error("didn't stored the decoder factory reference")
		default:
			client := sut.clientFactory()
			switch client.(type) {
			case *http.Client:
			default:
				t.Error("didn't stored a valid http client factory")
			}
		}
	})
}

func Test_ObsSourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(&config.Config{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(&config.Config{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(&config.Config{"type": config.UnknownSourceType}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		if !sut.Accept(&config.Config{"type": ObsType}) {
			t.Error("returned false")
		}
	})
}

func Test_ObsSourceStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("missing uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&config.Config{"format": "format", "timestampPath": "path", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, config.ErrInvalidSource):
			t.Errorf("returned the (%v) error when expecting (%v)", e, config.ErrInvalidSource)
		}
	})

	t.Run("missing timestamp path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&config.Config{"uri": "uri", "format": "format", "path": config.Config{"config": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, config.ErrInvalidSource):
			t.Errorf("returned the (%v) error when expecting (%v)", e, config.ErrInvalidSource)
		}
	})

	t.Run("missing config path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&config.Config{"uri": "uri", "format": "format", "path": config.Config{"timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, config.ErrInvalidSource):
			t.Errorf("returned the (%v) error when expecting (%v)", e, config.ErrInvalidSource)
		}
	})

	t.Run("non-string uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&config.Config{"uri": 123, "format": "format", "path": config.Config{"config": "path", "timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&config.Config{"uri": "uri", "format": 123, "path": config.Config{"config": "path", "timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-string timestamp path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&config.Config{"uri": "uri", "format": "format", "path": config.Config{"config": "path", "timestamp": 123}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-string config path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObsSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&config.Config{"uri": "uri", "format": "format", "path": config.Config{"config": 123, "timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("create the rest source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uri := "uri"
		format := "format"
		timestampPath := "timestamp"
		configPath := "path"
		field := "field"
		value := "value"
		expected := config.Config{field: value}
		respData := config.Config{"path": config.Config{"field": "value"}, "timestamp": "2000-01-01T00:00:00.000Z"}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create("format", gomock.Any()).Return(decoder, nil).Times(1)

		sut, _ := NewObsSourceStrategy(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.clientFactory = func() httpClient { return client }

		src, e := sut.Create(&config.Config{"uri": uri, "format": format, "path": config.Config{"config": configPath, "timestamp": timestampPath}})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *ObsSource:
				if !reflect.DeepEqual(s.Config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})

	t.Run("create the rest source defaulting format if not present in config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uri := "uri"
		timestampPath := "timestamp"
		configPath := "path"
		field := "field"
		value := "value"
		expected := config.Config{field: value}
		respData := config.Config{"path": config.Config{"field": "value"}, "timestamp": "2000-01-01T00:00:00.000Z"}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create("json", gomock.Any()).Return(decoder, nil).Times(1)

		sut, _ := NewObsSourceStrategy(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.clientFactory = func() httpClient { return client }

		src, e := sut.Create(&config.Config{"uri": uri, "path": config.Config{"config": configPath, "timestamp": timestampPath}})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *ObsSource:
				if !reflect.DeepEqual(s.Config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})
}
