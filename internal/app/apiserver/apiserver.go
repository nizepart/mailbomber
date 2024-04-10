package apiserver

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"github.com/nizepart/rest-go/internal/app"
	"github.com/nizepart/rest-go/internal/app/store/sqlstore"
	"net/http"
)

func Start() error {
	db, err := newDB(app.GetEnvString("host=$DB_HOST user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable", ""))
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(app.GetEnvString("SESSION_KEY", "")))
	s := newServer(store, sessionStore)

	defer s.emailService.Close()

	return http.ListenAndServe(app.GetEnvString("BIND_ADDR", ":8080"), s)
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
