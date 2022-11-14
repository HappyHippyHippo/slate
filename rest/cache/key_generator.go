package cache

import "net/http"

// KeyGenerator @todo doc
type KeyGenerator func(req *http.Request) string
