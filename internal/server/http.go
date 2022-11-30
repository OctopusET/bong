package server

import (
	"fmt"
	"net/http"
	"strconv"

	hg "github.com/npmania/bong/internal/htmlgen"
)

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

	data := hg.SearchParams{
		Title: "bong",
		Query: query,
	}

	fmt.Printf("%s %s %+v\n", r.Method, r.URL, r.Form)
	if err := hg.Search(w, data); err != nil {
		fmt.Println(err)
	}
}
