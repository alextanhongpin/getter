package examples

//go:generate go run ../main.go -type=RenameField
type RenameField struct {
	name string `get:"RealName"`
	age  int
}
