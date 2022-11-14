package cache

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

type cachedWriter struct {
	gin.ResponseWriter
	status  int
	written bool
	store   IStore
	expire  time.Duration
	key     string
}

var _ gin.ResponseWriter = &cachedWriter{}

func newCachedWriter(store IStore, expire time.Duration, writer gin.ResponseWriter, key string) *cachedWriter {
	return &cachedWriter{writer, 0, false, store, expire, key}
}

// WriteHeader @todo doc
func (w *cachedWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

// Status @todo doc
func (w *cachedWriter) Status() int {
	return w.ResponseWriter.Status()
}

// Written @todo doc
func (w *cachedWriter) Written() bool {
	return w.ResponseWriter.Written()
}

// Write @todo doc
func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil {
		store := w.store
		var cache responseCache
		if err := store.Get(w.key, &cache); err == nil {
			data = append(cache.Data, data...)
		}

		// cache responses with a status code < 300
		if w.Status() < 300 {
			val := responseCache{
				w.Status(),
				w.Header(),
				data,
			}
			err = store.Set(w.key, val, w.expire)
			if err != nil {
				// need logger
			}
		}
	}
	return ret, err
}

// WriteString @todo doc
func (w *cachedWriter) WriteString(data string) (n int, err error) {
	ret, err := w.ResponseWriter.WriteString(data)
	// cache responses with a status code < 300
	if err == nil && w.Status() < 300 {
		store := w.store
		val := responseCache{
			w.Status(),
			w.Header(),
			[]byte(data),
		}
		_ = store.Set(w.key, val, w.expire)
	}
	return ret, err
}
