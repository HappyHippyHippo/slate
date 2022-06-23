package srest

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/sfs"
	"github.com/happyhippyhippo/slate/slog"
	"github.com/happyhippyhippo/slate/smigration"
	"github.com/happyhippyhippo/slate/srdb"
)

// Application defines an object for a srest api project.
type Application struct {
	slate.Application
	engine Engine
}

// NewApplication used to instantiate a new application.
func NewApplication() *Application {
	app := &Application{
		Application: *slate.NewApplication(),
		engine:      gin.New(),
	}

	_ = app.Add(sfs.Provider{})
	_ = app.Add(sconfig.Provider{})
	_ = app.Add(slog.Provider{})
	_ = app.Add(srdb.Provider{})
	_ = app.Add(smigration.Provider{})

	_ = app.Container.Service(ContainerEngineID, func() (interface{}, error) {
		return app.engine, nil
	})

	app.engine.Use(
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	return app
}

// Engine returns the reference to the instantiated gin-gonic engine
func (a Application) Engine() Engine {
	return a.engine
}

// Run method will boot the application, if not yet, and the start
// the underlying gin server.
func (a Application) Run(addr ...string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()

	if err := a.Boot(); err != nil {
		return err
	}

	cfg, err := sconfig.GetConfig(a.Container)
	if err != nil {
		return err
	}

	port, err := cfg.Int(ConfigPortPath, 80)
	if err != nil {
		return err
	}

	logger, err := slog.GetLogger(a.Container)
	if err != nil {
		return err
	}

	_ = logger.Signal(LogChannel, slog.INFO, "App starting", map[string]interface{}{"port": port})
	if err = a.engine.Run(append([]string{fmt.Sprintf(":%d", port)}, addr...)...); err != nil {
		_ = logger.Signal(LogChannel, slog.FATAL, "App error", map[string]interface{}{"error": err.Error()})
		return err
	}

	return nil
}
