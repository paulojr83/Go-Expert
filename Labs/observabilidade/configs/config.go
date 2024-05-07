package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	Title              string `mapstructure:"TITLE"`
	BackgroundColor    string `mapstructure:"BACKGROUND_COLOR"`
	ResponseTime       int    `mapstructure:"RESPONSE_TIME"`
	ExternalCallUrl    string `mapstructure:"EXTERNAL_CALL_URL"`
	ExternalCallMethod string `mapstructure:"EXTERNAL_CALL_METHOD"`
	RequestNameOTEL    string `mapstructure:"OTEL_SERVICE_NAME"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
