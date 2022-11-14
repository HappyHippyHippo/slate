package error

import (
	"fmt"
)

var (
	// ErrTranslatorNotFound defines an error that signal that the
	// requested translator was not found.
	ErrTranslatorNotFound = fmt.Errorf("translator not found")

	// ErrInvalidKeyGeneratorType defines an error that signal an
	// unexpected/unknown key generator type.
	ErrInvalidKeyGeneratorType = fmt.Errorf("invalid key generator type")

	// ErrInvalidKeyGeneratorPartial defines an error that signal an
	// invalid key generator configuration Partial.
	ErrInvalidKeyGeneratorPartial = fmt.Errorf("invalid key generator config")

	// ErrInvalidStoreType defines an error that signal an
	// unexpected/unknown store type.
	ErrInvalidStoreType = fmt.Errorf("invalid store type")

	// ErrInvalidStorePartial defines an error that signal an
	// invalid store configuration Partial.
	ErrInvalidStorePartial = fmt.Errorf("invalid store config")

	// ErrCacheMiss defines an error that signals that the cache key
	// was not found in the cache store.
	ErrCacheMiss = fmt.Errorf("key not found")

	// ErrCacheNotStored defines an error that signals that the
	// element was not actually stored in the cache.
	ErrCacheNotStored = fmt.Errorf("element not stored")

	// ErrCacheOpNotSupport defines an error that signals that the
	// operation is not supported for the defined cache store.
	ErrCacheOpNotSupport = fmt.Errorf("op not supported")
)
