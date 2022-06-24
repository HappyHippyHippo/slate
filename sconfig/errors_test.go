package sconfig

import (
	"errors"
	"github.com/happyhippyhippo/slate/serror"
	"testing"
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

func Test_ErrConfigSourceNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config source not found : dummy argument"

		if err := errConfigSourceNotFound(arg); !errors.Is(err, serror.ErrConfigSourceNotFound) {
			t.Errorf("error not a instance of ErrSourceNotFound")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrDuplicateConfigSource(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config source already registered : dummy argument"

		if err := errDuplicateConfigSource(arg); !errors.Is(err, serror.ErrDuplicateConfigSource) {
			t.Errorf("error not a instance of ErrDuplicateSource")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrConfigPathNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config path not found : dummy argument"

		if err := errConfigPathNotFound(arg); !errors.Is(err, serror.ErrConfigPathNotFound) {
			t.Errorf("error not a instance of ErrPathNotFound")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrConfigRestPathNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "rest path not found : dummy argument"

		if err := errConfigRestPathNotFound(arg); !errors.Is(err, serror.ErrConfigRestPathNotFound) {
			t.Errorf("error not a instance of ErrRestConfigPathNotFound")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidConfigDecoderFormat(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := DecoderFormatUnknown
		expected := "invalid config decoder format : unknown"

		if err := errInvalidConfigDecoderFormat(arg); !errors.Is(err, serror.ErrInvalidConfigDecoderFormat) {
			t.Errorf("error not a instance of ErrInvalidDecoderFormat")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidConfigSourceType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := SourceTypeUnknown
		expected := "invalid config source type : unknown"

		if err := errInvalidConfigSourceType(arg); !errors.Is(err, serror.ErrInvalidConfigSourceType) {
			t.Errorf("error not a instance of ErrInvalidSourceType")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidConfigSourcePartial(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := Partial{"field": "dummy argument"}
		expected := "invalid config source config : &map[field:dummy argument]"

		if err := errInvalidConfigSourcePartial(&arg); !errors.Is(err, serror.ErrInvalidConfigSourcePartial) {
			t.Errorf("error not a instance of ErrInvalidSourceConfig")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}
