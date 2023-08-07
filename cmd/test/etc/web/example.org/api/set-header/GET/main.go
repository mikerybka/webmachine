package main

import "github.com/mikerybka/webmachine/pkg/web"

func main() {
	web.SetResponseHeader("test", "123")
	web.Return(200, "")
}
