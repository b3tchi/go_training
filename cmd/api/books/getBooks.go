package books

import (
	"context"

	"github.com/lib/pq"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	d "web-hello/internal/data"
	"web-hello/internal/db"
)

// Controller
func GetBooks() usecase.Interactor {
	u := usecase.NewInteractor(func(_ context.Context, _ struct{}, output *[]*d.Book) error {
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
func getAll() ([]*d.Book, error) {
	query := `
    SELECT *
  FROM books
  ORDER BY id`

	conn := db.GetDB()
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	// defer rows.Close()

	books := []*d.Book{}
	for rows.Next() {
		var book d.Book
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
