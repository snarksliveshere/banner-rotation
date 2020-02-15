package integration_test

import (
	"fmt"
	"github.com/snarksliveshere/banner-rotation/client/api/proto"
	"github.com/snarksliveshere/banner-rotation/client/cmd/grpc"
	"google.golang.org/grpc/status"
)

func (test *notifyTest) iSendRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience(banner, slot, audience string) error {
	c := grpc.Client(conf, slog)

	msg := proto.AddClickRequestMessage{
		Banner:   &proto.Banner{Id: banner},
		Audience: &proto.Audience{Id: audience},
		Slot:     &proto.Slot{Id: slot},
	}

	reply, err := c.AddClick(msg)
	if err != nil {
		return fmt.Errorf("error in method: %s : %s\n", "AddClick", status.Convert(err).Message())
	}
	test.response.responseStatus = reply.Response.Status.String()
	return nil
}

func (test *notifyTest) iSendErrorRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience(banner, slot, audience string) error {
	c := grpc.Client(conf, slog)

	msg := proto.AddClickRequestMessage{
		Banner:   &proto.Banner{Id: banner},
		Audience: &proto.Audience{Id: audience},
		Slot:     &proto.Slot{Id: slot},
	}

	_, err := c.AddClick(msg)

	if err != nil {
		test.errorGRPC = status.Convert(err).Message()
		return nil
	}
	return fmt.Errorf("there is no error in error method:%s\n", "iSendErrorRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience")
}
