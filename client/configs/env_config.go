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

type Rabbit struct {
	RabbitHost     string `envconfig:"CLIENT_RABBIT_HOST" required:"true"`
	RabbitPort     string `envconfig:"CLIENT_RABBIT_PORT" required:"true"`
	RabbitUser     string `envconfig:"CLIENT_RABBIT_USER" required:"true"`
	RabbitPassword string `envconfig:"CLIENT_RABBIT_PASSWORD" required:"true"`
}

type AppConfig struct {
	DbConfig
	Addr
	Rabbit
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
}
