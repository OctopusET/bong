package tohttps

import (
	"fmt"
	"net/http"
)

// Well, I really didn't want to do this, but some websites just blocked my
// plain request, resulting in failure checking if they support HTTPS.
const ua = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36`

func httpsSupported(url string) bool {
	strippedUrl := stripUrl(url)
	httpsUrl := "https://" + strippedUrl

	req, err := http.NewRequest("GET", httpsUrl, nil)
	if err != nil {
		fmt.Printf("error while checking url %s: %s\n\n", url, err.Error())
		return false
	}
	req.Header.Add("User-Agent", ua)

	client := new(http.Client)
	_, err = client.Do(req)

	if err != nil {
		fmt.Printf("error while checking url %s: %s\n\n", url, err.Error())
		return false
	}

	return true
}
