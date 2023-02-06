package validation

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/rest/envelope"
)

func Test_NewValidator(t *testing.T) {
	t.Run("nil validate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check, e := NewValidator(nil, NewMockParser(ctrl))
		switch {
		case check != nil:
			t.Error("return an unexpected valid validator instance")
		case e == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil parser", func(t *testing.T) {
		check, e := NewValidator(validator.New(), nil)
		switch {
		case check != nil:
			t.Error("return an unexpected valid validator instance")
		case e == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("construct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if check, e := NewValidator(validator.New(), NewMockParser(ctrl)); e != nil {
			t.Errorf("return the unexpected error (%v)", e)
		} else if check == nil {
			t.Error("didn't return the expected validation instance")
		}
	})
}

func Test_Validator_Call(t *testing.T) {
	t.Run("nil data", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		instance, _ := NewValidator(validator.New(), NewMockParser(ctrl))
		expected := errNilPointer("value")

		env, e := instance(nil)
		switch {
		case env != nil:
			t.Errorf("return the unexpected envelope (%v)", env)
		case e == nil:
			t.Error("didn't return an expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("no errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field1 int `validate:"gt=0,lte=10" vparam:"1"`
			Field2 int `validate:"gt=10,lte=20" vparam:"2"`
		}{Field1: 1, Field2: 11}
		instance, _ := NewValidator(validator.New(), NewMockParser(ctrl))

		if env, e := instance(data); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if env != nil {
			t.Errorf("returned the unexpected envelope (%v)", env)
		}
	})

	t.Run("error parsing errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field1 int `validate:"gt=0,lte=10" vparam:"1"`
			Field2 int `validate:"gt=10,lte=20" vparam:"2"`
		}{Field1: 11, Field2: 11}
		expected := fmt.Errorf("error message")
		parser := NewMockParser(ctrl)
		parser.EXPECT().Parse(data, gomock.Any()).Return(nil, expected).Times(1)

		instance, _ := NewValidator(validator.New(), parser)

		resp, e := instance(data)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		case resp != nil:
			t.Error("returned an unexpected instance of the response envelope")
		}
	})

	t.Run("parse errors", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field1 int `validate:"gt=0,lte=10" vparam:"1"`
			Field2 int `validate:"gt=10,lte=20" vparam:"2"`
		}{Field1: 11, Field2: 11}
		expected := envelope.NewEnvelope(http.StatusBadRequest, nil, nil)
		parser := NewMockParser(ctrl)
		parser.EXPECT().Parse(data, gomock.Any()).Return(expected, nil).Times(1)

		instance, _ := NewValidator(validator.New(), parser)

		if resp, e := instance(data); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if !reflect.DeepEqual(resp, expected) {
			t.Errorf("returned the (%v) envelope instead of the expected (%v)", resp, expected)
		}
	})
}
