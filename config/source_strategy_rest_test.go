package config

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewSourceStrategyRest(t *testing.T) {
	t.Run("nil decoder decoderFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceStrategyRest(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("new rest source decoderFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoderFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceStrategyRest(decoderFactory)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.(*sourceStrategyRest).decoderFactory != decoderFactory:
			t.Error("didn't stored the decoder decoderFactory reference")
		default:
			client := sut.(*sourceStrategyRest).clientFactory()
			switch client.(type) {
			case *http.Client:
			default:
				t.Error("didn't stored a valid http client decoderFactory")
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
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))
				if check := sut.Accept(scenario.sourceType); check != scenario.exp {
					t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
				}
			}
			test()
		}
	})
}

func Test_SourceStrategyRest_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		if !sut.AcceptFromConfig(&Partial{"type": SourceTypeRest}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyRest_Create(t *testing.T) {
	t.Run("missing uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("missing format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.Create("uri")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.Create("uri", "format")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("non-string uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.Create(123, "format", "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.Create("uri", 123, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.Create("uri", "format", 123)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
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
		respData := Partial{"path": Partial{"field": "value"}}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, gomock.Any()).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyRest(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.(*sourceStrategyRest).clientFactory = func() HTTPClient {
			return client
		}

		src, e := sut.Create(uri, format, path)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("missing uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"format": "format", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"uri": "path", "format": "format"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("non-string uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"uri": 123, "format": "format", "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"uri": "uri", "format": 123, "configPath": "path"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyRest(NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"uri": "uri", "format": "format", "configPath": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
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
		respData := Partial{"path": Partial{"field": "value"}}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatJSON, gomock.Any()).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyRest(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.(*sourceStrategyRest).clientFactory = func() HTTPClient {
			return client
		}

		src, e := sut.CreateFromConfig(&Partial{"uri": uri, "format": format, "configPath": path})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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
		respData := Partial{"path": Partial{"field": "value"}}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatJSON, gomock.Any()).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyRest(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.(*sourceStrategyRest).clientFactory = func() HTTPClient {
			return client
		}

		src, e := sut.CreateFromConfig(&Partial{"uri": uri, "configPath": path})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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
