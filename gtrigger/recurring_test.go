package gtrigger

import (
	"errors"
	"fmt"
	"github.com/happyhippyhippo/slate/gerror"
	"testing"
	"time"
)

func Test_NewRecurring(t *testing.T) {
	t.Run("nil callback", func(t *testing.T) {
		if _, err := NewRecurring(20*time.Millisecond, nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("new recurring trigger", func(t *testing.T) {
		if _, err := NewRecurring(20*time.Millisecond, func() error {
			return nil
		}); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		}
	})
}

func Test_Recurring_Close(t *testing.T) {
	t.Run("is the same as stopping it", func(t *testing.T) {
		called := false
		recurring, _ := NewRecurring(20*time.Millisecond, func() error {
			called = true
			return nil
		})
		_ = recurring.Close()
		time.Sleep(40 * time.Millisecond)

		if called {
			t.Error("didn't stop the recurring to be executed")
		}
	})
}

func Test_Recurring_Delay(t *testing.T) {
	t.Run("retrieves the trigger interval duration", func(t *testing.T) {
		duration := 20 * time.Millisecond
		recurring, _ := NewRecurring(duration, func() error {
			return nil
		})
		defer func() { _ = recurring.Close() }()

		if check := recurring.Delay(); check != duration {
			t.Errorf("returned (%v) interval duration", check)
		}
	})
}

func Test_Recurring(t *testing.T) {
	t.Run("run trigger multiple times", func(t *testing.T) {
		count := 0
		recurring, _ := NewRecurring(20*time.Millisecond, func() error {
			count++
			return nil
		})
		defer func() { _ = recurring.Close() }()
		time.Sleep(100 * time.Millisecond)

		if count <= 2 {
			t.Error("didn't recurrently called the callback function")
		}
	})

	t.Run("stop the trigger on callback error", func(t *testing.T) {
		count := 0
		recurring, _ := NewRecurring(20*time.Millisecond, func() error {
			count++
			return fmt.Errorf("__dummy_error__")
		})
		defer func() { _ = recurring.Close() }()
		time.Sleep(100 * time.Millisecond)

		if count != 1 {
			t.Error("didn't stop recursion calls after the first error")
		}
	})
}
