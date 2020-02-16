package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/banner-rotation/server/cmd/grpc"
	"github.com/snarksliveshere/banner-rotation/server/cmd/rabbit"
	"github.com/snarksliveshere/banner-rotation/server/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var conf configs.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	slog := loggerInit()
	rabbitConn := rabbit.RabbitCreateConnection(conf, slog)
	defer func() { _ = rabbitConn.Close() }()
	rabbitChannel := rabbit.RabbitCreateChannel(rabbitConn)
	defer func() { _ = rabbitChannel.Close() }()
	go func() { grpc.Server(conf, slog, rabbitChannel) }()
	go func() { rabbit.RabbitCreateServer(rabbitChannel) }()

	<-stopCh
}

func loggerInit() *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = customLevelEncoder
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err.Error())
	}
	slog := logger.Sugar()
	defer func() { _ = slog.Sync() }()
	slog.Info("GRPC Server Starts")
	return slog
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
