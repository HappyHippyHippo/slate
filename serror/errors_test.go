package serror

import "testing"

func Test_ErrNilPointer(t *testing.T) {
	expected := "invalid nil pointer"

	if chk := ErrNilPointer.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrConversion(t *testing.T) {
	expected := "invalid type conversion"

	if chk := ErrConversion.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrServiceNotFound(t *testing.T) {
	expected := "service not found"

	if chk := ErrServiceNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrConfigSourceNotFound(t *testing.T) {
	expected := "config source not found"

	if chk := ErrConfigSourceNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrDuplicateConfigSource(t *testing.T) {
	expected := "config source already registered"

	if chk := ErrDuplicateConfigSource.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrConfigPathNotFound(t *testing.T) {
	expected := "config path not found"

	if chk := ErrConfigPathNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrConfigRestPathNotFound(t *testing.T) {
	expected := "rest path not found"

	if chk := ErrConfigRestPathNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidConfigDecoderFormat(t *testing.T) {
	expected := "invalid config decoder format"

	if chk := ErrInvalidConfigDecoderFormat.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidConfigSourceType(t *testing.T) {
	expected := "invalid config source type"

	if chk := ErrInvalidConfigSourceType.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidConfigSourcePartial(t *testing.T) {
	expected := "invalid config source config"

	if chk := ErrInvalidConfigSourcePartial.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidLogFormat(t *testing.T) {
	expected := "invalid output format"

	if chk := ErrInvalidLogFormat.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidLogLevel(t *testing.T) {
	expected := "invalid logger level"

	if chk := ErrInvalidLogLevel.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrDuplicateLogStream(t *testing.T) {
	expected := "stream already registered"

	if chk := ErrDuplicateLogStream.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidLogStreamType(t *testing.T) {
	expected := "invalid stream type"

	if chk := ErrInvalidLogStreamType.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidLogStreamConfig(t *testing.T) {
	expected := "invalid log stream config"

	if chk := ErrInvalidLogStreamConfig.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrDatabaseConfigNotFound(t *testing.T) {
	expected := "database config not found"

	if chk := ErrDatabaseConfigNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrUnknownDatabaseDialect(t *testing.T) {
	expected := "unknown database dialect"

	if chk := ErrUnknownDatabaseDialect.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}
