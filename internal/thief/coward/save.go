package coward

import (
	"fmt"
	"time"

	"github.com/npmania/bong/internal/bong"
)

const timeFormat = "2006-01-02-15:04:05"

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

	filename := fmt.Sprintf("brave-%s.yaml", time.Now().Format(timeFormat))
	return saveAsBong(filename, bangs)
}
