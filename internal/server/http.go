package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/npmania/bong/internal/bong"
	"github.com/npmania/bong/internal/config"
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

	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	bongs, _ := bong.LoadBongs("bongs/duckduckgo-v260.yaml")

	indexHandler := handlers.IndexHandler{Config: config}
	searchHandler := handlers.SearchHandler{
		Config: config,
		Bongs:  bongs,
	}
	OpenSearchHandler := handlers.OpenSearchHandler{Config: config}

	mux := http.NewServeMux()

	mux.Handle("/", indexHandler)
	mux.Handle("/search", searchHandler)
	mux.Handle("/opensearch.xml", OpenSearchHandler)

	addr := ":" + strconv.Itoa(h.Port)

	fmt.Println("Starting http server at port", h.Port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		panic(err)
	}
}
