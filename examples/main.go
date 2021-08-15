package examples

import (
	"database/sql"
)

type FooType int64

//go:generate go run ../main.go -type=Foo,Bar
// Foo is an example struct.
type Foo struct {
	str           string
	i             int
	i32           int32
	i64           int64
	b             bool
	sliceStr      []string
	sliceInt      []int
	sliceBool     []bool
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
	skip     *string `read:"-"`
	aliasPtr *FooType
	alias    FooType
}

type Bar struct {
	foo     Foo
	version int64
}
