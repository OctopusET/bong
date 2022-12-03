package tohttps

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/npmania/bong/internal/bong"
)

func stripUrl(url string) string {
	return strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")
}

func hasInternetConnection() bool {
	// send an HTTP HEAD request to the specified URL
	resp, err := http.Head("http://www.example.com")
	if err != nil {
		return false
	}

	// check the response status code
	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func shuffle(bongs []bong.Bong) {
	for i := range bongs {
		j := rand.Intn(i + 1)
		bongs[i], bongs[j] = bongs[j], bongs[i]
	}
}

func urlToHttps(url string) string {
	return "https://" + stripUrl(url)
}

func urlToHttp(url string) string {
	return "http://" + stripUrl(url)
}
