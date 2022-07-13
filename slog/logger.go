package slog

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

// Logger defines a logging proxy for all the registered logging streams.
type Logger struct {
	mutex   sync.Locker
	streams map[string]IStream
}

var _ ILogger = &Logger{}

func newLogger() ILogger {
	return &Logger{
		mutex:   &sync.Mutex{},
		streams: map[string]IStream{},
	}
}

// Close will terminate all the logging stream associated to the logger.
func (l *Logger) Close() error {
	l.RemoveAllStreams()
	return nil
}

// Signal will propagate the channel filtered logging request
// to all stored logging streams.
func (l Logger) Signal(channel string, level Level, msg string, ctx map[string]interface{}) error {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	for _, stream := range l.streams {
		if err := stream.Signal(channel, level, msg, ctx); err != nil {
			return err
		}
	}
	return nil
}

// Broadcast will propagate the logging request to all stored logging streams.
func (l Logger) Broadcast(level Level, msg string, ctx map[string]interface{}) error {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	for _, stream := range l.streams {
		if err := stream.Broadcast(level, msg, ctx); err != nil {
			return err
		}
	}
	return nil
}

// HasStream check if a stream is registered with the requested id.
func (l Logger) HasStream(id string) bool {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	_, ok := l.streams[id]
	return ok
}

// ListStreams retrieve a list of id's of all registered streams on the logger.
func (l Logger) ListStreams() []string {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	var list []string
	for id := range l.streams {
		list = append(list, id)
	}
	return list
}

// AddStream registers a new stream into the logger instance.
func (l *Logger) AddStream(id string, stream IStream) error {
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
func (l *Logger) RemoveStream(id string) {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	if stream, ok := l.streams[id]; ok {
		if s, ok := stream.(io.Closer); ok {
			_ = s.Close()
		}
		delete(l.streams, id)
	}
}

// RemoveAllStreams will remove all registered streams from the logger.
func (l *Logger) RemoveAllStreams() {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	for id, stream := range l.streams {
		if s, ok := stream.(io.Closer); ok {
			_ = s.Close()
		}
		delete(l.streams, id)
	}
}

// Stream retrieve a stream from the logger that is registered with the
// requested id.
func (l Logger) Stream(id string) IStream {
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()

	if stream, ok := l.streams[id]; ok {
		return stream
	}
	return nil
}
