package web

import (
	"fmt"
)

func Redirect(url string) {
	SetResponseHeader("Location", url)
	fmt.Println("307")
}
