package main

import (
	"flag"
	"fmt"
)

func main() {
	var pathvar string
	var wildcard string
	flag.StringVar(&pathvar, "pathvar", "", "pathvar")
	flag.StringVar(&wildcard, "wildcard", "", "wildcard")
	flag.Parse()
	fmt.Println(pathvar)
	fmt.Println(wildcard)
}
