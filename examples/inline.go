package examples

//go:generate go run ../main.go -type=InlineParent -out=inline_gen.go
type InlineParent struct {
	InlineChildren `get:",inline"`
}

type InlineChildren struct {
	// May contain exported field.
	Name string

	// May contain unexported field.
	age int64
}
