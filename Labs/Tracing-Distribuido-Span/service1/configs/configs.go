package configs

type Conf struct {
	Service2        string `mapstructure:"SERVICE2_API"`
	RequestNameOTEL string `mapstructure:"OTEL_SERVICE_NAME"`
}
