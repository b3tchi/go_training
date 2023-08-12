package books

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/lib/pq"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"web-hello/internal/data"
	"web-hello/internal/db"
)

// Controller
func Read() usecase.Interactor {
	type getBookID struct {
		ID string `path:"id"`
	}
	u := usecase.NewInteractor(func(_ context.Context, input getBookID, output *data.Book) error {
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

		*output = *book
		return nil
	})
	u.SetTags("Books")
	return u
}

// Handler
func get(id int64) (*data.Book, error) {
	if id < 1 {
		return nil, errors.New("record not found")
	}

	query := `
  SELECT id, created_at, title, published, pages, genres, rating, version
  FROM books
  WHERE id = $1
  `

	var book data.Book

	conn := db.GetDB()
	err := conn.QueryRow(query, id).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.Title,
		&book.Published,
		&book.Pages,
		pq.Array(&book.Genres),
		&book.Rating,
		&book.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}
	return &book, nil
}
