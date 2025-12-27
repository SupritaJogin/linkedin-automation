package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Browser struct {
		UserAgent      string `yaml:"user_agent"`
		ViewportWidth  int    `yaml:"viewport_width"`
		ViewportHeight int    `yaml:"viewport_height"`
		Headless       bool   `yaml:"headless"`
	} `yaml:"browser"`

	Timing struct {
		MinDelay int `yaml:"min_delay_ms"`
		MaxDelay int `yaml:"max_delay_ms"`
	} `yaml:"timing"`

	Limits struct {
		DailyConnections int `yaml:"daily_connections"`
		DailyMessages    int `yaml:"daily_messages"`
	} `yaml:"limits"`

	Schedule struct {
		StartHour int `yaml:"start_hour"`
		EndHour   int `yaml:"end_hour"`
	} `yaml:"schedule"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
