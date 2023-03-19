package config

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const exampleFilename = "config.example.yaml"
const filename = "config.yaml"

type Config struct {
	Title         string `yaml:"Title"`
	BaseUrl       string `yaml:"Base URL"`
	BongFile      string `yaml:"Bong File"`
	Port          int    `yaml:"Port"`
	DefaultPrefix string `yaml:"Default Prefix"`
	Fallback      string `yaml:"Fallback,omitempty"`
}

func CopyIfNotExists() error {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		data, err := os.ReadFile(exampleFilename)
		if err != nil {
			return fmt.Errorf("failed reading %s: %s", exampleFilename, err.Error())
		}

		if err = os.WriteFile(filename, data, 0600); err != nil {
			return fmt.Errorf("failed writing to %s: %s", filename, err.Error())
		}
		log.Info("config.yaml does not exist. Copied example config")
	}

	return nil
}

// TODO: should return *Config
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
