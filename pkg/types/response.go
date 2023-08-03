package types

type Response struct {
	Request *Request          `json:"request"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
}
