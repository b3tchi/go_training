package main

import (
	"context"
	"errors"
	"strconv"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	d "web-hello/internal/data"
)

// Declare output port type.
func (app *application) Healthcheck() usecase.Interactor {
	type checkState struct {
		Status      string `json:"status"`
		Environment string `json:"environment"`
		Version     string `json:"version"`
	}

	u := usecase.NewInteractor(func(_ context.Context, _ struct{}, output *checkState) error {
		data := checkState{
			Status:      "available",
			Environment: app.config.env,
			Version:     version,
		}

		*output = data
		return nil
	})
	u.SetTags("Health Check")
	return u
}

func (app *application) GetBooks() usecase.Interactor {
	u := usecase.NewInteractor(func(_ context.Context, _ struct{}, output *[]*d.Book) error {
		books, err := app.models.Books.GetAll()
		if err != nil {
			return status.Internal
		}

		*output = books
		return nil
	})
	u.SetTags("Books")
	return u
}

func (app *application) CreateBook() usecase.Interactor {
	type newBook struct {
		Title     string   `json:"title" required:"true"`
		Genres    []string `json:"genres"`
		Published int      `json:"published"`
		Pages     int      `json:"pages"`
		Rating    float32  `json:"Rating"`
	}

	u := usecase.NewInteractor(func(_ context.Context, input newBook, output *d.Book) error {
		book := &d.Book{
			Title:     input.Title,
			Published: input.Published,
			Pages:     input.Pages,
			Genres:    input.Genres,
			Rating:    input.Rating,
		}

		err := app.models.Books.Insert(book)
		if err != nil {
			return status.Internal
		}

		// *output = envelope{"data": book}
		*output = *book
		return nil
	})

	u.SetTags("Book")
	u.SetExpectedErrors(status.Internal)

	return u
}

func (app *application) ReadBook() usecase.Interactor {
	type getBookID struct {
		ID string `path:"id"`
	}
	u := usecase.NewInteractor(func(_ context.Context, input getBookID, output *d.Book) error {
		id, err := strconv.ParseInt(input.ID, 10, 64)
		if err != nil {
			return status.Wrap(errors.New("bad request"), status.Unavailable)
		}
		book, err := app.models.Books.Get(id)
		if err != nil {
			switch {
			case err.Error() == "record not found":
				return status.NotFound
			default:
				return status.Internal
			}
		}

		// *output = envelope{"data": book}
		*output = *book
		return nil
	})
	u.SetTags("Book")
	return u
}

func (app *application) UpdateBook() usecase.Interactor {
	type updateBook struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Rating    *float32 `json:"Rating"`
		ID        string   `path:"id"`
		Genres    []string `json:"genres"`
	}
	u := usecase.NewInteractor(func(_ context.Context, input updateBook, output *d.Book) error {
		id, err := strconv.ParseInt(input.ID, 10, 64)
		if err != nil {
			return status.Wrap(errors.New("bad request"), status.Unavailable)
		}

		book, err := app.models.Books.Get(id)
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

		err = app.models.Books.Update(book)
		if err != nil {
			return status.Internal
		}

		*output = *book
		return nil
	})
	u.SetTags("Book")
	return u
}

func (app *application) DeleteBook() usecase.Interactor {
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

		err = app.models.Books.Delete(id)
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
	u.SetTags("Book")
	return u
}
