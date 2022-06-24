package srvalidation

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/pkg/errors"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		p := &Provider{}
		_ = p.Register(nil)
		expected := errNilPointer("container")

		if err := p.Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}

		err := p.Register(container)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case !container.Has(ContainerUniversalTranslatorID):
			t.Errorf("didn't registered the universal translator : %v", p)
		case !container.Has(ContainerTranslatorID):
			t.Errorf("didn't registered the translator : %v", p)
		case !container.Has(ContainerParserID):
			t.Errorf("didn't registered the error parser : %v", p)
		case !container.Has(ContainerID):
			t.Errorf("didn't registered the validator : %v", p)
		}
	})

	t.Run("retrieving universal translator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		translator, err := container.Get(ContainerUniversalTranslatorID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case translator == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch translator.(type) {
			case *ut.UniversalTranslator:
			default:
				t.Error("didn't returned the universal translator reference")
			}
		}
	})

	t.Run("error retrieving universal translator when retrieving translator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerUniversalTranslatorID, func() (interface{}, error) { return nil, expected })

		if _, err := container.Get(ContainerTranslatorID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid universal translator instance on retrieving the translator", func(t *testing.T) {
		expected := errConversion("string", "*ut.UniversalTranslator")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerUniversalTranslatorID, func() (interface{}, error) { return "string", nil })

		if _, err := container.Get(ContainerTranslatorID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error instantiating translator", func(t *testing.T) {
		locale := "unsupported"
		Locale = locale
		defer func() { Locale = "en" }()
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		expected := errTranslatorNotFound(locale)

		if _, err := container.Get(ContainerTranslatorID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("retrieving translator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		translator, err := container.Get(ContainerTranslatorID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case translator == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch translator.(type) {
			case ut.Translator:
			default:
				t.Error("didn't returned the translator reference")
			}
		}
	})

	t.Run("error instantiating translator when retrieving parser", func(t *testing.T) {
		locale := "unsupported"
		Locale = locale
		defer func() { Locale = "en" }()
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		expected := errTranslatorNotFound(locale)

		if _, err := container.Get(ContainerParserID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("retrieving parser", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		parser, err := container.Get(ContainerParserID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case parser == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch parser.(type) {
			case Parser:
			default:
				t.Error("didn't returned the translator reference")
			}
		}
	})

	t.Run("error instantiating translator when retrieving validator", func(t *testing.T) {
		locale := "unsupported"
		Locale = locale
		defer func() { Locale = "en" }()
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		expected := errTranslatorNotFound(locale)

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error instantiating parser when retrieving validator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		expected := fmt.Errorf("error message")
		_ = container.Service(ContainerParserID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("retrieving validator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		validator, err := container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case validator == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch validator.(type) {
			case Validator:
			default:
				t.Error("didn't returned the translator reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		p := &Provider{}
		container := slate.ServiceContainer{}
		_ = p.Register(container)

		if err := p.Boot(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("boot", func(t *testing.T) {
		p := &Provider{}
		container := slate.ServiceContainer{}
		_ = p.Register(container)

		if err := p.Boot(container); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
