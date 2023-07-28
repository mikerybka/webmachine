package ffi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func Run(header http.Header, body io.Reader, name string, args ...string) (status int, b []byte, err error) {
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ()
	for k := range header {
		key := k
		// key = strings.ReplaceAll(k, "-", "_")
		// key = strings.ToUpper(key)
		value := header.Get(k)
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Stdin = body
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	stdoutbytes, _ := io.ReadAll(stdout)
	stderrbytes, _ := io.ReadAll(stderr)
	if err != nil {
		status = http.StatusInternalServerError
		if len(stdoutbytes) > 0 {
			err = fmt.Errorf("%s\nstdout:\n%s", err, stdoutbytes)
		}
		if len(stderrbytes) > 0 {
			err = fmt.Errorf("%s\nstderr:\n%s", err, stderrbytes)
		}
		return status, nil, err
	}

	exitCode := cmd.ProcessState.ExitCode()
	fmt.Println(exitCode)
	if exitCode == 0 {
		status = 200
	} else if exitCode >= 100 && exitCode < 600 {
		status = exitCode
	} else {
		status = http.StatusInternalServerError
	}

	b = stdoutbytes

	return status, b, nil
}
