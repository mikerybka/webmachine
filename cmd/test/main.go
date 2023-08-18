package main

import (
	"net/http"
	"time"

	"github.com/mikerybka/webmachine/pkg/tests"
	"github.com/mikerybka/webmachine/pkg/webmachine"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create the server
	server := webmachine.Server{
		Dir: "cmd/test/etc/web",
	}

	// Start the server
	go func() {
		err := http.ListenAndServe(":3000", &server)
		check(err)
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Test
	err := tests.GetIndexHTML()
	check(err)
	err = tests.GetIndexJSON()
	check(err)
	err = tests.Get()
	check(err)
	err = tests.Post()
	check(err)
	err = tests.Put()
	check(err)
	err = tests.Delete()
	check(err)
	err = tests.Patch()
	check(err)
	err = tests.PathParams()
	check(err)
	err = tests.QueryParams()
	check(err)
	err = tests.RequestHeaders()
	check(err)
	err = tests.WildcardPathParam()
	check(err)
	err = tests.RequestBody()
	check(err)
	err = tests.ResponseBody()
	check(err)
	err = tests.ResponseStatusCode()
	check(err)
	err = tests.ResponseHeaders()
	check(err)
}
