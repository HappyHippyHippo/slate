package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_NewKennel(t *testing.T) {
	t.Run("nil factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewKennel(nil)

		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("valid instantiation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		factory := NewMockFactory(ctrl)
		sut, e := NewKennel(factory)

		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case sut.(*Kennel).factory != factory:
			t.Errorf("didn't store the given config instance")
		case sut.(*Kennel).regs == nil:
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
		factory := NewMockFactory(ctrl)
		sut, _ := NewKennel(factory)
		sut.(*Kennel).regs[service] = kennelReg{}

		e := sut.Add(process)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrDuplicateService):
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrDuplicateService)
		}
	})

	t.Run("error while watchdog creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		service := "service"
		process := &Process{service: service}
		factory := NewMockFactory(ctrl)
		factory.EXPECT().Create(service).Return(nil, expected).Times(1)
		sut, _ := NewKennel(factory)

		e := sut.Add(process)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("successful insertion", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service"
		process := &Process{service: service}
		wd := NewMockWatchdog(ctrl)
		factory := NewMockFactory(ctrl)
		factory.EXPECT().Create(service).Return(wd, nil).Times(1)
		sut, _ := NewKennel(factory)

		e := sut.Add(process)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		case len(sut.(*Kennel).regs) != 1:
			t.Error("didn't stored the process registry")
		case sut.(*Kennel).regs[service].process != process:
			t.Error("didn't stored the process info in the created registry")
		case sut.(*Kennel).regs[service].watchdog != wd:
			t.Error("didn't stored the generated watchdog in the created registry")
		}
	})
}

func Test_Kennel_Run(t *testing.T) {
	t.Run("no-op if no process was registered", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		factory := NewMockFactory(ctrl)
		sut, _ := NewKennel(factory)

		if e := sut.Run(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("simple process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "service"
		process := &Process{service: service}
		wd := NewMockWatchdog(ctrl)
		wd.EXPECT().Run(process).Return(nil).Times(1)
		factory := NewMockFactory(ctrl)
		factory.EXPECT().Create(service).Return(wd, nil).Times(1)
		sut, _ := NewKennel(factory)
		_ = sut.Add(process)

		if e := sut.Run(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("multiple processes", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		process1 := &Process{service: "service1"}
		process2 := &Process{service: "service2"}
		process3 := &Process{service: "service3"}
		wd1 := NewMockWatchdog(ctrl)
		wd1.EXPECT().Run(process1).Return(nil).Times(1)
		wd2 := NewMockWatchdog(ctrl)
		wd2.EXPECT().Run(process2).Return(nil).Times(1)
		wd3 := NewMockWatchdog(ctrl)
		wd3.EXPECT().Run(process3).Return(nil).Times(1)
		factory := NewMockFactory(ctrl)
		gomock.InOrder(
			factory.EXPECT().Create("service1").Return(wd1, nil).Times(1),
			factory.EXPECT().Create("service2").Return(wd2, nil).Times(1),
			factory.EXPECT().Create("service3").Return(wd3, nil).Times(1),
		)
		sut, _ := NewKennel(factory)
		_ = sut.Add(process1)
		_ = sut.Add(process2)
		_ = sut.Add(process3)

		if e := sut.Run(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("return process error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		process1 := &Process{service: "service1"}
		process2 := &Process{service: "service2"}
		process3 := &Process{service: "service3"}
		wd1 := NewMockWatchdog(ctrl)
		wd1.EXPECT().Run(process1).Return(nil).Times(1)
		wd2 := NewMockWatchdog(ctrl)
		wd2.EXPECT().Run(process2).Return(expected).Times(1)
		wd3 := NewMockWatchdog(ctrl)
		wd3.EXPECT().Run(process3).Return(nil).Times(1)
		factory := NewMockFactory(ctrl)
		gomock.InOrder(
			factory.EXPECT().Create("service1").Return(wd1, nil).Times(1),
			factory.EXPECT().Create("service2").Return(wd2, nil).Times(1),
			factory.EXPECT().Create("service3").Return(wd3, nil).Times(1),
		)
		sut, _ := NewKennel(factory)
		_ = sut.Add(process1)
		_ = sut.Add(process2)
		_ = sut.Add(process3)

		e := sut.Run()
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})
}
