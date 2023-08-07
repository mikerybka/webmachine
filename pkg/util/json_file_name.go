package util

import "path/filepath"

func JSONFileName(path ...string) string {
	return filepath.Join(path...) + ".json"
}
