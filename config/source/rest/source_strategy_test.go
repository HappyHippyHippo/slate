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

func Test_NewSourceStrategy(t *testing.T) {
	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewSourceStrategy(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new rest source factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoderFactory := config.NewDecoderFactory()

		sut, e := NewSourceStrategy(decoderFactory)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.decoderFactory != decoderFactory:
			t.Error("didn't stored the decoder factory reference")
		default:
			client := sut.clientFactory()
			switch client.(type) {
			case *http.Client:
			default:
				t.Error("didn't stored a valid http client")
			}
		}
	})
}

func Test_SourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		if sut.Accept(config.Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		if sut.Accept(config.Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		if sut.Accept(config.Partial{"type": config.UnknownSource}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		if !sut.Accept(config.Partial{"type": Type}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		src, e := sut.Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("missing uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		src, e := sut.Create(config.Partial{
			"format":     "format",
			"configPath": "path",
		})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, config.ErrInvalidSource):
			t.Errorf("(%v) when expecting (%v)", e, config.ErrInvalidSource)
		}
	})

	t.Run("missing config path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		src, e := sut.Create(config.Partial{
			"uri":    "path",
			"format": "format",
		})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, config.ErrInvalidSource):
			t.Errorf("(%v) when expecting (%v)", e, config.ErrInvalidSource)
		}
	})

	t.Run("non-string uri", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		src, e := sut.Create(config.Partial{
			"uri":        123,
			"format":     "format",
			"configPath": "path",
		})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		src, e := sut.Create(config.Partial{
			"uri":        "uri",
			"format":     123,
			"configPath": "path",
		})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewSourceStrategy(config.NewDecoderFactory())

		src, e := sut.Create(config.Partial{
			"uri":    "uri",
			"format": "format",
			"path": config.Partial{
				"config": 123,
			},
		})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("create the rest source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uri := "uri"
		format := "format"
		path := "path"
		field := "field"
		value := "value"
		expected := config.Partial{field: value}
		respData := config.Partial{"path": config.Partial{"field": "value"}}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, _ := NewSourceStrategy(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockRequester(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.clientFactory = func() requester { return client }

		src, e := sut.Create(config.Partial{
			"uri":    uri,
			"format": format,
			"path": config.Partial{
				"config": path,
			},
		})
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *Source:
				if !reflect.DeepEqual(s.Partial, expected) {
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
		expected := config.Partial{field: value}
		respData := config.Partial{"path": config.Partial{"field": "value"}}
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&respData, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("json").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, _ := NewSourceStrategy(decoderFactory)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "value"}}`))
		client := NewMockRequester(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		sut.clientFactory = func() requester { return client }

		src, e := sut.Create(config.Partial{
			"uri": uri,
			"path": config.Partial{
				"config": path,
			},
		})
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *Source:
				if !reflect.DeepEqual(s.Partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new rest source")
			}
		}
	})
}
