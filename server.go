package shorten

import (
	"github.com/gorilla/mux"
	"github.com/hectane/go-asyncserver"

	"fmt"
	"html/template"
	"net/http"
)

// Server responds to HTTP requests and routes them accordingly.
type Server struct {
	server   *server.AsyncServer
	config   *Config
	database *Database
	template *template.Template
}

// message represents a message to be displayed.
type message struct {
	Type string
	Body string
}

// adminHandler processes requests for the admin page. Access is restricted by
// HTTP basic auth.
func (s *Server) adminHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != s.config.AdminUsername || password != s.config.AdminPassword {
		w.Header().Set("WWW-Authenticate", "Basic realm=go-shorten")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	messages := make([]message, 0)
	defer func() {
		s.template.Execute(w, map[string]interface{}{
			"database": s.database,
			"messages": messages,
		})
	}()
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			messages = append(messages, message{"error", err.Error()})
		} else {
			switch r.Form.Get("action") {
			case "new":
				var (
					path        = r.Form.Get("path")
					destination = r.Form.Get("destination")
				)
				if len(path) == 0 || len(destination) == 0 {
					messages = append(messages, message{"error", "'path' and 'destination' cannot be empty"})
					return
				}
				s.database.Paths[path] = destination
				messages = append(messages, message{"info", "New item added"})
			case "delete":
				path := r.Form.Get("path")
				if len(path) == 0 {
					messages = append(messages, message{"error", "'path' cannot be empty"})
					return
				}
				delete(s.database.Paths, path)
				messages = append(messages, message{"info", fmt.Sprintf("'%s' deleted", path)})
			default:
				messages = append(messages, message{"error", fmt.Sprintf("Unrecognized action '%s'", r.Form.Get("action"))})
				return
			}
			if err := s.database.Save(); err != nil {
				messages = append(messages, message{"error", err.Error()})
				return
			}
		}
	}
}

// NewServer creates a new HTTP server.
func NewServer(cfg *Config, db *Database) (*Server, error) {
	var (
		r = mux.NewRouter()
		s = &Server{
			server:   server.New(cfg.Addr),
			config:   cfg,
			database: db,
			template: template.New("admin"),
		}
	)
	if len(cfg.AdminPassword) != 0 {
		r.HandleFunc(cfg.AdminPath, s.adminHandler)
	}
	r.NotFoundHandler = s
	s.server.Handler = r
	if _, err := s.template.Parse(adminTemplate); err != nil {
		return nil, err
	}
	if err := s.server.Start(); err != nil {
		return nil, err
	}
	return s, nil
}

// ServeHTTP processes requests for redirects.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dest, ok := s.database.Paths[r.URL.Path]
	if ok {
		http.Redirect(w, r, dest, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}

// Stop shuts down the HTTP server.
func (s *Server) Stop() {
	s.server.Stop()
}
