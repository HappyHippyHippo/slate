package glog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gerror"
	"io"
	"os"
	"testing"
	"time"
)

func Test_NewStreamFileRotateWriter(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		writer, err := NewStreamRotatingFileWriter(nil, "path")
		switch {
		case writer != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
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

		writer, err := NewStreamRotatingFileWriter(fileSystem, path)
		switch {
		case writer != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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

		if writer, err := NewStreamRotatingFileWriter(fileSystem, path); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if writer == nil {
			t.Error("didn't returned the expected writer reference")
		}
	})
}

func Test_RotateFileWriter_Write(t *testing.T) {
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
		writer, _ := NewStreamRotatingFileWriter(fileSystem, path)

		if _, err := writer.Write(output); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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
		writer, _ := NewStreamRotatingFileWriter(fileSystem, path)

		if written, err := writer.Write(output); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
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
		writer, _ := NewStreamRotatingFileWriter(fileSystem, path)

		writer.(*streamRotatingFileWriter).day++

		if _, err := writer.Write(output); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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
		writer, _ := NewStreamRotatingFileWriter(fileSystem, path)

		writer.(*streamRotatingFileWriter).day++

		if _, err := writer.Write(output); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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
		writer, _ := NewStreamRotatingFileWriter(fileSystem, path)

		writer.(*streamRotatingFileWriter).day++

		if written, err := writer.Write(output); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if written != count {
			t.Errorf("returned the unexpected number of written elements of (%v) when expecting (%v)", written, count)
		}
	})
}

func Test_RotateFileWriter_Close(t *testing.T) {
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
		writer, _ := NewStreamRotatingFileWriter(fileSystem, path)

		if err := writer.(io.Closer).Close(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})
}
