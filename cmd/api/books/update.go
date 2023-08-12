package books

import (
	"context"
	"errors"
	"strconv"

	"github.com/lib/pq"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"web-hello/internal/db"
	"web-hello/internal/dto"
)

// controler
func Update() usecase.Interactor {
	type updateBook struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Rating    *float32 `json:"Rating"`
		ID        string   `path:"id"`
		Genres    []string `json:"genres"`
	}
	u := usecase.NewInteractor(func(_ context.Context, input updateBook, output *dto.Book) error {
		id, err := strconv.ParseInt(input.ID, 10, 64)
		if err != nil {
			return status.Wrap(errors.New("bad request"), status.Unavailable)
		}

		book, err := get(id)
		if err != nil {
			switch {
			case err.Error() == "record not found":
				return status.NotFound
			default:
				return status.Internal
			}
		}

		if input.Title != nil {
			book.Title = *input.Title
		}
		if input.Published != nil {
			book.Published = *input.Published
		}
		if input.Pages != nil {
			book.Pages = *input.Pages
		}
		if input.Genres != nil {
			book.Genres = input.Genres
		}
		if input.Rating != nil {
			book.Rating = *input.Rating
		}

		err = update(book)
		if err != nil {
			return status.Internal
		}

		*output = *book
		return nil
	})
	u.SetTags("Books")
	return u
}

// handler
func update(book *dto.Book) error {
	query := `
    UPDATE books SET title = $1
      , published = $2
      , pages = $3
      , genres = $4
      , rating = $5
      , version = version + 1
    WHERE id = $6 
      AND version = $7
    RETURNING version
  `
	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.Genres), book.Rating, book.ID, book.Version}

	conn := db.GetDB()
	return conn.QueryRow(query, args...).Scan(&book.Version)
}
