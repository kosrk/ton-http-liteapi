package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

var Config = struct {
	LiteServer    string `env:"LITESERVER,required"`
	LiteServerKey string `env:"LITESERVER_KEY,required"`
	APIHost       string `env:"API_HOST" envDefault:"localhost:8081"`
}{}

func GetConfig() {
	err := env.Parse(&Config)
	if err != nil {
		log.Fatalf("Can not load config")
	}
}
