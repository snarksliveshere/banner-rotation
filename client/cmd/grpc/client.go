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

func Client(conf configs.AppConfig, log *zap.SugaredLogger) *GRPCConn {
	ctx, _ := context.WithTimeout(context.Background(), configs.GRPCTimeoutCancel*time.Second)
	//ctx := context.Background()
	cc, err := grpc.Dial(conf.ListenIP+":"+conf.GRPCPort, grpc.WithInsecure())
	if err != nil {
		log.DPanic(err.Error())
		log.Fatalf("could not connect: %v", err)
	}
	client := proto.NewBannerServiceClient(cc)
	return &GRPCConn{
		GConn:  cc,
		Client: client,
		Ctx:    ctx,
		log:    log,
	}
	// здесь какая-нибудь логика. т.к. мне клиент нужен только для интеграционных тестов, я ее не реализую
	//msg := proto.GetBannerRequestMessage{
	//	Audience: &proto.Audience{Id: "male_adult"},
	//	Slot:     &proto.Slot{Id: "top_slot_id"},
	//}
	//reply, err := grpcConn.GetBanner(msg)
	//if err != nil {
	//	log.Info(err.Error())
	//}
	//if reply == nil {
	//	log.Errorf("nil resp from GetBanner with audience:%v,slot:%v", msg.Audience.Id, msg.Slot.Id)
	//}
	//log.Info("banner ID:", reply.Banner.Id, "response status:", reply.Response.Status)

}

func (g *GRPCConn) GetHealthCheck(msg proto.Empty) (*proto.ResponseBannerMessage, error) {
	defer func() { _ = g.GConn.Close() }()
	resp, err := g.Client.SendHealthCheckMessage(g.Ctx, &msg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GRPCConn) GetBanner(msg proto.GetBannerRequestMessage) (*proto.GetBannerResponseMessage, error) {
	defer func() { _ = g.GConn.Close() }()
	resp, err := g.Client.SendGetBannerMessage(g.Ctx, &msg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GRPCConn) AddClick(msg proto.AddClickRequestMessage) (*proto.ResponseBannerMessage, error) {
	defer func() { _ = g.GConn.Close() }()
	resp, err := g.Client.SendAddClickBannerMessage(g.Ctx, &msg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GRPCConn) AddBannerToSlot(msg proto.AddBannerToSlotRequestMessage) (*proto.ResponseBannerMessage, error) {
	defer func() { _ = g.GConn.Close() }()
	resp, err := g.Client.SendAddBannerToSlotMessage(g.Ctx, &msg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GRPCConn) DeleteBannerFromSlot(msg proto.DeleteBannerFromSlotRequestMessage) (*proto.ResponseBannerMessage, error) {
	defer func() { _ = g.GConn.Close() }()
	resp, err := g.Client.SendDeleteBannerFromSlotMessage(g.Ctx, &msg)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.
