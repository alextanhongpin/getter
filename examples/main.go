package examples

import (
	"database/sql"

	uuid "github.com/google/uuid"
)

//go:generate go run ../main.go -type=Foo,Bar
// Foo is an example struct.
type Foo struct {
	name       string `access:"r"`
	Age        int64
	Extra      string         `json:"extra" access:"w"`
	hobby      sql.NullString `access:"r"`
	tags       []string       `access:"r"`
	permission *string        `access:"r,w"`
}

type Bar struct {
	id   uuid.UUID `access:"r,w"`
	name string    `access:"r,w"`
	age  int64     `access:"r,w"`
	url  string    `access:"r,w"`
}
