package bong

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadBongs(name string) (BongMap, error) {
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}

	raw, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	bongMap := make(BongMap)
	if err = yaml.Unmarshal(raw, &bongMap); err != nil {
		return nil, err
	}

	if err = bongMap.validate(); err != nil {
		return nil, err
	}

	return bongMap, nil
}
