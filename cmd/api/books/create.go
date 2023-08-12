package books

import (
	"context"

	"github.com/lib/pq"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"web-hello/internal/db"
	"web-hello/internal/dto"
)

// Controller
func Create() usecase.Interactor {
	type newBook struct {
		Title     string   `json:"title" required:"true"`
		Genres    []string `json:"genres"`
		Published int      `json:"published"`
		Pages     int      `json:"pages"`
		Rating    float32  `json:"Rating"`
	}

	u := usecase.NewInteractor(func(_ context.Context, input newBook, output *dto.Book) error {
		book := &dto.Book{
			Title:     input.Title,
			Published: input.Published,
			Pages:     input.Pages,
			Genres:    input.Genres,
			Rating:    input.Rating,
		}

		err := create(book)
		if err != nil {
			return status.Internal
		}

		// *output = envelope{"data": book}
		*output = *book
		return nil
	})

	u.SetTags("Books")
	u.SetExpectedErrors(status.Internal)

	return u
}

// Handler
func create(book *dto.Book) error {
	query := `
  INSERT INTO books (title, published, pages, genres, rating)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING id, created_at, version
  `

	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.Genres), book.Rating}
	conn := db.GetDB()
	return conn.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}
