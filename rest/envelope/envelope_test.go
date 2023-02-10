package envelope

import (
	"reflect"
	"testing"
)

func Test_NewEnvelope(t *testing.T) {
	t.Run("construct without list report", func(t *testing.T) {
		statusCode := 123
		data := "message"
		env := NewEnvelope(statusCode, data)

		if check := env.Status.Success; check != true {
			t.Error("initialized the status field as false")
		} else if len(env.Status.Errors) != 0 {
			t.Error("initialized the error list with some elements")
		} else if check := env.Data; !reflect.DeepEqual(check, data) {
			t.Errorf("stored the (%v) value in data field instead of expected (%v)", check, data)
		} else if env.ListReport != nil {
			t.Errorf("stored the unexpected (%v) list report", env.ListReport)
		}
	})

	t.Run("construct with nil list report", func(t *testing.T) {
		statusCode := 123
		data := "message"
		env := NewEnvelope(statusCode, data, nil)

		if check := env.Status.Success; check != true {
			t.Error("initialized the status field as false")
		} else if len(env.Status.Errors) != 0 {
			t.Error("initialized the error list with some elements")
		} else if check := env.Data; !reflect.DeepEqual(check, data) {
			t.Errorf("stored the (%v) value in data field instead of expected (%v)", check, data)
		} else if env.ListReport != nil {
			t.Errorf("stored the unexpected (%v) list report", env.ListReport)
		}
	})

	t.Run("construct with list report", func(t *testing.T) {
		statusCode := 123
		data := "message"
		report := NewListReport("search", 1, 2, 10)
		env := NewEnvelope(statusCode, data, report)

		if check := env.Status.Success; check != true {
			t.Error("initialized the status field as false")
		} else if len(env.Status.Errors) != 0 {
			t.Error("initialized the error list with some elements")
		} else if check := env.Data; !reflect.DeepEqual(check, data) {
			t.Errorf("stored the (%v) value in data field instead of expected (%v)", check, data)
		} else if check := env.ListReport; !reflect.DeepEqual(check, report) {
			t.Errorf("stored the (%v) value in report field instead of expected (%v)", check, report)
		}
	})
}

func Test_Envelope_GetStatusCode(t *testing.T) {
	t.Run("return stored value", func(t *testing.T) {
		expected := 123
		if check := NewEnvelope(expected, nil).GetStatusCode(); check != expected {
			t.Errorf("returned the (%v) status code when expecting (%v)", check, expected)
		}
	})
}

func Test_Envelope_AddError(t *testing.T) {
	statusCode := 123
	data := "message"

	t.Run("add single error", func(t *testing.T) {
		err := NewStatusError(123, "error message")
		env := NewEnvelope(statusCode, data).AddError(err)

		if env.Status.Success != false {
			t.Error("didn't assign the status as false")
		} else if len(env.Status.Errors) != 1 {
			t.Error("didn't stored the inserted error")
		} else if check := env.Status.Errors[0]; !reflect.DeepEqual(err, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err)
		}
	})

	t.Run("add multiple error", func(t *testing.T) {
		err1 := NewStatusError(123, "error message 1")
		err2 := NewStatusError(456, "error message 2")
		err3 := NewStatusError(789, "error message 3")
		env := NewEnvelope(statusCode, data).AddError(err1).AddError(err2).AddError(err3)

		if env.Status.Success != false {
			t.Error("didn't assign the status as false")
		} else if len(env.Status.Errors) != 3 {
			t.Error("didn't stored the inserted error")
		} else if check := env.Status.Errors[0]; !reflect.DeepEqual(err1, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err1)
		} else if check := env.Status.Errors[1]; !reflect.DeepEqual(err2, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err2)
		} else if check := env.Status.Errors[2]; !reflect.DeepEqual(err3, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err3)
		}
	})
}

func Test_Envelope_SetService(t *testing.T) {
	t.Run("assign to all stored error", func(t *testing.T) {
		service := 147
		statusCode := 123
		data := "message"
		err1 := NewStatusError(123, "error message 1")
		err2 := NewStatusError(456, "error message 2")
		err3 := NewStatusError(789, "error message 3")
		env := NewEnvelope(statusCode, data).AddError(err1).AddError(err2).AddError(err3)
		env = env.SetService(service)

		if check := env.Status.Errors[0]; check.Service != service {
			t.Errorf("the stored service (%v) differs from the expected (%v)", check.Service, service)
		} else if check := env.Status.Errors[1]; check.Service != service {
			t.Errorf("the stored service (%v) differs from the expected (%v)", check.Service, service)
		} else if check := env.Status.Errors[2]; check.Service != service {
			t.Errorf("the stored service (%v) differs from the expected (%v)", check.Service, service)
		}
	})
}

func Test_Envelope_SetEndpoint(t *testing.T) {
	t.Run("assign to all stored error", func(t *testing.T) {
		endpoint := 147
		statusCode := 123
		data := "message"
		err1 := NewStatusError(123, "error message 1")
		err2 := NewStatusError(456, "error message 2")
		err3 := NewStatusError(789, "error message 3")
		env := NewEnvelope(statusCode, data).AddError(err1).AddError(err2).AddError(err3)
		env = env.SetEndpoint(endpoint)

		if check := env.Status.Errors[0]; check.Endpoint != endpoint {
			t.Errorf("the stored endpoint (%v) differs from the expected (%v)", check.Endpoint, endpoint)
		} else if check := env.Status.Errors[1]; check.Endpoint != endpoint {
			t.Errorf("the stored endpoint (%v) differs from the expected (%v)", check.Endpoint, endpoint)
		} else if check := env.Status.Errors[2]; check.Endpoint != endpoint {
			t.Errorf("the stored endpoint (%v) differs from the expected (%v)", check.Endpoint, endpoint)
		}
	})
}

func Test_Envelope_SetListReport(t *testing.T) {
	t.Run("assign the list report", func(t *testing.T) {
		endpoint := 147
		statusCode := 123
		data := "message"
		report := NewListReport("search", 1, 2, 10)
		env := NewEnvelope(statusCode, data)
		env = env.SetEndpoint(endpoint)
		env.SetListReport(report)

		if check := env.Status.Success; check != true {
			t.Error("initialized the status field as false")
		} else if len(env.Status.Errors) != 0 {
			t.Error("initialized the error list with some elements")
		} else if check := env.Data; !reflect.DeepEqual(check, data) {
			t.Errorf("stored the (%v) value in data field instead of expected (%v)", check, data)
		} else if check := env.ListReport; !reflect.DeepEqual(check, report) {
			t.Errorf("stored the (%v) value in report field instead of expected (%v)", check, report)
		}
	})

	t.Run("assign/remove the list report if nil given", func(t *testing.T) {
		endpoint := 147
		statusCode := 123
		data := "message"
		report := NewListReport("search", 1, 2, 10)
		env := NewEnvelope(statusCode, data, report)
		env = env.SetEndpoint(endpoint)
		env.SetListReport(nil)

		if check := env.Status.Success; check != true {
			t.Error("initialized the status field as false")
		} else if len(env.Status.Errors) != 0 {
			t.Error("initialized the error list with some elements")
		} else if check := env.Data; !reflect.DeepEqual(check, data) {
			t.Errorf("stored the (%v) value in data field instead of expected (%v)", check, data)
		} else if env.ListReport != nil {
			t.Errorf("stored the unexpected (%v) list report", env.ListReport)
		}
	})
}
