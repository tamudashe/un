package main

import (
	"bytes"
	"fmt"
	"time"

	"net/http"
	"runtime/debug"
)

// render renders html templates
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, data *templateData) {
	templateSet, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("template %s does not exist", name))
		return
	}

	buffer := new(bytes.Buffer)

	// inject the default data while executing the template
	err := templateSet.Execute(buffer, app.addDefaultData(data, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	buffer.WriteTo(w)
}

// addDefaultData adds default data to a template
func (app *application) addDefaultData(data *templateData, r *http.Request) *templateData {
	// used to avoid nill pointer dereference
	if data == nil {
		data = &templateData{}
	}

	data.CurrentYear = time.Now().Year()

	return data
}

// serverError helper writes an error message and stack trace to the errorLogger then sends a generic 500 Internal
// Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError helper sends a specific status code and corresponding description to the user. We'll use this later
// in the book to send responses like 400 "Bad Request" when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is a convenience wrapper around clientError which sends a 404 Not Found response the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
