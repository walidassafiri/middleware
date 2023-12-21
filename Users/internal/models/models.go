package models

import (
	"github.com/gofrs/uuid"
)

type UserPublic struct {
	Id       *uuid.UUID `json:"id"`
	Name     string     `json:"name"`
	Username string     `json:"username"`
}
type User struct {
	Id       *uuid.UUID `json:"id"`
	Name     string     `json:"name"`
	Username string     `json:"username"`
}

type InsertUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
type UserId struct {
	Id string `json:"id"`
}
type UserName struct {
	Username string `json:"username"`
}
