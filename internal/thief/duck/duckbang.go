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

func parseBang(raw []byte) (bangs []duckBang, err error) {
	err = json.Unmarshal(raw, &bangs)
	return
}

func fixBangs(bangs []duckBang) (err error) {
	for i := range bangs {
		if string(bangs[i].Title[0]) == " " {
			bangs[i].Title = bangs[i].Title[1:]
		}

		bangs[i].MainUrl, err = url.QueryUnescape(bangs[i].MainUrl)
		if err != nil {
			return
		}

		bangs[i].BangUrl, err = url.QueryUnescape(bangs[i].BangUrl)
		if err != nil {
			return
		}

		bangs[i].MainUrl = strings.ReplaceAll(bangs[i].MainUrl, "{{{s}}}", "%s")
		bangs[i].BangUrl = strings.ReplaceAll(bangs[i].BangUrl, "{{{s}}}", "%s")

		bangs[i].MainUrl = "http://" + bangs[i].MainUrl
	}
	return
}

func toBongMap(bangs []duckBang) (bong.BongMap, error) {
	err := fixBangs(bangs)
	if err != nil {
		return nil, err
	}

	bm := make(bong.BongMap)

	for _, bang := range bangs {
		bm[bang.Bang] = bong.Bong{
			Title:   bang.Title,
			MainUrl: bang.MainUrl,
			BongUrl: bang.BangUrl,
			Bongus:  bang.Bang,
		}
	}

	return bm, nil
}
