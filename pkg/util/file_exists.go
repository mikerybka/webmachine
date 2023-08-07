package util

import "os"

func FileExists(filename string) bool {
	fi, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !fi.IsDir()
}
