package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/mikerybka/webmachine/pkg/webmachine"
	"github.com/pkg/browser"
	"golang.org/x/crypto/acme/autocert"
)

func usage() {}

func main() {
	command := os.Args[1]
	switch command {
	case "init":
		cmdInit()
	case "dev":
		cmdDev()
	case "deploy":
		cmdDeploy()
	case "serve":
		cmdServe()
	default:
		usage()
	}
}

func cmdInit() {
	flag.Parse()
	name := flag.Arg(1)
	var runtimesString string = flag.Arg(2)
	runtimes := []*webmachine.Runtime{}
	for _, id := range strings.Split(runtimesString, ",") {
		rt, ok := webmachine.Runtimes[id]
		if ok {
			runtimes = append(runtimes, rt)
		} else if id != "" {
			fmt.Println("Unknown runtime ID:", id)
		}
	}

	if name == "" {
		flag.Usage()
		return
	}

	err := webmachine.Init(name, runtimes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cmdDev() {
	server := webmachine.Server{
		Dir:     ".",
		DevMode: true,
		Logging: true,
	}
	browser.OpenURL("http://localhost:3000")
	err := http.ListenAndServe(":3000", &server)
	if err != nil {
		panic(err)
	}
}

func cmdDeploy() {
	cmd := exec.Command("rsync", "-avz", "--delete", "--exclude", ".git", ".", "/etc/web/")

	remoteHostname := flag.Arg(1)
	if remoteHostname != "" {
		cmd = exec.Command("rsync", "-avz", "--delete", "--exclude", ".git", ".", "root@"+remoteHostname+":/etc/web/")
	}

	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func cmdServe() {
	server := webmachine.Server{
		Dir:     "/etc/web",
		Logging: true,
	}
	var email string
	if len(os.Args) >= 3 {
		email = os.Args[2]
	}
	err := serveHTTPS(&server, email, "/etc/certs")
	if err != nil {
		panic(err)
	}
}

// Use Let's Encrypt to fetch and renew certificates on any domain.
// serveHTTPS binds to ports 80 and 443 and serves the given handler.
// It uses a special handler for port 80 that can handle ACME challenges.
func serveHTTPS(s http.Handler, email, certDir string) error {
	// Create a channel to receive errors from the HTTP servers.
	errChan := make(chan error)

	// Define the autocert manager.
	// See https://godoc.org/golang.org/x/crypto/acme/autocert#Manager for details.
	m := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(certDir),
		HostPolicy: func(ctx context.Context, host string) error {
			return nil
		},
		Email: email,
	}

	// Start the HTTP server.
	go func() {
		err := http.ListenAndServe(":80", m.HTTPHandler(s))
		if err != nil {
			errChan <- err
		}
	}()

	// Start the HTTPS server.
	go func() {
		err := http.Serve(m.Listener(), s)
		if err != nil {
			errChan <- err
		}
	}()

	// Wait for an error.
	return <-errChan
}
