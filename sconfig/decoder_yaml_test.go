package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"strings"
	"testing"
)

func Test_NewDecoderYAML(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		sut, err := newDecoderYAML(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new yaml decoder adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)

		if sut, err := newDecoderYAML(reader); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else {
			defer func() { _ = sut.Close() }()
			if err != nil {
				t.Errorf("returned the (%v) error", err)
			} else if sut.(*decoderYAML).reader != reader {
				t.Error("didn't store the reader reference")
			}
		}
	})
}

func Test_DecoderYAML_Close(t *testing.T) {
	t.Run("error while closing the reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Return(expected).Times(1)
		sut, _ := newDecoderYAML(reader)

		if err := sut.Close(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("call close method on reader only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := newDecoderYAML(reader)

		_ = sut.Close()
		_ = sut.Close()
	})
}

func Test_DecoderYAML_Decode(t *testing.T) {
	t.Run("return decode error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := newDecoderYAML(reader)
		defer func() { _ = sut.Close() }()
		yaml := NewMockYamler(ctrl)
		yaml.EXPECT().Decode(&Partial{}).DoAndReturn(func(p *Partial) error {
			return expected
		}).Times(1)
		sut.(*decoderYAML).decoder = yaml

		check, err := sut.Decode()
		switch {
		case check != nil:
			t.Error("returned an reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("redirect to the underlying decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{"node": "data"}
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := newDecoderYAML(reader)
		defer func() { _ = sut.Close() }()
		yaml := NewMockYamler(ctrl)
		yaml.EXPECT().Decode(&Partial{}).DoAndReturn(func(p *Partial) error {
			(*p)["node"] = data["node"]
			return nil
		}).Times(1)
		sut.(*decoderYAML).decoder = yaml

		check, err := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		case !reflect.DeepEqual(*check.(*Partial), data):
			t.Errorf("returned (%v)", check)
		}
	})

	t.Run("decode yaml string", func(t *testing.T) {
		yaml := "node:\n  subnode: data"
		expected := Partial{"node": Partial{"subnode": "data"}}
		reader := strings.NewReader(yaml)
		sut, _ := newDecoderYAML(reader)
		defer func() { _ = sut.Close() }()

		check, err := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		case !reflect.DeepEqual(*check.(*Partial), expected):
			t.Errorf("returned (%v)", check)
		}
	})
}
