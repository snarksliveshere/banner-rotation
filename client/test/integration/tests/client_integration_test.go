package integration_test

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/banner-rotation/client/api/proto"
	"github.com/snarksliveshere/banner-rotation/client/cmd/grpc"
	"github.com/snarksliveshere/banner-rotation/client/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/status"
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
	fmt.Println("waiting 3s")
	time.Sleep(3 * time.Second)
	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:        "pretty",
		Paths:         []string{"../features"},
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
	// HealthCheck
	s.Step(`^I send request to GRPC SendHealthCheckMessage$`, test.iSendRequestToGRPCSendHealthCheckMessage)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)

	// GetBanner
	s.Step(`^I send request to GRPC SendGetBannerMessage with audience "([^"]*)" and slot "([^"]*)"$`, test.iSendRequestToGRPCSendGetBannerMessageWithAudienceAndSlot)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	s.Step(`^The response bannerId should not be empty string$`, test.theResponseBannerIdShouldNotBeEmptyString)

	//AddClick
	s.Step(`^I send request to GRPC SendAddClickBannerMessage with banner "([^"]*)" and slot "([^"]*)" and audience "([^"]*)"$`, test.iSendRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)

}

func (test *notifyTest) iSendRequestToGRPCSendHealthCheckMessage() error {
	c := grpc.Client(conf, slog)
	reply, err := c.GetHealthCheck(proto.Empty{})
	if err != nil {
		return fmt.Errorf("error in method:%s:%s\n", "iSendRequestToGRPCSendHealthCheckMessage", status.Convert(err).Message())
	}
	test.response.responseStatus = reply.Response.Status.String()
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
