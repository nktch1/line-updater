package main

import (
	"context"
	"github.com/nikitych1w/softpro-task/internal"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	if err := internal.NewAPIServer(ctx).Start(); err != nil {
		logrus.Fatal(err)
	}
}
