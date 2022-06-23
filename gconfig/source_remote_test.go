package gconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gerror"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_NewSourceRemote(t *testing.T) {
	t.Run("nil client", func(t *testing.T) {
		factory := &(DecoderFactory{})

		src, err := NewSourceRemote(nil, "uri", DecoderFormatUnknown, factory, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := NewMockHTTPClient(ctrl)

		src, err := NewSourceRemote(client, "uri", DecoderFormatUnknown, nil, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("error while creating the request object", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`parse "\n": net/url: invalid control character in URL`)
		client := NewMockHTTPClient(ctrl)
		factory := &(DecoderFactory{})

		src, err := NewSourceRemote(client, "\n", DecoderFormatUnknown, factory, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error executing the http request", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`test exception`)
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(nil, expected).Times(1)
		factory := &(DecoderFactory{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatUnknown, factory, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("unable to get a format decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path"`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatUnknown, factory, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrInvalidConfigDecoderFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrInvalidConfigDecoderFormat)
		}
	})

	t.Run("invalid json body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`yaml: line 1: did not find expected ',' or '}'`)
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path"`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatYAML, factory, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("response path not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"other_path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatYAML, factory, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConfigRemotePathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConfigRemotePathNotFound)
		}
	})

	t.Run("response invalid path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatYAML, factory, "path.node")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConfigRemotePathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConfigRemotePathNotFound)
		}
	})

	t.Run("response path not pointing to a config Partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatYAML, factory, "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
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
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatYAML, factory, "path")
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error : %v", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceRemote:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't correctly stored the decoded Partial")
				}
			default:
				t.Error("didn't returned a new remote src")
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
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceRemote(client, "uri", DecoderFormatYAML, factory, "node..inner_node")
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error : %v", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceRemote:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't correctly stored the decoded Partial")
				}
			default:
				t.Error("didn't returned a new remote src")
			}
		}
	})
}
