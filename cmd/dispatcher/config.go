package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Kafka kafkaConfig `envconfig:"KAFKA"`
}

func configFromEnv() (*config, error) {
	var c config
	err := envconfig.Process("DISPATCHER", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type kafkaConfig struct {
	URL string `envconfig:"KAFKA_URL" default:"localhost:29092"`
}
