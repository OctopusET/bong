package bong

import (
	"os"
	"strings"

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

	for bg := range bongMap {
		b := bongMap[bg]
		b.BongUrl = strings.ReplaceAll(bongMap[bg].BongUrl, "%s", "%[1]s")
		bongMap[bg] = b
	}

	if err = bongMap.validate(); err != nil {
		return nil, err
	}

	return bongMap, nil
}
