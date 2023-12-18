package models

import (
	"github.com/gofrs/uuid"
)

type Song struct {
	Id      *uuid.UUID `json:"id"`
	Artist  string     `json:"artist"`
	Title   string     `json:"title"`
	Album   string     `json:"album"`
	Content string     `json:"content"`
}
