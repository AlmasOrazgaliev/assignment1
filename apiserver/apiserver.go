package apiserver

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	router *mux.Router
	db     *DB
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/home/", s.handleHome())
}

func (s *APIServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "home")
	}
}
