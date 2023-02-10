package log

import (
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_NewFileStream(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	writer := NewMockWriter(ctrl)
	writer.EXPECT().Close().Times(1)
	formatter := NewMockFormatter(ctrl)
	var channels []string
	level := WARNING

	t.Run("nil writer", func(t *testing.T) {
		sut, e := NewFileStream(nil, formatter, channels, level)
		switch {
		case sut != nil:
			_ = sut.(io.Closer).Close()
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil formatter", func(t *testing.T) {
		sut, e := NewFileStream(writer, nil, channels, level)
		switch {
		case sut != nil:
			_ = sut.(io.Closer).Close()
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new file stream", func(t *testing.T) {
		sut, e := NewFileStream(writer, formatter, []string{}, WARNING)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		default:
			_ = sut.(io.Closer).Close()
			if e != nil {
				t.Errorf("returned the (%v) error", e)
			}
		}
	})
}

func Test_FileStream_Close(t *testing.T) {
	t.Run("call the close on the writer only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		writer := NewMockWriter(ctrl)
		writer.EXPECT().Close().Times(1)
		sut, _ := NewFileStream(writer, NewMockFormatter(ctrl), []string{}, WARNING)

		_ = sut.(io.Closer).Close()
		_ = sut.(io.Closer).Close()
	})
}

func Test_FileStream_Signal(t *testing.T) {
	t.Run("signal message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				channel string
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
					message string
				}{
					level:   FATAL,
					channel: "dummy_channel",
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
					message string
				}{
					level:   DEBUG,
					channel: "dummy_channel",
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
					message string
				}{
					level:   FATAL,
					channel: "not_a_valid_dummy_channel",
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
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewFileStream(writer, formatter, scenario.state.channels, scenario.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Signal(scenario.call.channel, scenario.call.level, scenario.call.message); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})

	t.Run("signal message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				channel string
				ctx     Context
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
					ctx     Context
					message string
				}{
					level:   FATAL,
					channel: "dummy_channel",
					ctx:     Context{},
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
					ctx     Context
					message string
				}{
					level:   DEBUG,
					channel: "dummy_channel",
					ctx:     Context{},
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
					ctx     Context
					message string
				}{
					level:   FATAL,
					channel: "not_a_valid_dummy_channel",
					ctx:     Context{},
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
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.ctx).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewFileStream(writer, formatter, scenario.state.channels, scenario.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Signal(scenario.call.channel, scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}

func Test_FileStream_Broadcast(t *testing.T) {
	t.Run("broadcast message to the writer without context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
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
					message string
				}{
					level:   FATAL,
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
					message string
				}{
					level:   DEBUG,
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
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewFileStream(writer, formatter, scenario.state.channels, scenario.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Broadcast(scenario.call.level, scenario.call.message); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})

	t.Run("broadcast message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				ctx     Context
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
					ctx     Context
					message string
				}{
					level:   FATAL,
					ctx:     Context{},
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
					ctx     Context
					message string
				}{
					level:   DEBUG,
					ctx:     Context{},
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
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.ctx).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewFileStream(writer, formatter, scenario.state.channels, scenario.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Broadcast(scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}
