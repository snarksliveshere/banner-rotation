package integration_test

import (
	"fmt"
	"github.com/snarksliveshere/banner-rotation/client/api/proto"
	"github.com/snarksliveshere/banner-rotation/client/cmd/grpc"
	"google.golang.org/grpc/status"
)

func (test *notifyTest) iSendRequestToGRPCSendGetBannerMessageWithAudienceAndSlot(audience, slot string) error {
	c := grpc.Client(conf, slog)

	msg := proto.GetBannerRequestMessage{
		Audience: &proto.Audience{Id: audience},
		Slot:     &proto.Slot{Id: slot},
	}
	reply, err := c.GetBanner(msg)
	if err != nil {
		return fmt.Errorf("error in method:%s: %s\n", "iSendRequestToGRPCSendGetBannerMessageWithAudienceAndSlot", status.Convert(err).Message())
	}
	if reply == nil {
		errStr := fmt.Sprintf("nil resp from GetBanner with audience:%v,slot:%v", msg.Audience.Id, msg.Slot.Id)
		return fmt.Errorf(errStr)
	}
	test.response.responseStatus = reply.Response.Status.String()
	test.banner.id = reply.Banner.Id
	return nil
}
