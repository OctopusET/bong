package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/npmania/bong/internal/bong"
	hg "github.com/npmania/bong/internal/htmlgen"
)

var bongs []bong.Bong

type HttpServer struct {
	Port        int
	initialized bool
}

func (h *HttpServer) initialize() {
	if h.Port == 0 {
		h.Port = 1441
	}
}

func (h HttpServer) Start() {
	if !h.initialized {
		h.initialize()
	}
	bongs, _ = bong.LoadBongs("bongs/duckduckgo-v260.yaml")
	http.HandleFunc("/search", search)

	addr := ":" + strconv.Itoa(h.Port)

	fmt.Println("Starting http server at port", h.Port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	var query string

	if err := r.ParseForm(); err != nil {
		//TODO: when does this error get triggered?
		fmt.Println("wrong post data!")
	}

	query = r.FormValue("q")
	if string(query[0]) == "!" {
		bongRedirect(w, r, query[1:])
	}

	data := hg.SearchParams{
		Title: "bong",
		Query: query,
	}

	fmt.Printf("%s %s %+v\n", r.Method, r.URL, r.Form)
	if err := hg.Search(w, data); err != nil {
		fmt.Println(err)
	}
}

// TODO: add error handling
func bongRedirect(w http.ResponseWriter, r *http.Request, query string) error {
	splited := strings.Split(query, " ")
	bongus := splited[0]
	url := bong.FindBong(bongs, bongus).BongUrl
	target := fmt.Sprintf(url, strings.Join(splited[1:], " "))
	fmt.Printf("redirecting to %s\n", target)
	http.Redirect(w, r, target, http.StatusMovedPermanently)
	return nil
}
