package config

import (
	"strings"
)

// Convert will try to convert a source argument value
// into a config accepted value. This means that if the value
// is a map, then it will be converted into a Config instance
// (recursively)
func Convert(
	val interface{},
) interface{} {
	// recursive conversion call
	if v, ok := val.(Config); ok {
		// return the recursive conversion of the config
		p := Config{}
		for k, v := range v {
			// turn all string keys into lowercase
			sk, ok := k.(string)
			if ok {
				p[strings.ToLower(sk)] = Convert(v)
			} else {
				p[k] = Convert(v)
			}
		}
		return p
	}
	if v, ok := val.([]interface{}); ok {
		var p []interface{}
		for _, i := range v {
			p = append(p, Convert(i))
		}
		return p
	}
	// check if the value is a map that can be converted to a config
	if v, ok := val.(map[string]interface{}); ok {
		// return the recursive conversion of the config
		p := Config{}
		for k, i := range v {
			// turn all string keys into lowercase
			p[strings.ToLower(k)] = Convert(i)
		}
		return p
	}
	// check if the value is a map that can be converted to a config
	if v, ok := val.(map[interface{}]interface{}); ok {
		// return the recursive conversion of the config
		p := Config{}
		for k, i := range v {
			// turn all string keys into lowercase
			sk, ok := k.(string)
			if ok {
				p[strings.ToLower(sk)] = Convert(i)
			} else {
				p[k] = Convert(i)
			}
		}
		return p
	}
	// check if the value is a float64 but with an integer value
	if v, ok := val.(float64); ok {
		if float64(int(v)) == v {
			return int(v)
		}
	}
	// check if the value is a float32 but with an integer value
	if v, ok := val.(float32); ok {
		if float32(int(v)) == v {
			return int(v)
		}
	}
	return val
}
