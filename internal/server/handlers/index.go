package handlers

import (
	"fmt"
	"net/http"

	"github.com/npmania/bong/internal/config"
	tg "github.com/npmania/bong/internal/tmplgen"
)

type IndexHandler struct {
	Config config.Config
}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := tg.IndexParams{
		Title: h.Config.Title,
	}

	if err := tg.Index(w, data); err != nil {
		fmt.Println(err)
	}
}
