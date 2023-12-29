package dto

import (
	"time"
)

type Book struct {
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Genres    []string  `json:"genres,omitempty"`
	ID        int64     `json:"id"`
	Published int       `json:"published,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Version   int32     `json:"-"`
}
