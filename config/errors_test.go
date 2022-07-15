package config

import (
	"errors"
	"github.com/happyhippyhippo/slate/err"
	"testing"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("error not a instance of NilPointer")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrConversion(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		expected := "invalid type conversion : dummy value to dummy type"

		if e := errConversion(arg, typ); !errors.Is(e, err.ErrConversion) {
			t.Errorf("error not a instance of Conversion")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrConfigSourceNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config source not found : dummy argument"

		if e := errConfigSourceNotFound(arg); !errors.Is(e, err.ErrConfigSourceNotFound) {
			t.Errorf("error not a instance of ErrSourceNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrDuplicateConfigSource(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config source already registered : dummy argument"

		if e := errDuplicateConfigSource(arg); !errors.Is(e, err.ErrDuplicateConfigSource) {
			t.Errorf("error not a instance of ErrDuplicateSource")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrConfigPathNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "config path not found : dummy argument"

		if e := errConfigPathNotFound(arg); !errors.Is(e, err.ErrConfigPathNotFound) {
			t.Errorf("error not a instance of ErrPathNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrConfigRestPathNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "rest path not found : dummy argument"

		if e := errConfigRestPathNotFound(arg); !errors.Is(e, err.ErrConfigRestPathNotFound) {
			t.Errorf("error not a instance of ErrRestConfigPathNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidConfigDecoderFormat(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := DecoderFormatUnknown
		expected := "invalid config decoder format : unknown"

		if e := errInvalidConfigDecoderFormat(arg); !errors.Is(e, err.ErrInvalidConfigDecoderFormat) {
			t.Errorf("error not a instance of ErrInvalidDecoderFormat")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidConfigSourceType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := SourceTypeUnknown
		expected := "invalid config source type : unknown"

		if e := errInvalidConfigSourceType(arg); !errors.Is(e, err.ErrInvalidConfigSourceType) {
			t.Errorf("error not a instance of ErrInvalidSourceType")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidConfigSourcePartial(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := Partial{"field": "dummy argument"}
		expected := "invalid config source config : &map[field:dummy argument]"

		if e := errInvalidConfigSourcePartial(&arg); !errors.Is(e, err.ErrInvalidConfigSourcePartial) {
			t.Errorf("error not a instance of ErrInvalidSourceConfig")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}
