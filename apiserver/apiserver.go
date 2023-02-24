package apiserver

import (
	"database/sql"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	sessionStore := sessions.NewCookieStore([]byte("Test"))
	srv := newServer(db, sessionStore)
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
