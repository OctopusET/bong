package bong

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func SaveBongs(name string, bongs []Bong) error {
	makeBongDir()
	bongCol := toCollection(bongs)
	yamlBong, err := yaml.Marshal(bongCol)
	if err != nil {
		return err
	}

	bongPath := filepath.Join("bongs", name)
	return os.WriteFile(bongPath, yamlBong, 0600)
}

func makeBongDir() error {
	return os.MkdirAll("bongs", os.ModePerm)
}
