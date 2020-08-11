package main

import (
	"context"
	"github.com/nikitych1w/softpro-task/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	if err := pkg.NewAPIServer(ctx).Start(); err != nil {
		logrus.Fatal(err)
	}
}
