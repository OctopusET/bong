package tmplgen

import (
	"io"
	"os"
	"text/template"
)

type OpenSearchParams struct {
	ShortName string
	//TODO: add favicon, get/post, suggestions
	SearchUrl     string
	OpenSearchUrl string
}

func OpenSearch(w io.Writer, p OpenSearchParams) error {
	fs := os.DirFS("templates/common/")
	t, err := template.New("opensearch.xml").ParseFS(fs, "opensearch.xml")
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}
