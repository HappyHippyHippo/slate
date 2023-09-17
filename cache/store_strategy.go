package cache

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// UnknownStoreType defines the value to be used to
	// declare an unknown Store type.
	UnknownStoreType = "unknown"
)

// StoreStrategy @todo doc
type StoreStrategy interface {
	Accept(config.Partial) bool
	Create(config.Partial) (Store, error)
}
