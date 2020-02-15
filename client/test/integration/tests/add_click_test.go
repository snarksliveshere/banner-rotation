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
		return fmt.Errorf("error in method: %s : %s\n",
			"iSendRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience", status.Convert(err).Message())
	}
	test.response.responseStatus = reply.Response.Status.String()
	return nil
}
