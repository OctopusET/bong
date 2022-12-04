package duck

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/npmania/bong/internal/bong"
)

type duckBang struct {
	Category     string `json:"c"`
	MainUrl      string `json:"d"`
	SomeWeirdNum int    `json:"r"`
	Title        string `json:"s"`
	SubCategory  string `json:"sc"`
	Bang         string `json:"t"`
	BangUrl      string `json:"u"`
}

func parseRawBang(raw []byte) (bangs []duckBang, err error) {
	err = json.Unmarshal(raw, &bangs)
	return
}

func fixBangs(bangs []duckBang) (fixed []duckBang, err error) {
	for _, b := range bangs {
		if string(b.Title[0]) == " " {
			b.Title = b.Title[1:]
		}

		b.MainUrl, err = url.QueryUnescape(b.MainUrl)
		if err != nil {
			return nil, err
		}

		b.BangUrl, err = url.QueryUnescape(b.BangUrl)
		if err != nil {
			return nil, err
		}

		b.MainUrl = strings.ReplaceAll(b.MainUrl, "{{{s}}}", "%[1]s")
		b.BangUrl = strings.ReplaceAll(b.BangUrl, "{{{s}}}", "%[1]s")

		b.MainUrl = "http://" + b.MainUrl

		fixed = append(fixed, b)
	}
	return
}

func toBongMap(bangs []duckBang) (bong.BongMap, error) {
	fixed, err := fixBangs(bangs)
	if err != nil {
		return nil, err
	}

	bm := make(bong.BongMap)

	for _, bang := range fixed {
		bm[bang.Bang] = bong.Bong{
			Title:   bang.Title,
			MainUrl: bang.MainUrl,
			BongUrl: bang.BangUrl,
			Bongus:  bang.Bang,
		}
	}

	return bm, nil
}
