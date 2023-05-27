package rest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewObsSource(t *testing.T) {
	t.Run("nil client", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoderFactory := config.NewDecoderFactory()

		sut, e := NewObsSource(nil, "uri", "format", decoderFactory, "timestampPath", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := NewMockHTTPClient(ctrl)

		sut, e := NewObsSource(client, "uri", "format", nil, "timestampPath", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error while creating the request object", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf(`parse "\n": net/url: invalid control character in URL`)
		client := NewMockHTTPClient(ctrl)
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewObsSource(client, "\n", "format", decoderFactory, "timestampPath", "configPath")
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
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewObsSource(client, "uri", "format", decoderFactory, "timestampPath", "configPath")
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

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"path"`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewObsSource(client, "uri", "format", decoderFactory, "timestampPath", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, config.ErrInvalidFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", e, config.ErrInvalidFormat)
		}
	})

	t.Run("error decoding body", func(t *testing.T) {
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
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestampPath", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("response timestamp path not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"other_path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestampPath", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrTimestampNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrTimestampNotFound)
		}
	})

	t.Run("invalid timestamp value type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{"timestamp": 123}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
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
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{"timestamp": "abc"}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected:
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("response config path not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", other_path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z"}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "configPath")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrConfigNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrConfigNotFound)
		}
	})

	t.Run("response invalid path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z", "path": 123}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "path.node")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrConfigNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrConfigNotFound)
		}
	})

	t.Run("response path not pointing to a config Partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "path": 123}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z", "path": 123}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("correctly load", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := config.Partial{"field": "data"}
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "path": {"field": "data"}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z", "path": expected}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "path")
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case !reflect.DeepEqual(sut.Partial, expected):
			t.Error("didn't correctly stored the decoded Partial")
		}
	})

	t.Run("correctly load complex path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := config.Partial{"field": "data"}
		response := http.Response{}
		response.Body = io.NopCloser(strings.NewReader(`{"timestamp": "2000-01-01T00:00:00Z", "node": {"inner_node": {"field": "data"}}}`))
		client := NewMockHTTPClient(ctrl)
		client.EXPECT().Do(gomock.Any()).Return(&response, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z", "node": config.Partial{"inner_node": expected}}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("yaml").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "node..inner_node")
		switch {
		case e != nil:
			t.Errorf("returned the unexpected e : %v", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case !reflect.DeepEqual(sut.Partial, expected):
			t.Error("didn't correctly stored the decoded Partial")
		}
	})
}

func Test_ObsSource_Reload(t *testing.T) {
	t.Run("dont reload on same timestamp", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := config.Partial{"field": "data 1"}
		response1 := http.Response{}
		response1.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 1"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		response2 := http.Response{}
		response2.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 2"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		client := NewMockHTTPClient(ctrl)
		gomock.InOrder(
			client.EXPECT().Do(gomock.Any()).Return(&response1, nil),
			client.EXPECT().Do(gomock.Any()).Return(&response2, nil),
		)
		decoder1 := NewMockDecoder(ctrl)
		decoder1.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z", "node": expected}, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		decoder2 := NewMockDecoder(ctrl)
		decoder2.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z", "node": config.Partial{"field": "data 2"}}, nil).Times(1)
		decoder2.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		gomock.InOrder(
			decoderStrategy.EXPECT().Accept("yaml").Return(true),
			decoderStrategy.EXPECT().Accept("yaml").Return(true),
		)
		gomock.InOrder(
			decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder1, nil),
			decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder2, nil),
		)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, _ := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "node")

		loaded, e := sut.Reload()
		switch {
		case loaded != false:
			t.Error("unexpectedly reload the source config")
		case e != nil:
			t.Errorf("returned the eunexpected e : %v", e)
		case !reflect.DeepEqual(sut.Partial, expected):
			t.Error("didn't correctly stored the decoded Partial")
		}
	})

	t.Run("correctly reload config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := config.Partial{"field": "data 2"}
		response1 := http.Response{}
		response1.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 1"}, "timestamp": "2021-12-15T21:07:48.239Z"}`))
		response2 := http.Response{}
		response2.Body = io.NopCloser(strings.NewReader(`{"node": {"field": "data 2"}, "timestamp": "2021-12-15T21:07:48.240Z"}`))
		client := NewMockHTTPClient(ctrl)
		gomock.InOrder(
			client.EXPECT().Do(gomock.Any()).Return(&response1, nil),
			client.EXPECT().Do(gomock.Any()).Return(&response2, nil),
		)
		decoder1 := NewMockDecoder(ctrl)
		decoder1.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:00Z", "node": config.Partial{"field": "data 1"}}, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		decoder2 := NewMockDecoder(ctrl)
		decoder2.EXPECT().Decode().Return(&config.Partial{"timestamp": "2000-01-01T00:00:01Z", "node": expected}, nil).Times(1)
		decoder2.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		gomock.InOrder(
			decoderStrategy.EXPECT().Accept("yaml").Return(true),
			decoderStrategy.EXPECT().Accept("yaml").Return(true),
		)
		gomock.InOrder(
			decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder1, nil),
			decoderStrategy.EXPECT().Create(gomock.Any()).Return(decoder2, nil),
		)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, _ := NewObsSource(client, "uri", "yaml", decoderFactory, "timestamp", "node")

		loaded, e := sut.Reload()
		switch {
		case loaded != true:
			t.Error("didn't reload the source config")
		case e != nil:
			t.Errorf("returned the eunexpected e : %v", e)
		case !reflect.DeepEqual(sut.Partial, expected):
			t.Error("didn't correctly stored the decoded Partial")
		}
	})
}
