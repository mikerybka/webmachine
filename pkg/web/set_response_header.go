package web

import (
	"fmt"
	"os"
)

func SetResponseHeader(name, value string) {
	line := fmt.Sprintf("%s: %s\n", name, value)
	os.Stderr.WriteString(line)
}
