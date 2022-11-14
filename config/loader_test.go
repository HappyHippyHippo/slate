package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewLoader(t *testing.T) {
	t.Run("nil cfg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newLoader(nil, NewMockSourceFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("nil source decoderFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newLoader(NewMockManager(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("new loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := newLoader(NewMockManager(ctrl), NewMockSourceFactory(ctrl)); sut == nil {
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

		expected := fmt.Errorf("error message")
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(nil, expected).Times(1)
		sut, _ := newLoader(NewMockManager(ctrl), sourceFactory)

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
		src := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(expected).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("add base source into the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return([]interface{}{}, nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("invalid list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return([]interface{}{123}, nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on loaded missing id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partials := []interface{}{Partial{}}
		src := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConfigPathNotFound) {
			t.Errorf("returned (%v) error when expecting (%v)", e, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("error on loaded invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partials := []interface{}{Partial{"id": 12}}
		src := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error on loaded invalid priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partials := []interface{}{Partial{"id": "id", "priority": "string"}}
		src := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error on loaded source factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		srcPartial := Partial{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		sourceFactory.EXPECT().CreateFromConfig(&srcPartial).Return(nil, expected).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

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
		srcPartial := Partial{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		sourceFactory.EXPECT().CreateFromConfig(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(expected).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("register the loaded source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srcPartial := Partial{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		sourceFactory.EXPECT().CreateFromConfig(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourcePath = "config.yaml"
		defer func() { LoaderSourcePath = "config/config.yaml" }()

		srcPartial := Partial{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, "config.yaml", LoaderSourceFormat).Return(src, nil).Times(1)
		sourceFactory.EXPECT().CreateFromConfig(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined source list path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourceListPath = "config_list"
		defer func() { LoaderSourceListPath = "slate.config.list" }()

		srcPartial := Partial{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		sourceFactory.EXPECT().CreateFromConfig(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List("config_list").Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("load from defined format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourceFormat = "json"
		defer func() { LoaderSourceFormat = "yaml" }()

		srcPartial := Partial{"id": "id", "priority": 0}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, "json").Return(src, nil).Times(1)
		sourceFactory.EXPECT().CreateFromConfig(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("register the loaded source with default priority if missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srcPartial := Partial{"id": "id"}
		partials := []interface{}{srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sourceFactory := NewMockSourceFactory(ctrl)
		sourceFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		sourceFactory.EXPECT().CreateFromConfig(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().List(LoaderSourceListPath).Return(partials, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := newLoader(mockManager, sourceFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
