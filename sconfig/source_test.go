package sconfig

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_Source_Has(t *testing.T) {
	t.Run("lock and redirect to the stored Partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		search := "path"
		data := Partial{search: "value"}
		locker := NewMockLocker(ctrl)
		locker.EXPECT().Lock().Times(1)
		locker.EXPECT().Unlock().Times(1)

		sut := &source{mutex: locker, partial: data}

		if value := sut.Has(search); value != true {
			t.Errorf("returned the (%v) value", value)
		}
	})
}

func Test_Source_Get(t *testing.T) {
	t.Run("lock and redirect to the stored Partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		search := "path"
		expected := "value"
		data := Partial{search: expected}
		locker := NewMockLocker(ctrl)
		locker.EXPECT().Lock().Times(1)
		locker.EXPECT().Unlock().Times(1)

		sut := &source{mutex: locker, partial: data}

		if value, e := sut.Get(search); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if value != expected {
			t.Errorf("returned the (%v) value", value)
		}
	})
}
