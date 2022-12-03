package duck

import (
	"encoding/json"
	"net/url"
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
		b.BangUrl, _ = url.QueryUnescape(b.BangUrl)
		b.BangUrl = strings.ReplaceAll(b.BangUrl, "{{{s}}}", "%[1]s")
		b.MainUrl = "http://" + b.MainUrl
		fixed = append(fixed, b)
	}
	return
}

func toBongMap(bangs []DuckBang) bong.BongMap {
	fixed := fixBangs(bangs)
	bm := make(bong.BongMap)

	for _, bang := range fixed {
		bm[bang.Bang] = bong.Bong{
			Title:   bang.Title,
			MainUrl: bang.MainUrl,
			BongUrl: bang.BangUrl,
			Bongus:  bang.Bang,
		}
	}
	return bm
}
