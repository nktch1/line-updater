package rpcserver

import (
	"context"
	"fmt"
	"github.com/nikitych1w/softpro-task/pkg/store"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
)

type RPCServer struct {
	listener net.Listener
	srv      *grpc.Server
	logger   *logrus.Logger
	store    *store.Store
}

type reqParams struct {
	sportsToUpdate []string
	updTime        int
}

func NewRPCServer(lg *logrus.Logger, str *store.Store) *RPCServer {
	var s RPCServer

	s.logger = lg
	s.store = str

	s.srv = grpc.NewServer()
	RegisterLineProcessorServer(s.srv, &s)
	reflection.Register(s.srv)

	return &s
}

func (s *RPCServer) ListenAndServe(url string, ctx context.Context) error {
	var err error
	s.listener, err = net.Listen("tcp", url)
	if err != nil {
		logrus.Error(err)
	}

	return s.srv.Serve(s.listener)
}

func (s *RPCServer) subscribe(stream LineProcessor_SubscribeOnSportsLinesServer, pipe chan *Request) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("GRPC stream: (EOF)")
			return nil
		}

		if err != nil {
			fmt.Println("GRPC stream:", err)
			return err
		}

		logrus.Println(in.GetSports(), in.GetSports())

		pipe <- in
	}
}

func (s *RPCServer) broadcast(stream LineProcessor_SubscribeOnSportsLinesServer, pipe chan *Request) error {
	mtx := &sync.Mutex{}
	rp := &reqParams{}

	for {
		select {
		case req := <-pipe:
			rp.sportsToUpdate = req.GetSports()
			val, err := strconv.Atoi(req.GetTimeUpd())
			if err != nil {
				return err
			}

			rp.updTime = val
			respData := s.buildResponse(rp, mtx)

			err = stream.Send(&Response{
				Line: respData,
			})

			if err != nil {
				return err
			}

			s.logger.Info("\t\t ---> [GRPC] : NEW STREAM", rp, respData)

		default:
			if len(rp.sportsToUpdate) > 0 {
				for range time.Tick(time.Duration(rp.updTime) * time.Second) {
					respData := s.buildResponse(rp, mtx)
					err := stream.Send(&Response{
						Line: respData,
					})

					if err != nil {
						return err
					}
					s.logger.Info("\t ---> [GRPC] : SENT TO STREAM", rp, respData)
				}
			}
		}
	}
}

func (s *RPCServer) buildResponse(rp *reqParams, mtx *sync.Mutex) map[string]float32 {
	respData := map[string]float32{}

	for _, el := range rp.sportsToUpdate {
		val, err := s.store.GetLastValueByKey(el)
		if err != nil {
			logrus.Println(err)
		}
		mtx.Lock()
		respData[el] = val
		mtx.Unlock()
	}

	return respData
}

func (s *RPCServer) SubscribeOnSportsLines(stream LineProcessor_SubscribeOnSportsLinesServer) error {
	logrus.Println("SubscribeOnSportsLines")
	pipe := make(chan *Request)
	go s.subscribe(stream, pipe)
	s.broadcast(stream, pipe)

	return nil
}
