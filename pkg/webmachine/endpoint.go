package webmachine

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/webmachine/pkg/http2cli"
)

type Endpoint struct {
	Filepath string
	Args     map[string]string
}

func (e *Endpoint) CodePath(method, lang string) string {
	return filepath.Join(e.Filepath, method, "main."+lang)
}

func (e *Endpoint) IsFile() bool {
	fi, err := os.Stat(e.Filepath)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		fi, err = os.Stat(filepath.Join(e.Filepath, "index.html"))
		if err != nil {
			return false
		}
		if !fi.IsDir() {
			return true
		}
		return false
	}
	return true
}

func (e *Endpoint) File() io.Reader {
	fi, err := os.Stat(e.Filepath)
	if err != nil {
		return bytes.NewReader(nil)
	}
	if fi.IsDir() {
		path := filepath.Join(e.Filepath, "index.html")
		fi, err = os.Stat(path)
		if fi.IsDir() {
			return bytes.NewReader(nil)
		}
		f, err := os.Open(path)
		if err != nil {
			return bytes.NewReader(nil)
		}
		return f
	}
	f, err := os.Open(e.Filepath)
	if err != nil {
		return bytes.NewReader(nil)
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
	e.Run(r).WriteTo(w)
}

func (e *Endpoint) Run(r *http.Request) *http2cli.Response {
	cmdMap := map[string][]string{
		"go":  {"go", "run"},
		"rb":  {"ruby"},
		"py":  {"python3"},
		"js":  {"node"},
		"ts":  {"bun"},
		"tsx": {"bun"},
		"jsx": {"bun"},
	}
	exts := []string{
		"go",
		"rb",
		"py",
		"js",
		"ts",
		"tsx",
		"jsx",
	}
	for _, ext := range exts {
		codePath := e.CodePath(r.Method, ext)
		_, err := os.Stat(codePath)
		if err == nil {
			// Prepare command
			c := append(cmdMap[ext], codePath)

			// Run command
			return http2cli.Exec(r, c, e.Args)
		}
	}

	return &http2cli.Response{
		StatusCode: http.StatusNotFound,
		Body:       []byte("not found"),
	}
}
