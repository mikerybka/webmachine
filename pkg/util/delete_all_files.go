package util

import "os"

func DeleteAllFiles(dir string) error {
	return os.RemoveAll(dir)
}
