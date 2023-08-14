package page

import (
	"context"
	"html/template"
	"io"

	"github.com/swaggest/usecase"
)

type htmlResponseOutput struct {
	AntiHeader bool `header:"X-Anti-Header"`
	writer     io.Writer
}

func (o *htmlResponseOutput) Render(tmpl *template.Template) error {
	return tmpl.Execute(o.writer, o)
}

func HtmlResponse() usecase.Interactor {
	type htmlResponseInput struct {
		ID     int    `path:"id"`
		Filter string `query:"filter"`
		Header bool   `header:"X-Header"`
	}

	const tpl = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<a href="/html-response/{{.ID}}?filter={{.Filter}}">Next {{.Title}}</a><br />
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	tmpl, err := template.New("htmlResponse").Parse(tpl)
	if err != nil {
		panic(err)
	}

	u := usecase.NewInteractor(func(ctx context.Context, _ struct{}, out *htmlResponseOutput) (err error) {
		out.AntiHeader = true
		// out.Filter = in.Filter
		// out.ID = in.ID + 1
		// out.Title = "Foo"
		// out.Items = []string{"foo", "bar", "baz"}

		return out.Render(tmpl)
	})

	u.SetTitle("Request With HTML Response")
	u.SetDescription("Request with templated HTML response.")
	u.SetTags("Response")

	return u
}

// 	u := usecase.NewInteractor(func(ctx context.Context, in htmlResponseInput, out *htmlResponseOutput) (err error) {
// 		out.AntiHeader = !in.Header
// 		out.Filter = in.Filter
// 		out.ID = in.ID + 1
// 		out.Title = "Foo"
// 		out.Items = []string{"foo", "bar", "baz"}
//
// 		return out.Render(tmpl)
// 	})
//
// 	u.SetTitle("Request With HTML Response")
// 	u.SetDescription("Request with templated HTML response.")
// 	u.SetTags("Response")
//
// 	return u
// }
// func Page() usecase.Interactor {
// 	u := usecase.NewInteractor(func(_ context.Context, _ struct{}, output *http.ResponseWriter) error {
// 		templ := template.Must(template.ParseFiles("./cmd/api/page/index.html"))
// 		templ.Execute(*output, nil)
//
// 		// *output = books
// 		return nil
// 	})
// 	u.SetTags("Books")
// 	return u
// }
