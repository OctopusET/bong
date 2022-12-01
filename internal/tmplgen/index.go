package tmplgen

import (
	"html/template"
	"io"
	"os"
)

type IndexParams struct {
	Title string
}

func Index(w io.Writer, p IndexParams) error {
	fs := os.DirFS("templates/default/")
	t, err := template.New("layout.html").ParseFS(fs, "layout.html", "index.html")
	if err != nil {
		return err
	}

	return t.Execute(w, p)
}
