package validation

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_NewParser(t *testing.T) {
	t.Run("nil translator", func(t *testing.T) {
		parser, e := NewParser(nil)
		switch {
		case parser != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new parser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		translator := NewMockTranslator(ctrl)

		p, e := NewParser(translator)
		switch {
		case p == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("return the (%v) error", e)
		case p.(*parser).translator != translator:
			t.Error("didn't stored the translator reference")
		}
	})
}

func Test_Parser_Parse(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		translator := NewMockTranslator(ctrl)
		sut, _ := NewParser(translator)

		resp, e := sut.Parse(nil, []validator.FieldError{})
		switch {
		case resp != nil:
			t.Error("returned an unexpectedly valid instance of a response")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Error("returned the error is not of the expected nil pointer error")
		}
	})

	t.Run("no-op on nil error list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := struct{ Message string }{Message: "message"}
		translator := NewMockTranslator(ctrl)
		sut, _ := NewParser(translator)

		resp, e := sut.Parse(value, nil)
		switch {
		case e != nil:
			t.Errorf("returned the '%v' unexpected error", e)
		case resp != nil:
			t.Error("returned an unexpectedly valid instance of a response")
		}
	})

	t.Run("no-op if error list is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := struct{ Message string }{Message: "message"}
		translator := NewMockTranslator(ctrl)
		sut, _ := NewParser(translator)

		resp, e := sut.Parse(value, []validator.FieldError{})
		switch {
		case e != nil:
			t.Errorf("returned the '%v' unexpected error", e)
		case resp != nil:
			t.Error("returned an unexpectedly valid instance of a response")
		}
	})

	t.Run("invalid nil field error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field string `vparam:"string"`
		}{Field: "message"}
		translator := NewMockTranslator(ctrl)

		sut, _ := NewParser(translator)

		resp, e := sut.Parse(data, []validator.FieldError{nil})
		switch {
		case resp != nil:
			t.Error("returned an unexpectedly valid instance of a response")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Error("returned the error is not of the expected nil pointer error")
		}
	})

	t.Run("error retrieving field/param value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field string `vparam:"string"`
		}{Field: "message"}
		expected := `strconv.Atoi: parsing "string": invalid syntax`
		translator := NewMockTranslator(ctrl)
		fieldError := NewMockFieldError(ctrl)
		fieldError.EXPECT().StructField().Return("Field").Times(1)

		sut, _ := NewParser(translator)

		resp, e := sut.Parse(data, []validator.FieldError{fieldError})
		switch {
		case resp != nil:
			t.Error("returned an unexpectedly valid instance of a response")
		case e == nil:
			t.Error("didn't returned an expected error")
		case e.Error() != expected:
			t.Error("returned the error is not of the expected nil pointer error")
		}
	})

	t.Run("generating error for a non-tagged field", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field int `validate:"gt=0"`
		}{Field: 0}
		errMsg := "error message"
		expected := "c:89"
		translator := NewMockTranslator(ctrl)
		fieldError := NewMockFieldError(ctrl)
		fieldError.EXPECT().StructField().Return("Field").Times(1)
		fieldError.EXPECT().Translate(translator).Return(errMsg).Times(1)
		fieldError.EXPECT().Tag().Return("gt").Times(1)

		sut, _ := NewParser(translator)

		resp, e := sut.Parse(data, []validator.FieldError{fieldError})
		switch {
		case e != nil:
			t.Errorf("return the unexpected error (%v)", e)
		case len(resp.Status.Errors) == 0:
			t.Errorf("didn't stored the expected error on the response envelope")
		case resp.Status.Errors[0].GetCode() != expected:
			t.Errorf("returned the (%v) error code instead of the expected (%s)", resp.Status.Errors[0].GetCode(), expected)
		case resp.Status.Errors[0].GetMessage() != errMsg:
			t.Errorf("returned the (%v) error message instead of the expected (%v)", resp.Status.Errors[0].GetMessage(), errMsg)
		}
	})

	t.Run("generating error for an unrecognized error tag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field int `validate:"gt=0"`
		}{Field: 0}
		errMsg := "error message"
		expected := "c:0"
		translator := NewMockTranslator(ctrl)
		fieldError := NewMockFieldError(ctrl)
		fieldError.EXPECT().StructField().Return("Field").Times(1)
		fieldError.EXPECT().Translate(translator).Return(errMsg).Times(1)
		fieldError.EXPECT().Tag().Return("unrecognized").Times(1)

		sut, _ := NewParser(translator)

		resp, e := sut.Parse(data, []validator.FieldError{fieldError})
		switch {
		case e != nil:
			t.Errorf("return the unexpected error (%v)", e)
		case resp.Status.Errors[0].GetCode() != expected:
			t.Errorf("returned the (%v) error code instead of the expected (%s)", resp.Status.Errors[0].GetCode(), expected)
		case resp.Status.Errors[0].GetMessage() != errMsg:
			t.Errorf("returned the (%v) error message instead of the expected (%v)", resp.Status.Errors[0].GetMessage(), errMsg)
		}
	})

	t.Run("generating error with all information", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := struct {
			Field int `validate:"gt=0" vparam:"10"`
		}{Field: 0}
		expected := "p:10.c:89"
		errMsg := "error message"
		translator := NewMockTranslator(ctrl)
		fieldError := NewMockFieldError(ctrl)
		fieldError.EXPECT().StructField().Return("Field").Times(1)
		fieldError.EXPECT().Translate(translator).Return(errMsg).Times(1)
		fieldError.EXPECT().Tag().Return("gt").Times(1)

		sut, _ := NewParser(translator)

		resp, e := sut.Parse(data, []validator.FieldError{fieldError})
		switch {
		case e != nil:
			t.Errorf("return the unexpected error (%v)", e)
		case resp.Status.Errors[0].GetCode() != expected:
			t.Errorf("returned the (%v) error code instead of the expected (%s)", resp.Status.Errors[0].GetCode(), expected)
		case resp.Status.Errors[0].GetMessage() != errMsg:
			t.Errorf("returned the (%v) error message instead of the expected (%v)", resp.Status.Errors[0].GetMessage(), errMsg)
		}
	})
}

func Test_Parser_AddError(t *testing.T) {
	t.Run("adding a new error mapping value", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mappedErrorName := "unrecognized"
		mappedErrorCode := 10000
		expected := fmt.Sprintf("c:%d", mappedErrorCode)
		data := struct {
			Field int `validate:"gt=0"`
		}{Field: 0}
		errMsg := "error message"
		translator := NewMockTranslator(ctrl)
		fieldError := NewMockFieldError(ctrl)
		fieldError.EXPECT().StructField().Return("Field").Times(1)
		fieldError.EXPECT().Translate(translator).Return(errMsg).Times(1)
		fieldError.EXPECT().Tag().Return(mappedErrorName).Times(1)

		sut, _ := NewParser(translator)
		sut.AddError(mappedErrorName, mappedErrorCode)

		resp, e := sut.Parse(data, []validator.FieldError{fieldError})
		switch {
		case e != nil:
			t.Errorf("return the unexpected error (%v)", e)
		case resp.Status.Errors[0].GetCode() != expected:
			t.Errorf("returned the (%v) error code instead of the expected (%s)", resp.Status.Errors[0].GetCode(), expected)
		case resp.Status.Errors[0].GetMessage() != errMsg:
			t.Errorf("returned the (%v) error message instead of the expected (%v)", resp.Status.Errors[0].GetMessage(), errMsg)
		}
	})
}
