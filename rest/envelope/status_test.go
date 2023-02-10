package envelope

import (
	"reflect"
	"testing"
)

func Test_NewStatus(t *testing.T) {
	t.Run("construct", func(t *testing.T) {
		s := NewStatus()

		if check := s.Success; check != true {
			t.Error("initialized the status field as false")
		} else if len(s.Errors) != 0 {
			t.Error("initialized the error list with some elements")
		}
	})
}

func Test_Status_AddError(t *testing.T) {
	t.Run("add single error", func(t *testing.T) {
		err := NewStatusError(123, "error message")
		s := NewStatus().AddError(err)

		if s.Success != false {
			t.Error("didn't assign the status as false")
		} else if len(s.Errors) != 1 {
			t.Error("didn't stored the inserted error")
		} else if check := s.Errors[0]; !reflect.DeepEqual(err, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err)
		}
	})

	t.Run("add multiple error", func(t *testing.T) {
		err1 := NewStatusError(123, "error message 1")
		err2 := NewStatusError(456, "error message 2")
		err3 := NewStatusError(789, "error message 3")
		s := NewStatus().AddError(err1).AddError(err2).AddError(err3)

		if s.Success != false {
			t.Error("didn't assign the status as false")
		} else if len(s.Errors) != 3 {
			t.Error("didn't stored the inserted error")
		} else if check := s.Errors[0]; !reflect.DeepEqual(err1, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err1)
		} else if check := s.Errors[1]; !reflect.DeepEqual(err2, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err2)
		} else if check := s.Errors[2]; !reflect.DeepEqual(err3, check) {
			t.Errorf("the stored error (%v) differs from the inserted (%v) error", check, err3)
		}
	})
}

func Test_Status_SetService(t *testing.T) {
	t.Run("assign to all stored error", func(t *testing.T) {
		service := 147
		err1 := NewStatusError(123, "error message 1")
		err2 := NewStatusError(456, "error message 2")
		err3 := NewStatusError(789, "error message 3")
		s := NewStatus().AddError(err1).AddError(err2).AddError(err3)
		s = s.SetService(service)

		if check := s.Errors[0]; check.Service != service {
			t.Errorf("the stored service (%v) differs from the expected (%v)", check.Service, service)
		} else if check := s.Errors[1]; check.Service != service {
			t.Errorf("the stored service (%v) differs from the expected (%v)", check.Service, service)
		} else if check := s.Errors[2]; check.Service != service {
			t.Errorf("the stored service (%v) differs from the expected (%v)", check.Service, service)
		}
	})
}

func Test_Status_SetEndpoint(t *testing.T) {
	t.Run("assign to all stored error", func(t *testing.T) {
		endpoint := 147
		err1 := NewStatusError(123, "error message 1")
		err2 := NewStatusError(456, "error message 2")
		err3 := NewStatusError(789, "error message 3")
		s := NewStatus().AddError(err1).AddError(err2).AddError(err3)
		s = s.SetEndpoint(endpoint)

		if check := s.Errors[0]; check.Endpoint != endpoint {
			t.Errorf("the stored endpoint (%v) differs from the expected (%v)", check.Endpoint, endpoint)
		} else if check := s.Errors[1]; check.Endpoint != endpoint {
			t.Errorf("the stored endpoint (%v) differs from the expected (%v)", check.Endpoint, endpoint)
		} else if check := s.Errors[2]; check.Endpoint != endpoint {
			t.Errorf("the stored endpoint (%v) differs from the expected (%v)", check.Endpoint, endpoint)
		}
	})
}
