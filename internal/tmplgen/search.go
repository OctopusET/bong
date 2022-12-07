package tmplgen

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

type SearchParams struct {
	Title string
	Query string
}

func Search(w io.Writer, p SearchParams) error {
	fmt.Println("data got:", p.Query)

	path := filepath.Join("templates", "default")
	fs := os.DirFS(path)

	t, err := template.New("layout.html").ParseFS(fs, "layout.html", "search.html")
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}
