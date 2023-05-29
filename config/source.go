package config

// Source defines the base interface of a partial source.
type Source interface {
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
}

// ObsSource interface extends the Source interface with methods
// specific to sources that will be checked for updates in a regular
// periodicity defined in the config object where the source will be
// registered.
type ObsSource interface {
	Source
	Reload() (bool, error)
}
