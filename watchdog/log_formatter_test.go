package watchdog

import (
	"fmt"
	"testing"
)

func Test_LogFormatter_Start(t *testing.T) {
	t.Run("format message", func(t *testing.T) {
		service := "service name"
		expected := fmt.Sprintf(LogStartMessage, service)

		if check := (&LogFormatter{}).Start(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})

	t.Run("format message from env", func(t *testing.T) {
		prev := LogStartMessage
		message := "test message : %s"
		LogStartMessage = message
		defer func() { LogStartMessage = prev }()

		service := "service name"
		expected := fmt.Sprintf(LogStartMessage, service)

		if check := (&LogFormatter{}).Start(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})
}

func Test_LogFormatter_Error(t *testing.T) {
	t.Run("format message", func(t *testing.T) {
		service := "service name"
		err := fmt.Errorf("error message")
		expected := fmt.Sprintf(LogErrorMessage, service, err)

		if check := (&LogFormatter{}).Error(service, err); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})

	t.Run("format message from env", func(t *testing.T) {
		prev := LogErrorMessage
		message := "test message : %s - %e"
		LogErrorMessage = message
		defer func() { LogErrorMessage = prev }()

		service := "service name"
		err := fmt.Errorf("error message")
		expected := fmt.Sprintf(LogErrorMessage, service, err)

		if check := (&LogFormatter{}).Error(service, err); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})
}

func Test_LogFormatter_Done(t *testing.T) {
	t.Run("format message", func(t *testing.T) {
		service := "service name"
		expected := fmt.Sprintf(LogDoneMessage, service)

		if check := (&LogFormatter{}).Done(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})

	t.Run("format message from env", func(t *testing.T) {
		prev := LogDoneMessage
		message := "test message : %s"
		LogDoneMessage = message
		defer func() { LogDoneMessage = prev }()

		service := "service name"
		expected := fmt.Sprintf(LogDoneMessage, service)

		if check := (&LogFormatter{}).Done(service); check != expected {
			t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
		}
	})
}
