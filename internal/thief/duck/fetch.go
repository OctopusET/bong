package duck

import (
	"io"
	"net/http"
	"regexp"
	"strconv"
)

const (
	duck = "https://duckduckgo.com"

	bangVersionUrl = "/bv1.js"
	bangAddrExpr   = "/bang.v([0-9]+).js"
)

var (
	bangAddrRegexp, _ = regexp.Compile(bangAddrExpr)
)

type bangFetcher struct {
	bangVersionResp string
}

func (bf bangFetcher) fetch() ([]byte, error) {
	r, err := http.Get(bf.latestBangAddr())
	if err != nil {
		return nil, err
	}

	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func (bf bangFetcher) latestVersion() int {
	bangAddr := bf.latestBangAddr()

	verString := bangAddr[len(duck+"/bang.v"):]
	verString = verString[:len(verString)-len(".js")]

	version, err := strconv.Atoi(verString)
	if err != nil {
		panic(err)
	}

	return version
}

func (bf *bangFetcher) latestBangAddr() string {
	if bf.bangVersionResp == "" {
		r, err := http.Get(duck + bangVersionUrl)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		respData, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		bf.bangVersionResp = string(respData)
	}

	return duck + bangAddrRegexp.FindString(bf.bangVersionResp)
}
