package webmachine

import (
	"errors"
	"fmt"
	"os"
)

func Init(name string, runtimes []*Runtime) error {
	_, err := os.Stat(name)
	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%s already exists", name)
	}

	err = os.MkdirAll(name, 0755)
	if err != nil {
		return err
	}
	err = os.Chdir(name)
	if err != nil {
		return err
	}

	for _, runtime := range runtimes {
		err := runtime.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
