package trigger

import (
	"io"
	"time"
)

// ICallback used as a trigger execution process.
type ICallback func() error

// ITrigger defines the interface of a trigger used to execute a function call
// on determine time intervals.
type ITrigger interface {
	io.Closer
	Delay() time.Duration
}

type trigger struct {
	delay    time.Duration
	callback ICallback
}

// Delay will retrieve the time period associated to the trigger.
func (t trigger) Delay() time.Duration {
	return t.delay
}
