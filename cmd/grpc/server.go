package grpc

import (
	"context"
	"fmt"
	"github.com/snarksliveshere/banner-rotation/api/proto"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/config"
	"google.golang.org/grpc"
	"net"
)

type ServerBanner struct {
}

var (
	log     *logrus.Logger
	storage *pg_repository.Storage
)

func Server(logg *logrus.Logger, conf *config.AppConfig) {
	log = logg
	log.Infof("start grpc, conf #%v", conf)
	storage = pg_repository.CreateStorageInstance(log, conf)
	//stopCh := make(chan os.Signal, 1)
	//signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	goGRPC(conf)
	//<-stopCh
}

func goGRPC(conf *config.AppConfig) {
	listenAddr := conf.ListenIP + ":" + conf.GRPCPort
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen addr: %s, error: %v\n", listenAddr, err.Error())
	}

	grpcServer := grpc.NewServer()
	proto.RegisterBannerServiceServer(grpcServer, ServerBanner{})

	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Infof("Run grpc server on: %s\n", listenAddr)
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.
