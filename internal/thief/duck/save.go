package duck

import (
	"fmt"

	"github.com/npmania/bong/internal/bong"
	log "github.com/sirupsen/logrus"
)

func saveAsBong(name string, bangs []duckBang) error {
	bm, err := toBongMap(bangs)
	if err != nil {
		return err
	}
	return bong.SaveBongs(name, bm)
}

func UpdateBangs() error {
	bf := new(bangFetcher)

	rawBang, err := bf.fetch()
	if err != nil {
		return err
	}

	bangs, err := parseBang(rawBang)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("bongs/duckduckgo-v%d.yaml", bf.latestVersion())

	err = saveAsBong(filename, bangs)
	if err != nil {
		return err
	}

	log.Infof("Saved latest Duckduckgo bang to %s", filename)
	return nil
}
