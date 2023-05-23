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

	t.Run("error getting the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewLoader(NewConfig(), NewSourceFactory())

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrInvalidSource) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrInvalidSource)
		}
	})

	t.Run("error storing the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(Partial{}, nil).Times(1)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true).Times(1)
		sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil).Times(1)

		cfg := NewConfig()
		_ = cfg.AddSource(LoaderSourceID, 1, source)
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrDuplicateSource) {
			t.Errorf("returned (%v) error when expecting (%v)", e, ErrDuplicateSource)
		}
	})

	t.Run("add base source into the partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(1)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true).Times(1)
		sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil).Times(1)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("invalid list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": "string"}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(1)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true).Times(1)
		sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil).Times(1)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("invalid entry in list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourcePartial := Partial{}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(1)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(false),
		)
		sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil).Times(1)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrInvalidSource) {
			t.Errorf("returned (%v) error when expecting (%v)", e, ErrInvalidSource)
		}
	})

	t.Run("error on loaded invalid priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourcePartial := Partial{"priority": "string"}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(1)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true)
		sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil).Times(1)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error on loaded source factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		sourcePartial := Partial{"priority": 0}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(1)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(true),
		)
		gomock.InOrder(
			sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil),
			sourceStrategy.EXPECT().Create(&sourcePartial).Return(nil, expected),
		)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error on source registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourcePartial := Partial{"priority": 0}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(3)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(true),
		)
		gomock.InOrder(
			sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil),
			sourceStrategy.EXPECT().Create(&sourcePartial).Return(source, nil),
		)

		cfg := NewConfig()
		_ = cfg.AddSource("source", 0, source)
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, ErrDuplicateSource) {
			t.Errorf("returned (%v) error when expecting (%v)", e, ErrDuplicateSource)
		}
	})

	t.Run("register the loaded source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourcePartial := Partial{"priority": 0}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(3)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(true),
		)
		gomock.InOrder(
			sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil),
			sourceStrategy.EXPECT().Create(&sourcePartial).Return(source, nil),
		)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourcePath = "config.yaml"
		defer func() { LoaderSourcePath = "partial/sources.yaml" }()

		sourcePartial := Partial{"priority": 0}
		baseCreationPartial := &Partial{"type": "file", "path": "config.yaml", "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(3)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(true),
		)
		gomock.InOrder(
			sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil),
			sourceStrategy.EXPECT().Create(&sourcePartial).Return(source, nil),
		)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourceFormat = "json"
		defer func() { LoaderSourceFormat = "yaml" }()

		sourcePartial := Partial{"priority": 0}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": "json"}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(3)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(true),
		)
		gomock.InOrder(
			sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil),
			sourceStrategy.EXPECT().Create(&sourcePartial).Return(source, nil),
		)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source list path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourceListPath = "config_list"
		defer func() { LoaderSourceListPath = "slate.config.sources" }()

		sourcePartial := Partial{"priority": 0}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"config_list": Partial{"source": sourcePartial}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(3)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(true),
		)
		gomock.InOrder(
			sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil),
			sourceStrategy.EXPECT().Create(&sourcePartial).Return(source, nil),
		)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("register the loaded source with def priority if missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourcePartial := Partial{}
		baseCreationPartial := &Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}
		basePartial := Partial{"slate": Partial{"config": Partial{"sources": Partial{"source": sourcePartial}}}}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("").Return(basePartial, nil).Times(3)
		sourceStrategy := NewMockSourceStrategy(ctrl)
		gomock.InOrder(
			sourceStrategy.EXPECT().Accept(baseCreationPartial).Return(true),
			sourceStrategy.EXPECT().Accept(&sourcePartial).Return(true),
		)
		gomock.InOrder(
			sourceStrategy.EXPECT().Create(baseCreationPartial).Return(source, nil),
			sourceStrategy.EXPECT().Create(&sourcePartial).Return(source, nil),
		)

		cfg := NewConfig()
		sourceFactory := NewSourceFactory()
		_ = sourceFactory.Register(sourceStrategy)
		sut, _ := NewLoader(cfg, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
