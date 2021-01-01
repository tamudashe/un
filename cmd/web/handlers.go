package main

import (
	"fmt"
	"strconv"

	"github.com/tamudashe/un/pkg/repository"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.snippets.GetLatestSnippets()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.gohtml", &templateData{Snippets: s})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// pat doesn't strip the colon from the named capture key, so we need to get the value of ":id"
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.GetSnippetByID(id)
	if err == repository.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.gohtml", &templateData{Snippet: s})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.gohtml", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
