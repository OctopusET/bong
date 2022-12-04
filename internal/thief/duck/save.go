package duck

import (
	"fmt"

	"github.com/npmania/bong/internal/bong"
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
	return saveAsBong(filename, bangs)
}
