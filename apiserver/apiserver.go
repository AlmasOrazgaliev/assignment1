package apiserver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	srv := newServer(db)

	return http.ListenAndServe(config.Port, srv)
}

func newDB(DatabaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DatabaseURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

//func (s *APIServer) configureRouter() {
//	s.router.HandleFunc("/home/", s.handleHome())
//}
//
//func (s *APIServer) handleHome() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "home")
//	}
//}

//
//func  Start() error {
//	s.configureRouter()
//	err := s.configureStore()
//	if err != nil {
//		return err
//	}
//	return http.ListenAndServe(s.config.Port, s.router)
//}
//
//func (s *APIServer) configureStore() error {
//
//	st := db.New(s.config.Store)
//	err := st.Open()
//	if err != nil {
//		return err
//	}
//	s.store = st
//	return nil
//}
//
