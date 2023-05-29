package log

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"reflect"
	"sort"
	"testing"
)

func Test_NewLog(t *testing.T) {
	t.Run("new Log", func(t *testing.T) {
		if sut := NewLog(); sut == nil {
			t.Error("didn't returned a valid reference")
		}
	})
}

func Test_Log_Close(t *testing.T) {
	t.Run("execute close process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Close().Times(1)
		sut := NewLog()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)
		_ = sut.Close()

		if sut.HasStream(id1) {
			t.Error("didn't removed the stream")
		}
		if sut.HasStream(id2) {
			t.Error("didn't removed the stream")
		}
	})
}

func Test_Log_Signal(t *testing.T) {
	t.Run("propagate to all streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		channel := "channel"
		level := WARNING
		message := "message"
		ctx := Context{"field": "value"}
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Signal(channel, level, message, ctx).Return(nil).Times(1)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Signal(channel, level, message, ctx).Return(nil).Times(1)
		stream2.EXPECT().Close().Return(nil).Times(1)
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)

		if e := sut.Signal(channel, level, message, ctx); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("return on the first error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		channel := "channel"
		level := WARNING
		message := "message"
		ctx := Context{"field": "value"}
		expected := fmt.Errorf("error message")
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Signal(channel, level, message, ctx).Return(expected).AnyTimes()
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Signal(channel, level, message, ctx).Return(nil).AnyTimes()
		stream2.EXPECT().Close().Return(nil).Times(1)
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)

		if e := sut.Signal(channel, level, message, ctx); e == nil {
			t.Error("didn't returned the expected  error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})
}

func Test_Log_Broadcast(t *testing.T) {
	t.Run("propagate to all streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		level := WARNING
		message := "message"
		ctx := Context{"field": "value"}
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Broadcast(level, message, ctx).Return(nil).Times(1)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Broadcast(level, message, ctx).Return(nil).Times(1)
		stream2.EXPECT().Close().Return(nil).Times(1)
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)

		if e := sut.Broadcast(level, message, ctx); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("return on the first error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		level := WARNING
		ctx := Context{"field": "value"}
		message := "message"
		expected := fmt.Errorf("error")
		id1 := "stream.1"
		id2 := "stream.2"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Broadcast(level, message, ctx).Return(expected).AnyTimes()
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		stream2.EXPECT().Broadcast(level, message, ctx).Return(nil).AnyTimes()
		stream2.EXPECT().Close().Return(nil).Times(1)
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)

		if e := sut.Broadcast(level, message, ctx); e == nil {
			t.Error("didn't returned the expected  error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})
}

func Test_Log_HasStream(t *testing.T) {
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
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)

		if !sut.HasStream(id1) {
			t.Errorf("returned false")
		}
		if !sut.HasStream(id2) {
			t.Errorf("returned false")
		}
		if sut.HasStream(id3) {
			t.Errorf("returned true")
		}
	})
}

func Test_Log_ListStreams(t *testing.T) {
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
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)
		_ = sut.AddStream(id3, stream3)

		streams := sut.ListStreams()
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

func Test_Log_AddStream(t *testing.T) {
	t.Run("error if nil stream", func(t *testing.T) {
		sut := NewLog()
		defer func() { _ = sut.Close() }()

		if e := sut.AddStream("id", nil); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error if id is duplicate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		stream2 := NewMockStream(ctrl)
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id, stream1)

		if e := sut.AddStream(id, stream2); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, ErrDuplicateStream) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrDuplicateStream)
		}
	})

	t.Run("register a new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		sut := NewLog()
		defer func() { _ = sut.Close() }()

		if e := sut.AddStream(id, stream1); e != nil {
			t.Errorf("resulted the (%v) error", e)
		} else if check, e := sut.Stream(id); !reflect.DeepEqual(check, stream1) {
			t.Errorf("didn't stored the stream")
		} else if e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})
}

func Test_Log_RemoveStream(t *testing.T) {
	t.Run("unregister a stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id, stream1)
		sut.RemoveStream(id)

		if sut.HasStream(id) {
			t.Errorf("dnd't removed the stream")
		}
	})
}

func Test_Log_RemoveAllStreams(t *testing.T) {
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
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id1, stream1)
		_ = sut.AddStream(id2, stream2)
		_ = sut.AddStream(id3, stream3)
		sut.RemoveAllStreams()

		if check := sut.ListStreams(); len(check) != 0 {
			t.Errorf("returned the {%v} id's list instead of an empty list", check)
		}
	})
}

func Test_Log_Stream(t *testing.T) {
	t.Run("non-existing stream", func(t *testing.T) {
		sut := NewLog()
		defer func() { _ = sut.Close() }()

		result, e := sut.Stream("invalid id")
		switch {
		case result != nil:
			t.Errorf("returned a valid stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrStreamNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrStreamNotFound)
		}
	})

	t.Run("retrieve the requested stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := "stream"
		stream1 := NewMockStream(ctrl)
		stream1.EXPECT().Close().Return(nil).Times(1)
		sut := NewLog()
		defer func() { _ = sut.Close() }()
		_ = sut.AddStream(id, stream1)

		if check, e := sut.Stream(id); !reflect.DeepEqual(check, stream1) {
			t.Errorf("didn0t retrieved the stored stream")
		} else if e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})
}
