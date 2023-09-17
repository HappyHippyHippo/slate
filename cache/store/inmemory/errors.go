//go:build inmemory

package inmemory

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/cache"
	"github.com/happyhippyhippo/slate/config"
)

func errNilPointer(
	arg string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(slate.ErrNilPointer, arg, ctx...)
}

func errInvalidStore(
	cfg config.Partial,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(cache.ErrInvalidStore, fmt.Sprintf("%v", cfg), ctx...)
}

func errMiss(
	key string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(cache.ErrMiss, key, ctx...)
}

func errNotStored(
	key string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(cache.ErrNotStored, key, ctx...)
}
