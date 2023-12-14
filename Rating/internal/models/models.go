package models

import (
	"github.com/gofrs/uuid"
)

type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Content string     `json:"content"`
}

type Rating struct {
	Id      *uuid.UUID `json:"id"`
	Score   string     `json:"score"`
	IdUser  *uuid.UUID `json:"idUser"`	
	IdSong  *uuid.UUID `json:"idSong"`	
	Content string     `json:"content"`
}
type InsertRating struct {
	Score   string     `json:"score"`
	IdUser  *uuid.UUID `json:"idUser"`	
	IdSong  *uuid.UUID `json:"idSong"`	
	Content string     `json:"content"`
}
type UpdateRating struct {
	Score   string     `json:"score"`
	Content string     `json:"content"`
}