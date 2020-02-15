package integration_test

import (
	"fmt"
	"github.com/snarksliveshere/banner-rotation/client/api/proto"
	"github.com/snarksliveshere/banner-rotation/client/cmd/grpc"
	"google.golang.org/grpc/status"
)

func (test *notifyTest) iSendRequestToGRPCSendAddBannerToSlotMessageWithBannerAndSlot(banner, slot string) error {
	c := grpc.Client(conf, slog)

	msg := proto.AddBannerToSlotRequestMessage{
		Banner: &proto.Banner{Id: banner},
		Slot:   &proto.Slot{Id: slot},
	}
	reply, err := c.AddBannerToSlot(msg)
	if err != nil {
		return fmt.Errorf("error in method:%s: %s\n", "iSendRequestToGRPCSendAddBannerToSlotMessageWithBannerAndSlot", status.Convert(err).Message())
	}
	if reply == nil {
		errStr := fmt.Sprintf("nil resp from AddBannerToSlot with audience:%v,slot:%v", msg.Banner.Id, msg.Slot.Id)
		return fmt.Errorf(errStr)
	}
	test.response.responseStatus = reply.Response.Status.String()
	return nil
}

func (test *notifyTest) iSendErrorRequestToGRPCSendAddBannerToSlotMessageWithBannerAndSlot(banner, slot string) error {
	c := grpc.Client(conf, slog)

	msg := proto.AddBannerToSlotRequestMessage{
		Banner: &proto.Banner{Id: banner},
		Slot:   &proto.Slot{Id: slot},
	}
	_, err := c.AddBannerToSlot(msg)
	if err != nil {
		test.errorGRPC = status.Convert(err).Message()
		return nil
	}
	return fmt.Errorf("there is no error in error method:%s\n", "iSendErrorRequestToGRPCSendAddBannerToSlotMessageWithBannerAndSlot")
}
