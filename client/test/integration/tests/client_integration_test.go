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
	"sync"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/streadway/amqp"
)

var (
	conf configs.AppConfig
	slog *zap.SugaredLogger
)

type BannerStatistics struct {
	Type     string `json:"type"`
	Slot     string `json:"slot"`
	Audience string `json:"audience"`
	Banner   string `json:"banner"`
	Time     string `json:"time"`
}

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
	banner        banner
	response      response
	errorGRPC     string
	errorRabbit   string
	conn          *amqp.Connection
	ch            *amqp.Channel
	messages      [][]byte
	messagesMutex *sync.RWMutex
	stopSignal    chan struct{}
}

func TestMain(m *testing.M) {
	fmt.Println("waiting 5s.....")
	time.Sleep(5 * time.Second)
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

func (test *notifyTest) errorMustNotBeEmpty() error {
	if test.errorGRPC == "" {
		test.errorGRPC = ""
		return fmt.Errorf("there is no error in error method")
	}
	test.errorGRPC = ""
	return nil
}

func (test *notifyTest) rabbitErrorMustNotBeEmpty() error {
	if test.errorRabbit == "" {
		test.errorRabbit = ""
		return fmt.Errorf("there is no error rabbit in error method")
	}
	test.errorRabbit = ""
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := new(notifyTest)

	s.BeforeScenario(test.startConsuming)

	// HealthCheck
	s.Step(`^I send request to GRPC SendHealthCheckMessage$`, test.iSendRequestToGRPCSendHealthCheckMessage)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)

	// GetBanner
	s.Step(`^I send request to GRPC SendGetBannerMessage with audience "([^"]*)" and slot "([^"]*)"$`, test.iSendRequestToGRPCSendGetBannerMessageWithAudienceAndSlot)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	s.Step(`^The response bannerId should not be empty string$`, test.theResponseBannerIdShouldNotBeEmptyString)
	//error
	s.Step(`^I send error request to GRPC SendGetBannerMessage with audience "([^"]*)" and slot "([^"]*)"$`, test.iSendErrorRequestToGRPCSendGetBannerMessageWithAudienceAndSlot)
	s.Step(`^Error must not be empty$`, test.errorMustNotBeEmpty)

	// check notification after GetBanner
	s.Step(`^I send request to GRPC SendGetBannerMessage with audience "([^"]*)" and slot "([^"]*)"$`, test.iSendRequestToGRPCSendGetBannerMessageWithAudienceAndSlot)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	s.Step(`^The response bannerId should not be empty string$`, test.theResponseBannerIdShouldNotBeEmptyString)
	s.Step(`^Notification after SendGetBannerMessage must contain type "([^"]*)" and audience "([^"]*)" and slot "([^"]*)"$`, test.notificationAfterSendGetBannerMessageMustContainTypeAndAudienceAndSlot)

	//error
	s.Step(`^I send request to GRPC SendGetBannerMessage with audience "([^"]*)" and slot "([^"]*)"$`, test.iSendRequestToGRPCSendGetBannerMessageWithAudienceAndSlot)
	s.Step(`^Error Notification after SendGetBannerMessage must contain type "([^"]*)" and audience "([^"]*)" and slot "([^"]*)"$`, test.errorNotificationAfterSendGetBannerMessageMustContainTypeAndAudienceAndSlot)
	s.Step(`^Rabbit Error must not be empty$`, test.rabbitErrorMustNotBeEmpty)

	//AddClick
	s.Step(`^I send request to GRPC SendAddClickBannerMessage with banner "([^"]*)" and slot "([^"]*)" and audience "([^"]*)"$`, test.iSendRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	//error
	s.Step(`^I send error request to GRPC SendAddClickBannerMessage with banner "([^"]*)" and slot "([^"]*)" and audience "([^"]*)"$`, test.iSendErrorRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience)
	s.Step(`^Error must not be empty$`, test.errorMustNotBeEmpty)

	//check notification after addClick
	s.Step(`^I send request to GRPC SendAddClickBannerMessage with banner "([^"]*)" and slot "([^"]*)" and audience "([^"]*)"$`, test.iSendRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	s.Step(`^Notification SendAddClickBannerMessage must contain type "([^"]*)" and banner "([^"]*)" and slot "([^"]*)" and audience "([^"]*)"$`, test.notificationSendAddClickBannerMessageMustContainTypeAndBannerAndSlotAndAudience)
	//error
	s.Step(`^I send request to GRPC SendAddClickBannerMessage with banner "([^"]*)" and slot "([^"]*)" and audience "([^"]*)"$`, test.iSendRequestToGRPCSendAddClickBannerMessageWithBannerAndSlotAndAudience)
	s.Step(`^Error Notification SendAddClickBannerMessage must contain type "([^"]*)" and banner "([^"]*)" and slot "([^"]*)" and audience "([^"]*)"$`, test.errorNotificationSendAddClickBannerMessageMustContainTypeAndBannerAndSlotAndAudience)
	s.Step(`^Rabbit Error must not be empty$`, test.rabbitErrorMustNotBeEmpty)

	//AddBanner2Slot
	s.Step(`^I send request to GRPC sendAddBannerToSlotMessage with banner "([^"]*)" and slot "([^"]*)"$`, test.iSendRequestToGRPCSendAddBannerToSlotMessageWithBannerAndSlot)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	// error
	s.Step(`^I send error request to GRPC sendAddBannerToSlotMessage with banner "([^"]*)" and slot "([^"]*)"$`, test.iSendErrorRequestToGRPCSendAddBannerToSlotMessageWithBannerAndSlot)
	s.Step(`^Error must not be empty$`, test.errorMustNotBeEmpty)

	//DeleteBannerFromSlot
	s.Step(`^I send request to GRPC sendDeleteBannerFromSlotMessage with banner "([^"]*)" and slot "([^"]*)"$`, test.iSendRequestToGRPCSendDeleteBannerFromSlotMessageWithBannerAndSlot)
	s.Step(`^Status should be equal to success "([^"]*)"$`, test.statusShouldBeEqualToSuccess)
	//error
	s.Step(`^I send error request to GRPC sendDeleteBannerFromSlotMessage with banner "([^"]*)" and slot "([^"]*)"$`, test.iSendErrorRequestToGRPCSendDeleteBannerFromSlotMessageWithBannerAndSlot)
	s.Step(`^Error must not be empty$`, test.errorMustNotBeEmpty)

	s.AfterScenario(test.stopConsuming)
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

func createRabbitConn() *amqp.Connection {
	strDial := "amqp://" + conf.RabbitUser + ":" + conf.RabbitPassword + "@" + conf.RabbitHost + ":" + conf.RabbitPort + "/"
	for {
		conn, err := amqp.Dial(strDial)
		if err == nil {
			return conn
		} else {
			slog.Errorf("INFO:Failed to connect to RabbitMQ with %s", err.Error())
			time.Sleep(1 * time.Second)
		}
	}
}
func (test *notifyTest) startConsuming(interface{}) {
	test.messages = make([][]byte, 0)
	test.messagesMutex = new(sync.RWMutex)
	test.stopSignal = make(chan struct{})

	var err error
	test.conn = createRabbitConn()
	panicOnErr(err)

	test.ch, err = test.conn.Channel()
	panicOnErr(err)

	// Consume
	_, err = test.ch.QueueDeclare(configs.BannerTestEx, true, false, true, false, nil)
	panicOnErr(err)

	err = test.ch.QueueBind(configs.BannerTestEx, "", configs.BannerStatEx, false, nil)
	panicOnErr(err)

	events, err := test.ch.Consume(configs.BannerTestEx, "", true, true, false, false, nil)
	panicOnErr(err)

	go func(stop <-chan struct{}) {
		for {
			select {
			case <-stop:
				return
			case event := <-events:
				test.messagesMutex.Lock()
				test.messages = append(test.messages, event.Body)
				test.messagesMutex.Unlock()
			}
		}
	}(test.stopSignal)
}

func (test *notifyTest) stopConsuming(interface{}, error) {
	test.stopSignal <- struct{}{}

	panicOnErr(test.ch.Close())
	panicOnErr(test.conn.Close())
	test.messages = nil
}
func panicOnErr(err error) {
	if err != nil {
		slog.DPanic(err)
	}
}
