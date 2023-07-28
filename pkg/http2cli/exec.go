package http2cli

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func Exec(r *http.Request, command []string, pathArgs map[string]string) *Response {
	if len(command) == 0 {
		panic("no command")
	}

	// build args out of query params and args
	args := map[string]string{}
	for k, v := range pathArgs {
		args[k] = v
	}
	for k := range r.URL.Query() {
		args[k] = r.URL.Query().Get(k)
	}

	// Turn args to string list
	arglist := []string{}
	for k, v := range args {
		arglist = append(arglist, fmt.Sprintf("--%s=%s", k, v))
	}

	command = append(command, arglist...)

	// Build command
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = os.Environ()

	// Set headers
	for k, v := range r.Header {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v[0]))
	}

	// Set body
	cmd.Stdin = r.Body

	// Run command
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Run()
	exitCode := cmd.ProcessState.ExitCode()

	// Read response headers
	headers := map[string]string{}
	for _, line := range bytes.Split(stdout.Bytes(), []byte("\n")) {
		if len(line) == 0 {
			break
		}
		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) != 2 {
			continue
		}
		headers[string(parts[0])] = string(parts[1])
	}

	// Calculate status code
	statusCode := 500
	if exitCode == 0 {
		statusCode = 200
	} else if exitCode >= 100 && exitCode < 600 {
		statusCode = exitCode
	}

	return &Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       stdout.Bytes(),
	}
}
