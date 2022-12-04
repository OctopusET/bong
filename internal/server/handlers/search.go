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
	if ok := h.bongRedirect(w, r, query); ok {
		return
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
	var (
		realQuery string
		target    string
		bongus    string
	)

	splited := strings.Split(query, " ")
	if strings.HasPrefix(string(query), sh.Config.DefaultPrefix) {
		bongus = splited[0][len(sh.Config.DefaultPrefix):]
	}

	b, ok := sh.BongMap[bongus]

	if !ok {
		if sh.Config.Fallback == "" {
			return false
		} else if sh.Config.Fallback != "" {
			b, ok = sh.BongMap[sh.Config.Fallback]
			if !ok {
				return false
			}
			realQuery = strings.Join(splited, " ")
		}
	} else {
		realQuery = strings.Join(splited[1:], " ")
	}

	if realQuery != "" {
		target = fmt.Sprintf(b.BongUrl, realQuery)
	} else {
		target = fmt.Sprintf(b.MainUrl)
	}
	fmt.Printf("redirecting to %s\n", target)
	http.Redirect(w, r, target, http.StatusMovedPermanently)
	return true
}
