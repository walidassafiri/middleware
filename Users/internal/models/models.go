package models

import (
	"github.com/gofrs/uuid"
)

type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Content string     `json:"content"`
}
type UserPublic struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Mail string     `json:"mail"`
}
type User struct {
	Id       *uuid.UUID `json:"id"`
	Name     string     `json:"name"`
	Mail     string     `json:"mail"`
	Password string     `json:"password"`
}

type InsertUser struct {
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
type UserId struct {
	Id string `json:"id"`
}
