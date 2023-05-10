package stream

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/log"
)

func Test_BaseStream_HasChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sut := &Stream{
		Formatter: NewMockFormatter(ctrl),
		Channels:  []string{"channel.1", "channel.2"},
		Level:     log.WARNING,
		Writer:    nil,
	}

	t.Run("check the channel registration", func(t *testing.T) {
		switch {
		case !sut.HasChannel("channel.1"):
			t.Error("'channel.1' channel was not found")
		case !sut.HasChannel("channel.2"):
			t.Error("'channel.2' channel was not found")
		case sut.HasChannel("channel.3"):
			t.Error("'channel.3' channel was found")
		}
	})
}

func Test_BaseStream_ListChannels(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	channels := []string{"channel.1", "channel.2"}
	sut := &Stream{
		Formatter: NewMockFormatter(ctrl),
		Channels:  channels,
		Level:     log.WARNING,
		Writer:    nil,
	}

	t.Run("list the registered Channels", func(t *testing.T) {
		if check := sut.ListChannels(); !reflect.DeepEqual(check, channels) {
			t.Errorf("returned the (%v) list of Channels", check)
		}
	})
}

func Test_BaseStream_AddChannel(t *testing.T) {
	t.Run("register a new channel", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			channel  string
			expected []string
		}{
			{ // adding into an empty list
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{},
					level:    log.DEBUG,
				},
				channel:  "channel.1",
				expected: []string{"channel.1"},
			},
			{ // adding should keep sorting
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"channel.1", "channel.3"},
					level:    log.DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.2", "channel.3"},
			},
			{ // adding an already existent should result in a no-op
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"channel.1", "channel.2", "channel.3"},
					level:    log.DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.2", "channel.3"},
			},
		}

		for _, scn := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				sut := &Stream{
					Formatter: NewMockFormatter(ctrl),
					Channels:  scn.state.channels,
					Level:     scn.state.level,
					Writer:    nil,
				}
				sut.AddChannel(scn.channel)

				if check := sut.ListChannels(); !reflect.DeepEqual(check, scn.expected) {
					t.Errorf("returned the (%v) list of Channels", check)
				}
			}
			test()
		}
	})
}

func Test_BaseStream_RemoveChannel(t *testing.T) {
	t.Run("unregister a channel", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			channel  string
			expected []string
		}{
			{ // removing from an empty list
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{},
					level:    log.DEBUG,
				},
				channel:  "channel.1",
				expected: []string{},
			},
			{ // removing a non-existing channel
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"channel.1", "channel.3"},
					level:    log.DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.3"},
			},
			{ // removing an existing channel
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"channel.1", "channel.2", "channel.3"},
					level:    log.DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.3"},
			},
		}

		for _, scn := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				sut := &Stream{
					Formatter: NewMockFormatter(ctrl),
					Channels:  scn.state.channels,
					Level:     scn.state.level,
					Writer:    nil,
				}
				sut.RemoveChannel(scn.channel)

				if check := sut.ListChannels(); !reflect.DeepEqual(check, scn.expected) {
					t.Errorf("returned the (%v) list of Channels", check)
				}
			}
			test()
		}
	})
}

func Test_BaseStream_Format(t *testing.T) {
	t.Run("return message if there is no formatter", func(t *testing.T) {
		msg := "message"
		level := log.WARNING
		sut := &Stream{Formatter: nil,
			Channels: []string{},
			Level:    level,
			Writer:   nil,
		}

		if check := sut.Format(level, msg, log.Context{"field": "value"}); check != msg {
			t.Errorf("returned the (%v) formatted message", check)
		}
	})

	t.Run("return formatter response if formatter is present without context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msg := "message"
		expected := "formatted message"
		level := log.WARNING
		formatter := NewMockFormatter(ctrl)
		formatter.EXPECT().Format(level, msg).Return(expected).Times(1)
		sut := &Stream{
			Formatter: formatter,
			Channels:  []string{},
			Level:     level,
			Writer:    nil,
		}

		if check := sut.Format(level, msg); check != expected {
			t.Errorf("returned the (%v) formatted message", check)
		}
	})

	t.Run("return formatter response if formatter is present with context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msg := "message"
		ctx := log.Context{"field": "value"}
		expected := "formatted message"
		level := log.WARNING
		formatter := NewMockFormatter(ctrl)
		formatter.EXPECT().Format(level, msg, ctx).Return(expected).Times(1)
		sut := &Stream{
			Formatter: formatter,
			Channels:  []string{},
			Level:     level,
			Writer:    nil,
		}

		if check := sut.Format(level, msg, ctx); check != expected {
			t.Errorf("returned the (%v) formatted message", check)
		}
	})
}

func Test_BaseStream_Signal(t *testing.T) {
	t.Run("signal message to the writer without context", func(t *testing.T) {
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

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message).Return(scenario.expected).Times(scenario.callTimes)

				sut := &Stream{
					Formatter: formatter,
					Channels:  scenario.state.channels,
					Level:     scenario.state.level,
					Writer:    writer,
				}

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

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.ctx).Return(scenario.expected).Times(scenario.callTimes)

				sut := &Stream{
					Formatter: formatter,
					Channels:  scenario.state.channels,
					Level:     scenario.state.level,
					Writer:    writer,
				}

				if e := sut.Signal(scenario.call.channel, scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}

func Test_BaseStream_Broadcast(t *testing.T) {
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

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message).Return(scenario.expected).Times(scenario.callTimes)

				sut := &Stream{
					Formatter: formatter,
					Channels:  scenario.state.channels,
					Level:     scenario.state.level,
					Writer:    writer,
				}

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
					ctx:     log.Context{},
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

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.ctx).Return(scenario.expected).Times(scenario.callTimes)

				sut := &Stream{
					Formatter: formatter,
					Channels:  scenario.state.channels,
					Level:     scenario.state.level,
					Writer:    writer,
				}

				if e := sut.Broadcast(scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}
