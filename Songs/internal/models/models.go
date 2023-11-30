package models

import (
	"github.com/gofrs/uuid"
)

type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Content string     `json:"content"`

}
