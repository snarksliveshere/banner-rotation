package grpc

import (
	"context"
	"github.com/snarksliveshere/banner-rotation/server/api/proto"
	"github.com/snarksliveshere/banner-rotation/server/internal/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//type Response struct {
//	Date   entity.Date    `json:"day,omitempty"`
//	Event  entity.Event   `json:"event,omitempty"`
//	Events []entity.Event `json:"events,omitempty"`
//	Error  string         `json:"error,omitempty"`
//	Status string         `json:"status,omitempty"`
//	//Result     []string      `json:"result,omitempty"`
//}

func (s ServerBanner) SendGetBannerMessage(ctx context.Context, msg *proto.GetBannerRequestMessage) (*proto.GetBannerResponseMessage, error) {
	banner, err := task.ReturnBanner(s.db, s.log, msg.Audience.Id, msg.Slot.Id)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	reply := proto.GetBannerResponseMessage{Banner: &proto.Banner{Id: banner}}
	return &reply, nil
}

func (s ServerBanner) SendAddBannerToSlotMessage(context.Context, *proto.AddBannerToSlotRequestMessage) (*proto.ResponseBannerMessage, error) {
	panic("implement me")
}

func (s ServerBanner) SendDeleteBannerFromSlotMessage(context.Context, *proto.DeleteBannerFromSlotRequestMessage) (*proto.ResponseBannerMessage, error) {
	panic("implement me")
}

func (s ServerBanner) SendAddClickBannerMessage(context.Context, *proto.AddClickRequestMessage) (*proto.ResponseBannerMessage, error) {
	panic("implement me")
}

//func (s ServerCalendar) SendGetEventsForTimeIntervalMessage(ctx context.Context, msg *proto.GetEventsForTimeIntervalRequestMessage) (*proto.GetEventsForTimeIntervalResponseMessage, error) {
//	from, till, err := data_handlers.CheckGetEventByTimeIntervalFromProtoTimestamp(msg.From, msg.Till)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, "invalid incoming times")
//	}
//	events, err := storage.Actions.EventRepository.GetEventsByTimeInterval(from, till)
//	reply := proto.GetEventsForTimeIntervalResponseMessage{}
//
//	if err != nil {
//		reply.Status = config.StatusError
//		reply.Text = err.Error()
//		return &reply, nil
//	}
//	reply.Status = config.StatusSuccess
//	eventsJson, err := json.Marshal(&events)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, "invalid json marshall")
//	}
//	reply.Events = string(eventsJson)
//
//	return &reply, nil
//}