package coward

import (
	"io"
	"net/http"
)

const (
	coward   = "https://search.brave.com"
	bangAddr = "/bang/data"
)

type cowFetcher struct{}

func (cf cowFetcher) fetch() ([]byte, error) {
	r, err := http.Get(coward + bangAddr)
	if err != nil {
		return nil, err
	}

	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}
