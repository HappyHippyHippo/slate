package watchdog

import (
	"fmt"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

type configurer interface {
	Partial(path string, def ...config.Partial) (config.Partial, error)
}

type logFormatterCreator interface {
	Create(cfg *config.Partial) (LogFormatter, error)
}

// Factory defines an instance of a watchdog creator, used
// to create watchdogs related to a configuration entry.
type Factory struct {
	configurer          configurer
	log                 *log.Log
	logFormatterCreator logFormatterCreator
}

// NewFactory will generate a new watchdog creator instance.
func NewFactory(
	configurer *config.Config,
	logger *log.Log,
	logFormatterCreator *LogFormatterFactory,
) (*Factory, error) {
	// check config argument reference
	if configurer == nil {
		return nil, errNilPointer("config")
	}
	// check log argument reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check formatter creator reference
	if logFormatterCreator == nil {
		return nil, errNilPointer("logFormatterCreator")
	}
	// return the created watchdog creator instance
	return &Factory{
		configurer:          configurer,
		log:                 logger,
		logFormatterCreator: logFormatterCreator,
	}, nil
}

// Create will create a new watchdog instance for the required
// service with the name passed as argument.
func (f *Factory) Create(
	service string,
) (*Watchdog, error) {
	// get service watchdog configuration
	cfg, e := f.configurer.Partial(
		fmt.Sprintf("%s.%s", ConfigPathPrefix, service),
		config.Partial{},
	)
	if e != nil {
		return nil, e
	}
	// parse the retrieved configuration
	wc := struct {
		Name    string
		Channel string
		Level   struct {
			Start string
			Error string
			Done  string
		}
		Formatter string
	}{
		Channel: LogChannel,
		Level: struct {
			Start string
			Error string
			Done  string
		}{
			Start: LogStartLevel,
			Error: LogErrorLevel,
			Done:  LogDoneLevel,
		},
		Formatter: DefaultFormatter,
	}
	if _, e = cfg.Populate("", &wc); e != nil {
		return nil, e
	}
	// validate the logging levels read from config
	startLevel, ok := log.LevelMap[wc.Level.Start]
	if !ok {
		return nil, errConversion(wc.Level.Start, "log.Level")
	}
	errorLevel, ok := log.LevelMap[wc.Level.Error]
	if !ok {
		return nil, errConversion(wc.Level.Error, "log.Level")
	}
	doneLevel, ok := log.LevelMap[wc.Level.Done]
	if !ok {
		return nil, errConversion(wc.Level.Done, "log.Level")
	}
	// obtain the formatter for the watchdog log adapter
	formatter, e := f.logFormatterCreator.Create(&config.Partial{"type": wc.Formatter})
	if e != nil {
		return nil, e
	}
	// generate the watchdog log adapter
	la, _ := NewLogAdapter(wc.Name, wc.Channel, startLevel, errorLevel, doneLevel, f.log, formatter)
	// return the generated watchdog
	return NewWatchdog(la)
}
