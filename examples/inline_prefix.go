package examples

//go:generate go run ../main.go -type=InlinePrefixParent -out=inline_prefix_gen.go
type InlinePrefixParent struct {
	InlineChildren `get:",inline,Inline"`
}

type InlinePrefixChild struct {
	// May contain exported field.
	Name string

	// May contain unexported field.
	age int64
}
