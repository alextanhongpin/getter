package examples

import (
	"database/sql"
)

type FooType int64

//go:generate go run ../main.go -type=Foo,Bar
//go:generate go run ../main.go -type=Fizz -prefix=Get
// Foo is an example struct.
type Foo struct {
	str           string `get:"String"`
	i             int
	i32           int32
	i64           int64
	b             bool
	sliceStr      []string
	sliceInt      []int
	sliceBool     []bool
	sliceStrp     []*string
	sliceIntp     []*int
	sliceBoolp    []*bool
	sliceBoolFoo  []Foo
	sliceBoolFoop []*Foo
	boolsByString map[string][]bool
	boolByString  map[string]bool

	fooTypesByString    map[string][]FooType
	fooTypePtrsByString map[string][]*FooType
	stringByFooType     map[FooType]string
	stringByPtrFooType  map[*FooType]string

	strPtr  *string
	intPtr  *int
	boolPtr *bool

	imported sql.NullString
	skip     *string `get:"-"`
	aliasPtr *FooType
	alias    FooType

	fizz Fizz `get:",inline"`        // Flatten Fizz fields
	buzz Buzz `get:",inline,BuzzIt"` // Add a prefix for the nested fields.
}

type Bar struct {
	foo     Foo
	version int64
}

type Fizz struct {
	Remarks *string
	Ignore  *string `get:"-"`
	bar     *Bar    `get:"NestedBar"`
}

type Buzz struct {
	Name string
	Age  int64
}
