// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package server

import (
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	Text      string    `json:"text"`
}
