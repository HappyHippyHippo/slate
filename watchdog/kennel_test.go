package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_NewKennel(t *testing.T) {
	t.Run("nil creator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewKennel(nil)

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
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		factory := &Factory{}
		sut, e := NewKennel(factory)

		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case sut.creator != factory:
			t.Errorf("didn't store the given creator instance")
		case sut.regs == nil:
			t.Errorf("didn't initialize the kennel registration map")
		}
	})
}

func Test_Kennel_Add(t *testing.T) {
	t.Run("error on duplicate service name", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service"
		process := &Process{service: service}
		creator := NewMockCreator(ctrl)
		sut, _ := NewKennel(&Factory{})
		sut.creator = creator
		sut.regs[service] = kennelReg{}

		e := sut.Add(process)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrDuplicateService):
			t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateService)
		}
	})

	t.Run("error while watchdog creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		service := "service"
		process := &Process{service: service}
		creator := NewMockCreator(ctrl)
		creator.EXPECT().Create(service).Return(nil, expected).Times(1)
		sut, _ := NewKennel(&Factory{})
		sut.creator = creator

		e := sut.Add(process)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("successful insertion", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service"
		process := &Process{service: service}
		wd := &Watchdog{}
		creator := NewMockCreator(ctrl)
		creator.EXPECT().Create(service).Return(wd, nil).Times(1)
		sut, _ := NewKennel(&Factory{})
		sut.creator = creator

		e := sut.Add(process)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case len(sut.regs) != 1:
			t.Error("didn't stored the process registry")
		case sut.regs[service].process != process:
			t.Error("didn't stored the process info in the created registry")
		case sut.regs[service].watchdog != wd:
			t.Error("didn't stored the generated watchdog in the created registry")
		}
	})
}

func Test_Kennel_Run(t *testing.T) {
	t.Run("no-op if no process was registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewKennel(&Factory{})

		if e := sut.Run(); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("simple process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service"
		process := &Process{service: service, runner: func() error { return nil }}
		logAdapter := NewMockLogAdapter(ctrl)
		logAdapter.EXPECT().Start().Return(nil).Times(1)
		logAdapter.EXPECT().Done().Return(nil).Times(1)
		wd := &Watchdog{logAdapter: logAdapter}
		creator := NewMockCreator(ctrl)
		creator.EXPECT().Create(service).Return(wd, nil).Times(1)
		sut, _ := NewKennel(&Factory{})
		sut.creator = creator
		_ = sut.Add(process)

		if e := sut.Run(); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("multiple processes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		runner := func() error { return nil }

		process1 := &Process{service: "service1", runner: runner}
		logAdapter1 := NewMockLogAdapter(ctrl)
		logAdapter1.EXPECT().Start().Return(nil).Times(1)
		logAdapter1.EXPECT().Done().Return(nil).Times(1)
		wd1 := &Watchdog{logAdapter: logAdapter1}
		process2 := &Process{service: "service2", runner: runner}
		logAdapter2 := NewMockLogAdapter(ctrl)
		logAdapter2.EXPECT().Start().Return(nil).Times(1)
		logAdapter2.EXPECT().Done().Return(nil).Times(1)
		wd2 := &Watchdog{logAdapter: logAdapter2}
		process3 := &Process{service: "service3", runner: runner}
		logAdapter3 := NewMockLogAdapter(ctrl)
		logAdapter3.EXPECT().Start().Return(nil).Times(1)
		logAdapter3.EXPECT().Done().Return(nil).Times(1)
		wd3 := &Watchdog{logAdapter: logAdapter3}
		creator := NewMockCreator(ctrl)
		gomock.InOrder(
			creator.EXPECT().Create("service1").Return(wd1, nil).Times(1),
			creator.EXPECT().Create("service2").Return(wd2, nil).Times(1),
			creator.EXPECT().Create("service3").Return(wd3, nil).Times(1),
		)
		sut, _ := NewKennel(&Factory{})
		sut.creator = creator
		_ = sut.Add(process1)
		_ = sut.Add(process2)
		_ = sut.Add(process3)

		if e := sut.Run(); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("return process error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		runner := func() error { return nil }

		expected := fmt.Errorf("error message")
		panicError := fmt.Errorf("panic error")
		process1 := &Process{service: "service1", runner: runner}
		logAdapter1 := NewMockLogAdapter(ctrl)
		logAdapter1.EXPECT().Start().Return(nil).Times(1)
		logAdapter1.EXPECT().Done().Return(nil).Times(1)
		wd1 := &Watchdog{logAdapter: logAdapter1}
		count := 0
		process2 := &Process{service: "service2", runner: func() error {
			count++
			if count == 1 {
				panic(panicError)
			}
			return expected
		}}
		logAdapter2 := NewMockLogAdapter(ctrl)
		logAdapter2.EXPECT().Start().Return(nil).Times(1)
		logAdapter2.EXPECT().Error(panicError).Return(nil).Times(1)
		logAdapter2.EXPECT().Done().Return(nil).Times(1)
		wd2 := &Watchdog{logAdapter: logAdapter2}
		process3 := &Process{service: "service3", runner: runner}
		logAdapter3 := NewMockLogAdapter(ctrl)
		logAdapter3.EXPECT().Start().Return(nil).Times(1)
		logAdapter3.EXPECT().Done().Return(nil).Times(1)
		wd3 := &Watchdog{logAdapter: logAdapter3}
		creator := NewMockCreator(ctrl)
		gomock.InOrder(
			creator.EXPECT().Create("service1").Return(wd1, nil).Times(1),
			creator.EXPECT().Create("service2").Return(wd2, nil).Times(1),
			creator.EXPECT().Create("service3").Return(wd3, nil).Times(1),
		)
		sut, _ := NewKennel(&Factory{})
		sut.creator = creator
		_ = sut.Add(process1)
		_ = sut.Add(process2)
		_ = sut.Add(process3)

		e := sut.Run()
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})
}
