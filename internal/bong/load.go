package bong

import (
	"os"

	"gopkg.in/yaml.v3"
)

func loadBongs(name string) ([]Bong, error) {
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}

	raw, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	bongCol := BongCollection{}
	if err = yaml.Unmarshal(raw, &bongCol); err != nil {
		return nil, err
	}

	return bongCol.toBongs(), nil
}
