package console

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/log/formatter/json"
)

func Test_NewStreamStrategy(t *testing.T) {
	t.Run("nil formatter factory", func(t *testing.T) {
		sut, e := NewStreamStrategy(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new console stream factory strategy", func(t *testing.T) {
		if sut, e := NewStreamStrategy(log.NewFormatterFactory()); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}

func Test_StreamStrategy_Accept(t *testing.T) {
	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		partial := config.Partial{"type": config.Partial{}}
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		if sut.Accept(partial) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		partial := config.Partial{"type": "invalid type"}
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		if sut.Accept(partial) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		partial := config.Partial{"type": Type}
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		if !sut.Accept(partial) {
			t.Error("returned false")
		}
	})
}

func Test_StreamStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		src, e := sut.Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		partial := config.Partial{
			"type":   Type,
			"format": 123,
		}
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		stream, e := sut.Create(partial)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-log level name level", func(t *testing.T) {
		partial := config.Partial{
			"type":   Type,
			"format": json.Format,
			"level":  "invalid",
		}
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		stream, e := sut.Create(partial)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, log.ErrInvalidLevel):
			t.Errorf("(%v) when expecting (%v)", e, log.ErrInvalidLevel)
		}
	})

	t.Run("error creating the formatter", func(t *testing.T) {
		partial := config.Partial{
			"type":   Type,
			"format": json.Format,
			"level":  "fatal",
		}
		sut, _ := NewStreamStrategy(log.NewFormatterFactory())

		stream, e := sut.Create(partial)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, log.ErrInvalidFormat):
			t.Errorf("(%v) when expecting (%v)", e, log.ErrInvalidFormat)
		}
	})

	t.Run("new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{
			"type":     Type,
			"format":   "format",
			"level":    "fatal",
			"channels": []interface{}{"channel1", "channel2"}}
		formatter := NewMockFormatter(ctrl)
		formatterStrategy := NewMockFormatterStrategy(ctrl)
		formatterStrategy.EXPECT().Accept("format").Return(true).Times(1)
		formatterStrategy.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
		formatterFactory := log.NewFormatterFactory()
		_ = formatterFactory.Register(formatterStrategy)
		sut, _ := NewStreamStrategy(formatterFactory)

		stream, e := sut.Create(partial)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := stream.(type) {
			case *Stream:
				switch {
				case s.Level != log.FATAL:
					t.Errorf("invalid level (%s)", log.LevelMapName[s.Level])
				case len(s.Channels) != 2:
					t.Errorf("invalid channel list (%v)", s.Channels)
				}
			default:
				t.Error("didn't returned a new console stream")
			}
		}
	})
}
