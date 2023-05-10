package config

// ISource defines the base interface of a config source.
type ISource interface {
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
}

// IObsSource interface extends the ISource interface with methods
// specific to sources that will be checked for updates in a regular
// periodicity defined in the config object where the source will be
// registered.
type IObsSource interface {
	ISource
	Reload() (bool, error)
}
