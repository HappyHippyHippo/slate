package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewLogAdapter(t *testing.T) {
	t.Run("nil log", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		formatter := NewMockLogFormatter(ctrl)

		sut, e := NewLogAdapter("service", "channel", log.FATAL, log.FATAL, log.FATAL, nil, formatter)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil log", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		l := NewMockLog(ctrl)

		sut, e := NewLogAdapter("service", "channel", log.FATAL, log.FATAL, log.FATAL, l, nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("valid instantiation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		l := NewMockLog(ctrl)
		formatter := NewMockLogFormatter(ctrl)

		sut, e := NewLogAdapter("service", "channel", log.FATAL, log.ERROR, log.WARNING, l, formatter)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case sut.service != "service":
			t.Errorf("didn't store the given service name : %v", sut.service)
		case sut.channel != "channel":
			t.Errorf("didn't store the given channel name : %v", sut.channel)
		case sut.startLevel != log.FATAL:
			t.Errorf("didn't store the given start log level : %v", log.LevelMapName[sut.startLevel])
		case sut.errorLevel != log.ERROR:
			t.Errorf("didn't store the given error log level : %v", log.LevelMapName[sut.errorLevel])
		case sut.doneLevel != log.WARNING:
			t.Errorf("didn't store the given done log level : %v", log.LevelMapName[sut.doneLevel])
		case sut.logger != l:
			t.Errorf("didn't store the given log instance")
		case sut.formatter != formatter:
			t.Errorf("didn't store the given formatter instance")
		}
	})
}

func Test_LogAdapter_Start(t *testing.T) {
	t.Run("error while logging", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service name"
		channel := "channel name"
		message := "formatter message"
		expected := fmt.Errorf("error message")
		l := NewMockLog(ctrl)
		l.EXPECT().Signal(channel, log.FATAL, message, nil).Return(expected).Times(1)
		formatter := NewMockLogFormatter(ctrl)
		formatter.EXPECT().Start(service).Return(message).Times(1)

		sut, _ := NewLogAdapter(service, channel, log.FATAL, log.ERROR, log.WARNING, l, formatter)

		chk := sut.Start()
		switch {
		case chk == nil:
			t.Errorf("didn't returned the expected error")
		case chk.Error() != expected.Error():
			t.Errorf("returned the unexpected error (%v) while expecting : %v", chk, expected)
		}
	})

	t.Run("success logging", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service name"
		channel := "channel name"
		message := "formatter message"
		l := NewMockLog(ctrl)
		l.EXPECT().Signal(channel, log.FATAL, message, nil).Return(nil).Times(1)
		formatter := NewMockLogFormatter(ctrl)
		formatter.EXPECT().Start(service).Return(message).Times(1)

		sut, _ := NewLogAdapter(service, channel, log.FATAL, log.ERROR, log.WARNING, l, formatter)

		if chk := sut.Start(); chk != nil {
			t.Errorf("returned the unexpected error : %v", chk)
		}
	})
}

func Test_LogAdapter_Error(t *testing.T) {
	t.Run("error while logging", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service name"
		channel := "channel name"
		message := "formatter message"
		e := fmt.Errorf("error")
		expected := fmt.Errorf("error message")
		l := NewMockLog(ctrl)
		l.EXPECT().Signal(channel, log.ERROR, message, nil).Return(expected).Times(1)
		formatter := NewMockLogFormatter(ctrl)
		formatter.EXPECT().Error(service, e).Return(message).Times(1)

		sut, _ := NewLogAdapter(service, channel, log.FATAL, log.ERROR, log.WARNING, l, formatter)

		chk := sut.Error(e)
		switch {
		case chk == nil:
			t.Errorf("didn't returned the expected error")
		case chk.Error() != expected.Error():
			t.Errorf("returned the unexpected error (%v) while expecting : %v", chk, expected)
		}
	})

	t.Run("success logging", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service name"
		channel := "channel name"
		message := "formatter message"
		e := fmt.Errorf("error")
		l := NewMockLog(ctrl)
		l.EXPECT().Signal(channel, log.ERROR, message, nil).Return(nil).Times(1)
		formatter := NewMockLogFormatter(ctrl)
		formatter.EXPECT().Error(service, e).Return(message).Times(1)

		sut, _ := NewLogAdapter(service, channel, log.FATAL, log.ERROR, log.WARNING, l, formatter)

		if chk := sut.Error(e); chk != nil {
			t.Errorf("returned the unexpected error : %v", chk)
		}
	})
}

func Test_LogAdapter_Done(t *testing.T) {
	t.Run("error while logging", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service name"
		channel := "channel name"
		message := "formatter message"
		expected := fmt.Errorf("error message")
		l := NewMockLog(ctrl)
		l.EXPECT().Signal(channel, log.WARNING, message, nil).Return(expected).Times(1)
		formatter := NewMockLogFormatter(ctrl)
		formatter.EXPECT().Done(service).Return(message).Times(1)

		sut, _ := NewLogAdapter(service, channel, log.FATAL, log.ERROR, log.WARNING, l, formatter)

		chk := sut.Done()
		switch {
		case chk == nil:
			t.Errorf("didn't returned the expected error")
		case chk.Error() != expected.Error():
			t.Errorf("returned the unexpected error (%v) while expecting : %v", chk, expected)
		}
	})

	t.Run("success logging", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service name"
		channel := "channel name"
		message := "formatter message"
		l := NewMockLog(ctrl)
		l.EXPECT().Signal(channel, log.WARNING, message, nil).Return(nil).Times(1)
		formatter := NewMockLogFormatter(ctrl)
		formatter.EXPECT().Done(service).Return(message).Times(1)

		sut, _ := NewLogAdapter(service, channel, log.FATAL, log.ERROR, log.WARNING, l, formatter)

		if chk := sut.Done(); chk != nil {
			t.Errorf("returned the unexpected error : %v", chk)
		}
	})
}
