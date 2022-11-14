package error

import "testing"

func Test_ErrTranslatorNotFound(t *testing.T) {
	expected := "translator not found"

	if chk := ErrTranslatorNotFound.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidKeyGeneratorType(t *testing.T) {
	expected := "invalid key generator type"

	if chk := ErrInvalidKeyGeneratorType.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidKeyGeneratorPartial(t *testing.T) {
	expected := "invalid key generator config"

	if chk := ErrInvalidKeyGeneratorPartial.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidStoreType(t *testing.T) {
	expected := "invalid store type"

	if chk := ErrInvalidStoreType.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrInvalidStorePartial(t *testing.T) {
	expected := "invalid store config"

	if chk := ErrInvalidStorePartial.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrCacheMiss(t *testing.T) {
	expected := "key not found"

	if chk := ErrCacheMiss.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrCacheNotStored(t *testing.T) {
	expected := "element not stored"

	if chk := ErrCacheNotStored.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}

func Test_ErrCacheOpNotSupport(t *testing.T) {
	expected := "op not supported"

	if chk := ErrCacheOpNotSupport.Error(); chk != expected {
		t.Errorf("error with '%s' message when expecting '%s'", chk, expected)
	}
}
