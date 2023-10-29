package slate

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test_TriggerPulse(t *testing.T) {
	t.Run("NewTriggerPulse", func(t *testing.T) {
		t.Run("nil callback", func(t *testing.T) {
			if _, e := NewTriggerPulse(20*time.Millisecond, nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new TriggerPulse trigger", func(t *testing.T) {
			if _, e := NewTriggerPulse(20*time.Millisecond, func() error {
				return nil
			}); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("is the same as stopping it", func(t *testing.T) {
			called := false
			sut, _ := NewTriggerPulse(20*time.Millisecond, func() error {
				called = true
				return nil
			})
			_ = sut.Close()
			time.Sleep(40 * time.Millisecond)

			if called {
				t.Error("didn't prevented the TriggerPulse to execute")
			}
		})
	})

	t.Run("Delay", func(t *testing.T) {
		t.Run("retrieves the trigger delay time", func(t *testing.T) {
			duration := 20 * time.Millisecond
			sut, _ := NewTriggerPulse(duration, func() error {
				return nil
			})
			defer func() { _ = sut.Close() }()

			if check := sut.Delay(); check != duration {
				t.Errorf("returned (%v) wait duration", check)
			}
		})
	})

	t.Run("Execution", func(t *testing.T) {
		t.Run("only trigger execution once", func(t *testing.T) {
			count := 0
			sut, _ := NewTriggerPulse(20*time.Millisecond, func() error {
				count++
				return nil
			})
			defer func() { _ = sut.Close() }()
			time.Sleep(100 * time.Millisecond)

			if count == 0 {
				t.Error("didn't called the callback function once")
			} else if count > 1 {
				t.Error("recurrently called the callback function")
			}
		})
	})
}

func Test_TriggerRecurring(t *testing.T) {
	t.Run("NewTriggerRecurring", func(t *testing.T) {
		t.Run("nil callback", func(t *testing.T) {
			if _, e := NewTriggerRecurring(20*time.Millisecond, nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new TriggerRecurring trigger", func(t *testing.T) {
			if _, e := NewTriggerRecurring(20*time.Millisecond, func() error {
				return nil
			}); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("is the same as stopping it", func(t *testing.T) {
			called := false
			sut, _ := NewTriggerRecurring(20*time.Millisecond, func() error {
				called = true
				return nil
			})
			_ = sut.Close()
			time.Sleep(40 * time.Millisecond)

			if called {
				t.Error("didn't stop the TriggerRecurring to be executed")
			}
		})
	})

	t.Run("Delay", func(t *testing.T) {
		t.Run("retrieves the trigger interval duration", func(t *testing.T) {
			duration := 20 * time.Millisecond
			sut, _ := NewTriggerRecurring(duration, func() error {
				return nil
			})
			defer func() { _ = sut.Close() }()

			if check := sut.Delay(); check != duration {
				t.Errorf("returned (%v) interval duration", check)
			}
		})
	})

	t.Run("Execution", func(t *testing.T) {
		t.Run("run trigger multiple times", func(t *testing.T) {
			count := 0
			sut, _ := NewTriggerRecurring(20*time.Millisecond, func() error {
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
			sut, _ := NewTriggerRecurring(20*time.Millisecond, func() error {
				count++
				return fmt.Errorf("__dummy_error__")
			})
			defer func() { _ = sut.Close() }()
			time.Sleep(100 * time.Millisecond)

			if count != 1 {
				t.Error("didn't stop recursion calls after the first error")
			}
		})
	})
}
