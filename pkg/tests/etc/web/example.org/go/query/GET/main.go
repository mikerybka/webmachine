package main

import (
	"flag"
	"fmt"
)

func main() {
	var q string
	var b string
	flag.StringVar(&q, "q", "", "q")
	flag.StringVar(&b, "b", "", "b")
	flag.Parse()
	fmt.Printf("q: %s\n", q)
	fmt.Printf("b: %s\n", b)
}
