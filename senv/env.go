package senv

import (
	"os"
	"strconv"
	"strings"
)

// Bool will try to retrieve a boolean value from the environment variable.
func Bool(name string, def bool) bool {
	result := def
	if value, ok := os.LookupEnv(name); ok && value != "" {
		result, _ = strconv.ParseBool(value)
	}
	return result
}

// Int will try to retrieve an integer value from the environment variable.
func Int(name string, def int) int {
	result := def
	if value, ok := os.LookupEnv(name); ok && value != "" {
		var e error
		result, e = strconv.Atoi(value)
		if e != nil {
			return def
		}
	}
	return result
}

// String will try to retrieve a string value from the environment variable.
func String(name, def string) string {
	result := def
	if value, ok := os.LookupEnv(name); ok && value != "" {
		result = value
	}
	return result
}

// List will try to retrieve a list of strings values from the
// environment variable.
func List(name string, def []string) []string {
	result := def
	if value, ok := os.LookupEnv(name); ok && value != "" {
		result = strings.Split(value, ",")
	}
	return result
}
