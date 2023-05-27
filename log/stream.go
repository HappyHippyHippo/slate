package log

// Stream interface defines the interaction methods with a logging stream.
type Stream interface {
	Signal(channel string, level Level, message string, ctx ...Context) error
	Broadcast(level Level, message string, ctx ...Context) error

	HasChannel(channel string) bool
	ListChannels() []string
	AddChannel(channel string)
	RemoveChannel(channel string)
}
