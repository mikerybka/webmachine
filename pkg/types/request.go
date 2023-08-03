package types

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Timestamp   int64             `json:"timestamp"`
	Method      string            `json:"method"`
	Host        string            `json:"host"`
	Path        string            `json:"path"`
	QueryParams map[string]string `json:"queryParams"`
	Body        []byte            `json:"body"`
	Headers     map[string]string `json:"headers"`
}

func NewRequest(r *http.Request) *Request {
	now := time.Now()

	req := Request{
		Timestamp:   now.UnixNano(),
		Method:      r.Method,
		Host:        r.Host,
		Path:        r.URL.Path,
		QueryParams: map[string]string{},
		Body:        []byte{},
		Headers:     map[string]string{},
	}

	// Copy headers
	for k, v := range r.Header {
		req.Headers[k] = v[0]
	}

	// Copy query params
	for k := range r.URL.Query() {
		req.QueryParams[k] = r.URL.Query().Get(k)
	}

	// Copy body
	b, _ := io.ReadAll(r.Body)
	req.Body = b

	return &req
}

func (r *Request) Log(w io.Writer) {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	b = append(b, '\n')
	w.Write(b)
}
