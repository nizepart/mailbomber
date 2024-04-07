package apiserver

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/nizepart/rest-go/internal/app/email_service"
	"github.com/nizepart/rest-go/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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

	s.emailService.Start()
	s.startEmailScheduler()

	s.configureRouter()

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
