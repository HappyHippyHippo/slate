package source

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

func Test_Source_Has(t *testing.T) {
	t.Run("lock and redirect to the stored config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		search := "path"
		data := config.Partial{search: "value"}
		locker := NewMockLocker(ctrl)
		locker.EXPECT().Lock().Times(1)
		locker.EXPECT().Unlock().Times(1)

		sut := &Source{Mutex: locker, Partial: data}

		if value := sut.Has(search); value != true {
			t.Errorf("returned the (%v) value", value)
		}
	})
}

func Test_Source_Get(t *testing.T) {
	t.Run("lock and redirect to the stored config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		search := "path"
		expected := "value"
		data := config.Partial{search: expected}
		locker := NewMockLocker(ctrl)
		locker.EXPECT().Lock().Times(1)
		locker.EXPECT().Unlock().Times(1)

		sut := &Source{Mutex: locker, Partial: data}

		if value, e := sut.Get(search); e != nil {
			t.Errorf("unexpected (%v) error", e)
		} else if value != expected {
			t.Errorf("returned the (%v) value", value)
		}
	})
}
