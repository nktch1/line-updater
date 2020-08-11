package main

import (
	"github.com/nikitych1w/softpro-task/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := pkg.NewAPIServer().Start(); err != nil {
		logrus.Fatal(err)
	}
}
