package coward_test

import (
	"os"
	"testing"

	cow "github.com/npmania/bong/internal/thief/coward"
)

func TestUpdateBangs(t *testing.T) {
	if err := cow.UpdateBangs(); err != nil {
		panic(err)
	}
	os.RemoveAll("bongs")
}
