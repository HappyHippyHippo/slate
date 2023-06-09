package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_Watchdog(t *testing.T) {
	t.Run("nil logAdapter", func(t *testing.T) {
		sut, e := NewWatchdog(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("valid instantiation", func(t *testing.T) {
		logAdapter := &LogAdapter{}

		sut, e := NewWatchdog(logAdapter)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case sut.logAdapter != logAdapter:
			t.Errorf("didn't store the given logAdapter instance")
		}
	})
}

func Test_Watchdog_Run(t *testing.T) {
	t.Run("simple execution", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service"
		logAdapter := NewMockLogAdapter(ctrl)
		logAdapter.EXPECT().Start().Return(nil).Times(1)
		logAdapter.EXPECT().Error(gomock.Any()).Return(nil).Times(0)
		logAdapter.EXPECT().Done().Return(nil).Times(1)
		sut, _ := NewWatchdog(&LogAdapter{})
		sut.logAdapter = logAdapter

		count := 0
		p, _ := NewProcess(service, func() error {
			count++
			return nil
		})

		chk := sut.Run(p)
		switch {
		case count != 1:
			t.Errorf("didn't executed the process method")
		case chk != nil:
			t.Errorf("unexpected (%v) error", chk)
		}
	})

	t.Run("simple execution but return an error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service"
		logAdapter := NewMockLogAdapter(ctrl)
		logAdapter.EXPECT().Start().Return(nil).Times(1)
		logAdapter.EXPECT().Error(gomock.Any()).Return(nil).Times(0)
		logAdapter.EXPECT().Done().Return(nil).Times(1)
		sut, _ := NewWatchdog(&LogAdapter{})
		sut.logAdapter = logAdapter

		e := fmt.Errorf("error message")
		count := 0
		p, _ := NewProcess(service, func() error {
			count++
			return e
		})

		chk := sut.Run(p)
		switch {
		case count != 1:
			t.Errorf("didn't executed the process method")
		case chk == nil:
			t.Error("didn't returned the expected error")
		case chk.Error() != e.Error():
			t.Errorf("(%v) when expecting (%v)", chk, e)
		}
	})

	t.Run("panic recovery on execution", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		e := fmt.Errorf("error message")

		service := "service"
		logAdapter := NewMockLogAdapter(ctrl)
		logAdapter.EXPECT().Start().Return(nil).Times(1)
		logAdapter.EXPECT().Error(e).Return(nil).Times(1)
		logAdapter.EXPECT().Done().Return(nil).Times(1)
		sut, _ := NewWatchdog(&LogAdapter{})
		sut.logAdapter = logAdapter

		count := 0
		p, _ := NewProcess(service, func() error {
			count++
			if count == 1 {
				panic(e)
			}
			return nil
		})

		chk := sut.Run(p)
		switch {
		case count != 2:
			t.Errorf("didn't executed the process method two times")
		case chk != nil:
			t.Errorf("unexpected (%v) error", chk)
		}
	})

	t.Run("panic recovery on execution and error return", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		e := fmt.Errorf("error message")

		service := "service"
		logAdapter := NewMockLogAdapter(ctrl)
		logAdapter.EXPECT().Start().Return(nil).Times(1)
		logAdapter.EXPECT().Error(e).Return(nil).Times(1)
		logAdapter.EXPECT().Done().Return(nil).Times(1)
		sut, _ := NewWatchdog(&LogAdapter{})
		sut.logAdapter = logAdapter

		count := 0
		p, _ := NewProcess(service, func() error {
			count++
			if count == 1 {
				panic(e)
			}
			return e
		})

		chk := sut.Run(p)
		switch {
		case count != 2:
			t.Errorf("didn't executed the process method two times")
		case chk == nil:
			t.Error("didn't returned the expected error")
		case chk.Error() != e.Error():
			t.Errorf("(%v) when expecting (%v)", chk, e)
		}
	})
}
