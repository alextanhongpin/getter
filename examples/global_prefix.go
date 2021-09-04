package examples

//go:generate go run ../main.go -type=GlobalPrefix -out=global_prefix_gen.go -prefix=Get
type GlobalPrefix struct {
	name string
	age  int
}
