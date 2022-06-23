package glog

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gerror"
	"io"
	"testing"
)

func Test_NewStreamFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	writer := NewMockWriter(ctrl)
	writer.EXPECT().Close().Times(1)
	formatter := NewMockFormatter(ctrl)
	var channels []string
	level := WARNING

	t.Run("nil writer", func(t *testing.T) {
		stream, err := NewStreamFile(nil, formatter, channels, level)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("nil formatter", func(t *testing.T) {
		stream, err := NewStreamFile(writer, nil, channels, level)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("new file stream", func(t *testing.T) {
		stream, err := NewStreamFile(writer, formatter, []string{}, WARNING)
		switch {
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			_ = stream.(io.Closer).Close()
			if err != nil {
				t.Errorf("returned the (%v) error", err)
			}
		}
	})
}

func Test_StreamFile_Close(t *testing.T) {
	t.Run("call the close on the writer only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		writer := NewMockWriter(ctrl)
		writer.EXPECT().Close().Times(1)
		stream, _ := NewStreamFile(writer, NewMockFormatter(ctrl), []string{}, WARNING)

		_ = stream.(io.Closer).Close()
		_ = stream.(io.Closer).Close()
	})
}

func Test_StreamFile_Signal(t *testing.T) {
	t.Run("signal message to the writer", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				channel string
				fields  map[string]interface{}
				message string
			}
			callTimes int
			expected  string
		}{
			{ // signal through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					channel string
					fields  map[string]interface{}
					message string
				}{
					level:   FATAL,
					channel: "dummy_channel",
					fields:  map[string]interface{}{},
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with a filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					channel string
					fields  map[string]interface{}
					message string
				}{
					level:   DEBUG,
					channel: "dummy_channel",
					fields:  map[string]interface{}{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with an unregistered channel
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					channel string
					fields  map[string]interface{}
					message string
				}{
					level:   FATAL,
					channel: "not_a_valid_dummy_channel",
					fields:  map[string]interface{}{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Close().Times(1)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.fields).Return(scenario.expected).Times(scenario.callTimes)
				stream, _ := NewStreamFile(writer, formatter, scenario.state.channels, scenario.state.level)

				defer func() { _ = stream.(io.Closer).Close(); ctrl.Finish() }()

				if err := stream.Signal(scenario.call.channel, scenario.call.level, scenario.call.message, scenario.call.fields); err != nil {
					t.Errorf("returned the (%v) error", err)
				}
			}
			test()
		}
	})
}

func Test_StreamFile_Broadcast(t *testing.T) {
	t.Run("broadcast message to the writer", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				fields  map[string]interface{}
				message string
			}
			callTimes int
			expected  string
		}{
			{ // broadcast through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					fields  map[string]interface{}
					message string
				}{
					level:   FATAL,
					fields:  map[string]interface{}{},
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // broadcast through a valid channel with a filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					fields  map[string]interface{}
					message string
				}{
					level:   DEBUG,
					fields:  map[string]interface{}{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Close().Times(1)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.fields).Return(scenario.expected).Times(scenario.callTimes)
				stream, _ := NewStreamFile(writer, formatter, scenario.state.channels, scenario.state.level)

				defer func() { _ = stream.(io.Closer).Close(); ctrl.Finish() }()

				if err := stream.Broadcast(scenario.call.level, scenario.call.message, scenario.call.fields); err != nil {
					t.Errorf("returned the (%v) error", err)
				}
			}
			test()
		}
	})
}
