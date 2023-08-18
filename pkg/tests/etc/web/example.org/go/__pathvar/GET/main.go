package main

import (
	"flag"
	"fmt"
)

func main() {
	var pathvar string
	flag.StringVar(&pathvar, "pathvar", "", "pathvar")
	flag.Parse()
	fmt.Println(pathvar)
}
