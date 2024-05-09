package configs

type Conf struct {
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	ViaCepApi     string `mapstructure:"VIA_CEP_API"`
	WeatherApi    string `mapstructure:"WEATHER_API"`
}
