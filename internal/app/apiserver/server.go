package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/nizepart/mailbomber/internal/app"
	"github.com/nizepart/mailbomber/internal/app/email_service"
	"github.com/nizepart/mailbomber/internal/app/store"
	"github.com/sirupsen/logrus"
)

const (
	ctxKeyRequestID = iota
)

type server struct {
	router        *mux.Router
	logger        *logrus.Logger
	store         store.Store
	sessionsStore sessions.Store
	emailService  *email_service.Service
}

func newServer(store store.Store, sessionsStore sessions.Store) *server {
	s := &server{
		router:        mux.NewRouter(),
		logger:        logrus.New(),
		store:         store,
		sessionsStore: sessionsStore,
		emailService:  email_service.NewService(),
	}

	if err := s.configureLogger(); err != nil {
		s.logger.Fatalf("Error configuring logger: %v", err)
	}

	s.emailService.Start()
	s.startEmailScheduler()
	s.logger.Info("Email service started")

	s.configureRouter()

	s.logger.Info("Server successfully started")
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.LogRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")

	emails := private.PathPrefix("/email").Subrouter()
	emails.HandleFunc("/send", s.handleEmailSend()).Methods("POST")

	template := emails.PathPrefix("/template").Subrouter()
	template.HandleFunc("/schedule", s.handleEmailScheduleCreate()).Methods("POST")
	template.HandleFunc("/create", s.handleEmailTemplateCreate()).Methods("POST")
	template.HandleFunc("/{id}", s.handleEmailTemplateGet()).Methods("GET")
	template.HandleFunc("/{id}/send", s.handleEmailSend()).Methods("POST")
	//template.HandleFunc("/{id}/schedule", s.handleEmailSchedule()).Methods("POST")
}

func (s *server) configureLogger() error {
	level, err := logrus.ParseLevel(app.GetValue("LOG_LEVEL", "debug").String())
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	s.logger.SetOutput(os.Stdout)

	return nil
}

func (s *server) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof("completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
	})
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
