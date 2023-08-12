package books

import (
	"context"

	"github.com/lib/pq"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"web-hello/internal/data"
	"web-hello/internal/db"
)

// Controller
func List() usecase.Interactor {
	u := usecase.NewInteractor(func(_ context.Context, _ struct{}, output *[]*data.Book) error {
		books, err := getAll()
		if err != nil {
			return status.Internal
		}

		*output = books
		return nil
	})
	u.SetTags("Books")
	return u
}

// Handler
func getAll() ([]*data.Book, error) {
	query := `
    SELECT *
  FROM books
  ORDER BY id`

	conn := db.GetDB()
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []*data.Book{}
	for rows.Next() {
		var book data.Book
		err := rows.Scan(
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
			return nil, err
		}
		books = append(books, &book)

	}
	return books, nil
}
