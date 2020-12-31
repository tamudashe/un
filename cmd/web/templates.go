package main

import (
	"github.com/tamudashe/un/pkg/repository"
)

// templateData is the holding structure for any dynamic data to be passed to HTML templates.
type templateData struct {
	Snippet  *repository.Snippet
	Snippets []*repository.Snippet
}

// newTemplateCache caches the templates
//func newTemplateCache(dir string) (map[string]*template.Template, error) {
//	cache := map[string]*template.Template{}
//
//	pages, err := filepath.Glob(filepath.Join(dir, "*.page."))
//}