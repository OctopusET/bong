package duck

import (
	"os"
	"testing"
)

func TestUpdateBangs(t *testing.T) {
	if err := UpdateBangs(); err != nil {
		panic(err)
	}
	os.RemoveAll("bongs")
}
