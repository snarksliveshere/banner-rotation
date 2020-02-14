package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/banner-rotation/client/cmd/grpc"
	"github.com/snarksliveshere/banner-rotation/client/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var conf configs.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")

	grpc.Client(conf, loggerInit())
	//task.Run(db)
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
	//slog.Infow("failed to fetch URL",
	//	"url", "http://example.com",
	//	"attempt", 3,
	//	"backoff", time.Second,
	//)
	slog.Info("Start...")
	return slog
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
