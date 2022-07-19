package strigger

import (
	"time"
)

type recurring struct {
	trigger
	ticker      *time.Ticker
	isClosed    bool
	doneChannel chan struct{}
}

var _ ITrigger = &recurring{}

// NewRecurring instantiate a new strigger that will execute a
// callback method recurrently with a defined periodicity.
func NewRecurring(delay time.Duration, callback ICallback) (ITrigger, error) {
	if callback == nil {
		return nil, errNilPointer("callback")
	}

	t := &recurring{
		trigger: trigger{
			delay:    delay,
			callback: callback,
		},
		ticker:      time.NewTicker(delay),
		isClosed:    false,
		doneChannel: make(chan struct{}),
	}

	go func(t *recurring) {
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

func (t *recurring) Close() error {
	if !t.isClosed {
		t.doneChannel <- struct{}{}
		t.ticker.Stop()
	}
	return nil
}
