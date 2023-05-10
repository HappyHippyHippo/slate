package def

import (
	"fmt"
	"testing"

	"github.com/happyhippyhippo/slate/watchdog"
)

func Test_LogFormatterDefault_Start(t *testing.T) {
	t.Run("format message", func(t *testing.T) {
		service := "service name"
		expected := fmt.Sprintf(watchdog.LogStartMessage, service)

		if check := (&Formatter{}).Start(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})

	t.Run("format message from env", func(t *testing.T) {
		prev := watchdog.LogStartMessage
		message := "test message : %s"
		watchdog.LogStartMessage = message
		defer func() { watchdog.LogStartMessage = prev }()

		service := "service name"
		expected := fmt.Sprintf(watchdog.LogStartMessage, service)

		if check := (&Formatter{}).Start(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})
}

func Test_LogFormatterDefault_Error(t *testing.T) {
	t.Run("format message", func(t *testing.T) {
		service := "service name"
		err := fmt.Errorf("error message")
		expected := fmt.Sprintf(watchdog.LogErrorMessage, service, err)

		if check := (&Formatter{}).Error(service, err); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})

	t.Run("format message from env", func(t *testing.T) {
		prev := watchdog.LogErrorMessage
		message := "test message : %s - %e"
		watchdog.LogErrorMessage = message
		defer func() { watchdog.LogErrorMessage = prev }()

		service := "service name"
		err := fmt.Errorf("error message")
		expected := fmt.Sprintf(watchdog.LogErrorMessage, service, err)

		if check := (&Formatter{}).Error(service, err); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})
}

func Test_LogFormatterDefault_Done(t *testing.T) {
	t.Run("format message", func(t *testing.T) {
		service := "service name"
		expected := fmt.Sprintf(watchdog.LogDoneMessage, service)

		if check := (&Formatter{}).Done(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})

	t.Run("format message from env", func(t *testing.T) {
		prev := watchdog.LogDoneMessage
		message := "test message : %s"
		watchdog.LogDoneMessage = message
		defer func() { watchdog.LogDoneMessage = prev }()

		service := "service name"
		expected := fmt.Sprintf(watchdog.LogDoneMessage, service)

		if check := (&Formatter{}).Done(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})
}
