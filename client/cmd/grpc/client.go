package grpc

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/snarksliveshere/banner-rotation/client/api/proto"
	"github.com/snarksliveshere/banner-rotation/client/configs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type GRPCConn struct {
	GConn  *grpc.ClientConn
	Client proto.BannerServiceClient
	Ctx    context.Context
	log    *zap.SugaredLogger
}

func createTimeStampFromTimeString(timeStr string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(configs.EventTimeLayout, timeStr)
	if err != nil {
		fmt.Println("bad time format")
		return nil, err
	}
	return ptypes.TimestampProto(t)
}

func Client(conf configs.AppConfig, log *zap.SugaredLogger) {
	ctx, _ := context.WithTimeout(context.Background(), configs.GRPCTimeoutCancel*time.Second)
	cc, err := grpc.Dial(conf.ListenIP+":"+conf.GRPCPort, grpc.WithInsecure())
	if err != nil {
		log.DPanic(err.Error())
		log.Fatalf("could not connect: %v", err)
	}
	client := proto.NewBannerServiceClient(cc)
	grpcConn := GRPCConn{
		GConn:  cc,
		Client: client,
		Ctx:    ctx,
		log:    log,
	}
	// здесь какая-нибудь логика. т.к. мне клиент нужен только для интеграционных тестов, я ее не реализую
	msg := proto.GetBannerRequestMessage{
		Audience: &proto.Audience{Id: "male_adult"},
		Slot:     &proto.Slot{Id: "top_slot_id"},
	}
	grpcConn.GetBanner(msg)
	defer func() { _ = grpcConn.GConn.Close() }()
}

func (g *GRPCConn) GetBanner(msg proto.GetBannerRequestMessage) *proto.GetBannerResponseMessage {
	resp, err := g.Client.SendGetBannerMessage(g.Ctx, &msg)
	if err != nil {
		g.log.DPanic(err.Error())
	}
	return resp
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

//protoc ./proto/events.proto --go_out=plugins=grpc:.
