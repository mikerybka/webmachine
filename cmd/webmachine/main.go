package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/mikerybka/webmachine/pkg/webmachine"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	email := flag.String("email", "", "The email to share with Let's Encrypt.")
	dir := flag.String("dir", "/etc/webmachine", "The root directory to serve.")
	certDir := flag.String("cert-dir", "/etc/webmachine/certs", "The directory to store certificates in.")
	port := flag.String("port", "", "The port to listen on. Defaults to both 443 and 80. If a port is not provided, HTTPS is served on 443 and HTTP served on 80. If the port provided is 443, HTTPS is served on port 443. If any other port is provided, HTTP is served on that port.")
	devMode := flag.Bool("dev", false, "Run in development mode. This will serve files from the root directory instead of the subdomain directory.")
	flag.Parse()
	server := webmachine.Server{
		Dir:     *dir,
		DevMode: *devMode,
	}
	var err error
	if *port == "" {
		err = serveTLS(&server, *email, *certDir)
	} else {
		err = http.ListenAndServe(":"+*port, &server)
	}
	if err != nil {
		panic(err)
	}
}

// Use Let's Encrypt to fetch and renew certificates on any domain.
// serveTLS binds to ports 80 and 443 and serves the given handler.
// It uses a special handler for port 80 that can handle ACME challenges.
func serveTLS(s http.Handler, email, certDir string) error {
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
