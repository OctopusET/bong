package bong

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func LoadBongs(filename string) (BongMap, error) {
	log.WithFields(log.Fields{
		"filename": filename,
	}).Info("loading bong from file")

	if _, err := os.Stat(filename); err != nil {
		log.WithFields(log.Fields{
			"filename": filename,
			"error":    err,
		}).Error("failed file lookup")
		return nil, err
	}

	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	bongMap := make(BongMap)
	if err = yaml.Unmarshal(raw, &bongMap); err != nil {
		return nil, err
	}

	for bg := range bongMap {
		b := bongMap[bg]
		b.MainUrl = strings.ReplaceAll(bongMap[bg].MainUrl, "%s", "%[1]s")
		b.BongUrl = strings.ReplaceAll(bongMap[bg].BongUrl, "%s", "%[1]s")
		bongMap[bg] = b
	}

	if err = bongMap.validate(); err != nil {
		return nil, err
	}

	return bongMap, nil
}
