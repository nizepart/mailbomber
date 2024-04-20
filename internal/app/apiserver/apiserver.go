package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nizepart/rest-go/internal/app"
	"github.com/nizepart/rest-go/internal/app/store/sqlstore"
)

func Start() error {
	db, err := newDB(app.GetValue("DB_URL", "").String())
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(app.GetValue("SESSION_KEY", "").String()))
	s := newServer(store, sessionStore)

	defer s.emailService.Close()

	return http.ListenAndServe(app.GetValue("BIND_ADDR", ":8080").String(), s)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
