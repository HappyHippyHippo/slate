package srlog

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_NewResponseWriter(t *testing.T) {
	t.Run("error when missing writer", func(t *testing.T) {
		expected := errNilPointer("writer")

		value, err := newResponseWriter(nil)
		switch {
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		case value != nil:
			t.Error("returned a valid reference")
		}
	})

	t.Run("new log response writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		writer := NewMockResponseWriter(ctrl)

		if value, err := newResponseWriter(writer); err != nil {
			t.Errorf("return the (%v) error", err)
		} else if value == nil {
			t.Error("didn't returned a valid reference")
		}
	})
}

func Test_ResponseWriter_Write(t *testing.T) {
	t.Run("write to buffer and underlying writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		b := []byte{12, 34, 56}
		ginWriter := NewMockResponseWriter(ctrl)
		ginWriter.EXPECT().Write(b).Times(1)
		writer := &writer{body: &bytes.Buffer{}, ResponseWriter: ginWriter}
		_, _ = writer.Write(b)

		if !reflect.DeepEqual(writer.body.Bytes(), b) {
			t.Errorf("written (%v) bytes on buffer", writer.body)
		}
	})
}

func Test_ResponseWriter_Body(t *testing.T) {
	t.Run("write to buffer and underlying writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		b := []byte{12, 34, 56}
		ginWriter := NewMockResponseWriter(ctrl)
		writer := &writer{body: bytes.NewBuffer(b), ResponseWriter: ginWriter}

		if !reflect.DeepEqual(writer.Body(), b) {
			t.Errorf("written (%v) bytes on buffer", writer.body)
		}
	})
}
