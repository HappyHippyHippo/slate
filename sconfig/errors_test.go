package sconfig

import (
	"errors"
	"github.com/happyhippyhippo/slate/serr"
	"testing"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, serr.ErrNilPointer) {
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

		if e := errConversion(arg, typ); !errors.Is(e, serr.ErrConversion) {
			t.Errorf("error not a instance of Conversion")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrConfigSourceNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "sconfig source not found : dummy argument"

		if e := errConfigSourceNotFound(arg); !errors.Is(e, serr.ErrConfigSourceNotFound) {
			t.Errorf("error not a instance of ErrSourceNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrDuplicateConfigSource(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "sconfig source already registered : dummy argument"

		if e := errDuplicateConfigSource(arg); !errors.Is(e, serr.ErrDuplicateConfigSource) {
			t.Errorf("error not a instance of ErrDuplicateSource")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrConfigPathNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "sconfig path not found : dummy argument"

		if e := errConfigPathNotFound(arg); !errors.Is(e, serr.ErrConfigPathNotFound) {
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

		if e := errConfigRestPathNotFound(arg); !errors.Is(e, serr.ErrConfigRestPathNotFound) {
			t.Errorf("error not a instance of ErrRestConfigPathNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidConfigDecoderFormat(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := DecoderFormatUnknown
		expected := "invalid sconfig decoder format : unknown"

		if e := errInvalidConfigDecoderFormat(arg); !errors.Is(e, serr.ErrInvalidConfigDecoderFormat) {
			t.Errorf("error not a instance of ErrInvalidDecoderFormat")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidConfigSourceType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := SourceTypeUnknown
		expected := "invalid sconfig source type : unknown"

		if e := errInvalidConfigSourceType(arg); !errors.Is(e, serr.ErrInvalidConfigSourceType) {
			t.Errorf("error not a instance of ErrInvalidSourceType")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidConfigSourcePartial(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := Partial{"field": "dummy argument"}
		expected := "invalid sconfig source sconfig : &map[field:dummy argument]"

		if e := errInvalidConfigSourcePartial(&arg); !errors.Is(e, serr.ErrInvalidConfigSourcePartial) {
			t.Errorf("error not a instance of ErrInvalidSourceConfig")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}
