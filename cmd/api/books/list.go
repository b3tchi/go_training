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
func List() usecase.Interactor {
	u := usecase.NewInteractor(func(_ context.Context, _ struct{}, output *[]*dto.Book) error {
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
func getAll() ([]*dto.Book, error) {
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

	books := []*dto.Book{}
	for rows.Next() {
		var book dto.Book
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
