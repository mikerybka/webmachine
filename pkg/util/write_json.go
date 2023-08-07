package util

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func WriteJSON(filename string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		return err
	}
	b = append(b, '\n')
	return os.WriteFile(filename, b, 0644)
}
