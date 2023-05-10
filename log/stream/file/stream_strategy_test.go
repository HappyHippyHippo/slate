package file

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/log/formatter/json"
)

func Test_NewStreamStrategyFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewStreamStrategy(nil, NewMockFormatterFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil formatter factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewStreamStrategy(NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new file stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl)); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}

func Test_StreamStrategyFile_Accept(t *testing.T) {
	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				return nil, fmt.Errorf("dummy error")
			},
		).Times(1)

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if sut.Accept(config) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				data.Type = log.UnknownStream
				return data, nil
			},
		).Times(1)

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if sut.Accept(config) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				data.Type = Type
				return data, nil
			},
		).Times(1)

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if !sut.Accept(config) {
			t.Error("returned false")
		}
	})
}

func Test_StreamStrategyFile_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

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

	t.Run("error on config parsing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy error")
		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&streamConfig{})).DoAndReturn(
			func(_ string, data *streamConfig, _ ...bool) (interface{}, error) {
				return nil, expected
			},
		).Times(1)

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		src, e := sut.Create(config)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("non-log level name level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&streamConfig{})).DoAndReturn(
			func(_ string, data *streamConfig, _ ...bool) (interface{}, error) {
				data.Path = "path"
				data.Format = json.Format
				data.Level = "invalid"
				return data, nil
			},
		).Times(1)

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, log.ErrInvalidLevel):
			t.Errorf("returned the (%v) error when expecting (%v)", e, log.ErrInvalidLevel)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&streamConfig{})).DoAndReturn(
			func(_ string, data *streamConfig, _ ...bool) (interface{}, error) {
				data.Path = "path"
				data.Format = log.UnknownFormatterFormat
				data.Level = "fatal"
				return data, nil
			},
		).Times(1)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(log.UnknownFormatterFormat).Return(nil, expected).Times(1)

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), formatterFactory)

		stream, e := sut.Create(config)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error on opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&streamConfig{})).DoAndReturn(
			func(_ string, data *streamConfig, _ ...bool) (interface{}, error) {
				data.Path = "path"
				data.Format = json.Format
				data.Level = "fatal"
				return data, nil
			},
		).Times(1)

		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		formatter := NewMockFormatter(ctrl)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(json.Format).Return(formatter, nil).Times(1)

		sut, _ := NewStreamStrategy(fileSystem, formatterFactory)

		stream, e := sut.Create(config)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&streamConfig{})).DoAndReturn(
			func(_ string, data *streamConfig, _ ...bool) (interface{}, error) {
				data.Path = "path"
				data.Format = json.Format
				data.Level = "fatal"
				data.Channels = []interface{}{"channel1", "channel2"}
				return data, nil
			},
		).Times(1)

		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(nil).Times(1)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		formatter := NewMockFormatter(ctrl)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(json.Format).Return(formatter, nil).Times(1)

		sut, _ := NewStreamStrategy(fileSystem, formatterFactory)

		stream, e := sut.Create(config)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			_ = stream.(io.Closer).Close()
			switch s := stream.(type) {
			case *Stream:
				switch {
				case s.Level != log.FATAL:
					t.Error("didn't created a stream with the correct level")
				case len(s.Channels) != 2:
					t.Error("didn't created a stream with the correct channel list")
				}
			default:
				t.Error("didn't returned a new file stream")
			}
		}
	})
}
