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

func (e *Endpoint) ContentType() string {
	fi, err := os.Stat(e.Filepath)
	if err != nil {
		return ""
	}
	if fi.IsDir() {
		fi, err = os.Stat(filepath.Join(e.Filepath, "index.html"))
		if err != nil {
			return ""
		}
		if !fi.IsDir() {
			return "text/html"
		}
		return ""
	}
	switch filepath.Ext(e.Filepath) {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".pdf":
		return "application/pdf"
	case ".zip":
		return "application/zip"
	case ".tar":
		return "application/x-tar"
	case ".gz":
		return "application/gzip"
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".mp4":
		return "video/mp4"
	case ".mov":
		return "video/quicktime"
	case ".avi":
		return "video/x-msvideo"
	case ".png":
		return "image/png"
	case ".jpg":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	case ".ico":
		return "image/x-icon"
	case ".ttf":
		return "font/ttf"
	case ".otf":
		return "font/otf"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".eot":
		return "application/vnd.ms-fontobject"
	case ".csv":
		return "text/csv"
	case ".txt":
		return "text/plain"
	default:
		return "text/html"
	}
}

func (e *Endpoint) File() io.Reader {
	fi, err := os.Stat(e.Filepath)
	if err != nil {
		return bytes.NewReader(nil)
	}
	if fi.IsDir() {
		path := filepath.Join(e.Filepath, "index.html")
		fi, _ = os.Stat(path)
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
		w.Header().Set("Content-Type", e.ContentType())
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
