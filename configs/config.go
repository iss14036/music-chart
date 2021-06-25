package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName        string
	AppPort        string
	DBUri          string
	LogLevel       string
	Environment    string
	DbDriver       string
	DbLog          bool
	AuthKey        string
	AuthExpiration int
}

var configPath = "."

func Reader() *Config {
	var cfg Config

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("Fatal error configs file: %s \n", err))
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("Fatal error configs file: %s \n", err))
	}

	return &cfg
}
