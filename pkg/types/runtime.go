package types

type Runtime struct {
	FileName  string
	CmdPrefix []string
}

func (r *Runtime) Handle(req *Request) *Response {
	return nil
}
