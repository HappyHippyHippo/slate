package srvalidation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/happyhippyhippo/slate/srest/srenvelope"
	"net/http"
	"reflect"
	"testing"
)

func Test_NewValidator(t *testing.T) {
	t.Run("nil validate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		check, err := NewValidator(nil, NewMockParser(ctrl))
		switch {
		case check != nil:
			t.Error("return an unexpected valid validator instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil parser", func(t *testing.T) {
		check, err := NewValidator(validator.New(), nil)
		switch {
		case check != nil:
			t.Error("return an unexpected valid validator instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("construct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if check, err := NewValidator(validator.New(), NewMockParser(ctrl)); err != nil {
			t.Errorf("return the unexpected error (%v)", err)
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

		env, err := instance(nil)
		switch {
		case env != nil:
			t.Errorf("return the unexpected envelope (%v)", env)
		case err == nil:
			t.Error("didn't return an expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
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

		if env, err := instance(data); err != nil {
			t.Errorf("returned the unexpected error (%v)", err)
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

		resp, err := instance(data)
		switch {
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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
		expected := srenvelope.NewEnvelope(http.StatusBadRequest, nil, nil)
		parser := NewMockParser(ctrl)
		parser.EXPECT().Parse(data, gomock.Any()).Return(expected, nil).Times(1)

		instance, _ := NewValidator(validator.New(), parser)

		if resp, err := instance(data); err != nil {
			t.Errorf("returned the unexpected error (%v)", err)
		} else if !reflect.DeepEqual(resp, expected) {
			t.Errorf("returned the (%v) envelope instead of the expected (%v)", resp, expected)
		}
	})
}
