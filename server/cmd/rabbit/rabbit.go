package rabbit

import (
	"github.com/snarksliveshere/banner-rotation/server/configs"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
	"time"
)

func RabbitCMD(conf configs.AppConfig, slog *zap.SugaredLogger) {
	conn := createRabbitConn(conf, slog)
	defer func() { _ = conn.Close() }()
	ch := createChannel(conn)
	defer func() { _ = ch.Close() }()
	rabbitServer(ch)
}

func RabbitCreateServer(channel *amqp.Channel) {
	rabbitServer(channel)
}

func RabbitCreateConnection(conf configs.AppConfig, slog *zap.SugaredLogger) *amqp.Connection {
	return createRabbitConn(conf, slog)
}

func RabbitCreateChannel(conn *amqp.Connection) *amqp.Channel {
	return createChannel(conn)
}

func createRabbitConn(conf configs.AppConfig, slog *zap.SugaredLogger) *amqp.Connection {
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

func rabbitServer(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		configs.BannerStatEx, // name
		"fanout",             // type
		true,                 // durable
		false,                // auto-deleted
		false,                // internal
		false,                // no-wait
		nil,                  // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	_, err = ch.QueueDeclare(
		configs.BannerStatQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	err = ch.QueueBind(configs.BannerStatQueue, "", configs.BannerStatEx, false, nil)

	failOnError(err, "Failed to bind a queue")
}

func createChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
