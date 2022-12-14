package trigger

import (
	"time"
)

// Recurring defines an instance used to execute a callback function
// recurrently in a defined interval time.
type Recurring struct {
	trigger
	ticker      *time.Ticker
	isClosed    bool
	doneChannel chan struct{}
}

var _ ITrigger = &Recurring{}

// NewRecurring instantiate a new trigger that will execute a
// callback method recurrently with a defined periodicity.
func NewRecurring(
	delay time.Duration,
	callback ICallback,
) (*Recurring, error) {
	if callback == nil {
		return nil, errNilPointer("callback")
	}

	t := &Recurring{
		trigger: trigger{
			delay:    delay,
			callback: callback,
		},
		ticker:      time.NewTicker(delay),
		isClosed:    false,
		doneChannel: make(chan struct{}),
	}

	go func(t *Recurring) {
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
func (t *Recurring) Close() error {
	if !t.isClosed {
		t.doneChannel <- struct{}{}
		t.ticker.Stop()
	}
	return nil
}
