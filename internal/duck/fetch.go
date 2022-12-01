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
	bangVersionResp string

	bangAddrRegexp, _ = regexp.Compile(bangAddrExpr)
)

func Fetch() (string, error) {
	r, err := http.Get(duck + latestBangAddr())
	if err != nil {
		return "", err
	}

	respData, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(respData), nil
}

func latestBangVersion() int {
	bangAddr := latestBangAddr()

	verString := bangAddr[7:]
	verString = verString[:len(verString)-3]

	version, err := strconv.Atoi(verString)
	if err != nil {
		panic(err)
	}

	return version
}

func latestBangAddr() string {
	if bangVersionResp == "" {
		r, err := http.Get(duck + bangVersionUrl)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		respData, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		bangVersionResp = string(respData)
	}

	return bangAddrRegexp.FindString(bangVersionResp)
}
