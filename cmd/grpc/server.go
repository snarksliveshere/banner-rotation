package grpc

import (
	"context"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/api/proto"
	"github.com/snarksliveshere/banner-rotation/configs"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/config"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type ServerBanner struct {
	db *pg.DB
}

var (
	log *logrus.Logger
)

func Server(conf configs.AppConfig) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	goGRPC(conf)
	<-stopCh
}

func goGRPC(conf configs.AppConfig) {
	listenAddr := conf.ListenIP + ":" + conf.GRPCPort
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("failed to listen addr: %s, error: %v\n", listenAddr, err.Error())
	}

	grpcServer := grpc.NewServer()

	dbInst := configs.DB{Conf: &conf}
	serverBanner := ServerBanner{db: dbInst.CreatePgConn()}
	proto.RegisterBannerServiceServer(grpcServer, serverBanner)

	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Infof("Run grpc server on: %s\n", listenAddr)
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.
