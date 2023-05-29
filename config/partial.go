package config

import (
	"reflect"
	"strings"
)

// Partial defines a section of a configuration information
type Partial map[interface{}]interface{}

// Clone will instantiate an identical instance of the original Partial
func (p *Partial) Clone() Partial {
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
		// recursive partial scenario
		case Partial:
			return v.Clone()
		// simple scalar value
		default:
			return value
		}
	}
	// create the clone partial
	target := make(Partial)
	// clone the original partial elements to the target partial
	for key, value := range *p {
		target[key] = cloner(value)
	}
	return target
}

// Entries will retrieve the list of stored entries in the configuration.
func (p *Partial) Entries() []string {
	var entries []string
	for k := range *p {
		if key, ok := k.(string); ok {
			entries = append(entries, key)
		}
	}
	return entries
}

// Has will check if a requested path exists in the partial.
func (p *Partial) Has(
	path string,
) bool {
	_, e := p.path(path)
	return e == nil
}

// Get will retrieve the value stored in the requested path.
// If the path does not exist, then the value nil will be returned. Or, if
// a simple value was given as the optional extra argument, then it will
// be returned instead of the standard nil value.
func (p *Partial) Get(
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

// Bool will retrieve a value stored in the quested path cast to bool
func (p *Partial) Bool(
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
	if v, ok := val.(bool); ok {
		return v, nil
	}
	return false, errConversion(val, "bool")
}

// Int will retrieve a value stored in the quested path cast to int
func (p *Partial) Int(
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
	if v, ok := val.(int); ok {
		return v, nil
	}
	return 0, errConversion(val, "int")
}

// Float will retrieve a value stored in the quested path cast to float
func (p *Partial) Float(
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
	if v, ok := val.(float64); ok {
		return v, nil
	}
	return 0, errConversion(val, "float64")
}

// String will retrieve a value stored in the quested path cast to string
func (p *Partial) String(
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
	if v, ok := val.(string); ok {
		return v, nil
	}
	return "", errConversion(val, "string")
}

// List will retrieve a value stored in the quested path cast to list
func (p *Partial) List(
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
	if v, ok := val.([]interface{}); ok {
		return v, nil
	}
	return nil, errConversion(val, "[]interface{}")
}

// Partial will retrieve a value stored in the quested path cast to partial
func (p *Partial) Partial(
	path string,
	def ...Partial,
) (*Partial, error) {
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
	if v, ok := val.(Partial); ok {
		return &v, nil
	}
	return nil, errConversion(val, "config.Partial")
}

// Populate will try to populate the data argument with the data stored
// in the path partial location.
func (p *Partial) Populate(
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
	// retrieve the partial value
	v, e := p.Get(path)
	// error retrieving the path value
	if e != nil {
		return nil, e
	}
	// call recursive data population method
	dataType := reflect.TypeOf(data)
	if dataType.Kind() == reflect.Ptr {
		return p.populate(v, reflect.ValueOf(data).Elem(), ic)
	}
	return p.populate(v, reflect.New(dataType).Elem(), ic)
}

// Merge will increment the current partial instance with the
// information stored in another partial.
func (p *Partial) Merge(
	src Partial,
) {
	// try to Merge every source stored element into the target partial
	for key, value := range src {
		// if the key does not exist in the target partial, just store it
		if local, ok := (*p)[key]; !ok {
			(*p)[key] = value
		} else {
			// check if the 2 are partials
			typedLocal, okLocal := local.(Partial)
			typedValue, okValue := value.(Partial)
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

func (p *Partial) path(
	path string,
) (interface{}, error) {
	var ok bool
	var it interface{}

	// iterate through the path
	it = *p
	for _, part := range strings.Split(path, PathSeparator) {
		// if the iterated path is empty
		// (double occurrence of a separator), just continue
		if part == "" {
			continue
		}
		switch i := it.(type) {
		// check if the iterated part references a partial
		case Partial:
			// retrieve the part reference
			if it, ok = i[part]; !ok {
				return nil, errPathNotFound(path)
			}
		default:
			return nil, errPathNotFound(path)
		}
	}
	return it, nil
}

func (p *Partial) populate(
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
	case reflect.TypeOf(&Partial{}):
		return p.populate(*source.(*Partial), target, icase)
	case reflect.TypeOf(Partial{}):
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
					cfg := source.(Partial)
					data, e := cfg.Partial(path)
					if e != nil {
						continue
					}
					// recursive assignment
					if _, e := p.populate(data, fieldValue, icase); e != nil {
						return nil, e
					}
				default:
					// get the configuration value
					path := fieldType.Name
					if icase {
						path = strings.ToLower(path)
					}
					cfg := source.(Partial)
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
