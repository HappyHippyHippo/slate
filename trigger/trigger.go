package trigger

import (
	"time"
)

// Callback used as a trigger execution process.
type Callback func() error

type trigger struct {
	delay    time.Duration
	callback Callback
}

// Delay will retrieve the time period associated to the trigger.
func (t trigger) Delay() time.Duration {
	return t.delay
}
