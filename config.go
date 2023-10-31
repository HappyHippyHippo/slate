package slate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

// ----------------------------------------------------------------------------
// defs
// ----------------------------------------------------------------------------

const (
	// ConfigContainerID defines the id to be used as the provider
	// registration id of a config instance, as a base id of all other config
	// related connections registered in the application di Provider.
	ConfigContainerID = ContainerID + ".config"

	// ConfigParserContainerID defines the base if to all config
	// supplier parser related services
	ConfigParserContainerID = ConfigContainerID + ".parser"

	// ConfigParserCreatorTag defines the tag to be assigned to all
	// Provider parser creators services.
	ConfigParserCreatorTag = ConfigParserContainerID + ".creator"

	// ConfigYAMLDecoderCreatorContainerID defines the id of a YAML
	// content config parser creator service.
	ConfigYAMLDecoderCreatorContainerID = ConfigParserCreatorTag + ".yaml"

	// ConfigJSONDecoderCreatorContainerID defines the id of a JSON
	//	// content config parser creator service.
	ConfigJSONDecoderCreatorContainerID = ConfigParserCreatorTag + ".json"

	// ConfigAllParserCreatorsContainerID defines the id for an aggregation
	// of all registered services that are tagged with the config parser tag.
	ConfigAllParserCreatorsContainerID = ConfigParserCreatorTag + ".all"

	// ConfigParserFactoryContainerID defines the id to be used as the
	// Provider registration id of a config parser factory service.
	ConfigParserFactoryContainerID = ConfigParserContainerID + ".factory"

	// ConfigSupplierContainerID defines the base if to all config supplier
	// related services.
	ConfigSupplierContainerID = ConfigContainerID + ".supplier"

	// ConfigAggregateSupplierTag defines the tag to all config
	// suppliers registered in the container that should be injected in
	// the aggregate instance.
	ConfigAggregateSupplierTag = ConfigSupplierContainerID + ".aggregate"

	// ConfigAllAggregateSuppliersContainerID defines the id of the
	// list of config suppliers to be given to the config aggregation
	// supplier service.
	ConfigAllAggregateSuppliersContainerID = ConfigAggregateSupplierTag + ".all"

	// ConfigSupplierCreatorTag defines the tag to be assigned to all
	// Provider supplier creator services.
	ConfigSupplierCreatorTag = ConfigSupplierContainerID + ".creator"

	// ConfigAggregateSourceCreatorContainerID defines the id of a config
	// supplier service that retrieves the config data from an aggregation
	// of other suppliers.
	ConfigAggregateSourceCreatorContainerID = ConfigSupplierCreatorTag + ".aggregate"

	// ConfigEnvSourceCreatorContainerID defines the id of a config
	// supplier service that retrieves the config data from the environment
	// variables.
	ConfigEnvSourceCreatorContainerID = ConfigSupplierCreatorTag + ".env"

	// ConfigFileSourceCreatorContainerID defines the id of a config
	// supplier service that retrieves the config data from a file.
	ConfigFileSourceCreatorContainerID = ConfigSupplierCreatorTag + ".file"

	// ConfigObsFileSourceCreatorContainerID defines the id of a config
	// supplier service that retrieves the config data from a file that will
	// be periodically observed for changes.
	ConfigObsFileSourceCreatorContainerID = ConfigSupplierCreatorTag + ".obs-file"

	// ConfigDirSourceCreatorContainerID defines the id of a config
	// supplier service that retrieves the config data from a list of files
	// present in a directory (optionally recursive).
	ConfigDirSourceCreatorContainerID = ConfigSupplierCreatorTag + ".dir"

	// ConfigRestSourceCreatorContainerID defines the id of a config
	// supplier service that retrieves the config data from a REST
	// web service.
	ConfigRestSourceCreatorContainerID = ConfigSupplierCreatorTag + ".rest"

	// ConfigObsRestSourceCreatorContainerID defines the id of a config
	// supplier service that retrieves the config data from a REST
	// web service that will be periodically observed for changes.
	ConfigObsRestSourceCreatorContainerID = ConfigSupplierCreatorTag + ".obs-rest"

	// ConfigAllSupplierCreatorsContainerID defines the id for an aggregation
	// of all registered services that are tagged with the config supplier tag.
	ConfigAllSupplierCreatorsContainerID = ConfigSupplierCreatorTag + ".all"

	// ConfigSupplierFactoryContainerID defines the id to be used as the
	// Provider registration id config supplier factory service.
	ConfigSupplierFactoryContainerID = ConfigSupplierContainerID + ".factory"

	// ConfigLoaderContainerID defines the id to be used as the provider
	// registration id of the config loader service.
	ConfigLoaderContainerID = ConfigContainerID + ".loader"

	// ConfigEnvID defines the base environment variable name for all
	// config related environment variables.
	ConfigEnvID = EnvID + "_CONFIG"

	// ConfigFormatJSON defines the value to be used to declare
	// a JSON config supplier encoding format.
	ConfigFormatJSON = "json"

	// ConfigFormatYAML defines the value to be used to declare
	// a YAML config supplier encoding format.
	ConfigFormatYAML = "yaml"

	// ConfigTypeAggregate defines the value to be used to declare a
	// configs supplier type that provides config data from other suppliers.
	ConfigTypeAggregate = "aggregate"

	// ConfigTypeEnv defines the value to be used to
	// declare an environment config supplier type.
	ConfigTypeEnv = "env"

	// ConfigTypeFile defines the value to be used to declare a
	// file config supplier type.
	ConfigTypeFile = "file"

	// ConfigTypeObsFile defines the value to be used to
	// declare an observable file config supplier type.
	ConfigTypeObsFile = "observable-file"

	// ConfigTypeDir defines the value to be used to declare a
	// dir config supplier type.
	ConfigTypeDir = "dir"

	// ConfigTypeRest defines the value to be used to declare a
	// REST config supplier type.
	ConfigTypeRest = "rest"

	// ConfigTypeObsRest defines the value to be used to
	// declare an observable REST config supplier type.
	ConfigTypeObsRest = "observable-rest"
)

var (
	// ConfigDefaultFileFormat defines the file config supplier
	// format if the format is not present in the config.
	ConfigDefaultFileFormat = EnvString(ConfigEnvID+"_DEFAULT_FILE_FORMAT", "yaml")

	// ConfigDefaultRestFormat defines the rest config supplier
	// format if the format is not present in the config.
	ConfigDefaultRestFormat = EnvString(ConfigEnvID+"_DEFAULT_REST_FORMAT", "json")

	// ConfigPathSeparator defines the element(s) that will be used to split
	// a config path string into path elements.
	ConfigPathSeparator = EnvString(ConfigEnvID+"_PATH_SEPARATOR", ".")

	// ConfigLoaderActive defines if the config loader should be executed
	// while the provider boot
	ConfigLoaderActive = EnvBool(ConfigEnvID+"_LOADER_ACTIVE", true)

	// ConfigLoaderFileSupplierPath defines the loader config supplier path.
	ConfigLoaderFileSupplierPath = EnvString(ConfigEnvID+"_LOADER_FILE_SUPPLIER_PATH", "config/suppliers.yaml")

	// ConfigLoaderSupplierID defines the id to be used as the
	// entry config supplier id when loading the configurations.
	ConfigLoaderSupplierID = EnvString(ConfigEnvID+"_LOADER_SUPPLIER_ID", "_suppliers")

	// ConfigLoaderSupplierFormat defines the loader config supplier format.
	ConfigLoaderSupplierFormat = EnvString(ConfigEnvID+"_LOADER_SUPPLIER_FORMAT", "yaml")

	// ConfigLoaderSupplierListPath defines the loader config supplier
	// content path to be searched.
	ConfigLoaderSupplierListPath = EnvString(ConfigEnvID+"_LOADER_SUPPLIER_LIST_PATH", "slate.config.suppliers")

	// ConfigObserveFrequency defines the config observable suppliers
	// frequency time in milliseconds. Zero for no check.
	ConfigObserveFrequency = EnvInt(ConfigEnvID+"_OBSERVE_FREQUENCY", 0)
)

// ----------------------------------------------------------------------------
// errors
// ----------------------------------------------------------------------------

var (
	// ErrInvalidEmptyConfigPath defines an invalid empty path in config
	// partial assigning action.
	ErrInvalidEmptyConfigPath = fmt.Errorf("invalid empty config path")

	// ErrConfigPathNotFound defines an error that signals that a path
	// was not found when requesting a value from a partial or config.
	ErrConfigPathNotFound = fmt.Errorf("config path not found")

	// ErrInvalidConfigFormat defines an error that signals an
	// unexpected/unknown config supplier format.
	ErrInvalidConfigFormat = fmt.Errorf("invalid config format")

	// ErrInvalidConfigSupplier defines an error that signals an
	// unexpected/unknown config supplier type.
	ErrInvalidConfigSupplier = fmt.Errorf("invalid config supplier")

	// ErrConfigSupplierNotFound defines a supplier config supplier
	// not found error.
	ErrConfigSupplierNotFound = fmt.Errorf("config supplier not found")

	// ErrDuplicateConfigSupplier defines a duplicate config supplier
	// registration attempt.
	ErrDuplicateConfigSupplier = fmt.Errorf("config supplier already registered")
)

func errInvalidEmptyConfigPath(
	path string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrInvalidEmptyConfigPath, path, ctx...)
}

func errConfigPathNotFound(
	path string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrConfigPathNotFound, path, ctx...)
}

func errInvalidConfigFormat(
	format string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrInvalidConfigFormat, format, ctx...)
}

func errInvalidConfigSupplier(
	config ConfigPartial,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrInvalidConfigSupplier, fmt.Sprintf("%v", config), ctx...)
}

func errConfigSupplierNotFound(
	id string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrConfigSupplierNotFound, id, ctx...)
}

func errDuplicateConfigSupplier(
	id string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrDuplicateConfigSupplier, id, ctx...)
}

// ----------------------------------------------------------------------------
// config partial
// ----------------------------------------------------------------------------

// ConfigPartial defines a section of a configuration information
type ConfigPartial map[interface{}]interface{}

// Clone will instantiate an identical instance of the original partial
func (p *ConfigPartial) Clone() ConfigPartial {
	// recursive clone function declaration
	var cloner func(value interface{}) interface{}
	cloner = func(value interface{}) interface{} {
		switch typedValue := value.(type) {
		// recursive list scenario
		case []interface{}:
			var result []interface{}
			for _, i := range typedValue {
				result = append(result, cloner(i))
			}
			return result
		// recursive partial scenario
		case ConfigPartial:
			return typedValue.Clone()
		// scalar value
		default:
			return value
		}
	}
	// create the clone partial
	target := make(ConfigPartial)
	// clone the original partial elements to the target partial
	for key, value := range *p {
		target[key] = cloner(value)
	}
	return target
}

// Entries will retrieve the list of stored entries in the
// configuration partial.
func (p *ConfigPartial) Entries() []string {
	var entries []string
	for k := range *p {
		if key, ok := k.(string); ok {
			entries = append(entries, key)
		}
	}
	return entries
}

// Has will check if a requested path can be reached if
// request to the partial.
func (p *ConfigPartial) Has(
	path string,
) bool {
	_, e := p.path(path)
	return e == nil
}

// Set will store a value in the requested partial path.
func (p *ConfigPartial) Set(
	path string,
	value interface{},
) (*ConfigPartial, error) {
	// retrieve the path parts
	if path == "" {
		return nil, errInvalidEmptyConfigPath(path)
	}
	parts := strings.Split(path, ConfigPathSeparator)

	// iterate through the path
	it := p
	if len(parts) == 1 {
		(*it)[path] = value
		return p, nil
	}

	// path part partial generation
	generate := func(part string) {
		// check if the part exists and is a partial struct
		generate := false
		if next, ok := (*it)[part]; !ok {
			generate = true
		} else if _, ok := next.(ConfigPartial); !ok {
			generate = true
		}
		// generate the part if not present
		if generate == true {
			(*it)[part] = ConfigPartial{}
		}
	}

	// part walkthrough
	for _, part := range parts[:len(parts)-1] {
		// if the iterated path is empty
		// (double occurrence of a separator), just continue
		if part == "" {
			continue
		}
		// generate the part and advance the iteration
		generate(part)
		next := (*it)[part].(ConfigPartial)
		it = &next
	}

	// store the desired value
	part := parts[len(parts)-1:][0]
	generate(part)
	(*it)[part] = value

	return p, nil
}

// Get will retrieve the value stored in the requested path.
// If the path does not exist, then the value nil will be returned. Or, if
// a simple value was given as the optional extra argument, then it will
// be returned instead of the standard nil value.
func (p *ConfigPartial) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// retrieve the path element
	val, e := p.path(path)
	switch {
	// check for non-nil value
	case val != nil:
		return val, nil
	// check if is to return de simple value or not
	case e != nil:
		if len(def) > 0 {
			return def[0], nil
		}
		return nil, e
	// simple case : return nil
	default:
		return nil, nil
	}
}

// Bool will retrieve a value stored in the quested path
// casting it has a boolean
func (p *ConfigPartial) Bool(
	path string,
	def ...bool,
) (bool, error) {
	var val interface{}
	var e error

	// retrieve the partial value
	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return false, e
	}
	// result conversion
	if typedValue, ok := val.(bool); ok {
		return typedValue, nil
	}
	return false, errConversion(val, "bool")
}

// Int will retrieve a value stored in the quested path
// casting it has an integer
func (p *ConfigPartial) Int(
	path string,
	def ...int,
) (int, error) {
	var val interface{}
	var e error

	// retrieve the partial value
	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return 0, e
	}
	// result conversion
	if typedValue, ok := val.(int); ok {
		return typedValue, nil
	}
	return 0, errConversion(val, "int")
}

// Float will retrieve a value stored in the quested path
// casting it has a float
func (p *ConfigPartial) Float(
	path string,
	def ...float64,
) (float64, error) {
	var val interface{}
	var e error

	// retrieve the partial value
	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return 0, e
	}
	// result conversion
	if typedValue, ok := val.(float64); ok {
		return typedValue, nil
	}
	return 0, errConversion(val, "float64")
}

// String will retrieve a value stored in the quested path
// casting it has a string
func (p *ConfigPartial) String(
	path string,
	def ...string,
) (string, error) {
	var val interface{}
	var e error

	// retrieve the partial value
	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return "", e
	}
	// result conversion
	if typedValue, ok := val.(string); ok {
		return typedValue, nil
	}
	return "", errConversion(val, "string")
}

// List will retrieve a value stored in the quested path
// casting it has a list
func (p *ConfigPartial) List(
	path string,
	def ...[]interface{},
) ([]interface{}, error) {
	var val interface{}
	var e error

	// retrieve the partial value
	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return nil, e
	}
	// result conversion
	if typedValue, ok := val.([]interface{}); ok {
		return typedValue, nil
	}
	return nil, errConversion(val, "[]interface{}")
}

// Partial will retrieve a value stored in the quested path
// casting it has a config partial
func (p *ConfigPartial) Partial(
	path string,
	def ...ConfigPartial,
) (ConfigPartial, error) {
	var val interface{}
	var e error

	// retrieve the partial value
	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return nil, e
	}
	// result conversion
	if typedValue, ok := val.(ConfigPartial); ok {
		return typedValue, nil
	}
	return nil, errConversion(val, "ConfigPartial")
}

// Populate will try to populate the data argument with the data stored
// in the path partial location.
func (p *ConfigPartial) Populate(
	path string,
	data interface{},
	insensitive ...bool,
) (interface{}, error) {
	// check the case-insensitive flag
	isInsensitive := false
	if len(insensitive) == 0 || insensitive[0] == true {
		isInsensitive = true
		path = strings.ToLower(path)
	}
	// retrieve the partial value
	value, e := p.Get(path)
	// error retrieving the path value
	if e != nil {
		return nil, e
	}
	// call recursive data population method
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		return p.populate(value, reflect.ValueOf(data).Elem(), isInsensitive)
	}
	return p.populate(value, reflect.New(dataType).Elem(), isInsensitive)
}

// Merge will increment the current partial instance with the
// information stored in another partial.
func (p *ConfigPartial) Merge(
	src ConfigPartial,
) {
	// try to Merge every supplier stored element into the target partial
	for key, value := range src {
		// if the key does not exist in the target partial, just store it
		if local, ok := (*p)[key]; !ok {
			(*p)[key] = value
		} else {
			// check if the 2 are partials
			typedLocal, okLocal := local.(ConfigPartial)
			typedValue, okValue := value.(ConfigPartial)
			if okLocal && okValue {
				// Merge the both partials
				typedLocal.Merge(typedValue)
			} else {
				// just override the target value
				(*p)[key] = value
			}
		}
	}
}

func (p *ConfigPartial) path(
	path string,
) (interface{}, error) {
	var ok bool
	var it interface{}

	// iterate through the path
	it = *p
	for _, part := range strings.Split(path, ConfigPathSeparator) {
		// if the iterated path is empty
		// (double occurrence of a separator), just continue
		if part == "" {
			continue
		}
		switch typedIt := it.(type) {
		// check if the iterated part references a partial
		case ConfigPartial:
			// retrieve the part reference
			if it, ok = typedIt[part]; !ok {
				return nil, errConfigPathNotFound(path)
			}
		default:
			return nil, errConfigPathNotFound(path)
		}
	}
	return it, nil
}

func (p *ConfigPartial) populate(
	source interface{},
	target reflect.Value,
	insensitive bool,
) (interface{}, error) {
	// get the types of the supplier and target
	sourceType := reflect.TypeOf(source)
	targetType := target.Type()
	// if the types are the same, just return the supplier
	if sourceType == targetType {
		return source, nil
	}
	// supplier type action
	switch sourceType {
	case reflect.TypeOf(ConfigPartial{}):
		// iterate through all the target fields to be assigned
		for i := 0; i < target.NumField(); i++ {
			// get the field value and type
			fieldValue := target.Field(i)
			fieldType := target.Type().Field(i)
			// check if the field is exported
			if !fieldType.IsExported() {
				continue
			}
			// check if the retrieved configuration value can be
			// assigned to the field
			if fieldValue.CanAddr() {
				switch fieldValue.Kind() {
				case reflect.Struct:
					// get the configuration value
					path := fieldType.Name
					if insensitive {
						path = strings.ToLower(path)
					}
					config := source.(ConfigPartial)
					data, e := config.Partial(path)
					if e != nil {
						continue
					}
					// recursive assignment
					if _, e := p.populate(data, fieldValue, insensitive); e != nil {
						return nil, e
					}
				default:
					// get the configuration value
					path := fieldType.Name
					if insensitive {
						path = strings.ToLower(path)
					}
					config := source.(ConfigPartial)
					data, e := config.Get(path)
					if e != nil {
						continue
					}
					// assign the configuration value to the field
					if reflect.TypeOf(data) != fieldType.Type {
						return nil, errConversion(data, fieldType.Type.Name())
					}
					fieldValue.Set(reflect.ValueOf(data))
				}
			}
		}
	default:
		return nil, errConversion(source, targetType.Name())
	}
	return target.Interface(), nil
}

// ConfigConvert will convert the given value to a config partial
// normalized value. This means that the map based values will be
// cast as a config partial instance and for any other ones
// (scalar or lists), will be processed accordingly.
// The conversion is recursive, at all levels.
func ConfigConvert(
	val interface{},
) interface{} {
	// recursive conversion call
	if pValue, ok := val.(ConfigPartial); ok {
		// return the recursive conversion of the partial
		p := ConfigPartial{}
		for k, value := range pValue {
			// turn all string keys into lowercase
			stringKey, ok := k.(string)
			if ok {
				p[strings.ToLower(stringKey)] = ConfigConvert(value)
			} else {
				p[k] = ConfigConvert(value)
			}
		}
		return p
	}
	if lValue, ok := val.([]interface{}); ok {
		var result []interface{}
		for _, i := range lValue {
			result = append(result, ConfigConvert(i))
		}
		return result
	}
	// check if the value is a map that can be converted to a partial
	if mValue, ok := val.(map[string]interface{}); ok {
		// return the recursive conversion of the partial
		result := ConfigPartial{}
		for k, i := range mValue {
			// turn all string keys into lowercase
			result[strings.ToLower(k)] = ConfigConvert(i)
		}
		return result
	}
	// check if the value is a map that can be converted to a partial
	if mValue, ok := val.(map[interface{}]interface{}); ok {
		// return the recursive conversion of the partial
		result := ConfigPartial{}
		for k, i := range mValue {
			// turn all string keys into lowercase
			stringKey, ok := k.(string)
			if ok {
				result[strings.ToLower(stringKey)] = ConfigConvert(i)
			} else {
				result[k] = ConfigConvert(i)
			}
		}
		return result
	}
	// check if the value is a float64 but with an integer value
	if fValue, ok := val.(float64); ok {
		if float64(int(fValue)) == fValue {
			return int(fValue)
		}
	}
	// check if the value is a float32 but with an integer value
	if fValue, ok := val.(float32); ok {
		if float32(int(fValue)) == fValue {
			return int(fValue)
		}
	}
	return val
}

// ----------------------------------------------------------------------------
// config parser
// ----------------------------------------------------------------------------

// ConfigParser interface defines the interaction methods of a
// content parser used to parse/convert the config supplier content
// into an application usable config partial instance.
type ConfigParser interface {
	Parse() (*ConfigPartial, error)
}

// ----------------------------------------------------------------------------
// config parser creator
// ----------------------------------------------------------------------------

// ConfigParserCreator interface defines the methods of the parser
// factory creator that can validate creation requests and instantiation
// of a particular requested parser, if enable to do so.
type ConfigParserCreator interface {
	Accept(format string) bool
	Create(args ...interface{}) (ConfigParser, error)
}

// ----------------------------------------------------------------------------
// config parser factory
// ----------------------------------------------------------------------------

// ConfigParserFactory defines a config parser instantiation factory object.
type ConfigParserFactory []ConfigParserCreator

// NewConfigParserFactory will instantiate a new parser factory instance.
func NewConfigParserFactory(
	creators []ConfigParserCreator,
) *ConfigParserFactory {
	factory := &ConfigParserFactory{}
	for _, creator := range creators {
		*factory = append(*factory, creator)
	}
	return factory
}

// Create will instantiate the requested new parser capable to
// parse the formatted content into a usable config partial.
func (f *ConfigParserFactory) Create(
	format string,
	args ...interface{},
) (ConfigParser, error) {
	// find a stored creator that will accept the requested format
	for _, creator := range *f {
		if creator.Accept(format) {
			// return the parser instantiation
			return creator.Create(args...)
		}
	}
	// signal that no creator was found that would accept the
	// requested format
	return nil, errInvalidConfigFormat(format)
}

// ----------------------------------------------------------------------------
// config underlying decoder
// ----------------------------------------------------------------------------

// ConfigUnderlyingDecoder defines the interface to a content parser
// underlying decoder instance.
type ConfigUnderlyingDecoder interface {
	Decode(config interface{}) error
}

// ----------------------------------------------------------------------------
// config decoder
// ----------------------------------------------------------------------------

// ConfigDecoder defines the base structure to a config content
// decoder instance.
type ConfigDecoder struct {
	Reader            io.Reader
	UnderlyingDecoder ConfigUnderlyingDecoder
}

var _ ConfigParser = &ConfigDecoder{}

// NewConfigDecoder will instantiate a new base config decoder instance.
func NewConfigDecoder(
	reader io.Reader,
	decoder ConfigUnderlyingDecoder,
) *ConfigDecoder {
	return &ConfigDecoder{
		Reader:            reader,
		UnderlyingDecoder: decoder,
	}
}

// Close terminate the decoder, closing the associated underlying decoder.
func (d *ConfigDecoder) Close() error {
	// check if there is a jsonReader reference
	if d.Reader != nil {
		// check if the underlying decoder implements the closer interface
		if closer, ok := d.Reader.(io.Closer); ok {
			// close the underlying decoder
			if e := closer.Close(); e != nil {
				return e
			}
		}
		d.Reader = nil
	}
	return nil
}

// Parse the associated configuration supplier encoded content
// into a configuration partial instance.
func (d *ConfigDecoder) Parse() (*ConfigPartial, error) {
	// decode the referenced data
	data := map[string]interface{}{}
	if e := d.UnderlyingDecoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized partial
	result := ConfigConvert(data).(ConfigPartial)
	return &result, nil
}

// ----------------------------------------------------------------------------
// config json decoder
// ----------------------------------------------------------------------------

// ConfigJSONDecoder defines a config supplier JSON content decoder instance.
type ConfigJSONDecoder struct {
	ConfigDecoder
}

var _ ConfigParser = &ConfigJSONDecoder{}

// NewConfigJSONDecoder will instantiate a new JSON config content decoder.
func NewConfigJSONDecoder(
	reader io.Reader,
) (*ConfigJSONDecoder, error) {
	// validate the reader reference
	if reader == nil {
		return nil, errNilPointer("reader")
	}
	// return the new decoder reference
	return &ConfigJSONDecoder{
		ConfigDecoder: *NewConfigDecoder(reader, json.NewDecoder(reader)),
	}, nil
}

// ----------------------------------------------------------------------------
// config decoder json creator
// ----------------------------------------------------------------------------

// ConfigJSONDecoderCreator defines a config JSON decoder
// instantiation creator
type ConfigJSONDecoderCreator struct{}

var _ ConfigParserCreator = &ConfigJSONDecoderCreator{}

// NewConfigJSONDecoderCreator will instantiate a new JSON format
// config decoder creator
func NewConfigJSONDecoderCreator() *ConfigJSONDecoderCreator {
	return &ConfigJSONDecoderCreator{}
}

// Accept will check if the requested format is accepted by the created parser.
func (ConfigJSONDecoderCreator) Accept(
	format string,
) bool {
	// only accepts JSON format
	return format == ConfigFormatJSON
}

// Create will instantiate the desired decoder instance with the given JSON
// underlying decoder instance that will decode the supplier content.
func (ConfigJSONDecoderCreator) Create(
	args ...interface{},
) (ConfigParser, error) {
	// check for the existence of the mandatory reader argument
	if len(args) == 0 {
		return nil, errNilPointer("args[0]")
	}
	// validate the reader argument
	if reader, ok := args[0].(io.Reader); ok {
		return NewConfigJSONDecoder(reader)
	}
	return nil, errConversion(args[0], "io.Reader")
}

// ----------------------------------------------------------------------------
// config yaml decoder
// ----------------------------------------------------------------------------

// ConfigYAMLDecoder defines a config supplier YAML content decoder instance.
type ConfigYAMLDecoder struct {
	ConfigDecoder
}

var _ ConfigParser = &ConfigYAMLDecoder{}

// NewConfigYAMLDecoder will instantiate a new YAML config content decoder.
func NewConfigYAMLDecoder(
	reader io.Reader,
) (*ConfigYAMLDecoder, error) {
	// validate the reader reference
	if reader == nil {
		return nil, errNilPointer("reader")
	}
	// return the new decoder reference
	return &ConfigYAMLDecoder{
		ConfigDecoder: *NewConfigDecoder(reader, yaml.NewDecoder(reader)),
	}, nil
}

// Parse the associated config supplier encoded content
// into a config partial instance.
func (d ConfigYAMLDecoder) Parse() (*ConfigPartial, error) {
	// decode the referenced reader data
	data := ConfigPartial{}
	if e := d.UnderlyingDecoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized config
	result := ConfigConvert(data).(ConfigPartial)
	return &result, nil
}

// ----------------------------------------------------------------------------
// config decoder yaml creator
// ----------------------------------------------------------------------------

// ConfigYAMLDecoderCreator defines a YAML config decoder
// instantiation creator
type ConfigYAMLDecoderCreator struct{}

var _ ConfigParserCreator = &ConfigYAMLDecoderCreator{}

// NewConfigYAMLDecoderCreator will instantiate a new YAML format
// config decoder creator
func NewConfigYAMLDecoderCreator() *ConfigYAMLDecoderCreator {
	return &ConfigYAMLDecoderCreator{}
}

// Accept will check if the requested format is accepted by the created parser.
func (ConfigYAMLDecoderCreator) Accept(
	format string,
) bool {
	// only accepts YAML format
	return format == ConfigFormatYAML
}

// Create will instantiate the desired decoder instance with the given YAML
// underlying decoder instance that will decode the supplier content.
func (ConfigYAMLDecoderCreator) Create(
	args ...interface{},
) (ConfigParser, error) {
	// check for the existence of the mandatory reader argument
	if len(args) == 0 {
		return nil, errNilPointer("args[0]")
	}
	// validate the reader argument
	if reader, ok := args[0].(io.Reader); ok {
		return NewConfigYAMLDecoder(reader)
	}
	return nil, errConversion(args[0], "io.Reader")
}

// ----------------------------------------------------------------------------
// config supplier
// ----------------------------------------------------------------------------

// ConfigSupplier defines the interface of a config partial supplier.
type ConfigSupplier interface {
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
}

// ConfigObsSupplier interface extends the ConfigSupplier interface with
// methods specific to suppliers that will be checked for updates in a
// regular periodicity defined in the config object where the supplier
// will be registered.
type ConfigObsSupplier interface {
	ConfigSupplier
	Reload() (bool, error)
}

// ----------------------------------------------------------------------------
// config supplier creator
// ----------------------------------------------------------------------------

// ConfigSupplierCreator interface defines the methods of the supplier
// factory creator that will be used to instantiate a particular supplier
// responsible to handle a type of source.
type ConfigSupplierCreator interface {
	Accept(config *ConfigPartial) bool
	Create(config *ConfigPartial) (ConfigSupplier, error)
}

// ----------------------------------------------------------------------------
// config supplier factory
// ----------------------------------------------------------------------------

// ConfigSupplierFactory defines an object responsible to instantiate a
// new config supplier.
type ConfigSupplierFactory []ConfigSupplierCreator

// NewConfigSupplierFactory will instantiate a new config supplier
// factory instance.
func NewConfigSupplierFactory(
	creators []ConfigSupplierCreator,
) *ConfigSupplierFactory {
	factory := &ConfigSupplierFactory{}
	for _, creator := range creators {
		*factory = append(*factory, creator)
	}
	return factory
}

// Create will instantiate and return a new config supplier based on
// the given configuration data.
func (f *ConfigSupplierFactory) Create(
	config *ConfigPartial,
) (ConfigSupplier, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// find a creator that accepts the requested supplier type
	for _, creator := range *f {
		if creator.Accept(config) {
			// create the requested config supplier
			return creator.Create(config)
		}
	}
	return nil, errInvalidConfigSupplier(*config)
}

// ----------------------------------------------------------------------------
// config source
// ----------------------------------------------------------------------------

// ConfigSource defines a base config source structure used to implement the
// basics of a config source instance.
type ConfigSource struct {
	Mutex   sync.Locker
	Partial ConfigPartial
}

var _ ConfigSupplier = &ConfigSource{}

// NewConfigSource will instantiate a new config supplier base structure.
func NewConfigSource() *ConfigSource {
	return &ConfigSource{
		Mutex:   &sync.Mutex{},
		Partial: ConfigPartial{},
	}
}

// Has will check if the requested path is present in the supplier
// configuration content.
func (s *ConfigSource) Has(
	path string,
) bool {
	// lock the supplier for changes
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	// check if the supplier stored config has the requested path
	return s.Partial.Has(path)
}

// Get will retrieve the value stored in the requested path present in the
// configuration content.
// If the path does not exist, then the value nil will be returned, or if
// a default value is given, that will be the resulting value.
// This method will mostly be used by the config object to obtain the full
// content of the supplier to aggregate all the data into his internal storing
// config instance.
func (s *ConfigSource) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// lock the supplier for changes
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	// retrieve the supplier stored config path stored value
	return s.Partial.Get(path, def...)
}

// ----------------------------------------------------------------------------
// config aggregate source
// ----------------------------------------------------------------------------

// ConfigAggregateSource defines a config supplier that aggregates a list of
// config suppliers into a single aggregate provided supplier.
type ConfigAggregateSource struct {
	ConfigSource
	suppliers []ConfigSupplier
}

var _ ConfigSupplier = &ConfigAggregateSource{}

// NewConfigAggregateSource will instantiate a new config supplier
// that aggregate a list of suppliers connections.
func NewConfigAggregateSource(
	suppliers []ConfigSupplier,
) (*ConfigAggregateSource, error) {
	// instantiates the config supplier
	source := &ConfigAggregateSource{
		ConfigSource: *NewConfigSource(),
		suppliers:    suppliers,
	}
	// load the config information from the passed suppliers
	if e := source.load(); e != nil {
		return nil, e
	}
	return source, nil
}

func (c *ConfigAggregateSource) load() error {
	// merge all the config suppliers partials given into the local config
	partial := ConfigPartial{}
	for _, supplier := range c.suppliers {
		supplierPartial, e := supplier.Get("", ConfigPartial{})
		if e != nil {
			return e
		}
		partial.Merge(supplierPartial.(ConfigPartial))
	}
	// assign the merged suppliers to the supplier config
	c.Mutex.Lock()
	c.Partial = partial
	c.Mutex.Unlock()
	return nil
}

// ----------------------------------------------------------------------------
// config aggregate source creator
// ----------------------------------------------------------------------------

// ConfigAggregateSourceCreator defines a creator used to instantiate
// a config aggregation config supplier.
type ConfigAggregateSourceCreator struct {
	suppliers []ConfigSupplier
}

var _ ConfigSupplierCreator = &ConfigAggregateSourceCreator{}

// NewConfigAggregateSourceCreator instantiate the desired aggregation config
// creator instance.
func NewConfigAggregateSourceCreator(
	suppliers []ConfigSupplier,
) (*ConfigAggregateSourceCreator, error) {
	// instantiate the creator instance
	return &ConfigAggregateSourceCreator{
		suppliers: suppliers,
	}, nil
}

// Accept will check if the requested supplier can be instantiated by this
// creator by parsing the given config partial.
func (s ConfigAggregateSourceCreator) Accept(
	config *ConfigPartial,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == ConfigTypeAggregate
}

// Create will instantiate the desired aggregation supplier instance
// with the creator stored list of suppliers.
func (s ConfigAggregateSourceCreator) Create(
	_ *ConfigPartial,
) (ConfigSupplier, error) {
	// create the aggregate config supplier
	return NewConfigAggregateSource(s.suppliers)
}

// ----------------------------------------------------------------------------
// config env source
// ----------------------------------------------------------------------------

// ConfigEnvSource defines a config supplier that maps environment
// variables values to a config.
type ConfigEnvSource struct {
	ConfigSource
	mappings map[string]string
}

var _ ConfigSupplier = &ConfigEnvSource{}

// NewConfigEnvSource will instantiate a new configuration supplier
// that will map environmental variables to configuration
// path values.
func NewConfigEnvSource(
	mappings map[string]string,
) (*ConfigEnvSource, error) {
	// instantiate the supplier
	source := &ConfigEnvSource{
		ConfigSource: *NewConfigSource(),
		mappings:     mappings,
	}
	// load the supplier values from environment
	_ = source.load()
	return source, nil
}

func (s *ConfigEnvSource) load() error {
	// iterate through all the supplier mappings
	for key, path := range s.mappings {
		// retrieve the mapped value from the environment
		env := os.Getenv(key)
		if env == "" {
			continue
		}
		// navigate to the target storing path of the environment value
		step := s.Partial
		sections := strings.Split(path, ".")
		for i, section := range sections {
			if i != len(sections)-1 {
				// Convert the path section if is present and not a config
				if _, ok := step[section]; ok == false {
					step[section] = ConfigPartial{}
				}
				// create the section if not present
				// and iterate to the section
				step[section] = ConfigPartial{}
				step = step[section].(ConfigPartial)
			} else {
				// store the value in the target section
				step[section] = env
			}
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// config env source creator
// ----------------------------------------------------------------------------

// ConfigEnvSourceCreator defines a supplier creator used to instantiate an
// environment variable mapped config supplier.
type ConfigEnvSourceCreator struct{}

var _ ConfigSupplierCreator = &ConfigEnvSourceCreator{}

// NewConfigEnvSourceCreator instantiates a new environment config
// supplier creator.
func NewConfigEnvSourceCreator() *ConfigEnvSourceCreator {
	return &ConfigEnvSourceCreator{}
}

// Accept will check if the requested supplier can be instantiated by this
// creator by parsing the given config partial.
func (s ConfigEnvSourceCreator) Accept(
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
	return sConfig.Type == ConfigTypeEnv
}

// Create will instantiate the desired environment variable mapper
// supplier instance with the passed mappings.
func (s ConfigEnvSourceCreator) Create(
	config *ConfigPartial,
) (ConfigSupplier, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct{ Mappings ConfigPartial }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// create the mappings map
	mapping := make(map[string]string)
	for k, value := range sConfig.Mappings {
		typedKey, ok := k.(string)
		if !ok {
			return nil, errConversion(k, "string")
		}
		typedValue, ok := value.(string)
		if !ok {
			return nil, errConversion(value, "string")
		}
		mapping[typedKey] = typedValue
	}
	// create the config supplier
	return NewConfigEnvSource(mapping)
}

// ----------------------------------------------------------------------------
// config file source
// ----------------------------------------------------------------------------

// ConfigFileSource defines a config supplier that obtains the config
// content from a system file.
type ConfigFileSource struct {
	ConfigSource
	path          string
	format        string
	fileSystem    afero.Fs
	parserFactory *ConfigParserFactory
}

var _ ConfigSupplier = &ConfigFileSource{}

// NewConfigFileSource will instantiate a new configuration supplier
// that will read a file for its configuration info.
func NewConfigFileSource(
	path,
	format string,
	fileSystem afero.Fs,
	parserFactory *ConfigParserFactory,
) (*ConfigFileSource, error) {
	// check file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiates the config supplier
	source := &ConfigFileSource{
		ConfigSource:  *NewConfigSource(),
		path:          path,
		format:        format,
		fileSystem:    fileSystem,
		parserFactory: parserFactory,
	}
	// Load the file config content
	if e := source.load(); e != nil {
		return nil, e
	}
	return source, nil
}

func (s *ConfigFileSource) load() error {
	// open the supplier source file
	file, e := s.fileSystem.OpenFile(s.path, os.O_RDONLY, 0o644)
	if e != nil {
		return e
	}
	// creates the file content parser instance
	parser, e := s.parserFactory.Create(s.format, file)
	if e != nil {
		_ = file.Close()
		return e
	}
	defer func() {
		if closer, ok := parser.(io.Closer); ok {
			_ = closer.Close()
		}
	}()
	// decode the file content
	partial, e := parser.Parse()
	if e != nil {
		return e
	}
	// store the parsed content into the supplier local config
	s.Mutex.Lock()
	s.Partial = *partial
	s.Mutex.Unlock()
	return nil
}

// ----------------------------------------------------------------------------
// config file source creator
// ----------------------------------------------------------------------------

// ConfigFileSourceCreator defines a config supplier creator used to
// instantiate a file config supplier.
type ConfigFileSourceCreator struct {
	fileSystem    afero.Fs
	parserFactory *ConfigParserFactory
}

var _ ConfigSupplierCreator = &ConfigFileSourceCreator{}

// NewConfigFileSourceCreator instantiates a new file config supplier
// creator.
func NewConfigFileSourceCreator(
	fileSystem afero.Fs,
	parserFactory *ConfigParserFactory,
) (*ConfigFileSourceCreator, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiate the creator
	return &ConfigFileSourceCreator{
		fileSystem:    fileSystem,
		parserFactory: parserFactory,
	}, nil
}

// Accept will check if the requested supplier can be instantiated by this
// creator by parsing the given config partial.
func (s ConfigFileSourceCreator) Accept(
	config *ConfigPartial,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == ConfigTypeFile
}

// Create will instantiate the desired file supplier instance.
func (s ConfigFileSourceCreator) Create(
	config *ConfigPartial,
) (ConfigSupplier, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Path   string
		Format string
	}{
		Format: ConfigDefaultFileFormat,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	if sConfig.Path == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing path",
		})
	}
	// create the config supplier
	return NewConfigFileSource(
		sConfig.Path,
		sConfig.Format,
		s.fileSystem,
		s.parserFactory,
	)
}

// ----------------------------------------------------------------------------
// config observable file source
// ----------------------------------------------------------------------------

// ConfigObsFileSource defines a config supplier that read a file content
// and stores its config contents to be used as a config.
// The supplier will also be checked for changes recurrently, so it can
// update the stored configuration information.
type ConfigObsFileSource struct {
	ConfigFileSource
	timestamp time.Time
}

var _ ConfigSupplier = &ConfigObsFileSource{}
var _ ConfigObsSupplier = &ConfigObsFileSource{}

// NewConfigObsFileSource will instantiate a new configuration supplier
// that will read a file for configuration info, opening the
// possibility for on-the-fly update on supplier content change.
func NewConfigObsFileSource(
	path,
	format string,
	fileSystem afero.Fs,
	parserFactory *ConfigParserFactory,
) (*ConfigObsFileSource, error) {
	// check file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// return the requested observable config supplier instance
	source := &ConfigObsFileSource{
		ConfigFileSource: ConfigFileSource{
			ConfigSource:  *NewConfigSource(),
			path:          path,
			format:        format,
			fileSystem:    fileSystem,
			parserFactory: parserFactory,
		},
		timestamp: time.Unix(0, 0),
	}
	// Load the file config content
	if _, e := source.Reload(); e != nil {
		return nil, e
	}
	return source, nil
}

// Reload will check if the supplier has been updated, and, if so,
// reload the supplier config content.
func (s *ConfigObsFileSource) Reload() (bool, error) {
	// get the file stats, so we can store the modification time
	fileStats, e := s.fileSystem.Stat(s.path)
	if e != nil {
		return false, e
	}
	// check if the file modification time is greater than the stored one
	modTime := fileStats.ModTime()
	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(modTime) {
		// load the file content
		if e := s.load(); e != nil {
			return false, e
		}
		// update the stored config content modification time
		s.timestamp = modTime
		return true, nil
	}
	return false, nil
}

// ----------------------------------------------------------------------------
// config observable file source creator
// ----------------------------------------------------------------------------

// ConfigObsFileSourceCreator defines a supplier creator used to instantiate
// a new observable file supplier.
type ConfigObsFileSourceCreator struct {
	ConfigFileSourceCreator
}

var _ ConfigSupplierCreator = &ConfigObsFileSourceCreator{}

// NewConfigObsFileSourceCreator instantiates a new observable
// file config supplier creator instance.
func NewConfigObsFileSourceCreator(
	fileSystem afero.Fs,
	parserFactory *ConfigParserFactory,
) (*ConfigObsFileSourceCreator, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiate the creator
	return &ConfigObsFileSourceCreator{
		ConfigFileSourceCreator: ConfigFileSourceCreator{
			fileSystem:    fileSystem,
			parserFactory: parserFactory,
		},
	}, nil
}

// Accept will check if the requested supplier can be instantiated by this
// creator by parsing the given config partial.
func (s ConfigObsFileSourceCreator) Accept(
	config *ConfigPartial,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == ConfigTypeObsFile
}

// Create will instantiate the desired observable file supplier instance.
func (s ConfigObsFileSourceCreator) Create(
	config *ConfigPartial,
) (ConfigSupplier, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Path   string
		Format string
	}{
		Format: ConfigDefaultFileFormat,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	if sConfig.Path == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing path",
		})
	}
	// create the config supplier
	return NewConfigObsFileSource(
		sConfig.Path,
		sConfig.Format,
		s.fileSystem,
		s.parserFactory,
	)
}

// ----------------------------------------------------------------------------
// config dir source
// ----------------------------------------------------------------------------

// ConfigDirSource defines a config supplier that read all directory files,
// recursive or not, and parse each one and store all the read content
// as a config.
type ConfigDirSource struct {
	ConfigSource
	path          string
	format        string
	recursive     bool
	fileSystem    afero.Fs
	parserFactory *ConfigParserFactory
}

var _ ConfigSupplier = &ConfigDirSource{}

// NewConfigDirSource will instantiate a new configuration supplier
// that will read a directory files for configuration information.
func NewConfigDirSource(
	path,
	format string,
	recursive bool,
	fileSystem afero.Fs,
	parserFactory *ConfigParserFactory,
) (*ConfigDirSource, error) {
	// check file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiates the config supplier
	source := &ConfigDirSource{
		ConfigSource:  *NewConfigSource(),
		path:          path,
		format:        format,
		recursive:     recursive,
		fileSystem:    fileSystem,
		parserFactory: parserFactory,
	}
	// load the dir files config content
	if e := source.load(); e != nil {
		return nil, e
	}
	return source, nil
}

func (s *ConfigDirSource) load() error {
	// load the supplier directory contents
	partial, e := s.loadDir(s.path)
	if e != nil {
		return e
	}
	// store the parsed content into the supplier local config
	s.Mutex.Lock()
	s.Partial = *partial
	s.Mutex.Unlock()
	return nil
}

func (s *ConfigDirSource) loadDir(
	path string,
) (*ConfigPartial, error) {
	// open the directory stream
	dir, e := s.fileSystem.Open(path)
	if e != nil {
		return nil, e
	}
	defer func() { _ = dir.Close() }()
	// get the dir entry list
	files, e := dir.Readdir(0)
	if e != nil {
		return nil, e
	}
	// parse each founded entry
	loaded := &ConfigPartial{}
	for _, file := range files {
		// check if is an inner directory
		if file.IsDir() {
			// load the founded directory if the supplier is
			// configured to be recursive
			if s.recursive {
				partial, e := s.loadDir(path + "/" + file.Name())
				if e != nil {
					return nil, e
				}
				// merge the loaded content
				loaded.Merge(*partial)
			}
		} else {
			// load the file content
			partial, e := s.loadFile(path + "/" + file.Name())
			if e != nil {
				return nil, e
			}
			// merge the loaded content
			loaded.Merge(*partial)
		}
	}
	return loaded, nil
}

func (s *ConfigDirSource) loadFile(
	path string,
) (*ConfigPartial, error) {
	// open the file for reading
	file, e := s.fileSystem.OpenFile(path, os.O_RDONLY, 0o644)
	if e != nil {
		return nil, e
	}
	// get a parser instance to parse the file content
	parser, e := s.parserFactory.Create(s.format, file)
	if e != nil {
		_ = file.Close()
		return nil, e
	}
	defer func() {
		if closer, ok := parser.(io.Closer); ok {
			_ = closer.Close()
		}
	}()
	// decode the file content
	return parser.Parse()
}

// ----------------------------------------------------------------------------
// config dir source creator
// ----------------------------------------------------------------------------

// ConfigDirSourceCreator defines a supplier creator used to instantiate
// a dir config supplier.
type ConfigDirSourceCreator struct {
	fileSystem    afero.Fs
	parserFactory *ConfigParserFactory
}

var _ ConfigSupplierCreator = &ConfigDirSourceCreator{}

// NewConfigDirSourceCreator instantiates a new dir config
// supplier creator.
func NewConfigDirSourceCreator(
	fileSystem afero.Fs,
	parserFactory *ConfigParserFactory,
) (*ConfigDirSourceCreator, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiate the strategy
	return &ConfigDirSourceCreator{
		fileSystem:    fileSystem,
		parserFactory: parserFactory,
	}, nil
}

// Accept will check if the requested supplier can be instantiated by this
// creator by parsing the given config partial.
func (s ConfigDirSourceCreator) Accept(
	config *ConfigPartial,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == ConfigTypeDir
}

// Create will instantiate the desired dir supplier instance.
func (s ConfigDirSourceCreator) Create(
	config *ConfigPartial,
) (ConfigSupplier, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Path      string
		Format    string
		Recursive bool
	}{
		Format:    ConfigDefaultFileFormat,
		Recursive: false,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	if sConfig.Path == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing path",
		})
	}
	// create the dir source supplier
	return NewConfigDirSource(
		sConfig.Path,
		sConfig.Format,
		sConfig.Recursive,
		s.fileSystem,
		s.parserFactory,
	)
}

// ----------------------------------------------------------------------------
// config rest source
// ----------------------------------------------------------------------------

type configRestRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

// ConfigRestSource defines a config supplier that read a REST service and
// store a section of the response as the stored config.
type ConfigRestSource struct {
	ConfigSource
	client        configRestRequester
	uri           string
	format        string
	parserFactory *ConfigParserFactory
	configPath    string
}

var _ ConfigSupplier = &ConfigRestSource{}

// NewConfigRestSource will instantiate a new configuration supplier
// that will read a REST endpoint for configuration info.
func NewConfigRestSource(
	client configRestRequester,
	uri,
	format string,
	parserFactory *ConfigParserFactory,
	configPath string,
) (*ConfigRestSource, error) {
	// check client argument reference
	if client == nil {
		return nil, errNilPointer("client")
	}
	// check parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiates the config supplier
	source := &ConfigRestSource{
		ConfigSource:  *NewConfigSource(),
		client:        client,
		uri:           uri,
		format:        format,
		parserFactory: parserFactory,
		configPath:    configPath,
	}
	// load the config information from the REST service
	if e := source.load(); e != nil {
		return nil, e
	}
	return source, nil
}

func (s *ConfigRestSource) load() error {
	// get the REST service information
	config, e := s.request()
	if e != nil {
		return e
	}
	// retrieve the config information from the service response data
	partial, e := config.Partial(s.configPath)
	if e != nil {
		return e
	}
	// store the retrieved config
	s.Mutex.Lock()
	s.Partial = partial
	s.Mutex.Unlock()
	return nil
}

func (s *ConfigRestSource) request() (*ConfigPartial, error) {
	var e error
	// create the REST service config request
	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, s.uri, http.NoBody); e != nil {
		return nil, e
	}
	// call the REST service for the configuration information
	var res *http.Response
	if res, e = s.client.Do(req); e != nil {
		return nil, e
	}
	data, _ := io.ReadAll(res.Body)
	// gat a parser to parse the received service data
	parser, e := s.parserFactory.Create(s.format, bytes.NewReader(data))
	if e != nil {
		return nil, e
	}
	defer func() {
		if closer, ok := parser.(io.Closer); ok {
			_ = closer.Close()
		}
	}()
	// parse the data into a config instance
	return parser.Parse()
}

// ----------------------------------------------------------------------------
// config rest source creator
// ----------------------------------------------------------------------------

// ConfigRestSourceCreator defines a supplier creator used to instantiate
// a REST service config supplier.
type ConfigRestSourceCreator struct {
	clientFactory func() configRestRequester
	parserFactory *ConfigParserFactory
}

var _ ConfigSupplierCreator = &ConfigRestSourceCreator{}

// NewConfigRestSourceCreator instantiates a new REST service config
// supplier creator.
func NewConfigRestSourceCreator(
	parserFactory *ConfigParserFactory,
) (*ConfigRestSourceCreator, error) {
	// check the parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiate the strategy
	return &ConfigRestSourceCreator{
		clientFactory: func() configRestRequester { return &http.Client{} },
		parserFactory: parserFactory,
	}, nil
}

// Accept will check if the requested supplier can be instantiated by this
// creator by parsing the given config partial.
func (s ConfigRestSourceCreator) Accept(
	config *ConfigPartial,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == ConfigTypeRest
}

// Create will instantiate the desired rest supplier instance.
func (s ConfigRestSourceCreator) Create(
	config *ConfigPartial,
) (ConfigSupplier, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		URI    string
		Format string
		Path   struct {
			Config string
		}
	}{
		Format: ConfigDefaultRestFormat,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	if sConfig.URI == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing URI",
		})
	}
	if sConfig.Path.Config == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing response config path",
		})
	}
	// create the requested rest config supplier
	return NewConfigRestSource(
		s.clientFactory(),
		sConfig.URI,
		sConfig.Format,
		s.parserFactory,
		sConfig.Path.Config,
	)
}

// ----------------------------------------------------------------------------
// config observable rest source
// ----------------------------------------------------------------------------

// ConfigObsRestSource defines a config supplier that read a REST
// service and store a section of the response as the stored config.
// Also, the REST service will be periodically checked for updates.
type ConfigObsRestSource struct {
	ConfigRestSource
	timestampPath string
	timestamp     time.Time
}

var _ ConfigObsSupplier = &ConfigObsRestSource{}

// NewConfigObsRestSource will instantiate a new configuration supplier
// that will read a REST endpoint for configuration info, opening the
// possibility for on-the-fly update on supplier content change.
func NewConfigObsRestSource(
	client configRestRequester,
	uri,
	format string,
	parserFactory *ConfigParserFactory,
	timestampPath,
	configPath string,
) (*ConfigObsRestSource, error) {
	// check client argument reference
	if client == nil {
		return nil, errNilPointer("client")
	}
	// check parser factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiates the config supplier
	source := &ConfigObsRestSource{
		ConfigRestSource: ConfigRestSource{
			ConfigSource:  *NewConfigSource(),
			client:        client,
			uri:           uri,
			format:        format,
			parserFactory: parserFactory,
			configPath:    configPath,
		},
		timestampPath: timestampPath,
		timestamp:     time.Unix(0, 0),
	}
	// load the config information from the REST service
	if _, e := source.Reload(); e != nil {
		return nil, e
	}
	return source, nil
}

// Reload will check if the supplier has been updated, and, if so, reload the
// supplier configuration content.
func (s *ConfigObsRestSource) Reload() (bool, error) {
	// get the REST service information
	config, e := s.request()
	if e != nil {
		return false, e
	}
	// search for the response timestamp
	var timestamp time.Time
	if timestamp, e = s.searchTimestamp(config); e != nil {
		return false, e
	}
	// check if the response timestamp is greater than the locally stored
	// config information timestamp
	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(timestamp) {
		// get the response config information
		partial, e := config.Partial(s.configPath)
		if e != nil {
			return false, e
		}
		// store the loaded config information and response timestamp
		s.Mutex.Lock()
		s.Partial = partial
		s.timestamp = timestamp
		s.Mutex.Unlock()
		return true, nil
	}
	return false, nil
}

func (s *ConfigObsRestSource) searchTimestamp(
	config *ConfigPartial,
) (time.Time, error) {
	// retrieve the timestamp information from the parsed response data
	responseTimestamp, e := config.String(s.timestampPath)
	if e != nil {
		return time.Now(), e
	}
	// parse the timestamp string
	var timestamp time.Time
	if timestamp, e = time.Parse(time.RFC3339, responseTimestamp); e != nil {
		return time.Now(), e
	}
	return timestamp, nil
}

// ----------------------------------------------------------------------------
// config observable rest source creator
// ----------------------------------------------------------------------------

// ConfigObsRestSourceCreator defines a supplier creator used to instantiate
// an observable REST service config supplier.
type ConfigObsRestSourceCreator struct {
	ConfigRestSourceCreator
}

var _ ConfigSupplierCreator = &ConfigObsRestSourceCreator{}

// NewConfigObsRestSourceCreator instantiates a new observable REST
// config supplier creator service.
func NewConfigObsRestSourceCreator(
	parserFactory *ConfigParserFactory,
) (*ConfigObsRestSourceCreator, error) {
	// check the decoder factory argument reference
	if parserFactory == nil {
		return nil, errNilPointer("parserFactory")
	}
	// instantiate the strategy
	return &ConfigObsRestSourceCreator{
		ConfigRestSourceCreator: ConfigRestSourceCreator{
			clientFactory: func() configRestRequester { return &http.Client{} },
			parserFactory: parserFactory,
		},
	}, nil
}

// Accept will check if the requested supplier can be instantiated by this
// creator by parsing the given config partial.
func (s ConfigObsRestSourceCreator) Accept(
	config *ConfigPartial,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == ConfigTypeObsRest
}

// Create will instantiate the desired observable rest supplier instance.
func (s ConfigObsRestSourceCreator) Create(
	config *ConfigPartial,
) (ConfigSupplier, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		URI    string
		Format string
		Path   struct {
			Config    string
			Timestamp string
		}
	}{
		Format: ConfigDefaultRestFormat,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// validate configuration
	if sConfig.URI == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing URI",
		})
	}
	if sConfig.Path.Config == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing response config path",
		})
	}
	if sConfig.Path.Timestamp == "" {
		return nil, errInvalidConfigSupplier(*config, map[string]interface{}{
			"description": "missing response config timestamp",
		})
	}
	// create the observable rest config supplier
	return NewConfigObsRestSource(
		s.clientFactory(),
		sConfig.URI,
		sConfig.Format,
		s.parserFactory,
		sConfig.Path.Timestamp,
		sConfig.Path.Config,
	)
}

// ----------------------------------------------------------------------------
// config observer
// ----------------------------------------------------------------------------

// ConfigObserver callback function used to be called when an observed
// configuration path has changed.
type ConfigObserver func(old, new interface{})

// ----------------------------------------------------------------------------
// config
// ----------------------------------------------------------------------------

type configSupplierRef struct {
	id       string
	priority int
	supplier ConfigSupplier
}

type configSupplierRefSorter []configSupplierRef

func (supplier configSupplierRefSorter) Len() int {
	return len(supplier)
}

func (supplier configSupplierRefSorter) Swap(i, j int) {
	supplier[i], supplier[j] = supplier[j], supplier[i]
}

func (supplier configSupplierRefSorter) Less(i, j int) bool {
	return supplier[i].priority < supplier[j].priority
}

type configObserverRef struct {
	path     string
	current  interface{}
	callback ConfigObserver
}

// Config defines an object responsible to handle several config suppliers
// and enable config content observers by path.
type Config struct {
	suppliers []configSupplierRef
	observers []configObserverRef
	partial   *ConfigPartial
	mutex     sync.Locker
	observer  Trigger
}

// NewConfig instantiate a new configuration object.
// This object will handle a series of suppliers, alongside with the ability
// of registration of config path/values observer callbacks that will be
// called whenever the value has changed.
func NewConfig() *Config {
	// instantiate the config
	c := &Config{
		suppliers: []configSupplierRef{},
		observers: []configObserverRef{},
		partial:   &ConfigPartial{},
		mutex:     &sync.Mutex{},
		observer:  nil,
	}
	// check if there is a need to create the observable suppliers
	// trigger
	period := time.Duration(ConfigObserveFrequency) * time.Millisecond
	if period != 0 {
		// create the trigger used to poll the observable suppliers
		c.observer, _ = NewTriggerRecurring(period, func() error {
			return c.reload()
		})
	}
	return c
}

// Close terminates the config instance.
// This will stop the observer trigger and call close on
// all registered suppliers.
func (c *Config) Close() error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// iterate through all the config suppliers checking if we can close then
	for _, ref := range c.suppliers {
		// check if the iterated supplier implements the closer interface
		if source, ok := ref.supplier.(io.Closer); ok {
			// close the supplier
			if e := source.Close(); e != nil {
				return e
			}
		}
	}
	// check if a trigger was generated on creation
	// for observable suppliers polling
	if c.observer != nil {
		// terminate the trigger
		if e := c.observer.Close(); e != nil {
			return e
		}
		c.observer = nil
	}
	return nil
}

// Entries will retrieve the list of stored config entries.
func (c *Config) Entries() []string {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// retrieve the stored entries list
	return c.partial.Entries()
}

// Has will check if a path has been loaded.
// This means that if the values has been loaded by any registered supplier.
func (c *Config) Has(
	path string,
) bool {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// check if the requested path exists in the stored partial
	return c.partial.Has(path)
}

// Get will retrieve a configuration value loaded from a supplier.
func (c *Config) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve the requested value
	return c.partial.Get(path, def...)
}

// Bool will retrieve a bool configuration value loaded from a supplier.
func (c *Config) Bool(
	path string,
	def ...bool,
) (bool, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a boolean value from the local partial
	return c.partial.Bool(path, def...)
}

// Int will retrieve an integer configuration value loaded from a supplier.
func (c *Config) Int(
	path string,
	def ...int,
) (int, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve an integer value from the local partial
	return c.partial.Int(path, def...)
}

// Float will retrieve a floating point configuration value loaded
// from a supplier.
func (c *Config) Float(
	path string,
	def ...float64,
) (float64, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a float value from the local partial
	return c.partial.Float(path, def...)
}

// String will retrieve a string configuration value loaded from a supplier.
func (c *Config) String(
	path string,
	def ...string,
) (string, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a string value from the local partial
	return c.partial.String(path, def...)
}

// List will retrieve a list configuration value loaded from a supplier.
func (c *Config) List(
	path string,
	def ...[]interface{},
) ([]interface{}, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a list value from the local partial
	return c.partial.List(path, def...)
}

// Partial will retrieve partial values loaded from a supplier.
func (c *Config) Partial(
	path string,
	def ...ConfigPartial,
) (ConfigPartial, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a partial value from the local partial
	return c.partial.Partial(path, def...)
}

// Populate will try to populate a given value from a
// stored configuration path.
func (c *Config) Populate(
	path string,
	data interface{},
	insensitive ...bool,
) (interface{}, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to populate a value from the local partial
	return c.partial.Populate(path, data, insensitive...)
}

// HasSupplier check if a supplier with a specific id has been registered.
func (c *Config) HasSupplier(
	id string,
) bool {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find a supplier with the requested id
	for _, ref := range c.suppliers {
		if ref.id == id {
			return true
		}
	}
	return false
}

// AddSupplier register a new supplier with a specific id with a given priority.
func (c *Config) AddSupplier(
	id string,
	priority int,
	supplier ConfigSupplier,
) error {
	// check the supplier argument reference
	if supplier == nil {
		return errNilPointer("supplier")
	}
	// check if there is already a registered supplier with the given id
	if c.HasSupplier(id) {
		return errDuplicateConfigSupplier(id)
	}
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// add the supplier to the config and sort them so that the
	// data can be correctly merged
	c.suppliers = append(c.suppliers, configSupplierRef{id, priority, supplier})
	sort.Sort(configSupplierRefSorter(c.suppliers))
	// rebuild the local partial with the supplier's partial information
	c.rebuild()
	return nil
}

// RemoveSupplier remove a supplier from the registration list
// of the configuration. This will also update the configuration content and
// re-validate the observed paths.
func (c *Config) RemoveSupplier(
	id string,
) error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the requested supplier to be removed
	for i, ref := range c.suppliers {
		if ref.id != id {
			continue
		}
		// check if the supplier implements the closer interface
		if src, ok := ref.supplier.(io.Closer); ok {
			// close the removing supplier
			if e := src.Close(); e != nil {
				return e
			}
		}
		// remove the supplier from the config suppliers
		c.suppliers = append(c.suppliers[:i], c.suppliers[i+1:]...)
		// rebuild the local partial
		c.rebuild()
		return nil
	}
	return nil
}

// RemoveAllSuppliers remove all the registered suppliers from the registration
// list of the configuration. This will also update the configuration
// content and re-validate the observed paths.
func (c *Config) RemoveAllSuppliers() error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// iterate through all the stored suppliers
	for _, ref := range c.suppliers {
		// check if the iterated supplier implements the close interface
		if src, ok := ref.supplier.(io.Closer); ok {
			// close the supplier
			if e := src.Close(); e != nil {
				return e
			}
		}
	}
	// recreate the suppliers array and rebuild the local partial
	c.suppliers = []configSupplierRef{}
	c.rebuild()
	return nil
}

// Supplier retrieve a previously registered supplier with a requested id.
func (c *Config) Supplier(
	id string,
) (ConfigSupplier, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the requested supplier
	for _, ref := range c.suppliers {
		if ref.id == id {
			return ref.supplier, nil
		}
	}
	return nil, errConfigSupplierNotFound(id)
}

// SupplierPriority set a priority value of a previously registered
// supplier with the specified id. This may change the defined values if there
// was an override process of the configuration paths of the changing supplier.
func (c *Config) SupplierPriority(
	id string,
	priority int,
) error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the requested supplier to be updated
	for i, ref := range c.suppliers {
		if ref.id != id {
			continue
		}
		// redefine the stored supplier priority
		c.suppliers[i] = configSupplierRef{
			id:       ref.id,
			priority: priority,
			supplier: ref.supplier,
		}
		// sort the suppliers and rebuild the local partial
		sort.Sort(configSupplierRefSorter(c.suppliers))
		c.rebuild()
		return nil
	}
	return errConfigSupplierNotFound(id)
}

// HasObserver check if there is an observer to a configuration value path.
func (c *Config) HasObserver(
	path string,
) bool {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// check if the requested observer is registered
	for _, oreg := range c.observers {
		if oreg.path == path {
			return true
		}
	}
	return false
}

// AddObserver register a new observer to a configuration path.
func (c *Config) AddObserver(
	path string,
	callback ConfigObserver,
) error {
	// validate the callback argument reference
	if callback == nil {
		return errNilPointer("callback")
	}
	// check if the requested path is present
	val, e := c.Get(path)
	if e != nil {
		return e
	}
	// if the founded value is a partial, clone it, so
	// it can be used for update checks
	if v, ok := val.(ConfigPartial); ok {
		val = v.Clone()
	}
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// register the requested observer with the current path value
	c.observers = append(c.observers, configObserverRef{path, val, callback})
	return nil
}

// RemoveObserver remove an observer to a configuration path.
func (c *Config) RemoveObserver(
	path string,
) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the observer to be removed
	for i, observer := range c.observers {
		if observer.path == path {
			// remove the found observer
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			return
		}
	}
}

func (c *Config) reload() error {
	// iterate through all stores suppliers
	reloaded := false
	for _, ref := range c.suppliers {
		// check if the iterated supplier is an observable supplier
		if supplier, ok := ref.supplier.(ConfigObsSupplier); ok {
			// reload the supplier and update the reloaded flag if the request
			// resulted in a supplier info update
			updated, _ := supplier.Reload()
			reloaded = reloaded || updated
		}
	}
	// check if the iteration resulted in an update of any info
	if reloaded {
		// lock the config for handling
		c.mutex.Lock()
		defer c.mutex.Unlock()
		// rebuild the local partial with the new supplier info
		c.rebuild()
	}
	return nil
}

func (c *Config) rebuild() {
	// iterate through all the stored suppliers
	updated := ConfigPartial{}
	for _, ref := range c.suppliers {
		// retrieve the supplier stored partial information
		// and Merge it with all parsed suppliers
		config, _ := ref.supplier.Get("")
		updated.Merge(config.(ConfigPartial))
	}
	// store locally the resulting partial
	c.partial = &updated
	// iterate through all observers
	for id, observer := range c.observers {
		// retrieve the observer path value
		// and check if the current value differs from the previous one
		val, e := c.partial.Get(observer.path)
		if e == nil && !reflect.DeepEqual(observer.current, val) {
			// store the new value in the observer registry
			// and call the observer callback
			old := observer.current
			c.observers[id].current = val
			observer.callback(old, val)
		}
	}
}

// ----------------------------------------------------------------------------
// config loader
// ----------------------------------------------------------------------------

// ConfigLoader defines an object responsible to initialize a
// configuration manager.
type ConfigLoader struct {
	config          *Config
	supplierFactory *ConfigSupplierFactory
}

// NewConfigLoader instantiate a new configuration loader instance.
func NewConfigLoader(
	config *Config,
	supplierFactory *ConfigSupplierFactory,
) (*ConfigLoader, error) {
	// check config manager argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// check supplier factory argument reference
	if supplierFactory == nil {
		return nil, errNilPointer("supplierFactory")
	}
	// instantiate the loader
	return &ConfigLoader{
		config:          config,
		supplierFactory: supplierFactory,
	}, nil
}

// Load loads the configuration from a well-defined file.
func (l ConfigLoader) Load() error {
	// retrieve the loader entry file partial content
	supplier, e := l.supplierFactory.Create(&ConfigPartial{
		"type":   "file",
		"path":   ConfigLoaderFileSupplierPath,
		"format": ConfigLoaderSupplierFormat,
	})
	if e != nil {
		return e
	}
	// add the loaded entry file content into the manager
	if e := l.config.AddSupplier(ConfigLoaderSupplierID, 0, supplier); e != nil {
		return e
	}
	// retrieve from the loaded info the partial entries list
	suppliers, e := l.config.Partial(ConfigLoaderSupplierListPath)
	if e != nil {
		return nil
	}
	// iterate through the suppliers list
	for _, id := range suppliers.Entries() {
		// retrieve the source list entry
		if partial, e := suppliers.Partial(id); e == nil {
			// load the source
			if e := l.loadSupplier(id, partial); e != nil {
				return e
			}
		}
	}
	return nil
}

func (l ConfigLoader) loadSupplier(
	id string,
	config ConfigPartial,
) error {
	// parse the configuration
	sConfig := struct{ Priority int }{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return e
	}
	// create the requested config supplier
	supplier, e := l.supplierFactory.Create(&config)
	if e != nil {
		return e
	}
	// add the loaded supplier to the config manager
	return l.config.AddSupplier(id, sConfig.Priority, supplier)
}

// ----------------------------------------------------------------------------
// config service register
// ----------------------------------------------------------------------------

// ConfigServiceRegister defines the service provider to be used
// by the application to register the config services.
type ConfigServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &ConfigServiceRegister{}

// NewConfigServiceRegister will generate a new config related services
// registry instance
func NewConfigServiceRegister(
	app ...*App,
) *ConfigServiceRegister {
	return &ConfigServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the configuration module services in the
// application Provider.
func (sr ConfigServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// register the services
	_ = container.Add(ConfigYAMLDecoderCreatorContainerID, NewConfigYAMLDecoderCreator, ConfigParserCreatorTag)
	_ = container.Add(ConfigJSONDecoderCreatorContainerID, NewConfigJSONDecoderCreator, ConfigParserCreatorTag)
	_ = container.Add(ConfigAllParserCreatorsContainerID, sr.getParserCreators(container))
	_ = container.Add(ConfigParserFactoryContainerID, NewConfigParserFactory)
	_ = container.Add(ConfigAllAggregateSuppliersContainerID, sr.getAggregateSuppliers(container))
	_ = container.Add(ConfigAggregateSourceCreatorContainerID, NewConfigAggregateSourceCreator, ConfigSupplierCreatorTag)
	_ = container.Add(ConfigEnvSourceCreatorContainerID, NewConfigEnvSourceCreator, ConfigSupplierCreatorTag)
	_ = container.Add(ConfigFileSourceCreatorContainerID, NewConfigFileSourceCreator, ConfigSupplierCreatorTag)
	_ = container.Add(ConfigObsFileSourceCreatorContainerID, NewConfigObsFileSourceCreator, ConfigSupplierCreatorTag)
	_ = container.Add(ConfigDirSourceCreatorContainerID, NewConfigDirSourceCreator, ConfigSupplierCreatorTag)
	_ = container.Add(ConfigRestSourceCreatorContainerID, NewConfigRestSourceCreator, ConfigSupplierCreatorTag)
	_ = container.Add(ConfigObsRestSourceCreatorContainerID, NewConfigObsRestSourceCreator, ConfigSupplierCreatorTag)
	_ = container.Add(ConfigAllSupplierCreatorsContainerID, sr.getSupplierCreators(container))
	_ = container.Add(ConfigSupplierFactoryContainerID, NewConfigSupplierFactory)
	_ = container.Add(ConfigContainerID, NewConfig)
	_ = container.Add(ConfigLoaderContainerID, NewConfigLoader)
	return nil
}

// Boot will start the config services by calling the
// config loader initialization method.
func (sr ConfigServiceRegister) Boot(
	container *ServiceContainer,
) (e error) {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// check if the config loader is active
	if !ConfigLoaderActive {
		return nil
	}
	// execute the loading action
	loader, e := sr.getLoader(container)
	if e != nil {
		return e
	}
	return loader.Load()
}

func (ConfigServiceRegister) getLoader(
	container *ServiceContainer,
) (*ConfigLoader, error) {
	// retrieve the loader service from the provider
	entry, e := container.Get(ConfigLoaderContainerID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	if instance, ok := entry.(*ConfigLoader); ok {
		return instance, nil
	}
	return nil, errConversion(entry, "*ConfigLoader")
}

func (ConfigServiceRegister) getParserCreators(
	container *ServiceContainer,
) func() []ConfigParserCreator {
	return func() []ConfigParserCreator {
		// retrieve all the parser creators from the provider
		var creators []ConfigParserCreator
		entries, _ := container.Tag(ConfigParserCreatorTag)
		for _, entry := range entries {
			// type check the retrieved service
			creator, ok := entry.(ConfigParserCreator)
			if ok {
				creators = append(creators, creator)
			}
		}
		return creators
	}
}

func (ConfigServiceRegister) getAggregateSuppliers(
	container *ServiceContainer,
) func() []ConfigSupplier {
	return func() []ConfigSupplier {
		// retrieve all the suppliers from the provider
		var creators []ConfigSupplier
		entries, _ := container.Tag(ConfigAggregateSupplierTag)
		for _, entry := range entries {
			// type check the retrieved service
			creator, ok := entry.(ConfigSupplier)
			if ok {
				creators = append(creators, creator)
			}
		}
		return creators
	}
}

func (ConfigServiceRegister) getSupplierCreators(
	container *ServiceContainer,
) func() []ConfigSupplierCreator {
	return func() []ConfigSupplierCreator {
		// retrieve all the supplier creators from the provider
		var creators []ConfigSupplierCreator
		entries, _ := container.Tag(ConfigSupplierCreatorTag)
		for _, entry := range entries {
			// type check the retrieved service
			creator, ok := entry.(ConfigSupplierCreator)
			if ok {
				creators = append(creators, creator)
			}
		}
		return creators
	}
}
