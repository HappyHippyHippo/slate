package cache

import (
	"errors"
	"testing"

	sconfig "github.com/happyhippyhippo/slate/config"
	serror "github.com/happyhippyhippo/slate/error"
	srerror "github.com/happyhippyhippo/slate/rest/error"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if err := errNilPointer(arg); !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("error not a instance of ErrNilPointer")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrConversion(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		expected := "invalid type conversion : dummy value to dummy type"

		if err := errConversion(arg, typ); !errors.Is(err, serror.ErrConversion) {
			t.Errorf("error not a instance of ErrConversion")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidKeyGeneratorType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid key generator type : dummy argument"

		if err := errInvalidKeyGeneratorType(arg); !errors.Is(err, srerror.ErrInvalidKeyGeneratorType) {
			t.Errorf("error not a instance of ErrInvalidKeyGeneratorType")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidKeyGeneratorPartial(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := sconfig.Partial{"field": "dummy argument"}
		expected := "invalid key generator config : &map[field:dummy argument]"

		if err := errInvalidKeyGeneratorPartial(&arg); !errors.Is(err, srerror.ErrInvalidKeyGeneratorPartial) {
			t.Errorf("error not a instance of ErrInvalidKeyGeneratorPartial")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidStoreType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid store type : dummy argument"

		if err := errInvalidStoreType(arg); !errors.Is(err, srerror.ErrInvalidStoreType) {
			t.Errorf("error not a instance of ErrInvalidStoreType")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidStorePartial(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := sconfig.Partial{"field": "dummy argument"}
		expected := "invalid store config : &map[field:dummy argument]"

		if err := errInvalidStorePartial(&arg); !errors.Is(err, srerror.ErrInvalidStorePartial) {
			t.Errorf("error not a instance of ErrInvalidStorePartial")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrCacheMiss(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "key not found : dummy argument"

		if err := errCacheMiss(arg); !errors.Is(err, srerror.ErrCacheMiss) {
			t.Errorf("error not a instance of ErrCacheMiss")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrCacheNotStored(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "element not stored : dummy argument"

		if err := errCacheNotStored(arg); !errors.Is(err, srerror.ErrCacheNotStored) {
			t.Errorf("error not a instance of ErrCacheNotStored")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrCacheOpNotSupport(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "op not supported : dummy argument"

		if err := errCacheOpNotSupport(arg); !errors.Is(err, srerror.ErrCacheOpNotSupport) {
			t.Errorf("error not a instance of ErrCacheOpNotSupport")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}
