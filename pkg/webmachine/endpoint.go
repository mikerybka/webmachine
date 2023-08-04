package webmachine

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mikerybka/web/util"
	"github.com/mikerybka/webmachine/pkg/data"
	"github.com/mikerybka/webmachine/pkg/types"
)

type Endpoint struct {
	Filepath   string
	PathParams map[string]string
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
	rt := e.runtime(r.Method)

	// build params out of query params and path params
	params := map[string]string{}
	for k, v := range r.QueryParams {
		params[k] = v
	}
	for k, v := range e.PathParams {
		params[k] = v
	}

	// Build command
	command := []string{}
	for k, v := range params {
		command = append(command, fmt.Sprintf("--%s=%s", k, v))
	}
	path := filepath.Join(e.Filepath, r.Method, rt.FileName)
	command = append([]string{path}, command...)
	command = append(rt.CmdPrefix, command...)

	// Run command
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdin = bytes.NewReader(r.Body)
	cmd.Env = os.Environ()
	for k, v := range r.Headers {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Run()

	// Build response
	resp := types.NewResponse(r)
	for _, line := range bytes.Split(stdout.Bytes(), []byte("\n")) {
		if len(line) == 0 {
			break
		}
		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) != 2 {
			continue
		}
		resp.Headers[string(parts[0])] = string(parts[1])
	}
	status, body, ok := strings.Cut(stdout.String(), "\n")
	if !ok {
		resp.Status = 500
		resp.Body = append(stderr.Bytes(), stdout.Bytes()...)
		return resp
	}
	statusCode, err := strconv.Atoi(status)
	if err != nil {
		resp.Status = 500
		resp.Body = append(stderr.Bytes(), stdout.Bytes()...)
		return resp
	}
	resp.Status = statusCode
	resp.Body = []byte(body)
	return resp
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
