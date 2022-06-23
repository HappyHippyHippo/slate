package slog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"sort"
	"testing"
)

func Test_NewLogger(t *testing.T) {
	t.Run("new logger", func(t *testing.T) {
		if logger := NewLogger(); logger == nil {
			t.Error("didn't returned a valid reference")
		}
	})
}

func Test_Logger_Close(t *testing.T) {
	t.Run("execute close process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Close().Times(1)
		logger := NewLogger()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)
		_ = logger.Close()

		if logger.HasStream(id1) {
			t.Error("didn't removed the stream")
		}
		if logger.HasStream(id2) {
			t.Error("didn't removed the stream")
		}
	})
}

func Test_Logger_Signal(t *testing.T) {
	t.Run("propagate to all streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		channel := "channel"
		level := WARNING
		message := "message"
		fields := map[string]interface{}{"field": "value"}
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Signal(channel, level, message, fields).Return(nil).Times(1)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Signal(channel, level, message, fields).Return(nil).Times(1)
		stream2.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)

		if err := logger.Signal(channel, level, message, fields); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("return on the first error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		channel := "channel"
		level := WARNING
		message := "message"
		fields := map[string]interface{}{"field": "value"}
		expected := fmt.Errorf("error message")
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Signal(channel, level, message, fields).Return(expected).AnyTimes()
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Signal(channel, level, message, fields).Return(nil).AnyTimes()
		stream2.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)

		if err := logger.Signal(channel, level, message, fields); err == nil {
			t.Error("didn't returned the expected  error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})
}

func Test_Logger_Broadcast(t *testing.T) {
	t.Run("propagate to all streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		level := WARNING
		message := "message"
		fields := map[string]interface{}{"field": "value"}
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Broadcast(level, message, fields).Return(nil).Times(1)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Broadcast(level, message, fields).Return(nil).Times(1)
		stream2.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)

		if err := logger.Broadcast(level, message, fields); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("return on the first error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		level := WARNING
		fields := map[string]interface{}{"field": "value"}
		message := "message"
		expected := fmt.Errorf("error")
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Broadcast(level, message, fields).Return(expected).AnyTimes()
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Broadcast(level, message, fields).Return(nil).AnyTimes()
		stream2.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)

		if err := logger.Broadcast(level, message, fields); err == nil {
			t.Error("didn't returned the expected  error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})
}

func Test_Logger_HasStream(t *testing.T) {
	t.Run("check the registration of a stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "stream.1"
		id2 := "stream.2"
		id3 := "stream.3"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)

		if !logger.HasStream(id1) {
			t.Errorf("returned false")
		}
		if !logger.HasStream(id2) {
			t.Errorf("returned false")
		}
		if logger.HasStream(id3) {
			t.Errorf("returned true")
		}
	})
}

func Test_Logger_ListStreams(t *testing.T) {
	t.Run("retrieve the list of registered streams id's", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "stream.1"
		id2 := "stream.2"
		id3 := "stream.3"
		expected := []string{id1, id2, id3}
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Close().Return(nil).Times(1)
		stream3 := NewMockStream(ctrl)
		stream3.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)
		_ = logger.AddStream(id3, stream3)

		streams := logger.ListStreams()
		sort.Slice(streams, func(i, j int) bool {
			return streams[i] <= streams[j]
		})

		if sort.Search(len(streams), func(i int) bool {
			return streams[i] >= "id1"
		}) >= len(streams) {
			t.Errorf("returned the {%v} id's list instead of the expected: {%v}", streams, expected)
		}
		if sort.Search(len(streams), func(i int) bool {
			return streams[i] >= "id2"
		}) >= len(streams) {
			t.Errorf("returned the {%v} id's list instead of the expected: {%v}", streams, expected)
		}
		if sort.Search(len(streams), func(i int) bool {
			return streams[i] >= "id3"
		}) >= len(streams) {
			t.Errorf("returned the {%v} id's list instead of the expected: {%v}", streams, expected)
		}
	})
}

func Test_Logger_AddStream(t *testing.T) {
	t.Run("error if nil stream", func(t *testing.T) {
		logger := NewLogger()
		defer func() { _ = logger.Close() }()

		if err := logger.AddStream("id", nil); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error if id is duplicate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id, stream1)

		if err := logger.AddStream(id, stream2); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrDuplicateLogStream) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrDuplicateLogStream)
		}
	})

	t.Run("register a new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream := NewMockStream(ctrl)
		stream.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()

		if err := logger.AddStream(id, stream); err != nil {
			t.Errorf("resulted the (%v) error", err)
		} else if check := logger.Stream(id); !reflect.DeepEqual(check, stream) {
			t.Errorf("didn't stored the stream")
		}
	})
}

func Test_Logger_RemoveStream(t *testing.T) {
	t.Run("unregister a stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream := NewMockStream(ctrl)
		stream.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id, stream)
		logger.RemoveStream(id)

		if logger.HasStream(id) {
			t.Errorf("dnd't removed the stream")
		}
	})
}

func Test_Logger_RemoveAllStreams(t *testing.T) {
	t.Run("remove all registered streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "stream.1"
		id2 := "stream.2"
		id3 := "stream.3"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Close().Return(nil).Times(1)
		stream3 := NewMockStream(ctrl)
		stream3.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id1, stream1)
		_ = logger.AddStream(id2, stream2)
		_ = logger.AddStream(id3, stream3)
		logger.RemoveAllStreams()

		if check := logger.ListStreams(); len(check) != 0 {
			t.Errorf("returned the {%v} id's list instead of an empty list", check)
		}
	})
}

func Test_Logger_Stream(t *testing.T) {
	t.Run("nil on a non-existing stream", func(t *testing.T) {
		logger := NewLogger()
		defer func() { _ = logger.Close() }()

		if result := logger.Stream("invalid id"); result != nil {
			t.Errorf("returned a valid stream")
		}
	})

	t.Run("retrieve the requested stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream := NewMockStream(ctrl)
		stream.EXPECT().Close().Return(nil).Times(1)
		logger := NewLogger()
		defer func() { _ = logger.Close() }()
		_ = logger.AddStream(id, stream)

		if check := logger.Stream(id); !reflect.DeepEqual(check, stream) {
			t.Errorf("didn0t retrieved the stored stream")
		}
	})
}
