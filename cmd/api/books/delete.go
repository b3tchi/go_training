package books

import (
	"context"
	"errors"
	"strconv"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"web-hello/internal/db"
)

// Controller
func Delete() usecase.Interactor {
	type DeleteBookID struct {
		ID string `path:"id"`
	}
	type DeleteConfirm struct {
		Message string `json:"message"`
	}
	u := usecase.NewInteractor(func(_ context.Context, input DeleteBookID, output *DeleteConfirm) error {
		id, err := strconv.ParseInt(input.ID, 10, 64)
		if err != nil {
			return status.Internal
		}

		err = delete(id)
		if err != nil {
			switch {
			case err.Error() == "record not found":
				return status.NotFound
			default:
				return status.Internal
			}
		}

		*output = DeleteConfirm{Message: "succesfully deleted"}

		return nil
	})
	u.SetTags("Books")
	return u
}

// Handler
func delete(id int64) error {
	if id < 1 {
		return errors.New("record not found")
	}
	query := `
	 DELETE
	   FROM books
	  WHERE id = $1
	 `
	conn := db.GetDB()
	results, err := conn.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}
