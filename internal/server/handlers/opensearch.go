package handlers

import (
	"fmt"
	"net/http"

	"github.com/npmania/bong/internal/config"
	tg "github.com/npmania/bong/internal/server/tmplgen"
)

type OpenSearchHandler struct {
	Config config.Config
}

func (h OpenSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := tg.OpenSearchParams{
		ShortName:     h.Config.Title,
		SearchUrl:     h.Config.BaseUrl + "/search",
		OpenSearchUrl: h.Config.BaseUrl + "/opensearch.xml",
	}

	if err := tg.OpenSearch(w, data); err != nil {
		fmt.Println(err)
	}
}
