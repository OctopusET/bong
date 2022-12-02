package bong

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func SaveBongs(name string, bm BongMap) error {
	makeBongDir()
	yamlBong, err := yaml.Marshal(bm)
	if err != nil {
		return err
	}

	bongPath := filepath.Join("bongs", name)
	return os.WriteFile(bongPath, yamlBong, 0600)
}

func makeBongDir() error {
	return os.MkdirAll("bongs", os.ModePerm)
}
