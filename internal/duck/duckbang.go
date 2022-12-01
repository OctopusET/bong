package duck

import (
	"encoding/json"
	"strings"

	"github.com/npmania/bong/internal/bong"
)

type DuckBang struct {
	Category     string `json:"c"`
	MainUrl      string `json:"d"`
	SomeWeirdNum int    `json:"r"`
	Title        string `json:"s"`
	SubCategory  string `json:"sc"`
	Bang         string `json:"t"`
	BangUrl      string `json:"u"`
}

func parseRawBang(raw []byte) (bangs []DuckBang, err error) {
	err = json.Unmarshal(raw, &bangs)
	return
}

func fixBangs(bangs []DuckBang) (fixed []DuckBang) {
	for _, b := range bangs {
		if string(b.Title[0]) == " " {
			b.Title = b.Title[1:]
		}
		b.BangUrl = strings.Replace(b.BangUrl, "{{{s}}}", "%s", 1)
		fixed = append(fixed, b)
	}
	return
}

func toBongs(bangs []DuckBang) (bongs []bong.Bong) {
	fixed := fixBangs(bangs)
	for _, bang := range fixed {
		b := bong.Bong{
			Title:   bang.Title,
			MainUrl: bang.MainUrl,
			BongUrl: bang.BangUrl,
			Bongus:  bang.Bang,
		}
		bongs = append(bongs, b)
	}
	return
}
