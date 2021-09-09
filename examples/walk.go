package examples

//go:generate go run ../main.go -type TestWalk,Nested
type TestWalk struct {
	name   string
	age    int
	nested [3]Nested
	m1     map[string][]*Nested
	m2     map[string][][3]*map[string]*Nested
	m3     *Nested
	m4     [100]map[*Nested][]*map[*string]*Nested
}

type Nested struct {
	name string
}
