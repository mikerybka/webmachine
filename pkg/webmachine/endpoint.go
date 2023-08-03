package webmachine

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/web/util"
	"github.com/mikerybka/webmachine/pkg/data"
	"github.com/mikerybka/webmachine/pkg/types"
)

type Endpoint struct {
	Filepath string
	Args     map[string]string
}

func (e *Endpoint) Handle(req *types.Request) *types.Response {
	if req.Method == http.MethodGet && e.isFile() {
		return e.fileResponse(req)
	}
	return e.execResponse(req)
}

func (e *Endpoint) fileResponse(r *types.Request) *types.Response {
	res := types.NewResponse(r)
	res.Headers["Content-Type"] = e.contentType()
	b, err := io.ReadAll(e.file())
	if err != nil {
		panic(err)
	}
	res.Body = b
	return res
}

func (e *Endpoint) isFile() bool {
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

func (e *Endpoint) contentType() string {
	fi, err := os.Stat(e.Filepath)
	if err != nil {
		return ""
	}

	if fi.IsDir() {
		fi, err = os.Stat(filepath.Join(e.Filepath, "index.html"))
		if err != nil {
			return ""
		}
		if fi.IsDir() {
			return ""
		}
		return "text/html"
	}

	contentType, ok := data.ContentTypes[filepath.Ext(e.Filepath)]
	if !ok {
		return "text/plain"
	}

	return contentType
}

func (e *Endpoint) file() io.Reader {
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

func (e *Endpoint) execResponse(r *types.Request) *types.Response {
	return e.runtime(r.Method).Handle(r)
}

func (e *Endpoint) runtime(method string) *types.Runtime {
	for _, runtime := range data.Runtimes {
		handlerFile := filepath.Join(e.Filepath, method, runtime.FileName)
		if util.FileExists(handlerFile) {
			return &runtime
		}
	}
	return nil
}
