package duck

import (
	"fmt"

	"github.com/npmania/bong/internal/bong"
)

func grabDuckBangs() ([]DuckBang, error) {
	rawBang, err := fetch()
	if err != nil {
		return nil, err
	}

	return parseRawBang(rawBang)
}

func saveAsBong(name string, bangs []DuckBang) error {
	bongs := toBongs(bangs)
	return bong.SaveBongs(name, bongs)
}

func UpdateBangs() error {
	//TODO: check version before you grab bangs
	bangs, err := grabDuckBangs()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("duckduckgo-v%d.yaml", latestBangVersion())
	return saveAsBong(filename, bangs)
}
