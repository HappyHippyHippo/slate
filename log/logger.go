package log

import (
	"io"
	"sync"
)

// ILogger defines the interface of a logger instance.
type ILogger interface {
	io.Closer

	Signal(channel string, level Level, msg string, ctx map[string]interface{}) error
	Broadcast(level Level, msg string, ctx map[string]interface{}) error
	HasStream(id string) bool
	ListStreams() []string
	AddStream(id string, stream IStream) error
	RemoveStream(id string)
	RemoveAllStreams()
	Stream(id string) IStream
}

type logger struct {
	mutex   sync.Locker
	streams map[string]IStream
}

var _ ILogger = &logger{}

// NewLogger instantiate a new logger instance.
func NewLogger() ILogger {
	return &logger{
		mutex:   &sync.Mutex{},
		streams: map[string]IStream{},
	}
}

// Close will terminate all the logging stream associated to the logger.
func (l *logger) Close() error {
	l.RemoveAllStreams()
	return nil
}

// Signal will propagate the channel filtered logging request
// to all stored logging streams.
func (l logger) Signal(channel string, level Level, msg string, ctx map[string]interface{}) error {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	for _, s := range l.streams {
		if e := s.Signal(channel, level, msg, ctx); e != nil {
			return e
		}
	}
	return nil
}

// Broadcast will propagate the logging request to all stored logging streams.
func (l logger) Broadcast(level Level, msg string, ctx map[string]interface{}) error {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	for _, s := range l.streams {
		if e := s.Broadcast(level, msg, ctx); e != nil {
			return e
		}
	}
	return nil
}

// HasStream check if a stream is registered with the requested id.
func (l logger) HasStream(id string) bool {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	_, ok := l.streams[id]
	return ok
}

// ListStreams retrieve a list of id's of all registered streams on the logger.
func (l logger) ListStreams() []string {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	var list []string
	for id := range l.streams {
		list = append(list, id)
	}
	return list
}

// AddStream registers a new stream into the logger instance.
func (l *logger) AddStream(id string, stream IStream) error {
	if stream == nil {
		return errNilPointer("stream")
	}
	if l.HasStream(id) {
		return errDuplicateStream(id)
	}

	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	l.streams[id] = stream
	return nil
}

// RemoveStream will remove a registered stream with the requested id
// from the logger.
func (l *logger) RemoveStream(id string) {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	if s, ok := l.streams[id]; ok {
		if s, ok := s.(io.Closer); ok {
			_ = s.Close()
		}
		delete(l.streams, id)
	}
}

// RemoveAllStreams will remove all registered streams from the logger.
func (l *logger) RemoveAllStreams() {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	for id, s := range l.streams {
		if s, ok := s.(io.Closer); ok {
			_ = s.Close()
		}
		delete(l.streams, id)
	}
}

// Stream retrieve a stream from the logger that is registered with the
// requested id.
func (l logger) Stream(id string) IStream {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	if s, ok := l.streams[id]; ok {
		return s
	}
	return nil
}
