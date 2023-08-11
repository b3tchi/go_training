package main

import (
	"github.com/swaggest/rest/web"
	"github.com/swaggest/swgui/v4emb"

	"web-hello/cmd/api/books"
)

func (app *application) route() *web.Service {
	service := web.DefaultService()

	service.OpenAPISchema().SetTitle("Books Database")
	service.OpenAPISchema().SetDescription("database to manage books i read")
	service.OpenAPISchema().SetVersion(version)

	// healthcheck
	service.Get("/v1/healthcheck", app.Healthcheck())

	// collection
	// service.Get("/v1/books", app.GetBooks())
	service.Get("/v1/books", books.GetBooks())

	// item
	service.Post("/v1/books", app.CreateBook())
	service.Get("/v1/books/{id}", app.ReadBook())
	service.Put("/v1/books/{id}", app.UpdateBook())
	service.Delete("/v1/books/{id}", app.DeleteBook())

	// docs
	service.Docs("/docs", v4emb.New)

	return service
}
