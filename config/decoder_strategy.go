package config

const (
	// UnknownDecoder defines the value to be used to
	// declare an unknown partial source format.
	UnknownDecoder = "unknown"
)

// DecoderStrategy interface defines the methods of the decoder
// factory strategy that can validate creation requests and instantiation
// of a particular decoder.
type DecoderStrategy interface {
	Accept(format string) bool
	Create(args ...interface{}) (Decoder, error)
}
