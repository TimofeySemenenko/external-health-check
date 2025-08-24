package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Domains []string `yaml:"domains"`
}

func Load(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	return &cfg
}
