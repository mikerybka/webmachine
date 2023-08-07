package web

import (
	"fmt"
	"os"
)

func Return(statusCode int, body string) {
	fmt.Println(statusCode)
	fmt.Println(body)
	os.Exit(0)
}
