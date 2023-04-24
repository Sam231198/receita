package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samuel/receita/config"
	"github.com/spf13/viper"
)

func main() {
	server := gin.Default()

	config.Load()
	config.LoadMiddlewares(server)
	config.LoadControllers(server)

	server.Run(viper.GetString("app.port"))
}