package env

import (
	"os"
	"strings"
	"sync"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
)

// Source defines a config source that maps environment
// variables values to a config.
type Source struct {
	source.Source
	mappings map[string]string
}

var _ config.Source = &Source{}

// NewSource will instantiate a new configuration source
// that will map environmental variables to configuration
// path values.
func NewSource(
	mappings map[string]string,
) (*Source, error) {
	// instantiate the source
	s := &Source{
		Source: source.Source{
			Mutex:   &sync.Mutex{},
			Partial: config.Partial{},
		},
		mappings: mappings,
	}
	// load the source values from environment
	_ = s.load()
	return s, nil
}

func (s *Source) load() error {
	// iterate through all the source mappings
	for key, path := range s.mappings {
		// retrieve the mapped value from the environment
		if env := os.Getenv(key); env != "" {
			// navigate to the target storing path of the environment value
			step := s.Partial
			sections := strings.Split(path, ".")
			for i, section := range sections {
				if i != len(sections)-1 {
					// Convert the path section if is present and not a config
					if _, ok := step[section]; ok == false {
						step[section] = config.Partial{}
					}
					// create the section if not present
					// and iterate to the section
					step[section] = config.Partial{}
					step = step[section].(config.Partial)
				} else {
					// store the value in the target section
					step[section] = env
				}
			}
		}
	}
	return nil
}
