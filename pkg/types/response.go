package types

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	Request *Request          `json:"request"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
}

func NewResponse(req *Request) *Response {
	return &Response{
		Request: req,
		Status:  200,
		Headers: map[string]string{},
		Body:    []byte{},
	}
}

func (r *Response) Log(w io.Writer) {
	b, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	b = append(b, '\n')
	w.Write(b)
}

func (r *Response) Write(w http.ResponseWriter) {
	if r == nil {
		return
	}

	for k, v := range r.Headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(r.Status)
	w.Write(r.Body)
}
