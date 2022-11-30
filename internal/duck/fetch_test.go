package duck

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	_, err := Fetch()

	if err != nil {
		panic(err)
	}
}

func TestLatestBangVersion(t *testing.T) {
	ver := latestBangVersion()
	fmt.Println("current duckduckgo bang version is", ver)
}

func TestLatestBangAddr(t *testing.T) {
	addr := latestBangAddr()
	fmt.Println("latest duckduckgo bang is at", addr)
}
