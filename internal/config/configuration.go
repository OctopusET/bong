package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const filename = "config.yaml"

type Config struct {
	Title         string `yaml:"Title"`
	BaseUrl       string `yaml:"Base URL"`
	Port          int    `yaml:"Port"`
	DefaultPrefix string `yaml:"Default Prefix"`
}

func LoadConfig() (Config, error) {
	c := Config{}

	raw, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	if err = yaml.Unmarshal(raw, &c); err != nil {
		return Config{}, err
	}

	return c, nil
}
