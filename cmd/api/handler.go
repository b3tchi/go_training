package main

import (
	"context"
	"errors"
	"strconv"

	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"

	"web-hello/internal/data"
)

// Declare output port type.
func (app *application) Healthcheck() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, _ struct{}, output *envelope) error {
		data := map[string]string{
			"status":     "available",
			"enviroment": app.config.env,
			"version":    version,
		}

		*output = envelope{"healthcheck": data}

		return nil
	})
	return u
}

func (app *application) GetBooks() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, _ struct{}, output *envelope) error {
		books, err := app.models.Books.GetAll()
		if err != nil {
			return status.Internal
		}

		*output = envelope{"data": books}
		return nil
	})
	return u
}

func (app *application) CreateBook() usecase.Interactor {
	type newBook struct {
		Title     string   `json:"title" required:"true"`
		Published int      `json:"published"`
		Pages     int      `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    float32  `json:"Rating"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input newBook, output *envelope) error {
		book := &data.Book{
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

		*output = envelope{"data": book}
		return nil
	})
	return u
}

func (app *application) ReadBook() usecase.Interactor {
	type getBookID struct {
		ID string `path:"id"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input getBookID, output *envelope) error {
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

		*output = envelope{"data": book}
		return nil
	})
	return u
}

func (app *application) UpdateBook() usecase.Interactor {
	type updateBook struct {
		ID        string   `path:"id"`
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    *float32 `json:"Rating"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input updateBook, output *envelope) error {
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

		*output = envelope{"data": book}
		return nil
	})
	return u
}

func (app *application) DeleteBook() usecase.Interactor {
	type deleteBookID struct {
		ID string `path:"id"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input deleteBookID, output *envelope) error {
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
		*output = envelope{"message": "book succesfully deleted"}

		return nil
	})
	return u
}
