package slog

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

func newStreamRotatingFileWriter(fs afero.Fs, file string) (io.Writer, error) {
	if fs == nil {
		return nil, errNilPointer("sfs")
	}

	writer := &streamRotatingFileWriter{fs: fs, file: file}
	if e := writer.rotate(); e != nil {
		return nil, e
	}

	return writer, nil
}

// Write satisfies the io.Writer interface.
func (w *streamRotatingFileWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if e := w.checkRotation(); e != nil {
		return 0, e
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
	var e error

	if w.fp != nil {
		if e = w.fp.(io.Closer).Close(); e != nil {
			w.fp = nil
			return e
		}
	}

	now := time.Now()
	w.year = now.Year()
	w.month = now.Month()
	w.day = now.Day()
	w.current = fmt.Sprintf(w.file, now.Format("2006-01-02"))

	if w.fp, e = w.fs.OpenFile(w.current, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644); e != nil {
		return e
	}
	return nil
}
