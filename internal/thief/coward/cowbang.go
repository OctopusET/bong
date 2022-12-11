package coward

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/npmania/bong/internal/bong"
)

type cowBang struct {
	Bang string `json:"bang"`
	Meta struct {
		Scheme   string `json:"scheme"`
		Netloc   string `json:"netloc"`
		Hostname string `json:"hostname"`
		Favicon  string `json:"favicon"`
		Path     string `json:"path"`
	} `json:"meta"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Title       string `json:"title"`
	BangUrl     string `json:"url"`
}

func parseBang(raw []byte) (bangs []cowBang, err error) {
	err = json.Unmarshal(raw, &bangs)
	return
}

func fixBangs(bangs []cowBang) (fixed []cowBang, err error) {
	for _, b := range bangs {
		b.BangUrl, err = url.QueryUnescape(b.BangUrl)
		if err != nil {
			return nil, err
		}

		b.Meta.Hostname = strings.ReplaceAll(b.Meta.Hostname, "%", "%%")
		b.BangUrl = strings.ReplaceAll(b.BangUrl, "%", "%%")

		b.Meta.Hostname = strings.ReplaceAll(b.Meta.Hostname, "{query}", "%s")
		b.BangUrl = strings.ReplaceAll(b.BangUrl, "{query}", "%s")

		fixed = append(fixed, b)
	}

	return fixed, nil
}

func toBongMap(bangs []cowBang) (bong.BongMap, error) {
	bangs, err := fixBangs(bangs)
	if err != nil {
		return nil, err
	}

	bm := make(bong.BongMap)

	for _, bang := range bangs {
		bm[bang.Bang] = bong.Bong{
			Title:   bang.Title,
			MainUrl: bang.Meta.Scheme + "://" + bang.Meta.Hostname,
			BongUrl: bang.BangUrl,
			Bongus:  bang.Bang,
		}
	}

	return bm, nil
}
