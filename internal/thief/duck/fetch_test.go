package duck

import (
	"testing"
)

func TestFetch(t *testing.T) {
	bf := new(bangFetcher)
	_, err := bf.fetch()

	if err != nil {
		panic(err)
	}
}
