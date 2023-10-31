package slate

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/spf13/afero"
)

// ----------------------------------------------------------------------------
// defs
// ----------------------------------------------------------------------------

const (
	// LogContainerID defines the id to be used as the Provider
	// registration id of a logger instance, as a base id of all other logger
	// package connections registered in the application Provider.
	LogContainerID = ContainerID + ".log"

	// LogFormatterContainerID defines the base if to all log formatter
	// related services
	LogFormatterContainerID = LogContainerID + ".formatter"

	// LogFormatterCreatorTag defines the tag to be assigned to all
	// Provider formatter creators.
	LogFormatterCreatorTag = LogFormatterContainerID + ".creator"

	// LogJSONEncoderCreatorContainerID defines the Provider id to be used
	// to identifies an output formatter that encodes the logging message into
	// an JSON content.
	LogJSONEncoderCreatorContainerID = LogFormatterCreatorTag + ".json"

	// LogAllFormatterCreatorsContainerID defines the id to be used as the
	// Provider registration id of an aggregate formatter creators
	// retrieval function.
	LogAllFormatterCreatorsContainerID = LogFormatterCreatorTag + ".all"

	// LogFormatterFactoryContainerID defines the id to be used as the
	// Provider registration id of a logger formatter factory instance.
	LogFormatterFactoryContainerID = LogFormatterContainerID + ".factory"

	// LogWriterContainerID defines the base if to all log writers
	// related services
	LogWriterContainerID = LogContainerID + ".writer"

	// LogWriterCreatorTag defines the tag to be assigned to all
	// Provider writer creators.
	LogWriterCreatorTag = LogWriterContainerID + ".creator"

	// LogConsoleStreamCreatorContainerID defines the Provider id to be used
	// to identifies an output writer that sends the logging messages to the
	// standard console.
	LogConsoleStreamCreatorContainerID = LogWriterCreatorTag + ".console"

	// LogFileStreamCreatorContainerID defines the Provider id to be used
	// to identifies an output writer that sends the logging messages to a
	// specific file.
	LogFileStreamCreatorContainerID = LogWriterCreatorTag + ".file"

	// LogRotatingFileStreamCreatorContainerID defines the Provider id to be
	// used to identifies an output writer that sends the logging messages to
	// a specific date related rotating file.
	LogRotatingFileStreamCreatorContainerID = LogWriterCreatorTag + ".rotating-file"

	// LogAllWriterCreatorsContainerID defines the id to be used as the
	// Provider registration id of an aggregate writer creators
	// retrieval function.
	LogAllWriterCreatorsContainerID = LogWriterCreatorTag + ".all"

	// LogWriterFactoryContainerID defines the id to be used as the Provider
	// registration id of a logger writer factory instance.
	LogWriterFactoryContainerID = LogWriterContainerID + ".factory"

	// LogLoaderContainerID defines the id to be used as the Provider
	// registration id of a logger loader instance.
	LogLoaderContainerID = LogContainerID + ".loader"

	// LogEnvID defines the base environment variable name for all
	// log related environment variables.
	LogEnvID = EnvID + "_LOG"

	// LogFormatJSON defines the value to be used to declare a
	// JSON log formatter format.
	LogFormatJSON = "json"

	// LogTypeConsole defines the value to be used to declare a
	// console log writer type.
	LogTypeConsole = "console"

	// LogTypeFile defines the value to be used to declare a
	// file log writer type.
	LogTypeFile = "file"

	// LogTypeRotatingFile defines the value to be used to declare a
	// file log writer type that rotates regarding the current date.
	LogTypeRotatingFile = "rotating-file"
)

var (
	// LogFlushFrequency defines the log auto flushing action
	// frequency time in milliseconds.
	LogFlushFrequency = EnvInt(LogEnvID+"_FLUSH_FREQUENCY", 0)

	// LogLoaderActive defines the entry config source active flag
	// used to signal the config loader to load the writers or not
	LogLoaderActive = EnvBool(LogEnvID+"_LOADER_ACTIVE", true)

	// LogLoaderConfigPath defines the entry config source path
	// to be used as the loader entry.
	LogLoaderConfigPath = EnvString(LogEnvID+"_LOADER_CONFIG_PATH", "slate.log.writers")

	// LogLoaderObserveConfig defines the loader config observing flag
	// used to register in the config object an observer of the logger
	// config entries list, so it can reload the Log writers.
	LogLoaderObserveConfig = EnvBool(LogEnvID+"_LOADER_OBSERVE_CONFIG", true)
)

// ----------------------------------------------------------------------------
// errors
// ----------------------------------------------------------------------------

var (
	// ErrInvalidLogFormat defines an error that signal an invalid
	// logger format.
	ErrInvalidLogFormat = fmt.Errorf("invalid log format")

	// ErrInvalidLogLevel defines an error that signal an invalid
	// logger level.
	ErrInvalidLogLevel = fmt.Errorf("invalid log level")

	// ErrInvalidLogConfig defines an error that signal that the
	// given logger writer config was unable to be parsed correctly
	// enabling the logger writer generation.
	ErrInvalidLogConfig = fmt.Errorf("invalid log writer config")

	// ErrLogWriterNotFound defines an error that signal that the
	// given writer was not found in the logger manager.
	ErrLogWriterNotFound = fmt.Errorf("log writer not found")

	// ErrDuplicateLogWriter defines an error that signal that the
	// requested logger writer to be registered have an id of an already
	// registered logger writer.
	ErrDuplicateLogWriter = fmt.Errorf("log writer already registered")
)

func errInvalidLogFormat(
	format string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrInvalidLogFormat, format, ctx...)
}

func errInvalidLogLevel(
	level string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrInvalidLogLevel, level, ctx...)
}

func errInvalidLogConfig(
	config ConfigPartial,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrInvalidLogConfig, fmt.Sprintf("%v", config), ctx...)
}

func errLogWriterNotFound(
	id string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrLogWriterNotFound, id, ctx...)
}

func errDuplicateLogWriter(
	id string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrDuplicateLogWriter, id, ctx...)
}

// ----------------------------------------------------------------------------
// log level
// ----------------------------------------------------------------------------

// LogLevel identifies a value type that describes a logging level.
type LogLevel int

const (
	// FATAL defines a fatal logging level.
	FATAL LogLevel = 1 + iota
	// ERROR defines a error logging level.
	ERROR
	// WARNING defines a warning logging level.
	WARNING
	// NOTICE defines a notice logging level.
	NOTICE
	// INFO defines a info logging level.
	INFO
	// DEBUG defines a debug logging level.
	DEBUG
)

// LogLevelMap defines a relation between a human-readable string
// and a code level identifier of a logging level.
var LogLevelMap = map[string]LogLevel{
	"fatal":   FATAL,
	"error":   ERROR,
	"warning": WARNING,
	"notice":  NOTICE,
	"info":    INFO,
	"debug":   DEBUG,
}

// LogLevelMapName defines a relation between a code level identifier of a
// logging level and human-readable string representation of that level.
var LogLevelMapName = map[LogLevel]string{
	FATAL:   "fatal",
	ERROR:   "error",
	WARNING: "warning",
	NOTICE:  "notice",
	INFO:    "info",
	DEBUG:   "debug",
}

// ----------------------------------------------------------------------------
// log context
// ----------------------------------------------------------------------------

// LogContext defines a value map type used by the logging methods
// to give additional information to the logger entry
type LogContext map[string]interface{}

// ----------------------------------------------------------------------------
// log formatter
// ----------------------------------------------------------------------------

// LogFormatter interface defines the methods of a logging formatter instance
// responsible to parse a logging request into the output string.
type LogFormatter interface {
	Format(level LogLevel, message string, ctx ...LogContext) string
}

// ----------------------------------------------------------------------------
// log formatter creator
// ----------------------------------------------------------------------------

// LogFormatterCreator interface defines the methods of the formatter
// factory creator that can validate creation requests and instantiation
// of particular formatter.
type LogFormatterCreator interface {
	Accept(format string) bool
	Create(args ...interface{}) (LogFormatter, error)
}

// ----------------------------------------------------------------------------
// log formatter factory
// ----------------------------------------------------------------------------

// LogFormatterFactory defines the log formatter factory structure used to
// instantiate log formatters, based on registered instance creators.
type LogFormatterFactory []LogFormatterCreator

// NewLogFormatterFactory will instantiate a new formatter factory instance
func NewLogFormatterFactory(
	creators []LogFormatterCreator,
) *LogFormatterFactory {
	factory := &LogFormatterFactory{}
	for _, creator := range creators {
		*factory = append(*factory, creator)
	}
	return factory
}

// Create will instantiate and return a new content formatter.
func (f *LogFormatterFactory) Create(
	format string,
	args ...interface{},
) (LogFormatter, error) {
	// search in the factory creator pool for one that would accept
	// to generate the requested formatter with the requested format
	for _, creator := range *f {
		if creator.Accept(format) {
			// return the creation of the requested formatter
			return creator.Create(args...)
		}
	}
	return nil, errInvalidLogFormat(format)
}

// ----------------------------------------------------------------------------
// log JSON encoder
// ----------------------------------------------------------------------------

// LogJSONEncoder defines an instance used to format a log message into
// a JSON string.
type LogJSONEncoder struct{}

var _ LogFormatter = &LogJSONEncoder{}

// NewLogJSONEncoder will create a new JSON log message encoder.
func NewLogJSONEncoder() *LogJSONEncoder {
	return &LogJSONEncoder{}
}

// Format will create the output JSON string message formatted with the
// content of the passed level, message and context
func (f LogJSONEncoder) Format(
	level LogLevel,
	message string,
	ctx ...LogContext,
) string {
	// guarantee that the content context is a valid map reference,
	// so it can be used to compose the final formatted message
	// for that initialize an empty context map, and merge all the
	// given extra contexts
	data := LogContext{}
	for _, c := range ctx {
		for k, v := range c {
			data[k] = v
		}
	}
	// store the extra time, level and message in the request context
	data["time"] = time.Now().Format("2006-01-02T15:04:05.000-0700")
	data["level"] = strings.ToUpper(LogLevelMapName[level])
	data["message"] = message
	// compose the response JSON formatted string with the populated
	// context instance
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

// ----------------------------------------------------------------------------
// log JSON encoder creator
// ----------------------------------------------------------------------------

// LogJSONEncoderCreator defines a log message JSON formatter creator.
type LogJSONEncoderCreator struct{}

var _ LogFormatterCreator = &LogJSONEncoderCreator{}

// NewLogJSONEncoderCreator generates a new JSON formatter creator service.
func NewLogJSONEncoderCreator() *LogJSONEncoderCreator {
	return &LogJSONEncoderCreator{}
}

// Accept will check if the formatter factory creator can instantiate a
// formatter of the requested format.
func (LogJSONEncoderCreator) Accept(
	format string,
) bool {
	// only accept to create a JSON format formatter
	return format == LogFormatJSON
}

// Create will instantiate the desired formatter instance.
func (LogJSONEncoderCreator) Create(
	_ ...interface{},
) (LogFormatter, error) {
	// generate the JSON formatter
	return NewLogJSONEncoder(), nil
}

// ----------------------------------------------------------------------------
// log writer
// ----------------------------------------------------------------------------

// LogWriter interface defines the interaction methods with a logging writer
// responsible to output logging message to a defined output.
type LogWriter interface {
	Signal(channel string, level LogLevel, message string, ctx ...LogContext) error
	Broadcast(level LogLevel, message string, ctx ...LogContext) error
	Flush() error

	HasChannel(channel string) bool
	ListChannels() []string
	AddChannel(channel string)
	RemoveChannel(channel string)
}

// ----------------------------------------------------------------------------
// log writer creator
// ----------------------------------------------------------------------------

// LogWriterCreator interface defines the methods of the writer
// factory creator that can validate creation requests and instantiation
// of particular type of writer.
type LogWriterCreator interface {
	Accept(config *ConfigPartial) bool
	Create(config *ConfigPartial) (LogWriter, error)
}

// ----------------------------------------------------------------------------
// log writer factory
// ----------------------------------------------------------------------------

// LogWriterFactory is a logging writer generator based on a
// registered list of writer generation creators.
type LogWriterFactory []LogWriterCreator

// NewLogWriterFactory will instantiate a new writer factory instance
func NewLogWriterFactory(
	creators []LogWriterCreator,
) *LogWriterFactory {
	factory := &LogWriterFactory{}
	for _, creator := range creators {
		*factory = append(*factory, creator)
	}
	return factory
}

// Create will instantiate and return a new log writer instance
// based on the passed config.
func (f *LogWriterFactory) Create(
	config *ConfigPartial,
) (LogWriter, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// search in the factory creators pool for one that would accept
	// to generate the requested writer with the requested format defined
	// in the given config
	for _, creator := range *f {
		if creator.Accept(config) {
			// return the creation of the requested writer
			return creator.Create(config)
		}
	}
	return nil, errInvalidLogConfig(*config)
}

// ----------------------------------------------------------------------------
// log stream
// ----------------------------------------------------------------------------

// LogStream defines the base log stream instance used as starting point to all
// log stream/writers services.
type LogStream struct {
	Channels  []string
	Level     LogLevel
	Formatter LogFormatter
	Mutex     sync.Locker
	Buffer    []string
	Writer    io.Writer
}

var _ LogWriter = &LogStream{}

// NewLogStream will instantiate a new log stream instance
func NewLogStream(
	level LogLevel,
	formatter LogFormatter,
	channels []string,
	writer io.Writer,
) *LogStream {
	return &LogStream{
		Channels:  channels,
		Level:     level,
		Formatter: formatter,
		Mutex:     &sync.Mutex{},
		Writer:    writer,
	}
}

// Close will close the stream flushing all the stored messages.
func (s *LogStream) Close() error {
	// flush the message storing buffer
	return s.Flush()
}

// Signal will process the logging signal request and store the logging request
// into the underlying writer if passing the channel and level filtering.
func (s *LogStream) Signal(
	channel string,
	level LogLevel,
	msg string,
	ctx ...LogContext,
) error {
	// search if the requested channel is in the stream channel list
	i := sort.SearchStrings(s.Channels, channel)
	if i == len(s.Channels) || s.Channels[i] != channel {
		return nil
	}
	// write the message to the stream
	return s.Broadcast(level, msg, ctx...)
}

// Broadcast will process the logging signal request and store the logging
// request into the underlying writer if passing the level filtering.
func (s *LogStream) Broadcast(
	level LogLevel,
	msg string,
	ctx ...LogContext,
) error {
	// check if the request level is higher than the associated stream level
	if s.Level < level {
		return nil
	}
	// write the message after formatting to the simple writer (stdout)
	s.Buffer = append(s.Buffer, s.Format(level, msg, ctx...))
	return nil
}

// Flush will flush all buffered log entries to the stream writer
func (s *LogStream) Flush() error {
	// output the buffer content
	for _, line := range s.Buffer {
		_, e := fmt.Fprintln(s.Writer, line)
		if e != nil {
			return e
		}
	}
	// clear the buffer
	s.Buffer = nil
	return nil
}

// HasChannel will validate if the stream is listening to a specific
// logging channel.
func (s *LogStream) HasChannel(
	channel string,
) bool {
	// search the requested string in the already ordered
	// stream channel pool list
	i := sort.SearchStrings(s.Channels, channel)
	return i < len(s.Channels) && s.Channels[i] == channel
}

// ListChannels retrieves the list of channels that the stream is listening.
func (s *LogStream) ListChannels() []string {
	return s.Channels
}

// AddChannel register a channel to the list of channels that the
// stream is listening.
func (s *LogStream) AddChannel(
	channel string,
) {
	// check if the adding channel is not already in the stream
	// channel pool list
	if !s.HasChannel(channel) {
		// add the requested channel and sort the channel pool list
		s.Channels = append(s.Channels, channel)
		sort.Strings(s.Channels)
	}
}

// RemoveChannel removes a channel from the list of channels that the
// stream is listening.
func (s *LogStream) RemoveChannel(
	channel string,
) {
	// search for the channel pool position of the channel to be removed
	i := sort.SearchStrings(s.Channels, channel)
	// check if the channel was not found
	if i == len(s.Channels) || s.Channels[i] != channel {
		return
	}
	// remove the channel from the channel pool list
	s.Channels = append(s.Channels[:i], s.Channels[i+1:]...)
}

// Format will try to format a logging message.
func (s *LogStream) Format(
	level LogLevel,
	message string,
	ctx ...LogContext,
) string {
	// check if a valid formatter reference is present, if so, return
	// the formatter response of the message content
	if s.Formatter != nil {
		message = s.Formatter.Format(level, message, ctx...)
	}
	// return just the message if no formatter is present
	return message
}

// ----------------------------------------------------------------------------
// log stream creator
// ----------------------------------------------------------------------------

// LogStreamCreator defines a base log stream creator service.
type LogStreamCreator struct {
	formatterFactory *LogFormatterFactory
}

func newLogStreamCreator(
	formatterFactory *LogFormatterFactory,
) *LogStreamCreator {
	return &LogStreamCreator{
		formatterFactory: formatterFactory,
	}
}

func (LogStreamCreator) level(
	level string,
) (LogLevel, error) {
	// check if the retrieved level string can be mapped to a
	// level typed value
	level = strings.ToLower(level)
	if _, ok := LogLevelMap[level]; !ok {
		return FATAL, errInvalidLogLevel(level)
	}
	// return the level typed value of the retrieved level string
	return LogLevelMap[level], nil
}

func (LogStreamCreator) channels(
	list []interface{},
) []string {
	var result []string
	for _, channel := range list {
		if typedChannel, ok := channel.(string); ok {
			result = append(result, typedChannel)
		}
	}
	return result
}

// ----------------------------------------------------------------------------
// log console stream
// ----------------------------------------------------------------------------

// LogConsoleStream defines an instance to a console log output stream.
type LogConsoleStream struct {
	LogStream
}

var _ LogWriter = &LogConsoleStream{}

// NewLogConsoleStream generate a new console log stream instance.
func NewLogConsoleStream(
	level LogLevel,
	formatter LogFormatter,
	channels []string,
) (*LogConsoleStream, error) {
	// check formatter argument reference
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// instantiate the console stream with the stdout as the writer target
	stream := &LogConsoleStream{
		LogStream: *NewLogStream(level, formatter, channels, os.Stdout),
	}
	// sort the assigned channels list
	sort.Strings(stream.Channels)
	return stream, nil
}

// ----------------------------------------------------------------------------
// log console stream creator
// ----------------------------------------------------------------------------

// LogConsoleStreamCreator defines a console log stream creator service.
type LogConsoleStreamCreator struct {
	LogStreamCreator
}

var _ LogWriterCreator = &LogConsoleStreamCreator{}

// NewLogConsoleStreamCreator generates a new console log stream
// generation service.
func NewLogConsoleStreamCreator(
	formatterFactory *LogFormatterFactory,
) (*LogConsoleStreamCreator, error) {
	// check formatter factory argument reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// instantiates the console stream creator
	return &LogConsoleStreamCreator{
		LogStreamCreator: *newLogStreamCreator(formatterFactory),
	}, nil
}

// Accept will check if the writer factory creator can instantiate
// a writer where the output will be sent to the console.
func (s LogConsoleStreamCreator) Accept(
	config *ConfigPartial,
) bool {
	// check config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == LogTypeConsole
}

// Create will instantiate the desired stream instance where
// the writer output is the standard console stream.
func (s LogConsoleStreamCreator) Create(
	config *ConfigPartial,
) (LogWriter, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Format   string
		Channels []interface{}
		Level    string
	}{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	level, e := s.level(sConfig.Level)
	if e != nil {
		return nil, e
	}
	// create the formatter to be given to the console stream
	formatter, e := s.formatterFactory.Create(sConfig.Format)
	if e != nil {
		return nil, e
	}
	// instantiate the console stream
	return NewLogConsoleStream(
		level,
		formatter,
		s.channels(sConfig.Channels))
}

// ----------------------------------------------------------------------------
// log file stream
// ----------------------------------------------------------------------------

// LogFileStream defines an instance to a file log output stream.
type LogFileStream struct {
	LogStream
}

var _ LogWriter = &LogFileStream{}

// NewLogFileStream generate a new file log stream instance.
func NewLogFileStream(
	level LogLevel,
	formatter LogFormatter,
	channels []string,
	writer io.Writer,
) (*LogFileStream, error) {
	// check the formatter argument reference
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// check the writer argument reference
	if writer == nil {
		return nil, errNilPointer("writer")
	}
	// instantiate the file stream
	stream := &LogFileStream{
		LogStream: *NewLogStream(level, formatter, channels, writer),
	}
	// sort the assigned channels list
	sort.Strings(stream.Channels)
	return stream, nil
}

// Close will terminate the stream stored writer instance.
func (s *LogFileStream) Close() error {
	// flush the message storing buffer
	_ = s.Flush()
	// check if the stored writer implements the closer interface
	// and close it if so
	if s.Writer != nil {
		if closer, ok := s.Writer.(io.Closer); ok {
			_ = closer.Close()
		}
		s.Writer = nil
	}
	return nil
}

// ----------------------------------------------------------------------------
// log file stream creator
// ----------------------------------------------------------------------------

// LogFileStreamCreator defines a file log stream creator service.
type LogFileStreamCreator struct {
	LogStreamCreator
	fileSystem afero.Fs
}

var _ LogWriterCreator = &LogFileStreamCreator{}

// NewLogFileStreamCreator generates a new file log stream
// generation service.
func NewLogFileStreamCreator(
	fileSystem afero.Fs,
	formatterFactory *LogFormatterFactory,
) (*LogFileStreamCreator, error) {
	// check file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check formatter factory argument reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// instantiate the file stream creator instance
	return &LogFileStreamCreator{
		LogStreamCreator: *newLogStreamCreator(formatterFactory),
		fileSystem:       fileSystem,
	}, nil
}

// Accept will check if the stream factory creator can instantiate
// a stream where the writer will output to a file.
func (s LogFileStreamCreator) Accept(
	config *ConfigPartial,
) bool {
	// check config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == LogTypeFile
}

// Create will instantiate the desired stream instance where
// the output will be sent to a file.
func (s LogFileStreamCreator) Create(
	config *ConfigPartial,
) (LogWriter, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Path     string
		Format   string
		Channels []interface{}
		Level    string
	}{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	level, e := s.level(sConfig.Level)
	if e != nil {
		return nil, e
	}
	// create the stream formatter to be given to the stream
	formatter, e := s.formatterFactory.Create(sConfig.Format)
	if e != nil {
		return nil, e
	}
	// create the stream writer
	file, e := s.fileSystem.OpenFile(
		sConfig.Path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644)
	if e != nil {
		return nil, e
	}
	// instantiate the stream
	return NewLogFileStream(
		level,
		formatter,
		s.channels(sConfig.Channels),
		file)
}

// ----------------------------------------------------------------------------
// log rotating file writer
// ----------------------------------------------------------------------------

// LogRotatingFileWriter defines an output writer used by a file stream that
// will use a dated file for target output.
type LogRotatingFileWriter struct {
	lock       sync.Locker
	fileSystem afero.Fs
	fp         afero.File
	file       string
	current    string
	year       int
	month      time.Month
	day        int
}

var _ io.Writer = &LogRotatingFileWriter{}

// NewLogRotatingFileWriter generate a new rotating file writer instance.
func NewLogRotatingFileWriter(
	fileSystem afero.Fs,
	file string,
) (io.Writer, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// instantiate the rotating file writer instance
	writer := &LogRotatingFileWriter{
		lock:       &sync.Mutex{},
		fileSystem: fileSystem,
		file:       file,
	}
	// open the target rotated file
	if e := writer.rotate(); e != nil {
		return nil, e
	}
	return writer, nil
}

// Write satisfies the io.Writer interface.
func (w *LogRotatingFileWriter) Write(
	output []byte,
) (int, error) {
	// lock the file for interaction
	w.lock.Lock()
	defer w.lock.Unlock()
	// check if the file need rotation
	if e := w.checkRotation(); e != nil {
		return 0, e
	}
	// write the content to the target file
	return w.fp.Write(output)
}

// Close satisfies the Closable interface.
func (w *LogRotatingFileWriter) Close() error {
	// close the opened file handler
	return w.fp.(io.Closer).Close()
}

func (w *LogRotatingFileWriter) checkRotation() error {
	// check if the stored opened file date for the need of rotation
	now := time.Now()
	if now.Day() != w.day || now.Month() != w.month || now.Year() != w.year {
		// rotate the file handler
		return w.rotate()
	}
	return nil
}

func (w *LogRotatingFileWriter) rotate() error {
	var e error
	// close the currently opened file
	if w.fp != nil {
		if e = w.fp.(io.Closer).Close(); e != nil {
			w.fp = nil
			return e
		}
	}
	// store the opened file date and create the new target file name
	now := time.Now()
	w.year = now.Year()
	w.month = now.Month()
	w.day = now.Day()
	w.current = fmt.Sprintf(w.file, now.Format("2006-01-02"))
	// open the new target file
	if w.fp, e = w.fileSystem.OpenFile(
		w.current,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644,
	); e != nil {
		return e
	}
	return nil
}

// ----------------------------------------------------------------------------
// log rotating file stream creator
// ----------------------------------------------------------------------------

// LogRotatingFileStreamCreator define a new rotating file log
// stream creator.
type LogRotatingFileStreamCreator struct {
	LogFileStreamCreator
}

var _ LogWriterCreator = &LogRotatingFileStreamCreator{}

// NewLogRotatingFileStreamCreator generate a new rotating file
// log stream creator.
func NewLogRotatingFileStreamCreator(
	fileSystem afero.Fs,
	formatterFactory *LogFormatterFactory,
) (*LogRotatingFileStreamCreator, error) {
	// instantiate the rotating file stream creator
	stream, e := NewLogFileStreamCreator(fileSystem, formatterFactory)
	if e != nil {
		return nil, e
	}
	return &LogRotatingFileStreamCreator{
		LogFileStreamCreator: *stream,
	}, nil
}

// Accept will check if the stream factory creator can instantiate
// a stream where the output will be to a rotating file.
func (s LogRotatingFileStreamCreator) Accept(
	config *ConfigPartial,
) bool {
	// check config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == LogTypeRotatingFile
}

// Create will instantiate the desired stream instance where
// the output is a rotating file.
func (s LogRotatingFileStreamCreator) Create(
	config *ConfigPartial,
) (LogWriter, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Path     string
		Format   string
		Channels []interface{}
		Level    string
	}{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	level, e := s.level(sConfig.Level)
	if e != nil {
		return nil, e
	}
	// create the stream formatter to be given to the console stream
	formatter, e := s.formatterFactory.Create(sConfig.Format)
	if e != nil {
		return nil, e
	}
	// create the stream writer
	file, e := NewLogRotatingFileWriter(s.fileSystem, sConfig.Path)
	if e != nil {
		return nil, e
	}
	// instantiate the stream
	return NewLogFileStream(
		level,
		formatter,
		s.channels(sConfig.Channels),
		file)
}

// ----------------------------------------------------------------------------
// log
// ----------------------------------------------------------------------------

// Log defines a new logging manager used to centralize the logging
// messaging propagation to a list of registered output writers.
type Log struct {
	writers map[string]LogWriter
	mutex   sync.Locker
	flusher Trigger
}

// NewLog instantiate a new logging manager service.
func NewLog() *Log {
	// instantiate the logger
	log := &Log{
		mutex:   &sync.Mutex{},
		writers: map[string]LogWriter{},
	}
	// check if there is a need to create the flusher trigger
	period := time.Duration(LogFlushFrequency) * time.Millisecond
	if period != 0 {
		// create the trigger used to flush the writers
		log.flusher, _ = NewTriggerRecurring(period, func() error {
			// lock the logger for handling
			log.mutex.Lock()
			defer func() { log.mutex.Unlock() }()
			// flush all writers
			for _, writer := range log.writers {
				_ = writer.Flush()
			}
			return nil
		})
	}
	return log
}

// Close will terminate all the logging writers associated to the Log.
func (l *Log) Close() error {
	// clear the writer list, forcing the close call on all of them
	l.RemoveAllWriters()
	return nil
}

// Signal will propagate the channel filtered logging request
// to all stored logging writers.
func (l *Log) Signal(
	channel string,
	level LogLevel,
	msg string,
	ctx ...LogContext,
) error {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// propagate the signal request to all registered writers
	for _, writer := range l.writers {
		if e := writer.Signal(channel, level, msg, ctx...); e != nil {
			return e
		}
	}
	return nil
}

// Broadcast will propagate the logging request to all stored logging writers.
func (l *Log) Broadcast(
	level LogLevel,
	msg string,
	ctx ...LogContext,
) error {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// propagate the broadcast request to all registered writers
	for _, writer := range l.writers {
		if e := writer.Broadcast(level, msg, ctx...); e != nil {
			return e
		}
	}
	return nil
}

// Flush will try to flush all the writers from their buffered content
func (l *Log) Flush() error {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// propagate the flush request to all registered writers
	for _, writer := range l.writers {
		if e := writer.Flush(); e != nil {
			return e
		}
	}
	return nil
}

// HasWriter check if a writer is registered with the requested id.
func (l *Log) HasWriter(
	id string,
) bool {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// check if there is a registered writer with the requested id
	_, ok := l.writers[id]
	return ok
}

// ListWriters retrieve a list of id's of all registered writers on
// the log manager.
func (l *Log) ListWriters() []string {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// generate a list with all the registered writers id's
	var list []string
	for id := range l.writers {
		list = append(list, id)
	}
	return list
}

// AddWriter registers a new writer into the Log instance.
func (l *Log) AddWriter(
	id string,
	writer LogWriter,
) error {
	// check the writer argument reference
	if writer == nil {
		return errNilPointer("writer")
	}
	// check for writer id conflict
	if l.HasWriter(id) {
		return errDuplicateLogWriter(id)
	}
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// add the writer to the logger writer pool
	l.writers[id] = writer
	return nil
}

// RemoveWriter will remove a registered writer with the requested id
// from the logging manager.
func (l *Log) RemoveWriter(
	id string,
) {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// search for the requested removing writer
	if writer, ok := l.writers[id]; ok {
		// check if the writer implements the closer interface and
		// call it if so
		if closer, ok := writer.(io.Closer); ok {
			_ = closer.Close()
		}
		// remove the writer reference from the writer pool
		delete(l.writers, id)
	}
}

// RemoveAllWriters will remove all registered writers from
// the logging manager.
func (l *Log) RemoveAllWriters() {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// iterate through all the registered writers
	for id, writer := range l.writers {
		// check if the writer implements the closer interface and
		// call it if so
		if closer, ok := writer.(io.Closer); ok {
			_ = closer.Close()
		}
		// remove the writer reference from the writer pool
		delete(l.writers, id)
	}
}

// Writer retrieve a writer from the logging manager that is
// registered with the requested id.
func (l *Log) Writer(
	id string,
) (LogWriter, error) {
	// lock the logger for handling
	l.mutex.Lock()
	defer func() { l.mutex.Unlock() }()
	// retrieve the requested writer
	if writer, ok := l.writers[id]; ok {
		return writer, nil
	}
	return nil, errLogWriterNotFound(id)
}

// ----------------------------------------------------------------------------
// log loader
// ----------------------------------------------------------------------------

// LogLoader defines the logger instantiation and initialization of a new
// logging manager.
type LogLoader struct {
	config        *Config
	log           *Log
	writerFactory *LogWriterFactory
}

// NewLogLoader generates a new logger initialization instance.
func NewLogLoader(
	config *Config,
	log *Log,
	writerFactory *LogWriterFactory,
) (*LogLoader, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// check the logger argument reference
	if log == nil {
		return nil, errNilPointer("log")
	}
	// check the writer factory argument reference
	if writerFactory == nil {
		return nil, errNilPointer("writerFactory")
	}
	// instantiate the loader
	return &LogLoader{
		config:        config,
		log:           log,
		writerFactory: writerFactory,
	}, nil
}

// Load will parse the configuration and instantiates logging writers
// depending the data on the configuration.
func (l LogLoader) Load() error {
	// retrieve the logger entries from the config instance
	entries, e := l.config.Partial(LogLoaderConfigPath, ConfigPartial{})
	if e != nil {
		return e
	}
	// load the retrieved entries
	if e := l.load(entries); e != nil {
		return e
	}
	// check if the logger writers list should be observed for updates
	if LogLoaderObserveConfig {
		// add the observer to the given config
		_ = l.config.AddObserver(
			LogLoaderConfigPath,
			func(_ interface{}, newConfig interface{}) {
				// type check the new logger config with the logging streams
				config, ok := newConfig.(ConfigPartial)
				if !ok {
					return
				}
				// remove all the current registered writers
				l.log.RemoveAllWriters()
				// load the new writer entries into the logging manager
				_ = l.load(config)
			},
		)
	}
	return nil
}

func (l LogLoader) load(
	config ConfigPartial,
) error {
	// iterate through the given logger config writer list
	for _, id := range config.Entries() {
		// get the configuration
		entry, e := config.Partial(id)
		if e != nil {
			return e
		}
		// generate the new writer instance
		writer, e := l.writerFactory.Create(&entry)
		if e != nil {
			return e
		}
		// add the writer to the logger writer pool
		if e := l.log.AddWriter(id, writer); e != nil {
			return e
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// log service register
// ----------------------------------------------------------------------------

// LogServiceRegister defines a service provider to be used on
// the application initialization to register the logging service.
type LogServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &LogServiceRegister{}

// NewLogServiceRegister will generate a new logging services registry instance
func NewLogServiceRegister(
	app ...*App,
) *LogServiceRegister {
	return &LogServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the logging module services in the
// application Provider.
func (sr LogServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("Provider")
	}
	// register the services
	_ = container.Add(LogJSONEncoderCreatorContainerID, NewLogJSONEncoderCreator, LogFormatterCreatorTag)
	_ = container.Add(LogAllFormatterCreatorsContainerID, sr.getFormatterCreators(container))
	_ = container.Add(LogFormatterFactoryContainerID, NewLogFormatterFactory)
	_ = container.Add(LogConsoleStreamCreatorContainerID, NewLogConsoleStreamCreator, LogWriterCreatorTag)
	_ = container.Add(LogFileStreamCreatorContainerID, NewLogFileStreamCreator, LogWriterCreatorTag)
	_ = container.Add(LogRotatingFileStreamCreatorContainerID, NewLogRotatingFileStreamCreator, LogWriterCreatorTag)
	_ = container.Add(LogAllWriterCreatorsContainerID, sr.getWriterCreators(container))
	_ = container.Add(LogWriterFactoryContainerID, NewLogWriterFactory)
	_ = container.Add(LogContainerID, NewLog)
	_ = container.Add(LogLoaderContainerID, NewLogLoader)
	return nil
}

// Boot will start the logging services by calling the
// log loader initialization method.
func (sr LogServiceRegister) Boot(
	container *ServiceContainer,
) (e error) {
	// check container argument reference
	if container == nil {
		return errNilPointer("Provider")
	}
	// check if the logger loader is active
	if !LogLoaderActive {
		return nil
	}
	// execute the loader action
	loader, e := sr.getLoader(container)
	if e != nil {
		return e
	}
	return loader.Load()
}

func (LogServiceRegister) getLoader(
	container *ServiceContainer,
) (*LogLoader, error) {
	// retrieve the loader entry
	entry, e := container.Get(LogLoaderContainerID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	if instance, ok := entry.(*LogLoader); ok {
		return instance, nil
	}
	return nil, errConversion(entry, "*LogLoader")
}

func (LogServiceRegister) getFormatterCreators(
	container *ServiceContainer,
) func() []LogFormatterCreator {
	return func() []LogFormatterCreator {
		// retrieve all the formatter creators from the Provider
		var creators []LogFormatterCreator
		entries, _ := container.Tag(LogFormatterCreatorTag)
		for _, entry := range entries {
			// type check the retrieved service
			creator, ok := entry.(LogFormatterCreator)
			if ok {
				creators = append(creators, creator)
			}
		}
		return creators
	}
}

func (LogServiceRegister) getWriterCreators(
	container *ServiceContainer,
) func() []LogWriterCreator {
	return func() []LogWriterCreator {
		// retrieve all the writer creators from the Provider
		var creators []LogWriterCreator
		entries, _ := container.Tag(LogWriterCreatorTag)
		for _, entry := range entries {
			// type check the retrieved service
			creator, ok := entry.(LogWriterCreator)
			if ok {
				creators = append(creators, creator)
			}
		}
		return creators
	}
}
