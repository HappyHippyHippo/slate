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

func Test_NewSourceStrategyRest(t *testing.T) {
	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, err := NewSourceStrategyRest(nil)
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new rest source factory strategy", func(t *testing.T) {
		factory := &(DecoderFactory{})

		strategy, err := NewSourceStrategyRest(factory)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		case strategy.(*sourceStrategyRest).decoderFactory != factory:
			t.Error("didn't stored the decoder factory reference")
		default:
			client := strategy.(*sourceStrategyRest).clientFactory()
			switch client.(type) {
			case *http.Client:
			default:
				t.Error("didn't stored a valid http client factory")
			}
		}
	})
}

func Test_SourceStrategyRest_Accept(t *testing.T) {
	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			exp        bool
		}{
			{ // _test rest type
				sourceType: SourceTypeRest,
				exp:        true,
			},
			{ // _test non-rest type
				sourceType: SourceTypeUnknown,
				exp:        false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))
				if check := strategy.Accept(scenario.sourceType); check != scenario.exp {
					t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
				}
			}
			test()
		}
	})
}

func Test_SourceStrategyRest_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		if !strategy.AcceptFromConfig(&Partial{"type": SourceTypeRest}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyRest_Create(t *testing.T) {
	t.Run("missing uri", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

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
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

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

	t.Run("missing path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

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

	t.Run("non-string uri", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.Create(123, "format", "path")
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
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri", 123, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.Create("uri", "format", 123)
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
		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		factory := &(DecoderFactory{})
		_ = factory.Register(&decoderStrategyYAML{})
		strategy, _ := NewSourceStrategyRest(factory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		strategy.(*sourceStrategyRest).clientFactory = func() HTTPClient {
			return client
		}

		src, err := strategy.Create(uri, format, path)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest source")
			}
		}
	})
}

func Test_SourceStrategyRest_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

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
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"format": "format", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "path", "format": "format"})
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
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": 123, "format": "format", "configPath": "path"})
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
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "uri", "format": 123, "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		strategy, _ := NewSourceStrategyRest(&(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"uri": "uri", "format": "format", "configPath": 123})
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
		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		factory := &(DecoderFactory{})
		_ = factory.Register(&decoderStrategyJSON{})
		strategy, _ := NewSourceStrategyRest(factory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		strategy.(*sourceStrategyRest).clientFactory = func() HTTPClient {
			return client
		}

		src, err := strategy.CreateFromConfig(&Partial{"uri": uri, "format": format, "configPath": path})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest source")
			}
		}
	})

	t.Run("create the rest source defaulting format if not present in config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uri := "uri"
		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		factory := &(DecoderFactory{})
		_ = factory.Register(&decoderStrategyJSON{})
		strategy, _ := NewSourceStrategyRest(factory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		strategy.(*sourceStrategyRest).clientFactory = func() HTTPClient {
			return client
		}

		src, err := strategy.CreateFromConfig(&Partial{"uri": uri, "configPath": path})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest source")
			}
		}
	})
}