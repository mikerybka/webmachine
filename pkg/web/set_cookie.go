package web

import "fmt"

func SetCookie(name, value string) {
	SetResponseHeader("Set-Cookie", fmt.Sprintf("%s=%s", name, value))
}
