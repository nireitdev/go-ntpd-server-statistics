package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Config struct {
		Device    string `yaml:"device"`
		Ip        string `yaml:"ip"`
		Timerange int    `yaml:"timerange"`
	} `yaml:"config"`
}

func ReadConfig() *Config {
	cfg := &Config{}
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("Error parsing config: ", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalf("Error parsing config: ", err)
	}

	return cfg
}
