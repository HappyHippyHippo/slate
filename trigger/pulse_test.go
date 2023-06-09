package trigger

import (
	"errors"
	"testing"
	"time"

	"github.com/happyhippyhippo/slate"
)

func Test_NewPulse(t *testing.T) {
	t.Run("nil callback", func(t *testing.T) {
		if _, e := NewPulse(20*time.Millisecond, nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new Pulse trigger", func(t *testing.T) {
		if _, e := NewPulse(20*time.Millisecond, func() error {
			return nil
		}); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}

func Test_Pulse_Close(t *testing.T) {
	t.Run("is the same as stopping it", func(t *testing.T) {
		called := false
		sut, _ := NewPulse(20*time.Millisecond, func() error {
			called = true
			return nil
		})
		_ = sut.Close()
		time.Sleep(40 * time.Millisecond)

		if called {
			t.Error("didn't prevented the Pulse to execute")
		}
	})
}

func Test_Pulse_Delay(t *testing.T) {
	t.Run("retrieves the trigger delay time", func(t *testing.T) {
		duration := 20 * time.Millisecond
		sut, _ := NewPulse(duration, func() error {
			return nil
		})
		defer func() { _ = sut.Close() }()

		if check := sut.Delay(); check != duration {
			t.Errorf("returned (%v) wait duration", check)
		}
	})
}

func Test_Pulse(t *testing.T) {
	t.Run("only trigger execution once", func(t *testing.T) {
		count := 0
		sut, _ := NewPulse(20*time.Millisecond, func() error {
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
}
