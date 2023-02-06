package config

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_Source_Has(t *testing.T) {
	t.Run("lock and redirect to the stored config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		search := "path"
		data := Config{search: "value"}
		locker := NewMockLocker(ctrl)
		locker.EXPECT().Lock().Times(1)
		locker.EXPECT().Unlock().Times(1)

		sut := &Source{mutex: locker, config: data}

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
		data := Config{search: expected}
		locker := NewMockLocker(ctrl)
		locker.EXPECT().Lock().Times(1)
		locker.EXPECT().Unlock().Times(1)

		sut := &Source{mutex: locker, config: data}

		if value, e := sut.Get(search); e != nil {
			t.Errorf("returned the unexpected err : %v", e)
		} else if value != expected {
			t.Errorf("returned the (%v) value", value)
		}
	})
}
