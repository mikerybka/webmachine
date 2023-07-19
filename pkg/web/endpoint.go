package web

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/webmachine/pkg/ffi"
)

type Endpoint struct {
	Filepath string
}

func (e *Endpoint) CodePath(method, lang string) string {
	return filepath.Join(e.Filepath, method, "main."+lang)
}

func (e *Endpoint) IsFile() bool {
	fi, err := os.Stat(e.Filepath)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

func (e *Endpoint) File() io.Reader {
	f, err := os.Open(e.Filepath)
	if err != nil {
		panic(err)
	}
	return f
}

func (e *Endpoint) GoGET() string {
	path := filepath.Join(e.Filepath, "GET/main.go")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) GoPOST() string {
	path := filepath.Join(e.Filepath, "POST/main.go")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) GoPUT() string {
	path := filepath.Join(e.Filepath, "PUT/main.go")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) GoDELETE() string {
	path := filepath.Join(e.Filepath, "DELETE/main.go")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) RubyGET() string {
	path := filepath.Join(e.Filepath, "GET/main.rb")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) RubyPOST() string {
	path := filepath.Join(e.Filepath, "POST/main.rb")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) RubyPUT() string {
	path := filepath.Join(e.Filepath, "PUT/main.rb")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) RubyDELETE() string {
	path := filepath.Join(e.Filepath, "DELETE/main.rb")
	b, _ := os.ReadFile(path)
	return string(b)
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e == nil {
		http.NotFound(w, r)
		return
	}

	// Support raw files
	if r.Method == http.MethodGet && e.IsFile() {
		io.Copy(w, e.File())
		return
	}

	// Run request
	status, body, err := e.run(r)
	if err != nil {
		http.Error(w, err.Error()+"\n\n"+string(body), http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(status)
	w.Write(body)
}

func (e *Endpoint) run(r *http.Request) (status int, body []byte, err error) {
	method := r.Method
	input := r.Body

	codePath := e.CodePath(method, "go")
	_, err = os.Stat(codePath)
	if err == nil {
		return ffi.Run(input, "go", "run", codePath)
	}

	codePath = e.CodePath(method, "rb")
	_, err = os.Stat(codePath)
	if err == nil {
		return ffi.Run(input, "ruby", codePath)
	}

	codePath = e.CodePath(method, "py")
	_, err = os.Stat(codePath)
	if err == nil {
		return ffi.Run(input, "python3", codePath)
	}

	codePath = e.CodePath(method, "js")
	_, err = os.Stat(codePath)
	if err == nil {
		return ffi.Run(input, "node", codePath)
	}

	return http.StatusMethodNotAllowed, []byte("method not allowed"), nil
}
