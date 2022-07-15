package trigger

import (
	"time"
)

type pulse struct {
	trigger
	timer *time.Timer
}

var _ ITrigger = &pulse{}

// NewPulse instantiate a new pulse trigger that will execute a
// callback method after a determined amount of time.
func NewPulse(delay time.Duration, callback ICallback) (ITrigger, error) {
	if callback == nil {
		return nil, errNilPointer("callback")
	}

	t := &pulse{
		trigger: trigger{
			delay:    delay,
			callback: callback,
		},
		timer: time.NewTimer(delay),
	}

	go func(t *pulse) {
		<-t.timer.C
		_ = t.callback()
	}(t)

	return t, nil
}

func (t *pulse) Close() error {
	t.timer.Stop()
	return nil
}
