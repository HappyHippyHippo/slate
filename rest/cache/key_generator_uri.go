package cache

import (
	"bytes"
	"crypto/sha1"
	"io"
	"net/http"
	"net/url"
)

func keyGeneratorURI(req *http.Request) string {
	uri := req.URL.RequestURI()
	key := url.QueryEscape(uri)

	if len(key) > 200 {
		h := sha1.New()
		_, _ = io.WriteString(h, uri)
		key = string(h.Sum(nil))
	}

	var buffer bytes.Buffer
	buffer.WriteString(key)
	return buffer.String()
}
