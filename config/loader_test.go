package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewLoader(t *testing.T) {
	t.Run("nil manager", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewLoader(nil, NewMockSourceFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil source dFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewLoader(NewMockManager(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("new loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := NewLoader(NewMockManager(ctrl), NewMockSourceFactory(ctrl)); sut == nil {
			t.Error("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("return the (%v) error", e)
		}
	})
}

func Test_Loader_Load(t *testing.T) {
	LoaderSourceID = "base_source_id"
	LoaderSourcePath = "base_source_path"
	LoaderSourceFormat = DecoderFormatYAML
	defer func() {
		LoaderSourceID = "main"
		LoaderSourcePath = "config/config.yaml"
		LoaderSourceFormat = DecoderFormatYAML
	}()

	t.Run("error getting the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(nil, expected).Times(1)
		sut, _ := NewLoader(NewMockManager(ctrl), sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("error storing the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(expected).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("add base source into the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return([]interface{}{}, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("invalid list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(nil, expected).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		}
	})

	t.Run("invalid entry in list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return([]interface{}{123}, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		}
	})

	t.Run("error on loaded missing id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partials := []interface{}{Config{}}
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ConfigPathNotFound) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.ConfigPathNotFound)
		}
	})

	t.Run("error on loaded invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partials := []interface{}{Config{"id": 12}}
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("error on loaded invalid priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partials := []interface{}{Config{"id": "id", "priority": "string"}}
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("error on loaded source factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		srcPartial := Config{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(nil, expected).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("error on source registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("err message")
		srcPartial := Config{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(expected).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("register the loaded source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srcPartial := Config{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourcePath = "config.yaml"
		defer func() { LoaderSourcePath = "config/config.yaml" }()

		srcPartial := Config{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": "config.yaml", "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source list path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourceListPath = "config_list"
		defer func() { LoaderSourceListPath = "slate.config.list" }()

		srcPartial := Config{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List("config_list").Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourceFormat = "json"
		defer func() { LoaderSourceFormat = "yaml" }()

		srcPartial := Config{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": "json"}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("register the loaded source with default priority if missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srcPartial := Config{"id": "id"}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": SourceStrategyFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
