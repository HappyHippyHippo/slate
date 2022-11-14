package config

// IConfig defined an interface to an instance that holds
// configuration values
type IConfig interface {
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
	Bool(path string, def ...bool) (bool, error)
	Int(path string, def ...int) (int, error)
	Float(path string, def ...float64) (float64, error)
	String(path string, def ...string) (string, error)
	List(path string, def ...[]interface{}) ([]interface{}, error)
	Partial(path string, def ...Partial) (Partial, error)
}
