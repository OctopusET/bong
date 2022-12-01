package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/npmania/bong/internal/bong"
	"github.com/npmania/bong/internal/server/handlers"
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
	bongs, _ := bong.LoadBongs("bongs/duckduckgo-v260.yaml")

	indexHandler := handlers.IndexHandler{}
	searchHandler := handlers.SearchHandler{
		Bongs: bongs,
	}

	mux := http.NewServeMux()

	mux.Handle("/", indexHandler)
	mux.Handle("/search", searchHandler)

	addr := ":" + strconv.Itoa(h.Port)

	fmt.Println("Starting http server at port", h.Port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}
