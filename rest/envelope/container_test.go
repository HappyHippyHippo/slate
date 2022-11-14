package envelope

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	sconfig "github.com/happyhippyhippo/slate/config"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_GetMiddlewareGenerator(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetMiddlewareGenerator(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non middleware instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		s, err := GetMiddlewareGenerator(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid middleware instance returned", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		_ = c.Service(sconfig.ContainerID, func() (interface{}, error) {
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

		s, err := GetMiddlewareGenerator(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
