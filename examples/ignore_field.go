package examples

//go:generate go run ../main.go -type=IgnoreField
type IgnoreField struct {
	name string `get:"-"`
	age  int
}
