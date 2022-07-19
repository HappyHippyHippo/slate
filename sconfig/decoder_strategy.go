package sconfig

// IDecoderStrategy interface defines the methods of the decoder
// factory strategy that can validate creation requests and instantiation of a
// particular decoder.
type IDecoderStrategy interface {
	Accept(format string) bool
	Create(args ...interface{}) (IDecoder, error)
}
