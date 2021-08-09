package examples

import "database/sql"

//go:generate go run ../main.go -type=Foo
// Foo is an example struct.
type Foo struct {
	name       string `accessor:"r"`
	Age        int64
	Extra      string         `json:"extra" accessor:"w"`
	hobby      sql.NullString `accessor:"r"`
	tags       []string       `accessor:"r"`
	permission *string        `accessor:"r,w"`
}
