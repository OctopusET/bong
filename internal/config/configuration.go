package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

const exampleFilename = "config.example.yaml"
const filename = "config.yaml"

type Config struct {
	Title         string `yaml:"Title"`
	BaseUrl       string `yaml:"Base URL"`
	Port          int    `yaml:"Port"`
	DefaultPrefix string `yaml:"Default Prefix"`
}

func CopyIfNotExists() error {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		data, err := os.ReadFile(exampleFilename)
		if err != nil {
			return err
		}

		if err = os.WriteFile(filename, data, 0600); err != nil {
			return err
		}
	}

	return nil
}

func LoadConfig() (Config, error) {
	c := Config{}

	if err := CopyIfNotExists(); err != nil {
		return Config{}, err
	}

	raw, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	if err = yaml.Unmarshal(raw, &c); err != nil {
		return Config{}, err
	}

	return c, nil
}
