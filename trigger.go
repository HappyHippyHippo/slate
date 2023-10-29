package slate

import (
	"io"
	"time"
)

// ----------------------------------------------------------------------------
// trigger
// ----------------------------------------------------------------------------

// TriggerCallback used as a trigger execution process.
type TriggerCallback func() error

// Trigger defines an interface of all the time related future
// execution services.
type Trigger interface {
	io.Closer
	Delay() time.Duration
}

type trigger struct {
	delay    time.Duration
	callback TriggerCallback
}

// Delay will retrieve the time period associated to the trigger.
func (t trigger) Delay() time.Duration {
	return t.delay
}

// ----------------------------------------------------------------------------
// pulse trigger
// ----------------------------------------------------------------------------

// TriggerPulse defines an instance used to execute a callback function
// after a defined time.
type TriggerPulse struct {
	trigger
	timer *time.Timer
}

// NewTriggerPulse instantiate a new pulse trigger that will execute a
// callback method after a determined amount of time.
func NewTriggerPulse(
	delay time.Duration,
	callback TriggerCallback,
) (*TriggerPulse, error) {
	if callback == nil {
		return nil, errNilPointer("callback")
	}

	t := &TriggerPulse{
		trigger: trigger{
			delay:    delay,
			callback: callback,
		},
		timer: time.NewTimer(delay),
	}

	go func(t *TriggerPulse) {
		<-t.timer.C
		_ = t.callback()
	}(t)

	return t, nil
}

// Close will stop the trigger and release all the associated resources.
func (t *TriggerPulse) Close() error {
	t.timer.Stop()
	return nil
}

// ----------------------------------------------------------------------------
// recurring trigger
// ----------------------------------------------------------------------------

// TriggerRecurring defines an instance used to execute a callback function
// recurrently in a defined interval time.
type TriggerRecurring struct {
	trigger
	ticker      *time.Ticker
	isClosed    bool
	doneChannel chan struct{}
}

// NewTriggerRecurring instantiate a new trigger that will execute a
// callback method recurrently with a defined periodicity.
func NewTriggerRecurring(
	delay time.Duration,
	callback TriggerCallback,
) (*TriggerRecurring, error) {
	if callback == nil {
		return nil, errNilPointer("callback")
	}

	t := &TriggerRecurring{
		trigger: trigger{
			delay:    delay,
			callback: callback,
		},
		ticker:      time.NewTicker(delay),
		isClosed:    false,
		doneChannel: make(chan struct{}),
	}

	go func(t *TriggerRecurring) {
		for {
			select {
			case <-t.ticker.C:
				if t.callback() != nil {
					t.isClosed = true
					return
				}
			case <-t.doneChannel:
				t.isClosed = true
				return
			}
		}
	}(t)

	return t, nil
}

// Close will stop the trigger and release all the associated resources.
func (t *TriggerRecurring) Close() error {
	if !t.isClosed {
		t.doneChannel <- struct{}{}
		t.ticker.Stop()
	}
	return nil
}
