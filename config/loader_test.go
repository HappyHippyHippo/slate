package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
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
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", e, slate.ErrNilPointer)
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
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", e, slate.ErrNilPointer)
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
	LoaderSourceFormat = "format"
	defer func() {
		LoaderSourceID = "main"
		LoaderSourcePath = "config/sources.yaml"
		LoaderSourceFormat = "format"
	}()

	t.Run("error getting the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(nil, expected).Times(1)
		sut, _ := NewLoader(NewMockManager(ctrl), sFactory)

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
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(expected).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

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
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(&Config{}, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("invalid list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(nil, expected).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("invalid entry in list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(&Config{"entry": 123}, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("error on loaded invalid priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partials := &Config{"id": Config{"priority": "string"}}
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partials, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

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
		srcPartial := Config{"priority": 0}
		partial := &Config{"id": srcPartial}
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(nil, expected).Times(1)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partial, nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

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
		srcPartial := Config{"priority": 0}
		partial := &Config{"id": srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partial, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(expected).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("register the loaded source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srcPartial := Config{"priority": 0}
		partial := &Config{"id": srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partial, nil).Times(1)
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
		defer func() { LoaderSourcePath = "config/sources.yaml" }()

		srcPartial := Config{"priority": 0}
		partial := &Config{"id": srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": "config.yaml", "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partial, nil).Times(1)
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
		defer func() { LoaderSourceListPath = "slate.config.sources" }()

		srcPartial := Config{"priority": 0}
		partial := &Config{"id": srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partial, nil).Times(1)
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

		srcPartial := Config{"priority": 0}
		partial := &Config{"id": srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": "json"}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partial, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("register the loaded source with def priority if missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srcPartial := Config{"priority": 0}
		partial := &Config{"id": srcPartial}
		src := NewMockSource(ctrl)
		src1 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(&Config{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat}).Return(src, nil).Times(1)
		sFactory.EXPECT().Create(&srcPartial).Return(src1, nil)
		mockManager := NewMockManager(ctrl)
		mockManager.EXPECT().AddSource(LoaderSourceID, 0, src).Return(nil).Times(1)
		mockManager.EXPECT().Config(LoaderSourceListPath).Return(partial, nil).Times(1)
		mockManager.EXPECT().AddSource("id", 0, src1).Return(nil).Times(1)
		sut, _ := NewLoader(mockManager, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
