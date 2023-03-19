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

func (h *HttpServer) initialize(cfg config.Config) {
	if h.Port == 0 {
		h.Port = cfg.Port
	}
}

func (h HttpServer) Start() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	if !h.initialized {
		h.initialize(config)
	}

	bongs, _ := bong.LoadBongs(config.BongFile)

	indexHandler := handlers.IndexHandler{Config: config}
	searchHandler := handlers.SearchHandler{
		Config:  config,
		BongMap: bongs,
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
