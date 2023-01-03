package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
)

const Version = "dev"

type Config struct {
	HomeKit   HomeKitOpts  `yaml:"homekit"`
	Metrics   MetricsOpts  `yaml:"metrics"`
	Log       LogOpts      `yaml:"log"`
	HomeRelay HomeRelay    `yaml:"home-relay"`
	Daitem    DaitemConfig `yaml:"daitem"`
}

func LoadConfig() (*Config, error) {
	configPath := flag.String("config", "./config.yml", "config file path")
	flag.Parse()

	cfg := &Config{}

	err := cleanenv.ReadConfig(*configPath, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
