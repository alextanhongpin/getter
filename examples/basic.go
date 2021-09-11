package examples

import "github.com/google/uuid"

//go:generate go run ../main.go -type User -prefix Get
type User struct {
	id      uuid.UUID
	name    string  `get:"FullName"`
	remarks *string `get:"-"`
	acc     Account `get:",inline,Account"`
}

//go:generate go run ../main.go -type Account -prefix Get
type Account struct {
	id   int64
	name string
}
