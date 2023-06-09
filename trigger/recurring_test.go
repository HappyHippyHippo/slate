package trigger

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/happyhippyhippo/slate"
)

func Test_NewRecurring(t *testing.T) {
	t.Run("nil callback", func(t *testing.T) {
		if _, e := NewRecurring(20*time.Millisecond, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new Recurring trigger", func(t *testing.T) {
		if _, e := NewRecurring(20*time.Millisecond, func() error {
			return nil
		}); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}

func Test_Recurring_Close(t *testing.T) {
	t.Run("is the same as stopping it", func(t *testing.T) {
		called := false
		sut, _ := NewRecurring(20*time.Millisecond, func() error {
			called = true
			return nil
		})
		_ = sut.Close()
		time.Sleep(40 * time.Millisecond)

		if called {
			t.Error("didn't stop the Recurring to be executed")
		}
	})
}

func Test_Recurring_Delay(t *testing.T) {
	t.Run("retrieves the trigger interval duration", func(t *testing.T) {
		duration := 20 * time.Millisecond
		sut, _ := NewRecurring(duration, func() error {
			return nil
		})
		defer func() { _ = sut.Close() }()

		if check := sut.Delay(); check != duration {
			t.Errorf("returned (%v) interval duration", check)
		}
	})
}

func Test_Recurring(t *testing.T) {
	t.Run("run trigger multiple times", func(t *testing.T) {
		count := 0
		sut, _ := NewRecurring(20*time.Millisecond, func() error {
			count++
			return nil
		})
		defer func() { _ = sut.Close() }()
		time.Sleep(100 * time.Millisecond)

		if count <= 2 {
			t.Error("didn't recurrently called the callback function")
		}
	})

	t.Run("stop the trigger on callback error", func(t *testing.T) {
		count := 0
		sut, _ := NewRecurring(20*time.Millisecond, func() error {
			count++
			return fmt.Errorf("__dummy_error__")
		})
		defer func() { _ = sut.Close() }()
		time.Sleep(100 * time.Millisecond)

		if count != 1 {
			t.Error("didn't stop recursion calls after the first error")
		}
	})
}
