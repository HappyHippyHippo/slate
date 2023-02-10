package validation

import (
	"fmt"
	"testing"

	ut "github.com/go-playground/universal-translator"
	"github.com/happyhippyhippo/slate"
	"github.com/pkg/errors"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Register(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(UniversalTranslatorID):
			t.Errorf("didn't registered the universal translator : %v", sut)
		case !container.Has(TranslatorID):
			t.Errorf("didn't registered the translator : %v", sut)
		case !container.Has(ParserID):
			t.Errorf("didn't registered the error parser : %v", sut)
		case !container.Has(ID):
			t.Errorf("didn't registered the validator : %v", sut)
		}
	})

	t.Run("retrieving universal translator", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		translator, e := container.Get(UniversalTranslatorID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
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
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = container.Service(UniversalTranslatorID, func() (*ut.UniversalTranslator, error) { return nil, expected })

		if _, e := container.Get(TranslatorID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error instantiating translator", func(t *testing.T) {
		locale := "unsupported"
		Locale = locale
		defer func() { Locale = "en" }()
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		if _, e := container.Get(TranslatorID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving translator", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		translator, e := container.Get(TranslatorID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
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
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		if _, e := container.Get(ParserID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving parser", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		parser, e := container.Get(ParserID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case parser == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch parser.(type) {
			case IParser:
			default:
				t.Error("didn't returned the translator reference")
			}
		}
	})

	t.Run("error instantiating translator when retrieving validator", func(t *testing.T) {
		locale := "unsupported"
		Locale = locale
		defer func() { Locale = "en" }()
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error instantiating parser when retrieving validator", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		expected := fmt.Errorf("error message")
		_ = container.Service(ParserID, func() (IParser, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving validator", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		validator, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
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
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Boot(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("successful boot", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Provide(Provider{})

		if e := app.Boot(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
