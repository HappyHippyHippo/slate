package strigger

import (
	"io"
	"time"
)

// Callback used as a trigger execution process.
type Callback func() error

// Trigger defines the interface of a trigger used to execute a function call
// on determine time intervals.
type Trigger interface {
	io.Closer
	Delay() time.Duration
}

type trigger struct {
	delay    time.Duration
	callback Callback
}

// Delay will retrieve the time period associated to the trigger.
func (t trigger) Delay() time.Duration {
	return t.delay
}
