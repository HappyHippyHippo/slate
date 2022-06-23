package glog

import (
	"fmt"
	"github.com/spf13/afero"
	"io"
	"os"
	"sync"
	"time"
)

type streamRotatingFileWriter struct {
	lock    sync.Mutex
	fs      afero.Fs
	fp      afero.File
	file    string
	current string
	year    int
	month   time.Month
	day     int
}

var _ io.Writer = &streamRotatingFileWriter{}

// NewStreamRotatingFileWriter instantiates an io.Writer implementing structure that
// output to a file taking into account the file timestamp, so it can perform
// the rotation to a new file on day change.
func NewStreamRotatingFileWriter(fs afero.Fs, file string) (io.Writer, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}

	writer := &streamRotatingFileWriter{fs: fs, file: file}
	if err := writer.rotate(); err != nil {
		return nil, err
	}

	return writer, nil
}

// Write satisfies the io.Writer interface.
func (w *streamRotatingFileWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if err := w.checkRotation(); err != nil {
		return 0, err
	}

	return w.fp.Write(output)
}

// Close satisfies the Closable interface.
func (w *streamRotatingFileWriter) Close() error {
	return w.fp.(io.Closer).Close()
}

func (w *streamRotatingFileWriter) checkRotation() error {
	now := time.Now()
	if now.Day() != w.day || now.Month() != w.month || now.Year() != w.year {
		return w.rotate()
	}

	return nil
}

func (w *streamRotatingFileWriter) rotate() error {
	var err error

	if w.fp != nil {
		if err = w.fp.(io.Closer).Close(); err != nil {
			w.fp = nil
			return err
		}
	}

	now := time.Now()
	w.year = now.Year()
	w.month = now.Month()
	w.day = now.Day()
	w.current = fmt.Sprintf(w.file, now.Format("2006-01-02"))

	if w.fp, err = w.fs.OpenFile(w.current, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644); err != nil {
		return err
	}
	return nil
}
