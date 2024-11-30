package main

import (
	"log"
	"net/http"
)

type app struct {
	templateCache TemplateCache
}

func main() {

	// Cache all the templates
	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app := app{templateCache: templateCache}

	// Configure the server
	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	// and start the server
	log.Println("Listening on port 8080")
	srv.ListenAndServe()
}
