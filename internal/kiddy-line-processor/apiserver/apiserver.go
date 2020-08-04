package apiserver

import (
	"fmt"
	"github.com/nikitych1w/softpro-task/config"
	store "github.com/nikitych1w/softpro-task/internal/kiddy-line-processor/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	cfg    *config.Config
	logger *logrus.Logger
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		cfg:    config,
		logger: logrus.New(),
	}
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.cfg.Log.Level)
	if err != nil {
		return err
	}

	s.logger.Level = level

	return nil
}

func Start(s *APIServer) error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting API server at ", s.cfg.Server.Host, ":", s.cfg.Server.Port)

	st := store.NewStore(fmt.Sprintf("%s:%s", s.cfg.Database.Host, s.cfg.Database.Port),
		s.cfg.Database.Password)

	defer st.Client.Close()

	srv := newServer(st)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port), srv)
}
