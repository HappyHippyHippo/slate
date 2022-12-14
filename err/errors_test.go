package err

import (
	"testing"
)

func Test_NilPointer(t *testing.T) {
	expected := "invalid nil pointer"

	if chk := NilPointer.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_Conversion(t *testing.T) {
	expected := "invalid type conversion"

	if chk := Conversion.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_NonFunctionFactory(t *testing.T) {
	expected := "non-function factory"

	if chk := NonFunctionFactory.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_FactoryWithoutResult(t *testing.T) {
	expected := "factory without result"

	if chk := FactoryWithoutResult.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ServiceNotFound(t *testing.T) {
	expected := "service not found"

	if chk := ServiceNotFound.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_Container(t *testing.T) {
	expected := "service container error"

	if chk := Container.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ConfigPathNotFound(t *testing.T) {
	expected := "config path not found"

	if chk := ConfigPathNotFound.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidConfigFormat(t *testing.T) {
	expected := "invalid config format"

	if chk := InvalidConfigFormat.Error(); chk != expected {
		t.Errorf("errorwith '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidConfigSource(t *testing.T) {
	expected := "invalid config source type"

	if chk := InvalidConfigSource.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ConfigRestPathNotFound(t *testing.T) {
	expected := "rest path not found"

	if chk := ConfigRestPathNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidConfigSourceData(t *testing.T) {
	expected := "invalid config source data"

	if chk := InvalidConfigSourceData.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ConfigSourceNotFound(t *testing.T) {
	expected := "config source not found"

	if chk := ConfigSourceNotFound.Error(); chk != expected {
		t.Errorf("err with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_DuplicateConfigSource(t *testing.T) {
	expected := "config source already registered"

	if chk := DuplicateConfigSource.Error(); chk != expected {
		t.Errorf("err with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidLogFormat(t *testing.T) {
	expected := "invalid output log format"

	if chk := InvalidLogFormat.Error(); chk != expected {
		t.Errorf("err with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidLogLevel(t *testing.T) {
	expected := "invalid log level"

	if chk := InvalidLogLevel.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidLogStream(t *testing.T) {
	expected := "invalid log stream type"

	if chk := InvalidLogStream.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidLogConfig(t *testing.T) {
	expected := "invalid log config"

	if chk := InvalidLogConfig.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_DuplicateLogStream(t *testing.T) {
	expected := "stream already registered"

	if chk := DuplicateLogStream.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_DatabaseConfigNotFound(t *testing.T) {
	expected := "database config not found"

	if chk := DatabaseConfigNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_UnknownDatabaseDialect(t *testing.T) {
	expected := "unknown database dialect"

	if chk := UnknownDatabaseDialect.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_InvalidWatchdogConfig(t *testing.T) {
	expected := "invalid watchdog config"

	if chk := InvalidWatchdogConfig.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_DuplicateWatchdogService(t *testing.T) {
	expected := "duplicate watchdog service"

	if chk := DuplicateWatchdogService.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}
