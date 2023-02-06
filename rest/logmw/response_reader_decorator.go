package logmw

// ResponseReaderDecorator defines a function used to decorate a response
// reader output.
type ResponseReaderDecorator func(reader ResponseReader, model interface{}) (ResponseReader, error)
