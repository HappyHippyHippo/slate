package config

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewObservableRestSourceStrategy(t *testing.T) {
	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewObservableRestSourceStrategy(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("new observable rest source factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoderFactory := NewMockDecoderFactory(ctrl)

		sut, e := NewObservableRestSourceStrategy(decoderFactory)
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

func Test_SourceStrategyRestObservable_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{"type": SourceUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		if !sut.Accept(&Config{"type": SourceObservableRest}) {
			t.Error("returned false")
		}
	})
}

func Test_ObservableRestSourceStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("missing uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"format": "format", "timestampPath": "path", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.ConfigPathNotFound)
		}
	})

	t.Run("missing timestamp path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"uri": "uri", "format": "format", "path": Config{"config": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.ConfigPathNotFound)
		}
	})

	t.Run("missing config path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"uri": "uri", "format": "format", "path": Config{"timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.ConfigPathNotFound)
		}
	})

	t.Run("non-string uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"uri": 123, "format": "format", "path": Config{"config": "path", "timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"uri": "uri", "format": 123, "path": Config{"config": "path", "timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-string timestamp path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"uri": "uri", "format": "format", "path": Config{"config": "path", "timestamp": 123}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-string config path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableRestSourceStrategy(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"uri": "uri", "format": "format", "path": Config{"config": 123, "timestamp": "path"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("create the rest source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uri := "uri"
		format := FormatJSON
		timestampPath := "timestamp"
		configPath := "path"
		field := "field"
		value := "value"
		expected := Config{field: value}
		respData := Config{"path": Config{"field": "value"}, "timestamp": "2000-01-01T00:00:00.000Z"}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(FormatJSON, gomock.Any()).Return(decoder, nil).Times(1)

		sut, _ := NewObservableRestSourceStrategy(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.clientFactory = func() httpClient { return client }

		src, e := sut.Create(&Config{"uri": uri, "format": format, "path": Config{"config": configPath, "timestamp": timestampPath}})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *ObservableRestSource:
				if !reflect.DeepEqual(s.partial, expected) {
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
		expected := Config{field: value}
		respData := Config{"path": Config{"field": "value"}, "timestamp": "2000-01-01T00:00:00.000Z"}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(FormatJSON, gomock.Any()).Return(decoder, nil).Times(1)

		sut, _ := NewObservableRestSourceStrategy(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.clientFactory = func() httpClient { return client }

		src, e := sut.Create(&Config{"uri": uri, "path": Config{"config": configPath, "timestamp": timestampPath}})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *ObservableRestSource:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})
}
