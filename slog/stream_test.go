package slog

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_Stream_Level(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	level := WARNING
	stream := &stream{NewMockFormatter(ctrl), []string{}, level}

	t.Run("retrieve the filtering level", func(t *testing.T) {
		if check := stream.Level(); check != level {
			t.Errorf("returned the (%v) level", check)
		}
	})
}

func Test_Stream_HasChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := &stream{NewMockFormatter(ctrl), []string{"channel.1", "channel.2"}, WARNING}

	t.Run("check the channel registration", func(t *testing.T) {
		switch {
		case !stream.HasChannel("channel.1"):
			t.Error("'channel.1' channel was not found")
		case !stream.HasChannel("channel.2"):
			t.Error("'channel.2' channel was not found")
		case stream.HasChannel("channel.3"):
			t.Error("'channel.3' channel was found")
		}
	})
}

func Test_Stream_ListChannels(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	channels := []string{"channel.1", "channel.2"}
	stream := &stream{NewMockFormatter(ctrl), channels, WARNING}

	t.Run("list the registered channels", func(t *testing.T) {
		if check := stream.ListChannels(); !reflect.DeepEqual(check, channels) {
			t.Errorf("returned the (%v) list of channels", check)
		}
	})
}

func Test_Stream_AddChannel(t *testing.T) {
	t.Run("register a new channel", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			channel  string
			expected []string
		}{
			{ // adding into an empty list
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{},
					level:    DEBUG,
				},
				channel:  "channel.1",
				expected: []string{"channel.1"},
			},
			{ // adding should keep sorting
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"channel.1", "channel.3"},
					level:    DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.2", "channel.3"},
			},
			{ // adding an already existent should result in a no-op
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"channel.1", "channel.2", "channel.3"},
					level:    DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.2", "channel.3"},
			},
		}

		for _, scn := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				stream := &stream{NewMockFormatter(ctrl), scn.state.channels, scn.state.level}
				stream.AddChannel(scn.channel)

				if check := stream.ListChannels(); !reflect.DeepEqual(check, scn.expected) {
					t.Errorf("returned the (%v) list of channels", check)
				}
			}
			test()
		}
	})
}

func Test_Stream_RemoveChannel(t *testing.T) {
	t.Run("unregister a channel", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			channel  string
			expected []string
		}{
			{ // removing from an empty list
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{},
					level:    DEBUG,
				},
				channel:  "channel.1",
				expected: []string{},
			},
			{ // removing a non-existing channel
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"channel.1", "channel.3"},
					level:    DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.3"},
			},
			{ // removing an existing channel
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"channel.1", "channel.2", "channel.3"},
					level:    DEBUG,
				},
				channel:  "channel.2",
				expected: []string{"channel.1", "channel.3"},
			},
		}

		for _, scn := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				stream := &stream{NewMockFormatter(ctrl), scn.state.channels, scn.state.level}
				stream.RemoveChannel(scn.channel)

				if check := stream.ListChannels(); !reflect.DeepEqual(check, scn.expected) {
					t.Errorf("returned the (%v) list of channels", check)
				}
			}
			test()
		}
	})
}

func Test_Stream_Format(t *testing.T) {
	t.Run("return message if there is no formatter", func(t *testing.T) {
		msg := "message"
		level := WARNING
		stream := &stream{nil, []string{}, level}

		if check := stream.format(level, msg, map[string]interface{}{"field": "value"}); check != msg {
			t.Errorf("returned the (%v) formatted message", check)
		}
	})

	t.Run("return formatter response if formatter is present", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		msg := "message"
		ctx := map[string]interface{}{"field": "value"}
		expected := "formatted message"
		level := WARNING
		formatter := NewMockFormatter(ctrl)
		formatter.EXPECT().Format(level, msg, ctx).Return(expected).Times(1)
		stream := &stream{formatter, []string{}, level}

		if check := stream.format(level, msg, ctx); check != expected {
			t.Errorf("returned the (%v) formatted message", check)
		}
	})
}
