package main

import (
	"github.com/tamudashe/un/pkg/repository"
	"html/template"
	"path/filepath"
)

// templateData is the holding structure for any dynamic data to be passed to HTML templates.
type templateData struct {
	CurrentYear int
	Snippet     *repository.Snippet
	Snippets    []*repository.Snippet
}

// newTemplateCache caches the templates
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		if err != nil {
			return nil, err
		}

		cache[name] = templateSet
	}

	return cache, nil
}
