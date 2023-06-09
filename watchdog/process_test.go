package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/happyhippyhippo/slate"
)

func Test_NewProcess(t *testing.T) {
	t.Run("nil runner", func(t *testing.T) {
		sut, e := NewProcess("service", nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new process", func(t *testing.T) {
		service := "service name"
		runner := func() error { return nil }

		sut, e := NewProcess(service, runner)
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut.service != service:
			t.Errorf("(%v) service when expecting (%v)", sut.service, service)
		case fmt.Sprintf("%p", sut.runner) != fmt.Sprintf("%p", runner):
			t.Errorf("(%p) runner when expecting (%p)", sut.runner, runner)
		}
	})
}

func Test_Process_Service(t *testing.T) {
	t.Run("retrieve the service name", func(t *testing.T) {
		service := "service name"
		runner := func() error { return nil }
		sut, _ := NewProcess(service, runner)

		if chk := sut.Service(); chk != service {
			t.Errorf("(%v) service when expecting (%v)", sut.service, service)
		}
	})
}

func Test_Process_Runner(t *testing.T) {
	t.Run("retrieve the runner method", func(t *testing.T) {
		service := "service name"
		runner := func() error { return nil }
		sut, _ := NewProcess(service, runner)

		if chk := sut.Runner(); fmt.Sprintf("%p", chk) != fmt.Sprintf("%p", runner) {
			t.Errorf("(%p) runner when expecting (%p)", sut.runner, runner)
		}
	})
}
