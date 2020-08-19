package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nikitych1w/softpro-task/pkg/config"
	"github.com/nikitych1w/softpro-task/pkg/store"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// HTTPserver ...
type HTTPserver struct {
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
	config *config.Config
	Server *http.Server
	url    string
}

// NewHTTPServer конструктор для http сервера
func NewHTTPServer(cfg *config.Config, lg *logrus.Logger, store *store.Store) *HTTPserver {
	s := &HTTPserver{
		router: mux.NewRouter(),
		logger: lg,
		store:  store,
		config: cfg,
		Server: &http.Server{},
	}

	s.url = fmt.Sprintf("%s:%s", s.config.Server.Host, s.config.Server.Port)
	s.Server.Addr = s.url
	s.configureRouter()
	s.Server.Handler = s.router

	return s
}

// ServeHTTP ...
func (s *HTTPserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// добавляет обработчик GET запроса и middleware для логгирования
func (s *HTTPserver) configureRouter() {
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/ready", s.healthCheck()).Methods("GET")
}

// настройка параметров логгирования для http сервера
func (s *HTTPserver) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

// endpoint, который показывает статус хранилища
func (s *HTTPserver) healthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var status int
		if err := s.store.Ping(); err != nil {
			status = http.StatusServiceUnavailable
		} else {
			status = http.StatusOK
		}
		s.respond(w, r, status, status)
	}
}

func (s *HTTPserver) respond(w http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.logger.Errorf("encode error | [%s]", err.Error())
			return
		}
	}
}

// Shutdown корректно завершает работу http сервера
func (s *HTTPserver) Shutdown(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	s.logger.Infof("		========= [HTTP server is stopping...]")

	return nil
}
