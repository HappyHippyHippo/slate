package logmw

// RequestReaderDecorator defines a function used to decorate a
// request reader output.
type RequestReaderDecorator func(reader RequestReader, model interface{}) (RequestReader, error)
