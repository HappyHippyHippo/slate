package gtrigger

import (
	"errors"
	"github.com/happyhippyhippo/slate/gerror"
	"testing"
	"time"
)

func Test_NewPulse(t *testing.T) {
	t.Run("nil callback", func(t *testing.T) {
		if _, err := NewPulse(20*time.Millisecond, nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("new pulse trigger", func(t *testing.T) {
		if _, err := NewPulse(20*time.Millisecond, func() error {
			return nil
		}); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		}
	})
}

func Test_Pulse_Close(t *testing.T) {
	t.Run("is the same as stopping it", func(t *testing.T) {
		called := false
		pulse, _ := NewPulse(20*time.Millisecond, func() error {
			called = true
			return nil
		})
		_ = pulse.Close()
		time.Sleep(40 * time.Millisecond)

		if called {
			t.Error("didn't prevented the pulse to execute")
		}
	})
}

func Test_Pulse_Delay(t *testing.T) {
	t.Run("retrieves the trigger delay time", func(t *testing.T) {
		duration := 20 * time.Millisecond
		pulse, _ := NewPulse(duration, func() error {
			return nil
		})
		defer func() { _ = pulse.Close() }()

		if check := pulse.Delay(); check != duration {
			t.Errorf("returned (%v) wait duration", check)
		}
	})
}

func Test_Pulse(t *testing.T) {
	t.Run("only trigger execution once", func(t *testing.T) {
		count := 0
		pulse, _ := NewPulse(20*time.Millisecond, func() error {
			count++
			return nil
		})
		defer func() { _ = pulse.Close() }()
		time.Sleep(100 * time.Millisecond)

		if count == 0 {
			t.Error("didn't called the callback function once")
		} else if count > 1 {
			t.Error("recurrently called the callback function")
		}
	})
}
