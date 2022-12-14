package trigger

import (
	"time"
)

// Pulse defines an instance used to execute a callback function
// after a defined time.
type Pulse struct {
	trigger
	timer *time.Timer
}

var _ ITrigger = &Pulse{}

// NewPulse instantiate a new pulse trigger that will execute a
// callback method after a determined amount of time.
func NewPulse(
	delay time.Duration,
	callback ICallback,
) (*Pulse, error) {
	if callback == nil {
		return nil, errNilPointer("callback")
	}

	t := &Pulse{
		trigger: trigger{
			delay:    delay,
			callback: callback,
		},
		timer: time.NewTimer(delay),
	}

	go func(t *Pulse) {
		<-t.timer.C
		_ = t.callback()
	}(t)

	return t, nil
}

// Close will stop the trigger and release all the associated resources.
func (t *Pulse) Close() error {
	t.timer.Stop()
	return nil
}
