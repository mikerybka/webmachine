package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/paths"
)

type Server struct {
	Dir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(r.Host, r.URL.Path)

	// Log the request
	fmt.Println(r.Method, path)

	// Get the endpoint
	endpoint, err := s.Endpoint(path, r.Method)
	if errors.Is(err, os.ErrNotExist) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the endpoint
	endpoint.ServeHTTP(w, r)
}

func (s *Server) Endpoint(path, method string) (*Endpoint, error) {
	p := paths.Parse(path)

	// If this is not the root endpoint, find the next directory to work from and trim the path
	if len(p) > 0 {
		// Read the directory
		entries, err := os.ReadDir(s.Dir)
		if err != nil {
			return nil, err
		}
		catchall := ""
		for _, entry := range entries {
			// Ignore hidden files
			if entry.Name()[0] == '.' {
				continue
			}

			// Remember the catchall name
			if entry.Name()[0] == '[' && entry.Name()[len(entry.Name())-1] == ']' {
				catchall = entry.Name()
			}

			// If there is an exact match
			if entry.Name() == p[0] {
				path := filepath.Join(s.Dir, entry.Name())
				return (&Server{Dir: path}).Endpoint(paths.Join(p[1:]), method)
			}
		}

		// Return the catchall, if there is one
		if catchall != "" {
			path := filepath.Join(s.Dir, catchall)
			return (&Server{Dir: path}).Endpoint(paths.Join(p[1:]), method)
		}

		// If there is no catchall, 404
		return nil, os.ErrNotExist
	}

	return &Endpoint{Filepath: s.Dir}, nil
}
