package validation

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/envelope"
)

func Test_NewValidator(t *testing.T) {
	t.Run("nil validate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		translator := NewMockTranslator(ctrl)
		parser, _ := NewParser(translator)

		check, e := NewValidator(nil, parser)
		switch {
		case check != nil:
			t.Error("return an unexpected valid validator instance")
		case e == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil parser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		translator := NewMockTranslator(ctrl)

		check, e := NewValidator(translator, nil)
		switch {
		case check != nil:
			t.Error("return an unexpected valid validator instance")
		case e == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("construct", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		translator := NewMockTranslator(ctrl)
		translator.
			EXPECT().
			Add(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		translator.
			EXPECT().
			AddCardinal(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		parser, _ := NewParser(translator)

		if check, e := NewValidator(translator, parser); e != nil {
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

		expected := errNilPointer("value")
		translator := NewMockTranslator(ctrl)
		translator.
			EXPECT().
			Add(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		translator.
			EXPECT().
			AddCardinal(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		parser, _ := NewParser(translator)
		sut, _ := NewValidator(translator, parser)

		env, e := sut(nil)
		switch {
		case env != nil:
			t.Errorf("return the unexpected envelope (%v)", env)
		case e == nil:
			t.Error("didn't return an expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("no error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field1 int `validate:"gt=0,lte=10" vparam:"1"`
			Field2 int `validate:"gt=10,lte=20" vparam:"2"`
		}{Field1: 1, Field2: 11}
		translator := NewMockTranslator(ctrl)
		translator.
			EXPECT().
			Add(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		translator.
			EXPECT().
			AddCardinal(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		parser, _ := NewParser(translator)
		sut, _ := NewValidator(translator, parser)

		if env, e := sut(data); e != nil {
			t.Errorf("unexpected (%v) error", e)
		} else if env != nil {
			t.Errorf("returned the unexpected envelope (%v)", env)
		}
	})

	t.Run("error parsing error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field1 int `validate:"gt=0,lte=10" vparam:"string"`
			Field2 int `validate:"gt=10,lte=20" vparam:"2"`
		}{Field1: 11, Field2: 11}
		expected := "strconv.Atoi: parsing \"string\": invalid syntax"
		translator := NewMockTranslator(ctrl)
		translator.
			EXPECT().
			Add(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		translator.
			EXPECT().
			AddCardinal(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		parser, _ := NewParser(translator)
		sut, _ := NewValidator(translator, parser)

		resp, e := sut(data)
		switch {
		case resp != nil:
			t.Error("unexpected instance of the response envelope")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected:
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("parse error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field1 int `validate:"gt=0,lte=10" vparam:"1"`
			Field2 int `validate:"gt=10,lte=20" vparam:"2"`
		}{Field1: 11, Field2: 11}
		errMsg := "error message"
		expected := envelope.NewEnvelope(http.StatusBadRequest, nil, nil)
		expected.AddError(envelope.NewStatusError(92, errMsg).SetParam(1))
		translator := NewMockTranslator(ctrl)
		translator.
			EXPECT().
			Add(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		translator.
			EXPECT().
			AddCardinal(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()
		translator.
			EXPECT().
			FmtNumber(float64(10), uint64(0)).
			Return("10").
			Times(1)
		translator.
			EXPECT().
			T("lte-number", "Field1", gomock.Any()).
			Return(errMsg, nil).
			Times(1)
		parser, _ := NewParser(translator)
		sut, _ := NewValidator(translator, parser)

		if resp, e := sut(data); e != nil {
			t.Errorf("unexpected (%v) error", e)
		} else if !reflect.DeepEqual(resp, expected) {
			t.Errorf("(%v) when expecting (%v)", resp, expected)
		}
	})
}
