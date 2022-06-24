package sconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_NewSourceStrategyRestObservable(t *testing.T) {
	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, err := NewSourceStrategyObservableRest(nil)
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new observable rest source factory strategy", func(t *testing.T) {
		factory := &(DecoderFactory{})

		strategy, err := NewSourceStrategyObservableRest(factory)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		case strategy.(*sourceStrategyObservableRest).decoderFactory != factory:
			t.Error("didn't stored the decoder factory reference")
		default:
			client := strategy.(*sourceStrategyObservableRest).clientFactory()
			switch client.(type) {
			case *http.Client:
			default:
				t.Error("didn't stored a valid http client factory")
			}
		}
	})
}

func Test_SourceStrategyRestObservable_Accept(t *testing.T) {
	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			exp        bool
		}{
			{ // _test observable rest type
				sourceType: SourceTypeObservableRest,
				exp:        true,
			},
			{ // _test non-observable rest type
				sourceType: SourceTypeUnknown,
				exp:        false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))
				if check := strategy.Accept(scenario.sourceType); check != scenario.exp {
					t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
				}
			}
			test()
		}
	})
}

func Test_SourceStrategyRestObservable_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		if !strategy.AcceptFromConfig(&Partial{"type": SourceTypeObservableRest}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyRestObservable_Create(t *testing.T) {
	t.Run("missing uri", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing format", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing timestamp path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri", "format")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing config path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri", "format", "timestamp path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("non-string uri", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create(123, "format", "config path", "timestamp path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri", 123, "config path", "timestamp path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string timestamp path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri", "format", 123, "config path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string config path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri", "format", "timestamp path", 123)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("create the rest source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uri := "uri"
		format := "yaml"
		timestampPath := "timestamp_path"
		configPath := "config_path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		factory := &(DecoderFactory{})
		_ = factory.Register(&decoderStrategyYAML{})
		strategy, _ := NewSourceStrategyObservableRest(factory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"config_path": {"field": "value"}, "timestamp_path": "2000-01-01T00:00:00.000Z"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		strategy.(*sourceStrategyObservableRest).clientFactory = func() HTTPClient {
			return client
		}

		src, err := strategy.Create(uri, format, timestampPath, configPath)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})
}

func Test_SourceStrategyRestObservable_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing uri", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"format": "format", "timestampPath": "path", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("missing timestamp path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "uri", "format": "format", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("missing config path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "uri", "format": "format", "timestampPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("non-string uri", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": 123, "format": "format", "timestampPath": "path", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "uri", "format": 123, "timestampPath": "path", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string timestamp path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "uri", "format": "format", "timestampPath": 123, "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string config path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyObservableRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "uri", "format": "format", "timestampPath": "path", "configPath": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("create the rest source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uri := "uri"
		format := DecoderFormatJSON
		timestampPath := "timestamp"
		configPath := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		factory := &(DecoderFactory{})
		_ = factory.Register(&decoderStrategyJSON{})
		strategy, _ := NewSourceStrategyObservableRest(factory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		strategy.(*sourceStrategyObservableRest).clientFactory = func() HTTPClient { return client }

		src, err := strategy.CreateFromConfig(&Partial{"uri": uri, "format": format, "timestampPath": timestampPath, "configPath": configPath})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableRest:
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
		expected := Partial{field: value}
		factory := &(DecoderFactory{})
		_ = factory.Register(&decoderStrategyJSON{})
		strategy, _ := NewSourceStrategyObservableRest(factory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		strategy.(*sourceStrategyObservableRest).clientFactory = func() HTTPClient { return client }

		src, err := strategy.CreateFromConfig(&Partial{"uri": uri, "timestampPath": timestampPath, "configPath": configPath})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})
}
