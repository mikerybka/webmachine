package http2cli

import "net/http"

type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]string
}

func (r *Response) WriteTo(w http.ResponseWriter) {
	for k, v := range r.Headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(r.StatusCode)
	w.Write(r.Body)
}
