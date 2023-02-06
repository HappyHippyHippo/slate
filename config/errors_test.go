package config

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate/err"
)

func Test_errNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, err.NilPointer) {
			t.Errorf("error not a instance of err.NilPointer")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errConversion(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		expected := "invalid type conversion : dummy value to dummy type"

		if e := errConversion(arg, typ); !errors.Is(e, err.Conversion) {
			t.Errorf("error not a instance of err.Conversion")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errConfigPathNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config path not found : dummy argument"

		if e := errPathNotFound(arg); !errors.Is(e, err.ConfigPathNotFound) {
			t.Errorf("error not a instance of err.ConfigPathNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errInvalidFormat(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := DecoderFormatUnknown
		expected := "invalid config format : unknown"

		if e := errInvalidFormat(arg); !errors.Is(e, err.InvalidConfigFormat) {
			t.Errorf("error not a instance of err.InvalidConfigFormat")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errInvalidSource(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := SourceStrategyUnknown
		expected := "invalid config source type : unknown"

		if e := errInvalidSource(arg); !errors.Is(e, err.InvalidConfigSource) {
			t.Errorf("error not a instance of err.InvalidConfigSource")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errRestPathNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "rest path not found : dummy argument"

		if e := errRestPathNotFound(arg); !errors.Is(e, err.ConfigRestPathNotFound) {
			t.Errorf("error not a instance of err.ConfigRestPathNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errInvalidSourceData(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := Config{"field": "dummy argument"}
		expected := "invalid config source data : &map[field:dummy argument]"

		if e := errInvalidSourceData(&arg); !errors.Is(e, err.InvalidConfigSourceData) {
			t.Errorf("error not a instance of err.InvalidConfigSourceData")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errSourceNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config source not found : dummy argument"

		if e := errSourceNotFound(arg); !errors.Is(e, err.ConfigSourceNotFound) {
			t.Errorf("err not a instance of err.ConfigSourceNotFound")
		} else if e.Error() != expected {
			t.Errorf("err message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errDuplicateSource(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config source already registered : dummy argument"

		if e := errDuplicateSource(arg); !errors.Is(e, err.DuplicateConfigSource) {
			t.Errorf("err not a instance of err.DuplicateConfigSource")
		} else if e.Error() != expected {
			t.Errorf("err message (%v) not same as expected (%v)", e, expected)
		}
	})
}
