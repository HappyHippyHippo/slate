package log

import (
	"io"
	"sync"
)

// ILog defines the interface of a Log instance.
type ILog interface {
	io.Closer

	Signal(channel string, level Level, msg string, ctx ...Context) error
	Broadcast(level Level, msg string, ctx ...Context) error

	HasStream(id string) bool
	ListStreams() []string
	AddStream(id string, stream IStream) error
	RemoveStream(id string)
	RemoveAllStreams()
	Stream(id string) IStream
}

// Log defines a new logging manager used to centralize the logging
// messaging propagation to a list of registered output streams.
type Log struct {
	mutex   sync.Locker
	streams map[string]IStream
}

var _ ILog = &Log{}

// NewLog instantiate a new Log instance.
func NewLog() ILog {
	// instantiate the log
	return &Log{
		mutex:   &sync.Mutex{},
		streams: map[string]IStream{},
	}
}

// Close will terminate all the logging stream associated to the Log.
func (l *Log) Close() error {
	// clear the stream list, forcing the close call on all of them
	l.RemoveAllStreams()
	return nil
}

// Signal will propagate the channel filtered logging request
// to all stored logging streams.
func (l *Log) Signal(
	channel string,
	level Level,
	msg string,
	ctx ...Context,
) error {
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// propagate the signal request to all registered stream
	for _, s := range l.streams {
		if e := s.Signal(channel, level, msg, ctx...); e != nil {
			return e
		}
	}
	return nil
}

// Broadcast will propagate the logging request to all stored logging streams.
func (l *Log) Broadcast(
	level Level,
	msg string,
	ctx ...Context,
) error {
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// propagate the broadcast request to all registered stream
	for _, s := range l.streams {
		if e := s.Broadcast(level, msg, ctx...); e != nil {
			return e
		}
	}
	return nil
}

// HasStream check if a stream is registered with the requested id.
func (l *Log) HasStream(
	id string,
) bool {
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// check if there is a registered stream with the requested id
	_, ok := l.streams[id]
	return ok
}

// ListStreams retrieve a list of id's of all registered streams on the Log.
func (l *Log) ListStreams() []string {
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// generate a list with all the registered streams id's
	var list []string
	for id := range l.streams {
		list = append(list, id)
	}
	return list
}

// AddStream registers a new stream into the Log instance.
func (l *Log) AddStream(
	id string,
	stream IStream,
) error {
	// check the stream argument reference
	if stream == nil {
		return errNilPointer("stream")
	}
	// check for stream id conflict
	if l.HasStream(id) {
		return errDuplicateStream(id)
	}
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// add the stream to the log stream pool
	l.streams[id] = stream
	return nil
}

// RemoveStream will remove a registered stream with the requested id
// from the Log.
func (l *Log) RemoveStream(
	id string,
) {
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// search for the requested removing stream
	if s, ok := l.streams[id]; ok {
		// check if the stream implements the closer interface and
		// call it if so
		if s, ok := s.(io.Closer); ok {
			_ = s.Close()
		}
		// remove the stream reference from the stream pool
		delete(l.streams, id)
	}
}

// RemoveAllStreams will remove all registered streams from the Log.
func (l *Log) RemoveAllStreams() {
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// iterate through all the registered streams
	for id, s := range l.streams {
		// check if the stream implements the closer interface and
		// call it if so
		if s, ok := s.(io.Closer); ok {
			_ = s.Close()
		}
		// remove the stream reference from the stream pool
		delete(l.streams, id)
	}
}

// Stream retrieve a stream from the Log that is registered with the
// requested id.
func (l *Log) Stream(
	id string,
) IStream {
	// lock the log for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// retrieve the requested stream
	if s, ok := l.streams[id]; ok {
		return s
	}
	return nil
}
