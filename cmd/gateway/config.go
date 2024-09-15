package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HTTP  serverConfig `envconfig:"HTTP"`
	Kafka kafkaConfig  `envconfig:"KAFKA"`
}

func configFromEnv() (*config, error) {
	var c config
	err := envconfig.Process("GATEWAY", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type serverConfig struct {
	Host string `envconfig:"HOST" default:"0.0.0.0"` //localhost
	Port int    `envconfig:"PORT" default:"8080"`
}

type kafkaConfig struct {
	URL string `envconfig:"KAFKA_URL" default:"localhost:29092"`
}
