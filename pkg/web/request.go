package web

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Request(path, method string) *http.Request {
	// New URL
	host, path, ok := strings.Cut(path, "/")
	if !ok {
		host = path
		path = ""
	}
	url := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}

	// Set query string
	query := url.Query()
	for _, arg := range os.Args[1:] {
		_, arg, ok := strings.Cut(arg, "--")
		if ok {
			key, value, ok := strings.Cut(arg, "=")
			if ok {
				query.Set(key, value)
			}
		}
	}

	// New request
	req, err := http.NewRequest(method, url.String(), os.Stdin)
	if err != nil {
		panic(err)
	}

	// Set headers
	for _, env := range os.Environ() {
		key, value, ok := strings.Cut(env, "=")
		if ok {
			req.Header.Set(key, value)
		}
	}

	return req
}
