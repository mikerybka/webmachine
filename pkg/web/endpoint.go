package web

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/webmachine/pkg/golang"
	"github.com/mikerybka/webmachine/pkg/ruby"
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
	if r.Method == http.MethodGet && e.IsFile() {
		io.Copy(w, e.File())
		return
	}
	status, body, err := e.run(r.Method, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(body)
}

func (e *Endpoint) run(method string, input io.Reader) (status int, body []byte, err error) {
	_, err = os.Stat(e.CodePath(method, "go"))
	if err == nil {
		return golang.Run(e.CodePath(method, "go"), input)
	}
	_, err = os.Stat(e.CodePath(method, "rb"))
	if err == nil {
		return ruby.Run(e.CodePath(method, "rb"), input)
	}

	return http.StatusMethodNotAllowed, []byte("method not allowed"), nil
}
