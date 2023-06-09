package file

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/spf13/afero"
)

// RotatingWriter defines an output writer used by a file stream that will
// use a dated file for target output.
type RotatingWriter struct {
	lock    sync.Mutex
	fs      afero.Fs
	fp      afero.File
	file    string
	current string
	year    int
	month   time.Month
	day     int
}

var _ io.Writer = &RotatingWriter{}

// NewRotatingWriter generate a new rotating file writer instance.
func NewRotatingWriter(
	fs afero.Fs,
	file string,
) (io.Writer, error) {
	// check the file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// instantiate the rotating file writer instance
	writer := &RotatingWriter{fs: fs, file: file}
	// open the target rotated file
	if e := writer.rotate(); e != nil {
		return nil, e
	}
	return writer, nil
}

// Write satisfies the io.Writer interface.
func (w *RotatingWriter) Write(
	output []byte,
) (int, error) {
	// lock the file for interaction
	w.lock.Lock()
	defer w.lock.Unlock()
	// check if the file need rotation
	if e := w.checkRotation(); e != nil {
		return 0, e
	}
	// write the content to the target file
	return w.fp.Write(output)
}

// Close satisfies the Closable interface.
func (w *RotatingWriter) Close() error {
	// close the opened file handler
	return w.fp.(io.Closer).Close()
}

func (w *RotatingWriter) checkRotation() error {
	// check if the stored opened file date for the need of rotation
	now := time.Now()
	if now.Day() != w.day || now.Month() != w.month || now.Year() != w.year {
		// rotate the file handler
		return w.rotate()
	}
	return nil
}

func (w *RotatingWriter) rotate() error {
	var e error
	// close the currently opened file
	if w.fp != nil {
		if e = w.fp.(io.Closer).Close(); e != nil {
			w.fp = nil
			return e
		}
	}
	// store the opened file date and create the new target file name
	now := time.Now()
	w.year = now.Year()
	w.month = now.Month()
	w.day = now.Day()
	w.current = fmt.Sprintf(w.file, now.Format("2006-01-02"))
	// open the new target file
	if w.fp, e = w.fs.OpenFile(
		w.current,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644,
	); e != nil {
		return e
	}
	return nil
}
