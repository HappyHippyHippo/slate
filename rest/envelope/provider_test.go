package envelope

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	sconfig "github.com/happyhippyhippo/slate/config"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (&Provider{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}

		if err := p.Register(container); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !container.Has(ContainerID) {
			t.Errorf("didn't registered the middleware : %v", p)
		}
	})

	t.Run("error retrieving configuration when retrieving middleware generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid configuration instance when retrieving middleware generator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("retrieving middleware generator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&sconfig.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			cfgData1 := sconfig.Partial{
				"slate": sconfig.Partial{
					"service": sconfig.Partial{
						"id": 1,
					},
					"rest": sconfig.Partial{
						"accept": []interface{}{"application/json"},
					},
					"endpoints": sconfig.Partial{
						"index": sconfig.Partial{
							"id": 2,
						},
					},
				},
			}
			source1 := NewMockConfigManager(ctrl)
			source1.EXPECT().Get("").Return(cfgData1, nil).MinTimes(1)
			cfg := sconfig.NewManager(0)
			_ = cfg.AddSource("id1", 0, source1)

			return cfg, nil
		})

		generator, err := container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case generator == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch generator.(type) {
			case MiddlewareGenerator:
			default:
				t.Error("didn't returned a valid middleware generator")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (&Provider{}).Boot(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("boot", func(t *testing.T) {
		p := &Provider{}
		container := slate.ServiceContainer{}
		_ = p.Register(container)

		if err := p.Boot(container); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
