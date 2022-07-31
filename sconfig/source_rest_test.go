package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_NewSourceRest(t *testing.T) {
	t.Run("nil client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := NewSourceRest(nil, "uri", DecoderFormatUnknown, dFactory, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("nil decoder dFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := NewMockHTTPClient(ctrl)

		sut, e := NewSourceRest(client, "uri", DecoderFormatUnknown, nil, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("error while creating the request object", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`parse "\n": net/url: invalid control character in URL`)
		client := NewMockHTTPClient(ctrl)
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := NewSourceRest(client, "\n", DecoderFormatUnknown, dFactory, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error executing the http request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`test exception`)
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(nil, expected).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := NewSourceRest(client, "uri", DecoderFormatUnknown, dFactory, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("unable to get a format decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path"`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatUnknown, gomock.Any()).Return(nil, expected).Times(1)

		sut, e := NewSourceRest(client, "uri", DecoderFormatUnknown, dFactory, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid json body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`error message`)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path"`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(nil, expected).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, gomock.Any()).Return(decoder, nil).Times(1)

		sut, e := NewSourceRest(client, "uri", DecoderFormatYAML, dFactory, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("response path not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"other_path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, gomock.Any()).Return(decoder, nil).Times(1)

		sut, e := NewSourceRest(client, "uri", DecoderFormatYAML, dFactory, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConfigRestPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConfigRestPathNotFound)
		}
	})

	t.Run("response invalid path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{"path": 123}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, gomock.Any()).Return(decoder, nil).Times(1)

		sut, e := NewSourceRest(client, "uri", DecoderFormatYAML, dFactory, "path.node")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConfigRestPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConfigRestPathNotFound)
		}
	})

	t.Run("response path not pointing to a config Partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{"path": 123}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, gomock.Any()).Return(decoder, nil).Times(1)

		sut, e := NewSourceRest(client, "uri", DecoderFormatYAML, dFactory, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("correctly load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"field": "data"}
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": {"field": "data"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{"path": expected}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, gomock.Any()).Return(decoder, nil).Times(1)

		sut, e := NewSourceRest(client, "uri", DecoderFormatYAML, dFactory, "path")
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := sut.(type) {
			case *sourceRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't correctly stored the decoded Partial")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})

	t.Run("correctly load complex path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"field": "data"}
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"node": {"inner_node": {"field": "data"}}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{"node": Partial{"inner_node": expected}}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, gomock.Any()).Return(decoder, nil).Times(1)

		sut, e := NewSourceRest(client, "uri", DecoderFormatYAML, dFactory, "node..inner_node")
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := sut.(type) {
			case *sourceRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't correctly stored the decoded Partial")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})
}
