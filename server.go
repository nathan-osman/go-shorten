package main

import (
	"github.com/gorilla/mux"
	"github.com/hectane/go-asyncserver"

	"net/http"
)

// Server responds to HTTP requests and routes them accordingly.
type Server struct {
	server   *server.AsyncServer
	database *Database
}

// adminHandler processes requests for the admin page.
func (s *Server) adminHandler(w http.ResponseWriter, r *http.Request) {
	//...
}

// NewServer creates a new HTTP server.
func NewServer(cfg *Config, db *Database) (*Server, error) {
	var (
		r = mux.NewRouter()
		s = &Server{
			server:   server.New(cfg.Addr),
			database: db,
		}
	)
	r.HandleFunc(cfg.AdminPath, s.adminHandler)
	r.NotFoundHandler = s
	if err := s.server.Start(); err != nil {
		return nil, err
	}
	return s, nil
}

// ServeHTTP processes requests for redirects.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//...
}

// Stop shuts down the HTTP server.
func (s *Server) Stop() {
	s.server.Stop()
}
