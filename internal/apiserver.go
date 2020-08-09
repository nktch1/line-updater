package internal

import (
	"context"
	"fmt"
	"github.com/nikitych1w/softpro-task/internal/config"
	"github.com/nikitych1w/softpro-task/internal/httpserver"
	"github.com/nikitych1w/softpro-task/internal/model"
	"github.com/nikitych1w/softpro-task/internal/rpcserver"
	"github.com/nikitych1w/softpro-task/internal/workers"
	"github.com/nikitych1w/softpro-task/pkg/logger"
	store "github.com/nikitych1w/softpro-task/pkg/store"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type APIServer struct {
	ctx    context.Context
	cfg    *config.Config
	logger *logrus.Logger
	store  *store.Store
}

func NewAPIServer(ctx context.Context) *APIServer {
	var as APIServer
	as.ctx = ctx
	as.cfg = config.New()
	as.logger = logger.New(as.cfg)
	as.store = store.New(as.cfg)

	return &as
}

func (s *APIServer) Start() error {
	var wg = &sync.WaitGroup{}
	defer wg.Wait()

	rpcURL := fmt.Sprintf("%s:%s", s.cfg.RPCServer.Host, s.cfg.RPCServer.Port)
	httpURL := fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)

	ctx, _ := context.WithCancel(s.ctx)

	rpc := rpcserver.NewRPCServer(s.logger, s.store)
	srv := httpserver.NewHTTPServer(s.store)

	w := workers.New(s.cfg, []model.Sport{
		model.NewSport("soccer"),
		model.NewSport("football"),
		model.NewSport("baseball"),
	}, s.store)

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.logger.Infof("		========= [workers are starting]")

		if err := w.RunWorkers(ctx); err != nil {
			s.logger.Error(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.logger.Infof("		========= [HTTP server is starting at '%s']", httpURL)

		if err := http.ListenAndServe(httpURL, srv); err != nil {
			s.logger.Error(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.logger.Infof("		========= [GRPC server is starting at '%s']", rpcURL)

		if err := rpc.ListenAndServe(rpcURL, ctx); err != nil {
			s.logger.Error(err)
		}
	}()

	return nil
}
