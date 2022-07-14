package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_NewSourceStrategyFileObservable(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		strategy, err := newSourceStrategyObservableFile(nil, &decoderFactory{})
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil decoder dFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, err := newSourceStrategyObservableFile(NewMockFs(ctrl), nil)
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new file source dFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fs := NewMockFs(ctrl)
		dFactory := &decoderFactory{}

		strategy, err := newSourceStrategyObservableFile(fs, dFactory)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		case strategy.(*sourceStrategyObservableFile).fs != fs:
			t.Error("didn't stored the file system adapter reference")
		case strategy.(*sourceStrategyObservableFile).factory != dFactory:
			t.Error("didn't stored the decoder dFactory reference")
		}
	})
}

func Test_SourceStrategyFileObservable_Accept(t *testing.T) {
	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			exp        bool
		}{
			{ // _test observable file type
				sourceType: SourceTypeObservableFile,
				exp:        true,
			},
			{ // _test non-observable file type
				sourceType: SourceTypeUnknown,
				exp:        false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()
				strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})
				if check := strategy.Accept(scenario.sourceType); check != scenario.exp {
					t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
				}
			}
			test()
		}
	})
}

func Test_SourceStrategyFileObservable_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		if strategy.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		if strategy.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		if strategy.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		if strategy.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		if !strategy.AcceptFromConfig(&Partial{"type": SourceTypeObservableFile}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyFileObservable_Create(t *testing.T) {
	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.Create("path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.Create(123, "format")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.Create("path", 123)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("create the observable file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%s: %s", field, value))
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		dFactory := &decoderFactory{}
		_ = dFactory.Register(&decoderStrategyYAML{})
		strategy, _ := newSourceStrategyObservableFile(fs, dFactory)

		src, err := strategy.Create(path, DecoderFormatYAML)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}

func Test_SourceStrategyFileObservable_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.CreateFromConfig(&Partial{"format": "format"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.CreateFromConfig(&Partial{"path": 123, "format": "format"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), &decoderFactory{})

		src, err := strategy.CreateFromConfig(&Partial{"path": "path", "format": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("create the observable file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) { copy(buf, fmt.Sprintf("%s: %s", field, value)); return 12, io.EOF }).Times(1)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		dFactory := &decoderFactory{}
		_ = dFactory.Register(&decoderStrategyYAML{})
		strategy, _ := newSourceStrategyObservableFile(fs, dFactory)

		src, err := strategy.CreateFromConfig(&Partial{"path": path, "format": DecoderFormatYAML})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})

	t.Run("create the observable file source defaulting format if not present in config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) { copy(buf, fmt.Sprintf("%s: %s", field, value)); return 12, io.EOF }).Times(1)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		dFactory := &decoderFactory{}
		_ = dFactory.Register(&decoderStrategyYAML{})
		strategy, _ := newSourceStrategyObservableFile(fs, dFactory)

		src, err := strategy.CreateFromConfig(&Partial{"path": path})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}
