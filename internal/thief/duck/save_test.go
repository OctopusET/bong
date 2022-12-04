package duck_test

import (
	"os"
	"testing"

	"github.com/npmania/bong/internal/thief/duck"
)

func TestUpdateBangs(t *testing.T) {
	if err := duck.UpdateBangs(); err != nil {
		panic(err)
	}
	os.RemoveAll("bongs")
}
