package config

const (
	// UnknownDecoderFormat defines the value to be used to
	// declare an unknown config source format.
	UnknownDecoderFormat = "unknown"
)

// IDecoderStrategy interface defines the methods of the decoder
// factory strategy that can validate creation requests and instantiation
// of a particular decoder.
type IDecoderStrategy interface {
	Accept(format string) bool
	Create(args ...interface{}) (IDecoder, error)
}
