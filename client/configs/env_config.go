package configs

type DbConfig struct {
	DBDriver   string `envconfig:"DB_DRIVER_CLIENT" default:"postgres"`
	DBHost     string `envconfig:"DB_HOST_CLIENT" required:"true"`
	DBPort     string `envconfig:"DB_PORT_CLIENT" required:"true"`
	DBUser     string `envconfig:"DB_USER_CLIENT" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD_CLIENT" required:"true"`
	DBName     string `envconfig:"DB_NAME_CLIENT" required:"true"`
}

type Addr struct {
	GRPCPort string `envconfig:"GRPC_PORT" required:"true"`
	ListenIP string `envconfig:"LISTEN_IP"`
}

type AppConfig struct {
	DbConfig
	Addr
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
}
