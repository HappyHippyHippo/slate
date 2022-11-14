package rest

import (
	"testing"

	"github.com/gin-gonic/gin"
	sconfig "github.com/happyhippyhippo/slate/config"
	sfs "github.com/happyhippyhippo/slate/fs"
	slog "github.com/happyhippyhippo/slate/log"
	smigration "github.com/happyhippyhippo/slate/migration"
	srdb "github.com/happyhippyhippo/slate/rdb"
)

func Test_NewApplication(t *testing.T) {
	t.Run("correctly initialize the application", func(t *testing.T) {
		app := NewApplication()

		if app == nil {
			t.Error("didn't returned a valid application reference")
		} else if !app.Container.Has(sfs.ContainerID) {
			t.Errorf("didn't registered the sfs provider")
		} else if !app.Container.Has(sconfig.ContainerID) {
			t.Errorf("didn't registered the sconfig provider")
		} else if !app.Container.Has(slog.ContainerID) {
			t.Errorf("didn't registered the slog provider")
		} else if !app.Container.Has(srdb.ContainerID) {
			t.Errorf("didn't registered the srdb provider")
		} else if !app.Container.Has(smigration.ContainerID) {
			t.Errorf("didn't registered the smigration provider")
		} else if !app.Container.Has(ContainerEngineID) {
			t.Errorf("didn't registered the gin engine service")
		} else if engine, err := app.Container.Get(ContainerEngineID); err != nil {
			t.Errorf("returned the unexcepected '%v' error while retrieveing the gin engine", err)
		} else if _, ok := engine.(*gin.Engine); !ok {
			t.Error("the engine service didn't returned a valid gin engine instance")
		}
	})
}

func Test_Application_Engine(t *testing.T) {
	t.Run("correctly initialize the application", func(t *testing.T) {
		app := NewApplication()

		if app.engine != app.Engine() {
			t.Error("didn't returned the application registered gin engine")
		}
	})
}
