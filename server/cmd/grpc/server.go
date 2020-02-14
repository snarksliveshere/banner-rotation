package grpc

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/api/proto"
	"github.com/snarksliveshere/banner-rotation/server/configs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type ServerBanner struct {
	db  *pg.DB
	log *zap.SugaredLogger
}

func Server(conf configs.AppConfig, log *zap.SugaredLogger) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	go func() { goGRPC(conf, log) }()
	<-stopCh
}

func goGRPC(conf configs.AppConfig, log *zap.SugaredLogger) {
	listenAddr := conf.ListenIP + ":" + conf.GRPCPort
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.DPanicf("failed to listen addr: %s, error: %v\n", listenAddr, err.Error())
	}
	grpcServer := grpc.NewServer()
	dbInst := configs.DB{Conf: &conf}
	serverBanner := ServerBanner{
		db:  dbInst.CreatePgConn(),
		log: log,
	}
	proto.RegisterBannerServiceServer(grpcServer, serverBanner)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.DPanic(err.Error())
	}
	log.Infof("Run GRPC server on: %s\n", listenAddr)
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.
