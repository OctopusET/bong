package bong

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func SaveBongs(name string, bm BongMap) error {
	if err := bm.validate(); err != nil {
		return err
	}

	for bg := range bm {
		b := bm[bg]
		b.MainUrl = strings.ReplaceAll(b.MainUrl, "%[1]s", "%s")
		b.BongUrl = strings.ReplaceAll(b.BongUrl, "%[1]s", "%s")
		bm[bg] = b
	}

	makeBongDir()
	yamlBong, err := yaml.Marshal(bm)
	if err != nil {
		return err
	}

	bongPath := name
	return os.WriteFile(bongPath, yamlBong, 0600)
}

func makeBongDir() error {
	return os.MkdirAll("bongs", os.ModePerm)
}
