package workers

import (
	"fmt"
	"github.com/nikitych1w/softpro-task/config"
	"github.com/nikitych1w/softpro-task/pkg/model"
	"github.com/sirupsen/logrus"
	"strings"
)

type worker struct {
	url     string
	updTime int
}

type BackgroundWorkers struct {
	workers []worker
}

func New(cfg *config.Config, w []model.Sport) *BackgroundWorkers {
	var bckWorkers BackgroundWorkers
	for _, el := range w {
		var updTime int
		var url string

		switch el {
		case model.Soccer:
			updTime = cfg.Request.UpdateIntervalSoccer
		case model.Football:
			updTime = cfg.Request.UpdateIntervalFootball
		case model.Baseball:
			updTime = cfg.Request.UpdateIntervalBaseball
		}

		url = fmt.Sprintf("%s/%s", cfg.LineProvider.URL, strings.ToLower(el.String()))

		bckWorkers.workers = append(bckWorkers.workers, worker{
			url:     url,
			updTime: updTime,
		})

		logrus.Infof("url: [%s]; updTime: [%dms]", url, updTime)
	}

	return &bckWorkers
}

func (w *BackgroundWorkers) RunWorkers() error {
	// TO DO
	return nil
}
