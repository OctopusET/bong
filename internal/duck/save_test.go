package duck

import (
	"fmt"
	"os"
	"testing"
)

func TestGrabDuckBang(t *testing.T) {
	bangs, err := grabDuckBangs()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", bangs[0])
}

func TestUpdateBangs(t *testing.T) {
	if err := UpdateBangs(); err != nil {
		panic(err)
	}
	os.RemoveAll("bongs")
}
