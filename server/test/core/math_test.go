package core_test

import (
	"github.com/go-pg/pg"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/banner-rotation/server/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"testing"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var (
	slog *zap.SugaredLogger
	db   *pg.DB
)

func init() {
	var conf configs.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	dbInst := configs.DB{Conf: &conf}
	db = dbInst.CreatePgConn()
	slog = loggerInit()
}

func TestGetBanner(t *testing.T) {

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
