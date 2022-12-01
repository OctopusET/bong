package handlers

import (
	"fmt"
	"net/http"

	tg "github.com/npmania/bong/internal/tmplgen"
)

type IndexHandler struct{}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := tg.IndexParams{
		Title: "bong",
	}

	if err := tg.Index(w, data); err != nil {
		fmt.Println(err)
	}
}
