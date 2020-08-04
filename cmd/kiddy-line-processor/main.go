package main

import (
	"github.com/nikitych1w/softpro-task/config"
	"github.com/nikitych1w/softpro-task/internal/kiddy-line-processor/apiserver"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()
	as := apiserver.New(cfg)
	err := apiserver.Start(as)

	if err != nil {
		logrus.Fatal(err)
	}
}
