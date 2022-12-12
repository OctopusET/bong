package tmplgen

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
)

type OpenSearchParams struct {
	ShortName string
	//TODO: add favicon, get/post, suggestions
	SearchUrl     string
	OpenSearchUrl string
}

func OpenSearch(w io.Writer, p OpenSearchParams) error {
	path := filepath.Join("templates", "common")
	fs := os.DirFS(path)
	t, err := template.New("opensearch.xml").ParseFS(fs, "opensearch.xml")
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}
