package file

import (
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewStream(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	writer := NewMockWriter(ctrl)
	writer.EXPECT().Close().Times(1)
	formatter := NewMockFormatter(ctrl)
	var channels []string
	level := log.WARNING

	t.Run("nil writer", func(t *testing.T) {
		sut, e := NewStream(nil, formatter, channels, level)
		switch {
		case sut != nil:
			_ = sut.(io.Closer).Close()
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil formatter", func(t *testing.T) {
		sut, e := NewStream(writer, nil, channels, level)
		switch {
		case sut != nil:
			_ = sut.(io.Closer).Close()
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new file stream", func(t *testing.T) {
		sut, e := NewStream(writer, formatter, []string{}, log.WARNING)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		default:
			_ = sut.(io.Closer).Close()
			if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		}
	})
}

func Test_Stream_Close(t *testing.T) {
	t.Run("call the close on the writer only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		writer := NewMockWriter(ctrl)
		writer.EXPECT().Close().Times(1)
		sut, _ := NewStream(writer, NewMockFormatter(ctrl), []string{}, log.WARNING)

		_ = sut.(io.Closer).Close()
		_ = sut.(io.Closer).Close()
	})
}

func Test_Stream_Signal(t *testing.T) {
	t.Run("signal message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			call struct {
				level   log.Level
				channel string
				message string
			}
			callTimes int
			expected  string
		}{
			{ // signal through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					message string
				}{
					level:   log.FATAL,
					channel: "dummy_channel",
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					message string
				}{
					level:   log.DEBUG,
					channel: "dummy_channel",
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with an unregistered channel
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					message string
				}{
					level:   log.FATAL,
					channel: "not_a_valid_dummy_channel",
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, s := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Close().Times(1)
				writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.
					EXPECT().
					Format(s.call.level, s.call.message).
					Return(s.expected).
					Times(s.callTimes)

				sut, _ := NewStream(writer, formatter, s.state.channels, s.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Signal(s.call.channel, s.call.level, s.call.message); e != nil {
					t.Errorf("unexpected (%v) error", e)
				}
			}
			test()
		}
	})

	t.Run("signal message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			call struct {
				level   log.Level
				channel string
				ctx     log.Context
				message string
			}
			callTimes int
			expected  string
		}{
			{ // signal through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					ctx     log.Context
					message string
				}{
					level:   log.FATAL,
					channel: "dummy_channel",
					ctx:     log.Context{},
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					ctx     log.Context
					message string
				}{
					level:   log.DEBUG,
					channel: "dummy_channel",
					ctx:     log.Context{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with an unregistered channel
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					ctx     log.Context
					message string
				}{
					level:   log.FATAL,
					channel: "not_a_valid_dummy_channel",
					ctx:     log.Context{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, s := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Close().Times(1)
				writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.
					EXPECT().
					Format(s.call.level, s.call.message, s.call.ctx).
					Return(s.expected).
					Times(s.callTimes)

				sut, _ := NewStream(writer, formatter, s.state.channels, s.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Signal(s.call.channel, s.call.level, s.call.message, s.call.ctx); e != nil {
					t.Errorf("unexpected (%v) error", e)
				}
			}
			test()
		}
	})
}

func Test_Stream_Broadcast(t *testing.T) {
	t.Run("broadcast message to the writer without context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			call struct {
				level   log.Level
				message string
			}
			callTimes int
			expected  string
		}{
			{ // broadcast through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					message string
				}{
					level:   log.FATAL,
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // broadcast through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					message string
				}{
					level:   log.DEBUG,
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, s := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Close().Times(1)
				writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.
					EXPECT().
					Format(s.call.level, s.call.message).
					Return(s.expected).
					Times(s.callTimes)

				sut, _ := NewStream(writer, formatter, s.state.channels, s.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Broadcast(s.call.level, s.call.message); e != nil {
					t.Errorf("unexpected (%v) error", e)
				}
			}
			test()
		}
	})

	t.Run("broadcast message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			call struct {
				level   log.Level
				ctx     log.Context
				message string
			}
			callTimes int
			expected  string
		}{
			{ // broadcast through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					ctx     log.Context
					message string
				}{
					level:   log.FATAL,
					ctx:     log.Context{},
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // broadcast through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					ctx     log.Context
					message string
				}{
					level:   log.DEBUG,
					ctx:     log.Context{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, s := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Close().Times(1)
				writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.
					EXPECT().
					Format(s.call.level, s.call.message, s.call.ctx).
					Return(s.expected).
					Times(s.callTimes)

				sut, _ := NewStream(writer, formatter, s.state.channels, s.state.level)
				defer func() {
					_ = sut.(io.Closer).Close()
					ctrl.Finish()
				}()

				if e := sut.Broadcast(s.call.level, s.call.message, s.call.ctx); e != nil {
					t.Errorf("unexpected (%v) error", e)
				}
			}
			test()
		}
	})
}
