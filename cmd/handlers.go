package main

import (
	"net/http"
)

func (app *app) Home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "hello.page.html", nil)
}
func (app *app) About(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.page.html", nil)
}
func (app *app) Contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact.page.html", nil)
}
