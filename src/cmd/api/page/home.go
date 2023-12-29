package page

import (
	"context"
	"html/template"
	"io"

	"github.com/swaggest/usecase"
)

type htmlResponseOutput struct {
	writer     io.Writer
	Filter     string
	Title      string
	Items      []string
	ID         int
	AntiHeader bool `header:"X-Anti-Header"`
}

func (o *htmlResponseOutput) SetWriter(w io.Writer) {
	o.writer = w
}

func (o *htmlResponseOutput) Render(tmpl *template.Template) error {
	return tmpl.Execute(o.writer, o)
}

func HtmlResponse() usecase.Interactor {
	type htmlResponseInput struct {
		Filter string `query:"filter"`
		ID     int    `path:"id"`
		Header bool   `header:"X-Header"`
	}

	tmpl, err := template.New("index.html").ParseFiles("./cmd/api/page/index.html")
	if err != nil {
		panic(err)
	}

	u := usecase.NewInteractor(func(_ context.Context, in htmlResponseInput, out *htmlResponseOutput) (err error) {
		out.AntiHeader = !in.Header
		out.Filter = in.Filter
		out.ID = in.ID + 1
		out.Title = "Foo"
		out.Items = []string{"foo", "bar", "baz"}

		return out.Render(tmpl)
	})

	u.SetTitle("Request With HTML Response")
	u.SetDescription("Request with templated HTML response.")
	u.SetTags("Response")

	return u
}
