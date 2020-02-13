package main

import (
	"github.com/snarksliveshere/banner-rotation/cmd/grpc"
	"github.com/snarksliveshere/banner-rotation/task"
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/snarksliveshere/banner-rotation/configs"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var conf configs.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")

	grpc.Server(conf)
	task.Run(db)
}
