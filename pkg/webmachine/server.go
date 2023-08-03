package webmachine

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/paths"
	"github.com/mikerybka/webmachine/pkg/types"
)

type Server struct {
	Dir     string
	Args    map[string]string
	DevMode bool
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := types.NewRequest(r)
	req.Log(os.Stdout)
	path := s.path(r)
	endpoint, err := s.endpoint(path, r.Method)
	if errors.Is(err, os.ErrNotExist) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		panic(err)
	}
	res := endpoint.Handle(req)
	res.Write(w)
	res.Log(os.Stdout)
}

func (s *Server) path(r *http.Request) string {
	path := filepath.Join(r.Host, r.URL.Path)
	if s.DevMode {
		path = r.URL.Path[1:]
	}
	return path
}

func (s *Server) endpoint(path, method string) (*Endpoint, error) {
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

			// Ignore HTTP method directories
			if entry.Name() == "GET" || entry.Name() == "POST" || entry.Name() == "PUT" || entry.Name() == "DELETE" {
				continue
			}

			// Remember the catchall name
			if len(entry.Name()) > 2 && entry.Name()[0] == '_' && entry.Name()[1] == '_' {
				catchall = entry.Name()
			}

			// If there is an exact match
			if entry.Name() == p[0] {
				newDir := filepath.Join(s.Dir, entry.Name())
				s := Server{Dir: newDir, Args: s.Args}
				return s.endpoint(paths.Join(p[1:]), method)
			}
		}

		if catchall != "" {
			newDir := filepath.Join(s.Dir, catchall)
			s := Server{Dir: newDir, Args: s.Args}

			// Record the arg
			key := catchall[2:]
			value := p[0]
			if s.Args == nil {
				s.Args = map[string]string{}
			}
			s.Args[key] = value

			return s.endpoint(paths.Join(p[1:]), method)
		}

		// If there is no catchall, 404
		return nil, os.ErrNotExist
	}

	return &Endpoint{Filepath: s.Dir, PathParams: s.Args}, nil
}
