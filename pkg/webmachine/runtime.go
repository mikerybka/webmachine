package webmachine

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Runtime struct {
	InitScript        string
	InstallDepsScript string
	AddDepScript      string
}

func (r *Runtime) Init() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	name := filepath.Base(dir)
	script := strings.ReplaceAll(r.InitScript, "{{.Name}}", name)
	cmd := exec.Command("bash", "-c", script)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
