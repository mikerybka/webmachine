package main

import (
	"fmt"

	"github.com/mikerybka/webmachine/pkg/web"
)

func main() {
	accept := web.Header("Accept")
	fmt.Println(accept)
}
