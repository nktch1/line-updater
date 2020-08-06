package main

import (
	"github.com/nikitych1w/softpro-task/config"
	"github.com/nikitych1w/softpro-task/pkg/apiserver"
	"github.com/nikitych1w/softpro-task/pkg/model"
	"github.com/nikitych1w/softpro-task/pkg/rpcserver"
	"github.com/nikitych1w/softpro-task/pkg/workers"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()

	w := workers.New(cfg, []model.Sport{
		model.NewSport("SOCCER"),
		model.NewSport("FOOTBALL"),
		model.NewSport("BASEBALL"),
	})

	err := w.RunWorkers()
	if err != nil {
		logrus.Fatal(err)
	}

	as := apiserver.New(cfg)
	err = apiserver.Start(as)

	rpc, _ := rpcserver.New(cfg)
	err = rpc.Start()

	if err != nil {
		logrus.Fatal(err)
	}
}
