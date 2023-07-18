package golang

import (
	"io"
	"net/http"
	"os/exec"
)

func Run(path string, input io.Reader) (status int, body []byte, err error) {
	cmd := exec.Command("go", "run", path)
	cmd.Stdin = input
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	err = cmd.Start()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	out, err := io.ReadAll(stdout)
	if err != nil {
		return http.StatusInternalServerError, out, err
	}
	errout, err := io.ReadAll(stderr)
	if err != nil {
		return http.StatusInternalServerError, out, err
	}
	err = cmd.Wait()
	if err != nil {
		return http.StatusInternalServerError, out, err
	}
	if len(errout) > 0 {
		switch string(errout) {
		case "200":
			return http.StatusOK, out, nil
		case "201":
			return http.StatusCreated, out, nil
		case "204":
			return http.StatusNoContent, out, nil
		case "400":
			return http.StatusBadRequest, out, nil
		case "401":
			return http.StatusUnauthorized, out, nil
		case "403":
			return http.StatusForbidden, out, nil
		case "404":
			return http.StatusNotFound, out, nil
		case "405":
			return http.StatusMethodNotAllowed, out, nil
		case "409":
			return http.StatusConflict, out, nil
		case "500":
			return http.StatusInternalServerError, out, nil
		case "501":
			return http.StatusNotImplemented, out, nil
		case "503":
			return http.StatusServiceUnavailable, out, nil
		case "504":
			return http.StatusGatewayTimeout, out, nil
		case "507":
			return http.StatusInsufficientStorage, out, nil
		default:
			return http.StatusInternalServerError, out, err
		}
	}
	return http.StatusOK, out, nil
}
