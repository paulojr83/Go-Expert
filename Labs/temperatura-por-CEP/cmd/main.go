package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/temperatura-por-CEP/internal/api/wep/controller"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AutomaticEnv()
	/*viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}*/
}

func main() {
	router := gin.Default()
	router.GET("/cep/:cep", controller.FindCep)

	PORT := viper.GetString("HTTP_PORT")
	err := router.Run(PORT)
	log.Println("Starting server on port: ", PORT)
	if err != nil {
		panic(err)
		return
	}
}
