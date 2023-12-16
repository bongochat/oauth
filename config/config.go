package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port       string
	Host       string
	Keyspace   string
	UserAPIURL string
}

var conf Config

func init() {
	if _, err := toml.DecodeFile("./conf.toml", &conf); err != nil {
		log.Println("Error getting configuration", err)
	}
}

func GetConfig() Config {
	return conf
}
