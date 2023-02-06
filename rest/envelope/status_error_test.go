package envelope

import (
	"fmt"
	"testing"
)

func Test_NewStatusError(t *testing.T) {
	t.Run("construct", func(t *testing.T) {
		code := 123
		msg := "message"
		e := NewStatusError(code, msg)
		if check := e.Service; check != 0 {
			t.Errorf("stored the (%v) service value instead of zero", check)
		} else if check := e.Endpoint; check != 0 {
			t.Errorf("stored the (%v) endpoint value instead of zero", check)
		} else if check := e.Param; check != 0 {
			t.Errorf("stored the (%v) param value instead of zero", check)
		} else if check := e.Error; check != fmt.Sprintf("%d", code) {
			t.Errorf("stored the (%v) error instead of (%v)", check, code)
		} else if check := e.Code; check != fmt.Sprintf("c:%d", code) {
			t.Errorf("stored the (%v) code instead of (%v)", check, fmt.Sprintf("c:%d", code))
		} else if check := e.Message; check != msg {
			t.Errorf("stored the (%v) message instead of (%v)", check, msg)
		}
	})

	t.Run("construct with string error code", func(t *testing.T) {
		code := "error code"
		msg := "message"
		e := NewStatusError(code, msg)
		if check := e.Service; check != 0 {
			t.Errorf("stored the (%v) service value instead of zero", check)
		} else if check := e.Endpoint; check != 0 {
			t.Errorf("stored the (%v) endpoint value instead of zero", check)
		} else if check := e.Param; check != 0 {
			t.Errorf("stored the (%v) param value instead of zero", check)
		} else if check := e.Error; check != code {
			t.Errorf("stored the (%v) error instead of (%v)", check, code)
		} else if check := e.Code; check != code {
			t.Errorf("stored the (%v) code instead of (%v)", check, code)
		} else if check := e.Message; check != msg {
			t.Errorf("stored the (%v) message instead of (%v)", check, msg)
		}
	})
}

func Test_StatusError_SetService(t *testing.T) {
	t.Run("assign", func(t *testing.T) {
		service := 123
		code := 456
		msg := "message"
		expected := fmt.Sprintf("s:%d.c:%d", service, code)
		err1 := NewStatusError(code, msg)
		err2 := err1.SetService(service)

		if check := err2.Service; check != service {
			t.Errorf("stored the (%v) service value instead of expected(%v)", check, service)
		} else if check := err2.Code; check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})
}

func Test_StatusError_SetEndpoint(t *testing.T) {
	t.Run("assign", func(t *testing.T) {
		endpoint := 123
		code := 456
		msg := "message"
		expected := fmt.Sprintf("e:%d.c:%d", endpoint, code)
		err1 := NewStatusError(code, msg)
		err2 := err1.SetEndpoint(endpoint)

		if check := err2.Endpoint; check != endpoint {
			t.Errorf("stored the (%v) service value instead of expected(%v)", check, endpoint)
		} else if check := err2.Code; check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})
}

func Test_StatusError_SetParam(t *testing.T) {
	t.Run("assign", func(t *testing.T) {
		param := 123
		code := 456
		msg := "message"
		expected := fmt.Sprintf("p:%d.c:%d", param, code)
		err1 := NewStatusError(code, msg)
		err2 := err1.SetParam(param)

		if check := err2.Param; check != param {
			t.Errorf("stored the (%v) param value instead of expected(%v)", check, param)
		} else if check := err2.Code; check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})
}

func Test_StatusError_SetError(t *testing.T) {
	t.Run("assign", func(t *testing.T) {
		newCode := 123
		code := 456
		expected := fmt.Sprintf("c:%d", newCode)
		msg := "message"
		err1 := NewStatusError(code, msg)
		err2 := err1.SetError(newCode)

		if check := err2.Error; check != fmt.Sprintf("%d", newCode) {
			t.Errorf("stored the (%v) error value instead of expected(%v)", check, newCode)
		} else if check := err2.Code; check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})

	t.Run("assign with string error code", func(t *testing.T) {
		code := 456
		newCode := "new error code"
		msg := "message"
		err1 := NewStatusError(code, msg)
		err2 := err1.SetError(newCode)

		if check := err2.Error; check != newCode {
			t.Errorf("stored the (%v) error value instead of expected(%v)", check, newCode)
		} else if check := err2.Code; check != newCode {
			t.Errorf("stored the (%v) code instead of (%v)", check, newCode)
		}
	})
}

func Test_StatusError_SetMessage(t *testing.T) {
	t.Run("assign", func(t *testing.T) {
		code := 456
		msg := "message"
		newMsg := "new message"
		err1 := NewStatusError(code, msg)
		err2 := err1.SetMessage(newMsg)

		if check := err2.Message; check != newMsg {
			t.Errorf("stored the (%v) message instead of expected(%v)", check, newMsg)
		}
	})
}

func Test_StatusError_GetCode(t *testing.T) {
	t.Run("retrieval", func(t *testing.T) {
		service := 12
		endpoint := 34
		param := 56
		code := 78
		expected := fmt.Sprintf("s:%d.e:%d.p:%d.c:%d", service, endpoint, param, code)
		msg := "message"
		e := NewStatusError(1, msg).SetService(service).SetEndpoint(endpoint).SetParam(param).SetError(code)

		if check := e.GetCode(); check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})
}

func Test_StatusError_GetMessage(t *testing.T) {
	t.Run("retrieval", func(t *testing.T) {
		msg := "message"
		e := NewStatusError(123, msg)

		if check := e.GetMessage(); check != msg {
			t.Errorf("stored the (%v) message value instead of expected(%v)", check, msg)
		}
	})
}

func Test_StatusError(t *testing.T) {
	t.Run("assign param error", func(t *testing.T) {
		service := 12
		endpoint := 34
		param := 56
		code := 78
		expected := fmt.Sprintf("s:%d.e:%d.p:%d.c:%d", service, endpoint, param, code)
		originalCode := 1
		msg := "message"
		err1 := NewStatusError(originalCode, msg)
		err2 := err1.SetService(service).SetEndpoint(endpoint).SetParam(param).SetError(code)

		if check := err2.Service; check != service {
			t.Errorf("stored the (%v) service value instead of expected(%v)", check, service)
		} else if check := err2.Endpoint; check != endpoint {
			t.Errorf("stored the (%v) endpoint value instead of expected(%v)", check, endpoint)
		} else if check := err2.Param; check != param {
			t.Errorf("stored the (%v) param value instead of expected(%v)", check, param)
		} else if check := err2.Error; check != fmt.Sprintf("%d", code) {
			t.Errorf("stored the (%v) code value instead of expected(%v)", check, code)
		} else if check := err2.Code; check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})

	t.Run("assign with string error code without param", func(t *testing.T) {
		service := 12
		endpoint := 34
		code := "error code"
		expected := fmt.Sprintf("s:%d.e:%d.%s", service, endpoint, code)
		originalCode := 1
		msg := "message"
		err1 := NewStatusError(originalCode, msg)
		err2 := err1.SetService(service).SetEndpoint(endpoint).SetError(code)

		if check := err2.Service; check != service {
			t.Errorf("stored the (%v) service value instead of expected(%v)", check, service)
		} else if check := err2.Endpoint; check != endpoint {
			t.Errorf("stored the (%v) endpoint value instead of expected(%v)", check, endpoint)
		} else if check := err2.Param; check != 0 {
			t.Errorf("stored the (%v) param value instead of expected(%v)", check, 0)
		} else if check := err2.Error; check != code {
			t.Errorf("stored the (%v) code value instead of expected(%v)", check, code)
		} else if check := err2.Code; check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})

	t.Run("assign with string error code with param", func(t *testing.T) {
		service := 12
		endpoint := 34
		param := 56
		code := "error code"
		expected := fmt.Sprintf("s:%d.e:%d.p:%d.%s", service, endpoint, param, code)
		originalCode := 1
		msg := "message"
		err1 := NewStatusError(originalCode, msg)
		err2 := err1.SetService(service).SetEndpoint(endpoint).SetParam(param).SetError(code)

		if check := err2.Service; check != service {
			t.Errorf("stored the (%v) service value instead of expected(%v)", check, service)
		} else if check := err2.Endpoint; check != endpoint {
			t.Errorf("stored the (%v) endpoint value instead of expected(%v)", check, endpoint)
		} else if check := err2.Param; check != param {
			t.Errorf("stored the (%v) param value instead of expected(%v)", check, param)
		} else if check := err2.Error; check != code {
			t.Errorf("stored the (%v) code value instead of expected(%v)", check, code)
		} else if check := err2.Code; check != expected {
			t.Errorf("stored the (%v) code instead of (%v)", check, expected)
		}
	})
}
