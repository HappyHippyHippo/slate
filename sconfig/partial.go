package sconfig

import "strings"

// Partial defines a section of a configuration information
type Partial map[interface{}]interface{}

var _ IConfig = &Partial{}

// Clone will instantiate an identical instance of the original Partial
func (p Partial) Clone() Partial {
	var cloner func(value interface{}) interface{}
	cloner = func(value interface{}) interface{} {
		switch v := value.(type) {
		case []interface{}:
			result := []interface{}{}
			for _, i := range v {
				result = append(result, cloner(i))
			}
			return result
		case Partial:
			return v.Clone()
		default:
			return value
		}
	}

	target := make(Partial)
	for key, value := range p {
		target[key] = cloner(value)
	}
	return target
}

// Has will check if a requested path exists in the config Partial.
func (p *Partial) Has(path string) bool {
	_, e := p.path(path)
	return e == nil
}

// Get will retrieve the value stored in the requested path.
// If the path does not exist, then the value nil will be returned. Or, if
// a default value was given as the optional extra argument, then it will
// be returned instead of the standard nil value.
func (p *Partial) Get(path string, def ...interface{}) (interface{}, error) {
	val, e := p.path(path)

	switch {
	case val != nil:
		return val, nil
	case e != nil:
		if len(def) > 0 {
			return def[0], nil
		}
		return nil, e
	default:
		return nil, nil
	}
}

// Bool will retrieve a value stored in the quested path cast to bool
func (p *Partial) Bool(path string, def ...bool) (bool, error) {
	var val interface{}
	var e error

	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	if e != nil {
		return false, e
	}

	v, ok := val.(bool)
	if !ok {
		return false, errConversion(val, "bool")
	}

	return v, nil
}

// Int will retrieve a value stored in the quested path cast to int
func (p *Partial) Int(path string, def ...int) (int, error) {
	var val interface{}
	var e error

	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	if e != nil {
		return 0, e
	}

	v, ok := val.(int)
	if !ok {
		return 0, errConversion(val, "int")
	}

	return v, nil
}

// Float will retrieve a value stored in the quested path cast to float
func (p *Partial) Float(path string, def ...float64) (float64, error) {
	var val interface{}
	var e error

	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	if e != nil {
		return 0, e
	}

	v, ok := val.(float64)
	if !ok {
		return 0, errConversion(val, "float64")
	}

	return v, nil
}

// String will retrieve a value stored in the quested path cast to string
func (p *Partial) String(path string, def ...string) (string, error) {
	var val interface{}
	var e error

	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	if e != nil {
		return "", e
	}

	v, ok := val.(string)
	if !ok {
		return "", errConversion(val, "string")
	}

	return v, nil
}

// List will retrieve a value stored in the quested path cast to list
func (p *Partial) List(path string, def ...[]interface{}) ([]interface{}, error) {
	var val interface{}
	var e error

	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	if e != nil {
		return nil, e
	}

	v, ok := val.([]interface{})
	if !ok {
		return nil, errConversion(val, "[]interface{}")
	}

	return v, nil
}

// Partial will retrieve a value stored in the quested path cast to config
func (p *Partial) Partial(path string, def ...Partial) (Partial, error) {
	var val interface{}
	var e error

	if len(def) > 0 {
		val, e = p.Get(path, def[0])
	} else {
		val, e = p.Get(path)
	}
	if e != nil {
		return nil, e
	}

	v, ok := val.(Partial)
	if !ok {
		return nil, errConversion(val, "config.Partial")
	}

	return v, nil
}

func (p *Partial) merge(src Partial) {
	for key, value := range src {
		if local, ok := (*p)[key]; !ok {
			(*p)[key] = value
		} else {
			typedLocal, okLocal := local.(Partial)
			typedValue, okValue := value.(Partial)
			if okLocal && okValue {
				typedLocal.merge(typedValue)
			} else {
				(*p)[key] = value
			}
		}
	}
}

func (p *Partial) path(path string) (interface{}, error) {
	var ok bool
	var it interface{}

	it = *p
	for _, part := range strings.Split(path, PathSeparator) {
		if part == "" {
			continue
		}

		switch i := it.(type) {
		case Partial:
			it, ok = i[part]
			if !ok {
				return nil, errConfigPathNotFound(path)
			}
		default:
			return nil, errConfigPathNotFound(path)
		}
	}
	return it, nil
}

func (Partial) convert(val interface{}) interface{} {
	if jm, ok := val.(map[string]interface{}); ok {
		p := Partial{}
		for k, v := range jm {
			p[k] = p.convert(v)
		}
		return p
	}
	if f, ok := val.(float64); ok {
		if float64(int(f)) == f {
			return int(f)
		}
	}
	if f, ok := val.(float32); ok {
		if float32(int(f)) == f {
			return int(f)
		}
	}
	return val
}
