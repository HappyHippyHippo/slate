package sconfig

import (
	"os"
	"strings"
	"sync"
)

type sourceEnv struct {
	source
	mappings map[string]string
}

var _ ISource = &sourceEnv{}

// NewSourceEnv will instantiate a new configuration source
// that will map environmental variables to configuration
// path values.
func NewSourceEnv(mappings map[string]string) (ISource, error) {
	s := &sourceEnv{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		mappings: mappings,
	}

	_ = s.load()

	return s, nil
}

func (s *sourceEnv) load() error {
	for key, path := range s.mappings {
		if env := os.Getenv(key); env != "" {
			step := s.partial
			sections := strings.Split(path, ".")
			for i, section := range sections {
				if i != len(sections)-1 {
					if _, ok := step[section]; ok == false {
						step[section] = Partial{}
					}

					switch step[section].(type) {
					case IConfig:
					default:
						step[section] = Partial{}
					}

					step = step[section].(Partial)
				} else {
					step[section] = env
				}
			}
		}
	}

	return nil
}
