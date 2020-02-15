package integration_test

import (
	"fmt"
	"github.com/snarksliveshere/banner-rotation/client/api/proto"
	"github.com/snarksliveshere/banner-rotation/client/cmd/grpc"
	"google.golang.org/grpc/status"
)

func (test *notifyTest) iSendRequestToGRPCSendGetBannerMessageWithAudienceAndSlot(audience, slot string) error {
	for i := 0; i < 10; i++ {
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
	}
	return nil
}

func (test *notifyTest) iSendErrorRequestToGRPCSendGetBannerMessageWithAudienceAndSlot(audience, slot string) error {
	c := grpc.Client(conf, slog)

	msg := proto.GetBannerRequestMessage{
		Audience: &proto.Audience{Id: audience},
		Slot:     &proto.Slot{Id: slot},
	}
	_, err := c.GetBanner(msg)
	if err != nil {
		test.errorGRPC = status.Convert(err).Message()
		return nil
	}
	return fmt.Errorf("there is no error in error method:%s\n", "iSendErrorRequestToGRPCSendGetBannerMessageWithAudienceAndSlot")
}
