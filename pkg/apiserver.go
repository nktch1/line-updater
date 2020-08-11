package pkg

import (
	"context"
	"github.com/nikitych1w/softpro-task/pkg/config"
	"github.com/nikitych1w/softpro-task/pkg/httpserver"
	"github.com/nikitych1w/softpro-task/pkg/logger"
	"github.com/nikitych1w/softpro-task/pkg/model"
	"github.com/nikitych1w/softpro-task/pkg/rpcserver"
	store "github.com/nikitych1w/softpro-task/pkg/store"
	"github.com/nikitych1w/softpro-task/pkg/workers"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

// APIServer ...
type APIServer struct {
	cfg    *config.Config
	logger *logrus.Logger
	store  *store.Store
}

// NewAPIServer конструктор api сервера
func NewAPIServer() *APIServer {
	var as APIServer
	as.cfg = config.New()
	as.logger = logger.New(as.cfg)
	as.store = store.New(as.cfg)

	return &as
}

// Start отвечает за инициализацию воркеров и серверов, дожидается окончания их работы и корректно
// завершает их
func (s *APIServer) Start() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	rpcServer := rpcserver.NewRPCServer(s.cfg, s.logger, s.store)
	httpServer := httpserver.NewHTTPServer(s.cfg, s.logger, s.store)

	// тест корректного завершения всех воркеров, rpc сервера и http сервера
	//time.AfterFunc(20*time.Second, func() {
	//	stop <- os.Signal(os.Interrupt)
	//})

	w := workers.New(s.cfg, s.logger, s.store,
		[]model.Sport{
			model.NewSport("soccer"),
			model.NewSport("football"),
			model.NewSport("baseball"),
		})

	go func() {
		s.logger.Infof("		========= [workers are starting]")
		if err := w.RunWorkers(); err != nil {
			s.logger.Error(err)
		}
	}()

	go func() {
		s.logger.Infof("		========= [HTTP server is starting...]")
		if err := httpServer.Server.ListenAndServe(); err != nil {
			s.logger.Error(err)
		}
	}()

	for {
		if err := s.store.Ping(); err == nil {
			s.logger.Infof("		========= [HTTP server and database are available!]")
			go func() {
				s.logger.Infof("		========= [RPC server is starting...]")
				if err := rpcServer.ListenAndServe(); err != nil {
					s.logger.Error(err)
				}
			}()
			break
		}
	}

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		return err
	}

	if err := w.Shutdown(ctx); err != nil {
		return err
	}

	if err := rpcServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
