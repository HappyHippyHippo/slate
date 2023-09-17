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

var _ watchdog.Processor = &Process{}

// NewProcess will try to instantiate an REST watchdog process.
func NewProcess(
	cfg *config.Config,
	logger *log.Log,
	engine Engine,
) (*Process, error) {
	// check the config reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// check the log reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check the engine reference
	if engine == nil {
		return nil, errNilPointer("engine")
	}
	// retrieve the rest configuration
	c, e := cfg.Partial(ConfigPath, config.Partial{})
	if e != nil {
		return nil, e
	}
	// parse the retrieved configuration
	wc := struct {
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
	}{
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
	if _, e := c.Populate("", &wc); e != nil {
		return nil, e
	}
	// validate the logging level read from config
	logLevel, ok := log.LevelMap[wc.Log.Level]
	if !ok {
		return nil, errConversion(wc.Log.Level, "log.Level")
	}
	// generate the watchdog process instance
	proc, _ := watchdog.NewProcess(wc.Watchdog, func() error {
		_ = logger.Signal(
			wc.Log.Channel,
			logLevel,
			wc.Log.Message.Start,
			log.Context{"port": wc.Port},
		)

		if e := engine.Run(fmt.Sprintf(":%d", wc.Port)); e != nil {
			_ = logger.Signal(
				wc.Log.Channel,
				log.FATAL,
				wc.Log.Message.Error,
				log.Context{"error": e.Error()},
			)
			return e
		}
		_ = logger.Signal(wc.Log.Channel, logLevel, wc.Log.Message.End)
		return nil
	})
	// return a locally defined instance of the watchdog process
	return &Process{
		Process: proc,
	}, nil
}
