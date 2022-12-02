package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/npmania/bong/internal/bong"
	"github.com/npmania/bong/internal/config"
	tg "github.com/npmania/bong/internal/tmplgen"
)

type SearchHandler struct {
	Config  config.Config
	BongMap bong.BongMap
}

func (h SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var query string

	if err := r.ParseForm(); err != nil {
		//TODO: when does this error get triggered?
		fmt.Println("wrong post data!")
	}

	query = r.FormValue("q")
	if strings.HasPrefix(string(query), h.Config.DefaultPrefix) {
		if ok := h.bongRedirect(w, r, query[len(h.Config.DefaultPrefix):]); ok {
			return
		}
	}

	data := tg.SearchParams{
		Title: "bong",
		Query: query,
	}

	fmt.Printf("%s %s %+v\n", r.Method, r.URL, r.Form)
	if err := tg.Search(w, data); err != nil {
		fmt.Println(err)
	}
}

func (sh SearchHandler) bongRedirect(w http.ResponseWriter, r *http.Request, query string) bool {
	splited := strings.Split(query, " ")
	bongus := splited[0]

	b, ok := sh.BongMap[bongus]
	if !ok {
		return false
	}

	target := fmt.Sprintf(b.BongUrl, strings.Join(splited[1:], " "))
	fmt.Printf("redirecting to %s\n", target)
	http.Redirect(w, r, target, http.StatusMovedPermanently)
	return true
}
