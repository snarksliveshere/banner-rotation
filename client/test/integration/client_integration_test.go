package integration

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/banner-rotation/client/api/proto"
	"github.com/snarksliveshere/banner-rotation/client/cmd/grpc"
	"github.com/snarksliveshere/banner-rotation/client/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
)

var (
	conf configs.AppConfig
	slog *zap.SugaredLogger
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	slog = loggerInit()
}

func loggerInit() *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = customLevelEncoder
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err.Error())
	}
	sulog := logger.Sugar()
	defer func() { _ = sulog.Sync() }()
	return sulog
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

type banner struct {
	id string
}

type response struct {
	responseStatus string
	message        string
}

type notifyTest struct {
	banner   banner
	response response
}

func TestMain(m *testing.M) {
	fmt.Println("waiting 5s")
	time.Sleep(5 * time.Second)
	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:        "pretty",
		Paths:         []string{"features"},
		Randomize:     0,
		StopOnFailure: true,
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func FeatureContext(s *godog.Suite) {
	test := new(notifyTest)
	s.Step(`^I send request to GRPC server SendGetBannerMessage$`, test.iSendRequestToGRPCServerSendGetBannerMessage)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	s.Step(`^The response bannerId should not be empty string$`, test.theResponseBannerIdShouldNotBeEmptyString)
}

func (test *notifyTest) iSendRequestToGRPCServerSendGetBannerMessage() error {
	c := grpc.Client(conf, slog)

	msg := proto.GetBannerRequestMessage{
		Audience: &proto.Audience{Id: "male_adult"},
		Slot:     &proto.Slot{Id: "top_slot_id"},
	}
	reply, err := c.GetBanner(msg)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if reply == nil {
		errStr := fmt.Sprintf("nil resp from GetBanner with audience:%v,slot:%v", msg.Audience.Id, msg.Slot.Id)
		return fmt.Errorf(errStr)
	}
	test.response.responseStatus = reply.Response.Status.String()
	test.banner.id = reply.Banner.Id
	return nil
}

func (test *notifyTest) statusShouldBeEqualToSuccess(status string) error {
	if status != test.response.responseStatus {
		return fmt.Errorf("unexpected status: %s != %s", test.response.responseStatus, status)
	}
	return nil
}

func (test *notifyTest) theResponseBannerIdShouldNotBeEmptyString() error {
	if test.banner.id == "" {
		return fmt.Errorf("unexpected empty string instead banner id")
	}
	return nil
}

//func (test *notifyTest) statusShouldBeEqualToSuccess(status string) error {
//	if test.responseStruct.Status != status {
//		return fmt.Errorf("status must be: %s, not %s", status, test.responseStruct.Status)
//	}
//	return nil
//}
//
//func (test *notifyTest) statusShouldBeEqualToError(status string) error {
//	if test.responseStruct.Status != status {
//		return fmt.Errorf("status must be: %s, not %s", status, test.responseStruct.Status)
//	}
//	return nil
//}
