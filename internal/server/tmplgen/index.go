package tmplgen

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

type IndexParams struct {
	Title string
}

func Index(w io.Writer, p IndexParams) error {
	path := filepath.Join("templates/default")
	fs := os.DirFS(path)
	t, err := template.New("layout.html").ParseFS(fs, "layout.html", "index.html")
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}
