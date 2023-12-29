package app

import (
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/web"
	"github.com/swaggest/swgui/v4emb"

	"web-hello/cmd/api/books"
	"web-hello/cmd/api/page"
)

func (app *application) Route() *web.Service {
	service := web.DefaultService()

	service.OpenAPISchema().SetTitle("Books Database")
	service.OpenAPISchema().SetDescription("database to manage books i read")
	service.OpenAPISchema().SetVersion(app.Version)

	service.Get("/books/{id}", page.HtmlResponse(), nethttp.SuccessfulResponseContentType("text/html"))
	// healthcheck
	service.Get("/v1/healthcheck", healthcheck())

	// collection
	// service.Get("/v1/books", app.GetBooks())
	service.Get("/v1/books", books.List())

	// item
	service.Post("/v1/books", books.Create())
	service.Get("/v1/books/{id}", books.Read())
	service.Put("/v1/books/{id}", books.Update())
	service.Delete("/v1/books/{id}", books.Delete())

	// docs
	service.Docs("/docs", v4emb.New)

	return service
}

func htmlResponse() {
	panic("unimplemented")
}
