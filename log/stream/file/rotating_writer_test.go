package file

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_NewRotatingFileWriter(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		sut, e := NewRotatingWriter(nil, "path")
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error opening file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)

		sut, e := NewRotatingWriter(fileSystem, path)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("create valid writer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file := NewMockFile(ctrl)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)

		if sut, e := NewRotatingWriter(fileSystem, path); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if sut == nil {
			t.Error("didn't returned the expected writer reference")
		}
	})
}

func Test_RotatingFileWriter_Write(t *testing.T) {
	t.Run("error while writing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		output := []byte("message")
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file := NewMockFile(ctrl)
		file.EXPECT().Write(output).Return(0, expected).Times(1)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)

		sut, _ := NewRotatingWriter(fileSystem, path)

		if _, e := sut.Write(output); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("no rotation if write done in same day of opened file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		output := []byte("message")
		count := 123
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file := NewMockFile(ctrl)
		file.EXPECT().Write(output).Return(count, nil).Times(1)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)

		sut, _ := NewRotatingWriter(fileSystem, path)

		if written, e := sut.Write(output); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if written != count {
			t.Errorf("returned the unexpected number of written elements of (%v) when expecting (%v)", written, count)
		}
	})

	t.Run("error while closing rotated file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		output := []byte("message")
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(expected)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)

		sut, _ := NewRotatingWriter(fileSystem, path)
		sut.(*RotatingWriter).day++

		if _, e := sut.Write(output); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error while opening rotating file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		output := []byte("message")
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(nil)
		fileSystem := NewMockFs(ctrl)
		gomock.InOrder(
			fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil),
			fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected),
		)

		sut, _ := NewRotatingWriter(fileSystem, path)
		sut.(*RotatingWriter).day++

		if _, e := sut.Write(output); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("rotate file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		output := []byte("message")
		count := 123
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file1 := NewMockFile(ctrl)
		file1.EXPECT().Close().Return(nil)
		file2 := NewMockFile(ctrl)
		file2.EXPECT().Write(output).Return(count, nil).Times(1)
		fileSystem := NewMockFs(ctrl)
		gomock.InOrder(
			fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file1, nil),
			fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file2, nil),
		)

		sut, _ := NewRotatingWriter(fileSystem, path)
		sut.(*RotatingWriter).day++

		if written, e := sut.Write(output); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if written != count {
			t.Errorf("returned the unexpected number of written elements of (%v) when expecting (%v)", written, count)
		}
	})
}

func Test_RotatingFileWriter_Close(t *testing.T) {
	t.Run("call close on the underlying file pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(expected)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)

		sut, _ := NewRotatingWriter(fileSystem, path)

		if e := sut.(io.Closer).Close(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})
}
