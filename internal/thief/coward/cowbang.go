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

func fixBangs(bangs []cowBang) error {
	var err error

	for i := range bangs {
		bangs[i].BangUrl, err = url.QueryUnescape(bangs[i].BangUrl)
		if err != nil {
			return err
		}

		bangs[i].BangUrl = strings.ReplaceAll(bangs[i].BangUrl, "{query}", "%s")
	}

	return nil
}

func toBongMap(bangs []cowBang) (bong.BongMap, error) {
	err := fixBangs(bangs)
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
