package config

import (
	"os"
	"strings"
	"sync"
)

// EnvSource defines a config source that maps environment
// variables values to a config.
type EnvSource struct {
	Source
	mappings map[string]string
}

var _ ISource = &EnvSource{}

// NewEnvSource will instantiate a new configuration source
// that will map environmental variables to configuration
// path values.
func NewEnvSource(
	mappings map[string]string,
) (*EnvSource, error) {
	// instantiate the source
	s := &EnvSource{
		Source: Source{
			mutex:   &sync.Mutex{},
			partial: Config{},
		},
		mappings: mappings,
	}
	// load the source values from environment
	_ = s.load()
	return s, nil
}

func (s *EnvSource) load() error {
	// iterate through all the source mappings
	for key, path := range s.mappings {
		// retrieve the mapped value from the environment
		if env := os.Getenv(key); env != "" {
			// navigate to the target storing path of the environment value
			step := s.partial
			sections := strings.Split(path, ".")
			for i, section := range sections {
				if i != len(sections)-1 {
					// convert the path section if is present and not a config
					if _, ok := step[section]; ok == false {
						step[section] = Config{}
					}
					switch step[section].(type) {
					case IConfig:
					default:
						// create the section if not present
						step[section] = Config{}
					}
					// iterate to the section
					step = step[section].(Config)
				} else {
					// store the value in the target section
					step[section] = env
				}
			}
		}
	}
	return nil
}
