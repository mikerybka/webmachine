package util

import "os"

func DeleteFile(filename string) error {
	return os.Remove(filename)
}
