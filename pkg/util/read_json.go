package util

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadJSON(filename string, v any) error {
	f, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(v)
}
