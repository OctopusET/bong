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
	if strings.HasPrefix(string(query[0]), h.Config.DefaultPrefix) {
		h.bongRedirect(w, r, query[1:])
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

// TODO: add error handling
func (sh SearchHandler) bongRedirect(w http.ResponseWriter, r *http.Request, query string) error {
	splited := strings.Split(query, " ")
	bongus := splited[0]
	url := sh.BongMap[bongus].BongUrl
	target := fmt.Sprintf(url, strings.Join(splited[1:], " "))
	fmt.Printf("redirecting to %s\n", target)
	http.Redirect(w, r, target, http.StatusMovedPermanently)
	return nil
}
