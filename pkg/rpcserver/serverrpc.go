package rpcserver

import (
	"context"
	"fmt"
	"github.com/nikitych1w/softpro-task/config"
	"github.com/nikitych1w/softpro-task/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type ServerRPC struct {
	listener net.Listener
	srv      *grpc.Server
	port     string
}

func New(cfg *config.Config) (*ServerRPC, error) {
	var s ServerRPC
	var err error
	s.port = cfg.ServerRPC.Port

	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return nil, err
	}

	s.srv = grpc.NewServer()
	proto.RegisterLineProcessorServer(s.srv, &s)
	reflection.Register(s.srv)
	return &s, nil
}

func (s *ServerRPC) Start() error {
	logrus.Info("rpc started at port ", s.port)
	return s.srv.Serve(s.listener)
}

func (s *ServerRPC) SubscribeOnSportsLines(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	return nil, nil
}
