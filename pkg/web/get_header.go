package web

import "os"

func Header(key string) string {
	return os.Getenv(key)
}
