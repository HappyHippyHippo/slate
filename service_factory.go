package slate

// IServiceFactory is a callback function used to instantiate an object used by
// the application container when a not yet instantiated object is requested.
type IServiceFactory func() (any, error)
