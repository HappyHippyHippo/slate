package rest

import (
	"fmt"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/watchdog"
)

// Process defines the REST watchdog process instance.
type Process struct {
	*watchdog.Process
}

var _ watchdog.IProcess = &Process{}

type processConfig struct {
	Watchdog string
	Port     int
	Log      struct {
		Level   string
		Channel string
		Message struct {
			Start string
			Error string
			End   string
		}
	}
}

// NewProcess will try to instantiate an REST watchdog process.
func NewProcess(
	cfgManager config.IManager,
	logger log.ILog,
	engine Engine,
) (*Process, error) {
	// check the config reference
	if cfgManager == nil {
		return nil, errNilPointer("cfgManager")
	}
	// check the log reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check the engine reference
	if engine == nil {
		return nil, errNilPointer("engine")
	}
	// get service watchdog process configuration
	cfg, e := cfgManager.Config(ConfigPath, config.Config{})
	if e != nil {
		return nil, e
	}
	// parse the retrieved configuration
	wc := processConfig{
		Watchdog: WatchdogName,
		Port:     Port,
		Log: struct {
			Level   string
			Channel string
			Message struct {
				Start string
				Error string
				End   string
			}
		}{
			Level:   LogLevel,
			Channel: LogChannel,
			Message: struct {
				Start string
				Error string
				End   string
			}{
				Start: LogStartMessage,
				Error: LogErrorMessage,
				End:   LogEndMessage,
			},
		},
	}
	_, e = cfg.Populate("", &wc)
	if e != nil {
		return nil, e
	}
	// validate the logging level read from config
	logLevel, ok := log.LevelMap[wc.Log.Level]
	if !ok {
		return nil, errConversion(wc.Log.Level, "log.Level")
	}
	// generate the watchdog process instance
	proc, _ := watchdog.NewProcess(wc.Watchdog, func() error {
		_ = logger.Signal(wc.Log.Channel, logLevel, wc.Log.Message.Start, log.Context{"port": wc.Port})
		if e = engine.Run(fmt.Sprintf(":%d", wc.Port)); e != nil {
			_ = logger.Signal(wc.Log.Channel, log.FATAL, wc.Log.Message.Error, log.Context{"error": e.Error()})
			return e
		}
		_ = logger.Signal(wc.Log.Channel, logLevel, wc.Log.Message.End)
		return nil
	})
	// return a locally defines instance of the watchdog process
	return &Process{
		Process: proc,
	}, nil
}
