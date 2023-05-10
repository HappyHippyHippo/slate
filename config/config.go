package config

import (
	"reflect"
	"strings"
)

// IConfig defined an interface to an instance that holds
// configuration values
type IConfig interface {
	Entries() []string
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
	Bool(path string, def ...bool) (bool, error)
	Int(path string, def ...int) (int, error)
	Float(path string, def ...float64) (float64, error)
	String(path string, def ...string) (string, error)
	List(path string, def ...[]interface{}) ([]interface{}, error)
	Config(path string, def ...Config) (IConfig, error)
	Populate(path string, data any, icase ...bool) (any, error)
}

// Config defines a section of a configuration information
type Config map[interface{}]interface{}

var _ IConfig = &Config{}

// Clone will instantiate an identical instance of the original Config
func (c *Config) Clone() Config {
	// recursive clone function declaration
	var cloner func(value interface{}) interface{}
	cloner = func(value interface{}) interface{} {
		switch v := value.(type) {
		// recursive list scenario
		case []interface{}:
			var result []interface{}
			for _, i := range v {
				result = append(result, cloner(i))
			}
			return result
		// recursive config scenario
		case Config:
			return v.Clone()
		// def scalar value
		default:
			return value
		}
	}
	// create the clone config
	target := make(Config)
	// clone the original config elements to the target config
	for key, value := range *c {
		target[key] = cloner(value)
	}
	return target
}

// Entries will retrieve the list of stored entries in the configuration.
func (c *Config) Entries() []string {
	var entries []string
	for k := range *c {
		key, ok := k.(string)
		if ok {
			entries = append(entries, key)
		}
	}
	return entries
}

// Has will check if a requested path exists in the config.
func (c *Config) Has(
	path string,
) bool {
	_, e := c.path(path)
	return e == nil
}

// Get will retrieve the value stored in the requested path.
// If the path does not exist, then the value nil will be returned. Or, if
// a def value was given as the optional extra argument, then it will
// be returned instead of the standard nil value.
func (c *Config) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// retrieve the path element
	val, e := c.path(path)
	switch {
	// check for non-nil value
	case val != nil:
		return val, nil
	// check if is to return de def value or not
	case e != nil:
		if len(def) > 0 {
			return def[0], nil
		}
		return nil, e
	// def case : return nil
	default:
		return nil, nil
	}
}

// Bool will retrieve a value stored in the quested path cast to bool
func (c *Config) Bool(
	path string,
	def ...bool,
) (bool, error) {
	var val interface{}
	var e error

	// retrieve the config value
	if len(def) > 0 {
		val, e = c.Get(path, def[0])
	} else {
		val, e = c.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return false, e
	}
	// result conversion
	v, ok := val.(bool)
	if !ok {
		return false, errConversion(val, "bool")
	}
	return v, nil
}

// Int will retrieve a value stored in the quested path cast to int
func (c *Config) Int(
	path string,
	def ...int,
) (int, error) {
	var val interface{}
	var e error

	// retrieve the config value
	if len(def) > 0 {
		val, e = c.Get(path, def[0])
	} else {
		val, e = c.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return 0, e
	}
	// result conversion
	v, ok := val.(int)
	if !ok {
		return 0, errConversion(val, "int")
	}
	return v, nil
}

// Float will retrieve a value stored in the quested path cast to float
func (c *Config) Float(
	path string,
	def ...float64,
) (float64, error) {
	var val interface{}
	var e error

	// retrieve the config value
	if len(def) > 0 {
		val, e = c.Get(path, def[0])
	} else {
		val, e = c.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return 0, e
	}
	// result conversion
	v, ok := val.(float64)
	if !ok {
		return 0, errConversion(val, "float64")
	}
	return v, nil
}

// String will retrieve a value stored in the quested path cast to string
func (c *Config) String(
	path string,
	def ...string,
) (string, error) {
	var val interface{}
	var e error

	// retrieve the config value
	if len(def) > 0 {
		val, e = c.Get(path, def[0])
	} else {
		val, e = c.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return "", e
	}
	// result conversion
	v, ok := val.(string)
	if !ok {
		return "", errConversion(val, "string")
	}
	return v, nil
}

// List will retrieve a value stored in the quested path cast to list
func (c *Config) List(
	path string,
	def ...[]interface{},
) ([]interface{}, error) {
	var val interface{}
	var e error

	// retrieve the config value
	if len(def) > 0 {
		val, e = c.Get(path, def[0])
	} else {
		val, e = c.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return nil, e
	}
	// result conversion
	v, ok := val.([]interface{})
	if !ok {
		return nil, errConversion(val, "[]interface{}")
	}
	return v, nil
}

// Config will retrieve a value stored in the quested path cast to config
func (c *Config) Config(
	path string,
	def ...Config,
) (IConfig, error) {
	var val interface{}
	var e error

	// retrieve the config value
	if len(def) > 0 {
		val, e = c.Get(path, def[0])
	} else {
		val, e = c.Get(path)
	}
	// error retrieving the path value
	if e != nil {
		return nil, e
	}
	// result conversion
	v, ok := val.(Config)
	if !ok {
		return nil, errConversion(val, "config.Config")
	}
	return &v, nil
}

// Populate will try to populate the data argument with the data stored
// in the path config location.
func (c *Config) Populate(
	path string,
	data interface{},
	icase ...bool,
) (interface{}, error) {
	// check the case-insensitive flag
	ic := false
	if len(icase) == 0 || icase[0] == true {
		ic = true
		path = strings.ToLower(path)
	}
	// retrieve the config value
	v, e := c.Get(path)
	// error retrieving the path value
	if e != nil {
		return nil, e
	}
	// call recursive data population method
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		return c.populate(v, reflect.ValueOf(data).Elem(), ic)
	}
	return c.populate(v, reflect.New(dataType).Elem(), ic)
}

// Merge will increment the current config instance with the
// information stored in another config.
func (c *Config) Merge(
	src Config,
) {
	// try to Merge every source stored element into the target config
	for key, value := range src {
		// if the key does not exist in the target config, just store it
		if local, ok := (*c)[key]; !ok {
			(*c)[key] = value
		} else {
			// check if the 2 are partials
			typedLocal, okLocal := local.(Config)
			typedValue, okValue := value.(Config)
			if okLocal && okValue {
				// Merge the both partials
				typedLocal.Merge(typedValue)
			} else {
				// just override the target value
				(*c)[key] = value
			}
		}
	}
}

func (c *Config) path(
	path string,
) (interface{}, error) {
	var ok bool
	var it interface{}

	// iterate through the path
	it = *c
	for _, part := range strings.Split(path, PathSeparator) {
		// if the iterated path is empty
		// (double occurrence of a separator), just continue
		if part == "" {
			continue
		}
		switch i := it.(type) {
		// check if the iterated part references a config
		case Config:
			// retrieve the part reference
			it, ok = i[part]
			if !ok {
				return nil, errPathNotFound(path)
			}
		default:
			return nil, errPathNotFound(path)
		}
	}
	return it, nil
}

func (c *Config) populate(
	source interface{},
	target reflect.Value,
	icase bool,
) (interface{}, error) {
	// get the types of the source and target
	sourceType := reflect.TypeOf(source)
	targetType := target.Type()
	// if the types are the same, just return the source
	if sourceType == targetType {
		return source, nil
	}
	// source type action
	switch sourceType {
	case reflect.TypeOf(&Config{}):
		return c.populate(*source.(*Config), target, icase)
	case reflect.TypeOf(Config{}):
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
					if icase {
						path = strings.ToLower(path)
					}
					cfg := source.(Config)
					data, e := cfg.Config(path)
					if e != nil {
						continue
					}
					// recursive assignment
					_, e = c.populate(data, fieldValue, icase)
					if e != nil {
						return nil, e
					}
				default:
					// get the configuration value
					path := fieldType.Name
					if icase {
						path = strings.ToLower(path)
					}
					cfg := source.(Config)
					data, e := cfg.Get(path)
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
