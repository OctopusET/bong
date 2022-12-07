package coward

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/npmania/bong/internal/bong"
	log "github.com/sirupsen/logrus"
)

const timeFormat = "20060102150405"

func saveAsBong(name string, bangs []cowBang) error {
	bm, err := toBongMap(bangs)
	if err != nil {
		return err
	}
	return bong.SaveBongs(name, bm)
}

func UpdateBangs() error {
	cf := new(cowFetcher)

	rawBang, err := cf.fetch()
	if err != nil {
		return err
	}

	bangs, err := parseBang(rawBang)
	if err != nil {
		return err
	}

	filename := filepath.Join("bongs", fmt.Sprintf("brave-%s.yaml", time.Now().Format(timeFormat)))
	err = saveAsBong(filename, bangs)
	if err != nil {
		return err
	}

	log.Infof("Saved latest Brave Search bang to %s", filename)
	return nil
}
