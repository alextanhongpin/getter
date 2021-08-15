package examples

import (
	"database/sql"

	uuid "github.com/google/uuid"
)

type Age int64

//go:generate go run ../main.go -type=Foo,Bar
// Foo is an example struct.
type Foo struct {
	name         string
	Age          Age
	realAge      Age
	Extra        string
	hobby        sql.NullString
	tags         []string
	permission   *string
	skip         *string `read:"-"`
	ages         []Age
	nulls        []sql.NullInt64
	validByAge   map[Age]bool
	ageByString  map[string]Age
	agesByString map[string][]Age
}

type Bar struct {
	id   uuid.UUID
	name string
	age  int64 `read:"-"`
	url  string
}
