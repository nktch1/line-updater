package rpcserver

import (
	"context"
	"fmt"
	"github.com/nikitych1w/softpro-task/internal/config"
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
	Server   *grpc.Server
	logger   *logrus.Logger
	store    *store.Store
	prevResp map[string]float32
	mtx      *sync.Mutex
	wg       *sync.WaitGroup
	ctx      context.Context
	cfg      *config.Config
	url      string
	done     chan struct{}
}

type reqParams struct {
	sportsToUpdate []string
	updTime        int
}

type rawToDelta struct {
	raw, delta float32
}

type requestAndPreviousDelta struct {
	req  *Request
	prev map[string]rawToDelta
}

// конструктор для rpc сервера
func NewRPCServer(cfg *config.Config, lg *logrus.Logger, str *store.Store) *RPCServer {
	var s RPCServer

	s.logger = lg
	s.store = str
	s.prevResp = make(map[string]float32)
	s.mtx = &sync.Mutex{}
	s.Server = grpc.NewServer()
	s.wg = &sync.WaitGroup{}
	s.cfg = cfg
	s.url = fmt.Sprintf("%s:%s", s.cfg.RPCServer.Host, s.cfg.RPCServer.Port)
	s.done = make(chan struct{})

	// регистрация сервера для обработки grpc запросов
	RegisterLineProcessorServer(s.Server, &s)
	reflection.Register(s.Server)

	return &s
}

// запускает прослушивание порта
func (s *RPCServer) ListenAndServe() error {
	var err error
	s.listener, err = net.Listen("tcp", s.url)
	if err != nil {
		logrus.Error(err)
	}

	return s.Server.Serve(s.listener)
}

// регистрирует запросы и обрабатывает их
func (s *RPCServer) process(stream LineProcessor_SubscribeOnSportsLinesServer) error {
	subscribeRequests := make(chan requestAndPreviousDelta)
	prevResp := make(map[string]rawToDelta)

	// запись подписок в канал
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("GRPC stream: (EOF)")
				return
			}
			if err != nil {
				fmt.Println("GRPC stream:", err)
				return
			}

			subscribeRequests <- requestAndPreviousDelta{
				req:  in,
				prev: prevResp,
			}
		}
	}()

	// чтение из канала, сравнение с предыдущим значение и высчитывание дельты
	for request := range subscribeRequests {
		var val int
		var err error

		rp := reqParams{}
		rp.sportsToUpdate = request.req.GetSports()

		if val, err = strconv.Atoi(request.req.GetTimeUpd()); err != nil {
			s.logger.Errorf("GRPC stream: (can'requestAndPreviousDelta convert interval value | [%s])", err.Error())
			return nil
		}
		rp.updTime = val

		s.wg.Add(1)
		go func(rp reqParams, prevResp map[string]rawToDelta) {
			defer s.wg.Done()

			for {
				select {
				case <-s.done:
					return
				default:
					data := s.buildResponse(rp, prevResp)
					respData := make(map[string]float32)
					for k, v := range data {
						respData[k] = v.delta
					}

					if err := stream.Send(&Response{Line: respData}); err != nil {
						s.logger.Errorf("GRPC stream: (streaming error | [%s])", err.Error())
						return
					}

					s.logger.Info("\t ---> [GRPC] : SENT TO STREAM ", rp, respData)

					s.mtx.Lock()
					prevResp = data
					s.mtx.Unlock()

					time.Sleep(time.Duration(rp.updTime) * time.Second)
				}
			}
		}(rp, request.prev)
	}

	s.wg.Wait()

	return nil
}

// создает ответ для отдачи в стрим
func (s *RPCServer) buildResponse(rp reqParams, prevResp map[string]rawToDelta) map[string]rawToDelta {
	currResp := make(map[string]rawToDelta)

	for _, el := range rp.sportsToUpdate {
		val, err := s.store.GetLastValueByKey(el)
		if err != nil {
			s.logger.Errorf("GRPC stream: (getting from store error | [%s])", err.Error())
			return nil
		}

		var res float32
		if len(prevResp) > 0 {
			res = val - prevResp[el].delta
		} else {
			res = val
		}

		s.mtx.Lock()
		currResp[el] = rawToDelta{
			raw:   val,
			delta: res,
		}
		s.mtx.Unlock()
	}

	return currResp
}

// endpoint из .proto
func (s *RPCServer) SubscribeOnSportsLines(stream LineProcessor_SubscribeOnSportsLinesServer) error {
	err := s.process(stream)
	if err != nil {
		s.logger.Errorf("SubscribeOnSportsLines method error | [%s]", err.Error())
		return err
	}
	return nil
}

// корректное завершение работы GRPC сервера
func (s *RPCServer) Shutdown(ctx context.Context) error {
	s.logger.Infof("		========= [RPC server is stopping...]")

	for {
		s.done <- struct{}{}
	}

	s.Server.GracefulStop()
	s.listener.Close()

	return nil
}
