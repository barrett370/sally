package templates

import (
	"embed"
	"html/template"
)

var (
	//go:embed *.html
	templateFiles embed.FS
	Templates     = template.Must(template.ParseFS(templateFiles, "*.html"))
)
