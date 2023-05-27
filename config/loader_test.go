package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_NewLoader(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		sut, e := NewLoader(nil, NewSourceFactory())
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil source dFactory", func(t *testing.T) {
		sut, e := NewLoader(NewConfig(), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new observer", func(t *testing.T) {
		if sut, e := NewLoader(NewConfig(), NewSourceFactory()); sut == nil {
			t.Error("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("return the (%v) error", e)
		}
	})
}

func Test_Loader_Load(t *testing.T) {
	LoaderSourceID = "base_source_id"
	LoaderSourcePath = "base_source_path"
	LoaderSourceFormat = "format"
	defer func() {
		LoaderSourceID = "main"
		LoaderSourcePath = "partial/sources.yaml"
		LoaderSourceFormat = "format"
	}()
	baseSourcePartial := &Partial{"type": "file", "path": "base_source_path", "format": "format"}

	t.Run("error getting the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(baseSourcePartial).Return(nil, expected).Times(1)
		cfg := NewMockConfigurer(ctrl)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error storing the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		source := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(baseSourcePartial).Return(source, nil).Times(1)
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().AddSource("base_source_id", 0, source).Return(expected).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("add base source into the partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		source := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(baseSourcePartial).Return(source, nil).Times(1)
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().AddSource("base_source_id", 0, source).Return(nil).Times(1)
		cfg.EXPECT().Partial("slate.config.sources").Return(&Partial{}, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("invalid list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &Partial{"sources": "string"}
		source := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(baseSourcePartial).Return(source, nil).Times(1)
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().AddSource("base_source_id", 0, source).Return(nil).Times(1)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error parsing source entry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &Partial{"sources": Partial{"priority": "string"}}
		source := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(baseSourcePartial).Return(source, nil).Times(1)
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().AddSource("base_source_id", 0, source).Return(nil).Times(1)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error creating the source entry", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		sourcePartial := Partial{"type": "my type"}
		partial := &Partial{"source": sourcePartial}
		source := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		gomock.InOrder(
			sourceFactory.EXPECT().Create(baseSourcePartial).Return(source, nil),
			sourceFactory.EXPECT().Create(&sourcePartial).Return(nil, expected),
		)
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().AddSource("base_source_id", 0, source).Return(nil).Times(1)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error on source registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		sourcePartial := Partial{"type": "my type", "priority": 101}
		partial := &Partial{"source": sourcePartial}
		source1 := NewMockSource(ctrl)
		source2 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		gomock.InOrder(
			sourceFactory.EXPECT().Create(baseSourcePartial).Return(source1, nil),
			sourceFactory.EXPECT().Create(&sourcePartial).Return(source2, nil),
		)
		cfg := NewMockConfigurer(ctrl)
		gomock.InOrder(
			cfg.EXPECT().AddSource("base_source_id", 0, source1).Return(nil),
			cfg.EXPECT().AddSource("source", 101, source2).Return(expected),
		)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("register the loaded source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourcePartial := Partial{"type": "my type", "priority": 101}
		partial := &Partial{"source": sourcePartial}
		source1 := NewMockSource(ctrl)
		source2 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		gomock.InOrder(
			sourceFactory.EXPECT().Create(baseSourcePartial).Return(source1, nil),
			sourceFactory.EXPECT().Create(&sourcePartial).Return(source2, nil),
		)
		cfg := NewMockConfigurer(ctrl)
		gomock.InOrder(
			cfg.EXPECT().AddSource("base_source_id", 0, source1).Return(nil),
			cfg.EXPECT().AddSource("source", 101, source2).Return(nil),
		)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LoaderSourcePath
		LoaderSourcePath = "config.yaml"
		defer func() { LoaderSourcePath = prev }()

		basePartial := &Partial{"type": "file", "path": "config.yaml", "format": "format"}
		sourcePartial := Partial{"type": "my type", "priority": 101}
		partial := &Partial{"source": sourcePartial}
		source1 := NewMockSource(ctrl)
		source2 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		gomock.InOrder(
			sourceFactory.EXPECT().Create(basePartial).Return(source1, nil),
			sourceFactory.EXPECT().Create(&sourcePartial).Return(source2, nil),
		)
		cfg := NewMockConfigurer(ctrl)
		gomock.InOrder(
			cfg.EXPECT().AddSource("base_source_id", 0, source1).Return(nil),
			cfg.EXPECT().AddSource("source", 101, source2).Return(nil),
		)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LoaderSourceFormat
		LoaderSourceFormat = "json"
		defer func() { LoaderSourceFormat = prev }()

		basePartial := &Partial{"type": "file", "path": "base_source_path", "format": "json"}
		sourcePartial := Partial{"type": "my type", "priority": 101}
		partial := &Partial{"source": sourcePartial}
		source1 := NewMockSource(ctrl)
		source2 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		gomock.InOrder(
			sourceFactory.EXPECT().Create(basePartial).Return(source1, nil),
			sourceFactory.EXPECT().Create(&sourcePartial).Return(source2, nil),
		)
		cfg := NewMockConfigurer(ctrl)
		gomock.InOrder(
			cfg.EXPECT().AddSource("base_source_id", 0, source1).Return(nil),
			cfg.EXPECT().AddSource("source", 101, source2).Return(nil),
		)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source list path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LoaderSourceListPath
		LoaderSourceListPath = "config_list"
		defer func() { LoaderSourceListPath = prev }()

		basePartial := &Partial{"type": "file", "path": "base_source_path", "format": "format"}
		sourcePartial := Partial{"type": "my type", "priority": 101}
		partial := &Partial{"source": sourcePartial}
		source1 := NewMockSource(ctrl)
		source2 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		gomock.InOrder(
			sourceFactory.EXPECT().Create(basePartial).Return(source1, nil),
			sourceFactory.EXPECT().Create(&sourcePartial).Return(source2, nil),
		)
		cfg := NewMockConfigurer(ctrl)
		gomock.InOrder(
			cfg.EXPECT().AddSource("base_source_id", 0, source1).Return(nil),
			cfg.EXPECT().AddSource("source", 101, source2).Return(nil),
		)
		cfg.EXPECT().Partial("config_list").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("register the loaded source with simple priority if missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourcePartial := Partial{"type": "my type"}
		partial := &Partial{"source": sourcePartial}
		source1 := NewMockSource(ctrl)
		source2 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		gomock.InOrder(
			sourceFactory.EXPECT().Create(baseSourcePartial).Return(source1, nil),
			sourceFactory.EXPECT().Create(&sourcePartial).Return(source2, nil),
		)
		cfg := NewMockConfigurer(ctrl)
		gomock.InOrder(
			cfg.EXPECT().AddSource("base_source_id", 0, source1).Return(nil),
			cfg.EXPECT().AddSource("source", 0, source2).Return(nil),
		)
		cfg.EXPECT().Partial("slate.config.sources").Return(partial, nil).Times(1)

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())
		sut.config = cfg
		sut.sourceFactory = sourceFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
