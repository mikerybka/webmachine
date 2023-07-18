package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/mikerybka/paths"
)

func main() {
	s := Server{
		Backend: "localhost:4000",
	}
	http.ListenAndServe(":3000", &s)
}

type Server struct {
	Backend string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := paths.Parse(r.URL.Path)
	if len(parts) == 0 {
		http.NotFound(w, r)
		return
	}
	host := parts[0]
	proxy := httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.Out.URL.Scheme = "http"
			r.Out.URL.Host = s.Backend
			r.Out.URL.Path = paths.Join(parts[1:])
			r.Out.Host = host
		},
	}
	proxy.ServeHTTP(w, r)
}
