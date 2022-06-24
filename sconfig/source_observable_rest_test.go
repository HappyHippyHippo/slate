package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_NewSourceRestObservable(t *testing.T) {
	t.Run("nil client", func(t *testing.T) {
		factory := &(DecoderFactory{})

		src, err := NewSourceObservableRest(nil, "uri", "format", factory, "timestampPath", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := NewMockHTTPClient(ctrl)

		src, err := NewSourceObservableRest(client, "uri", "format", nil, "timestampPath", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error while creating the request object", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`parse "\n": net/url: invalid control character in URL`)
		client := NewMockHTTPClient(ctrl)
		factory := &(DecoderFactory{})

		src, err := NewSourceObservableRest(client, "\n", "format", factory, "timestampPath", "configPath")
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

		src, err := NewSourceObservableRest(client, "uri", "format", factory, "timestampPath", "configPath")
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

		src, err := NewSourceObservableRest(client, "uri", "format", factory, "timestampPath", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidConfigDecoderFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidConfigDecoderFormat)
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

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestampPath", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("response timestamp path not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"other_path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestampPath", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigRestPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigRestPathNotFound)
		}
	})

	t.Run("invalid timestamp value type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("invalid timestamp value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := "parsing time \"abc\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"abc\" as \"2006\""
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "abc"}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected:
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("response config path not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", other_path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "configPath")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigRestPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigRestPathNotFound)
		}
	})

	t.Run("response invalid path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "path.node")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigRestPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigRestPathNotFound)
		}
	})

	t.Run("response path not pointing to a config Partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("correctly load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"field": "data"}
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "path": {"field": "data"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "path")
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error : %v", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableRest:
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
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "node": {"inner_node": {"field": "data"}}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, err := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "node..inner_node")
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error : %v", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't correctly stored the decoded Partial")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})
}

func Test_SourceRestObservable_Reload(t *testing.T) {
	t.Run("dont reload on same timestamp", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"field": "data 1"}
		response1 := http.Response{}
		response1.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 1"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		response2 := http.Response{}
		response2.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 2"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		gomock.InOrder(
			client.EXPECT().Do(gomock.Any()).Return(&response1, nil),
			client.EXPECT().Do(gomock.Any()).Return(&response2, nil),
		)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, _ := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "node")

		loaded, err := src.Reload()
		switch {
		case loaded != false:
			t.Error("unexpectedly reload the source config")
		case err != nil:
			t.Errorf("returned the eunexpected error : %v", err)
		default:
			switch s := src.(type) {
			case *sourceObservableRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't correctly stored the decoded Partial")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})

	t.Run("correctly reload config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := Partial{"field": "data 2"}
		response1 := http.Response{}
		response1.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 1"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		response2 := http.Response{}
		response2.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 2"}, "timestamp": "2021-12-15T21:07:48.240Z"}`))
		client := NewMockHTTPClient(ctrl)
		gomock.InOrder(
			client.EXPECT().Do(gomock.Any()).Return(&response1, nil),
			client.EXPECT().Do(gomock.Any()).Return(&response2, nil),
		)
		factory := &(DecoderFactory{})
		_ = factory.Register(&DecoderStrategyYAML{})

		src, _ := NewSourceObservableRest(client, "uri", "yaml", factory, "timestamp", "node")

		loaded, err := src.Reload()
		switch {
		case loaded != true:
			t.Error("didn't reload the source config")
		case err != nil:
			t.Errorf("returned the eunexpected error : %v", err)
		default:
			switch s := src.(type) {
			case *sourceObservableRest:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't correctly stored the decoded Partial")
				}
			default:
				t.Error("didn't returned a new rest src")
			}
		}
	})
}
