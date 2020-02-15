package grpc

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/banner-rotation/server/api/proto"
	"github.com/snarksliveshere/banner-rotation/server/configs"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type ServerBanner struct {
	db      *pg.DB
	log     *zap.SugaredLogger
	channel *amqp.Channel
}

func Server(conf configs.AppConfig, log *zap.SugaredLogger, channel *amqp.Channel) {
	goGRPC(conf, log, channel)
}

func goGRPC(conf configs.AppConfig, log *zap.SugaredLogger, channel *amqp.Channel) {
	listenAddr := conf.ListenIP + ":" + conf.GRPCPort
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.DPanicf("failed to listen addr: %s, error: %v\n", listenAddr, err.Error())
	}
	grpcServer := grpc.NewServer()
	dbInst := configs.DB{Conf: &conf}
	serverBanner := ServerBanner{
		db:      dbInst.CreatePgConn(),
		log:     log,
		channel: channel,
	}
	proto.RegisterBannerServiceServer(grpcServer, serverBanner)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.DPanic(err.Error())
	}
	log.Infof("Run GRPC server on: %s\n", listenAddr)
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.
