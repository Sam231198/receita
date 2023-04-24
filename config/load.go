package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print("Error ao ler o arquivo de configuração")
	}

	viper.WatchConfig()
}