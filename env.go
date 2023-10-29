package slate

import (
	"os"
	"strconv"
	"strings"
)

// EnvBool will try to retrieve a boolean value from the
// environment variable.
func EnvBool(
	name string,
	def ...bool,
) bool {
	// set default value
	result := false
	if len(def) > 0 {
		result = def[0]
	}
	// read value from env
	if value, ok := os.LookupEnv(name); ok && value != "" {
		parsed, e := strconv.ParseBool(value)
		if e != nil {
			return result
		}
		result = parsed
	}
	return result
}

// EnvInt will try to retrieve an integer value from the
// environment variable.
func EnvInt(
	name string,
	def ...int,
) int {
	// set default value
	result := 0
	if len(def) > 0 {
		result = def[0]
	}
	// read value from env
	if value, ok := os.LookupEnv(name); ok && value != "" {
		parsed, e := strconv.Atoi(value)
		if e != nil {
			return result
		}
		result = parsed
	}
	return result
}

// EnvString will try to retrieve a string value from the
// environment variable.
func EnvString(
	name string,
	def ...string,
) string {
	// set default value
	result := ""
	if len(def) > 0 {
		result = def[0]
	}
	// read value from env
	if value, ok := os.LookupEnv(name); ok && value != "" {
		result = value
	}
	return result
}

// EnvList will try to retrieve a list of strings values from the
// environment variable.
func EnvList(
	name string,
	def ...[]string,
) []string {
	// set default value
	var result []string
	if len(def) > 0 {
		result = def[0]
	}
	// read value from env
	if value, ok := os.LookupEnv(name); ok && value != "" {
		result = strings.Split(value, ",")
	}
	return result
}
