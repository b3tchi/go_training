package main

import (
	"net/http"

	"github.com/swaggest/rest/web"
)

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	// mux.HandleFunc("/v1/books", app.getCreateBooksHandler)
	// mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler)
	return mux
}

func (app *application) route2() *web.Service {
	service := web.DefaultService()

	// healthcheck
	service.Get("/v1/healthcheck", app.Healthcheck())

	// collection
	service.Get("/v1/books", app.GetBooks())

	// item
	service.Post("/v1/books", app.CreateBook())
	service.Get("/v1/books/{id}", app.ReadBook())
	service.Put("/v1/books/{id}", app.UpdateBook())
	service.Delete("/v1/books/{id}", app.DeleteBook())

	return service
}
