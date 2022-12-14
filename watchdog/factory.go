package watchdog

import (
	"fmt"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

// IFactory defines an interface to a watchdog factory instance.
type IFactory interface {
	Create(service string) (IWatchdog, error)
}

// Factory defines an instance of a watchdog factory, used
// to create watchdogs related to a configuration entry.
type Factory struct {
	config           config.IManager
	log              log.ILog
	formatterFactory ILogFormatterFactory
}

var _ IFactory = &Factory{}

// NewFactory will generate a new watchdog factory instance.
func NewFactory(
	cfg config.IManager,
	logger log.ILog,
	formatterFactory ILogFormatterFactory,
) (IFactory, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// check log argument reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check formatter factory reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// return the created watchdog factory instance
	return &Factory{
		config:           cfg,
		log:              logger,
		formatterFactory: formatterFactory,
	}, nil
}

// Create will create a new watchdog instance for the required
// service with the name passed as argument.
func (f *Factory) Create(
	service string,
) (IWatchdog, error) {
	// get service watchdog configuration
	cfg, e := f.config.Config(fmt.Sprintf("%s.%s", ConfigPathPrefix, service), config.Config{})
	if e != nil {
		return nil, e
	}
	// parse the retrieved configuration
	wc := struct {
		Service string
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
		Formatter: FormatterDefault,
	}
	_, e = cfg.Populate("", &wc)
	if e != nil {
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
	formatter, e := f.formatterFactory.Create(&config.Config{"type": wc.Formatter})
	if e != nil {
		return nil, e
	}
	// generate the watchdog log adapter
	la, _ := NewLogAdapter(wc.Service, wc.Channel, startLevel, errorLevel, doneLevel, f.log, formatter)
	// return the generated watchdog
	return NewWatchdog(la)
}
